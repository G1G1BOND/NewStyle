package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-v1/config"
	"go-v1/database"
	"go-v1/dto"
	"go-v1/message"
	"go-v1/model"
	"go-v1/utils"
	"log"
)

type UserService struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Avatar   string `form:"Avatar"`
	Code     string `form:"code"`
	Token    string `form:"token"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SendCode() *dto.Result {
	//创建redis客户端
	redisClient := config.NewRedisClient()

	code := message.Success

	//检验email正确性
	if isTrue := utils.VerifyEmailFormat(s.Email); !isTrue {
		code = message.InvalidEmail
		return dto.Fail(code, nil)
	}

	//检查验证码是否重复发送
	if cnt := redisClient.Exists(message.VerificationCodeKey + s.Email).Val(); cnt == 1 {
		//60s内已发送过
		code := message.RepeatSending
		return dto.Fail(code, nil)
	}

	//获取随机验证码并发送
	//s.Email = "1340274713@qq.com"
	vCode := utils.RandomCode(6)
	redisClient.Set(message.VerificationCodeKey+s.Email, vCode, message.VerificationCodeKeyTTL)
	utils.SendCode(vCode, s.Email)

	return dto.Success(message.Success, "发送成功")
}

func (s *UserService) Register(ctx context.Context) *dto.Result {
	//获取redis客户端
	redisClient := config.NewRedisClient()
	//获取数据库指针
	db := database.GetDB()
	//检验密码格式
	if isTrue := utils.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(message.WrongPasswordFormat, nil)
	}
	//验证码校验
	vCode := redisClient.Get(message.VerificationCodeKey + s.Email).Val()
	if vCode != s.Code || vCode == "" {
		return dto.Fail(message.WrongCode, nil)
	}
	//判断邮箱是否注册过
	if db.Where("email=?", s.Email).Find(&model.User{}).RowsAffected > 0 {
		return dto.Fail(message.RepeatEmail, nil)
	}
	//用户名为空
	if s.Name == "" {
		return dto.Fail(message.NilName, nil)
	}
	//判断用户名是否使用过
	if db.Where("name=?", s.Name).Find(&model.User{}).RowsAffected > 0 {
		return dto.Fail(message.RepeatName, nil)
	}
	//密码加密
	hashPassword, _ := utils.Encrypt(s.Password)
	//创建用户并持久化
	db.Create(&model.User{
		Name:     s.Name,
		Email:    s.Email,
		Password: hashPassword,
		//Avatar:   url,
		//Password: password,
	})

	return dto.Success(message.Success, "注册成功")
}

func (s *UserService) LoginByCode(ctx *gin.Context) *dto.Result {
	//获取数据库指针
	db := database.GetDB()
	//获取redis客户端
	redisClient := config.NewRedisClient()
	//检验邮箱
	if (db.Where("email=?", s.Email).Find(&model.User{}).RowsAffected == 0) {
		//邮箱未注册
		return dto.Fail(message.WrongAccountOrPassword, nil)
	}
	//检验验证码
	if vCode := redisClient.Get(message.VerificationCodeKey + s.Email).Val(); vCode != s.Code || vCode == "" {
		//验证码错误
		return dto.Fail(message.WrongCode, nil)
	}
	var user model.User
	db.Where("email = ?", s.Email).Find(&user)
	//生成token
	token, _ := utils.ReleaseToken(user)
	//将token保存到redis
	userPointer := &model.User{}
	db.Where("email = ?", s.Email).Find(userPointer)
	userDto := dto.BuildUser(userPointer, token)
	redisClient.Del(message.UserLoginInfo + user.Name) //防止用户重复登录导致生成多个token
	redisClient.HMSet(message.UserLoginInfo+user.Name, utils.StructToMap(userDto))
	redisClient.Expire(message.UserLoginInfo+user.Name, message.UserLoginInfoTTL)
	//返回用户信息
	return dto.Success(message.Success, userDto)
}

func (s *UserService) LoginByPassword(ctx *gin.Context) *dto.Result {

	//获取数据库指针
	db := database.GetDB()
	redisClient := config.NewRedisClient()

	//检验密码账号
	//检验邮箱
	if (db.Where("email=?", s.Email).Find(&model.User{}).RowsAffected == 0) {
		//邮箱未注册
		return dto.Fail(message.WrongAccountOrPassword, nil)
	}
	var user model.User
	db.Where("email = ?", s.Email).Find(&user)
	// 密码验证
	if !utils.Decode(user.Password, s.Password) {
		return dto.Fail(message.WrongAccountOrPassword, nil)
	}

	//生成token
	token, _ := utils.ReleaseToken(user)

	//将token保存到redis
	userPointer := &model.User{}
	db.Where("email = ?", s.Email).Find(userPointer)
	userDto := dto.BuildUser(userPointer, token)
	redisClient.Del(message.UserLoginInfo + user.Name) //防止用户重复登录导致生成多个token
	redisClient.HMSet(message.UserLoginInfo+user.Name, utils.StructToMap(userDto))
	redisClient.Expire(message.UserLoginInfo+user.Name, message.UserLoginInfoTTL)

	//返回用户信息
	return dto.Success(message.Success, userDto)
}

func (s *UserService) UploadAvatar(ctx *gin.Context) *dto.Result {
	//获取数据库指针
	db := database.GetDB()
	//redisClient := config.NewRedisClient()

	iconFile, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return dto.Fail(message.InvalidParam, err)
	}
	//校验文件
	if header.Size > (8 << 18) {
		return dto.Fail(message.IconTooBig, nil)
	}
	if typ := header.Header.Get("Content-Type"); typ != "image/png" &&
		typ != "image/gif" &&
		typ != "image/jpeg" &&
		typ != "image/jpg" &&
		typ != "image/bmp" {
		return dto.Fail(message.WrongPictureFormat, nil)
	}
	//若原先头像不是默认头像的话删除头像
	userI, _ := ctx.Get("user")
	if icon := userI.(model.User).Avatar; icon != "http://rsa9yybad.hn-bkt.clouddn.com/FmEya-O5gDPkQ55efxU2dSnOZ7UI" {
		if err := utils.DelImg(icon); err != nil {
			return dto.Fail(message.Error, err)
		}
	}
	//上传图片
	url, _ := utils.UpLoadFile(iconFile, header.Size)
	if err != nil {
		return dto.Fail(message.Error, err)
	}
	//修改数据库
	db.Model(&model.User{}).Where("name = ?", s.Name).Update("avatar", url)
	//if err = db.UploadIcon(s.NickName, url); err != nil {
	//	dto.Fail(message.Error, nil)
	//}
	//修改缓存
	//if err = redisClient.HSet(message.UserLoginInfo+s.Name, "Icon", url).Err(); err != nil {
	//	return dto.Fail(message.Error, nil)
	//}
	//ctx
	return dto.Success(message.Success, url)
}

// 修改姓名
func (s *UserService) UpdateName(newName string) *dto.Result {
	////获取redis客户端
	//redisClient := config.NewRedisClient()
	db := database.GetDB()
	//判断昵称是否合法
	if newName == "" {
		return dto.Fail(message.NilNickName, nil)
	}
	if newName == s.Name {
		return dto.Success(message.Success, "修改成功")
	}
	//检查昵称是否已存在
	if db.Where("name=?", newName).Find(&model.User{}).RowsAffected > 0 {
		return dto.Fail(message.RepeatName, nil)
	}
	//修改昵称
	//修改redis中的昵称
	//redisClient.HSet(message.UserLoginInfo+s.Name, "NickName", newName)
	//重命名key
	//redisClient.Rename(message.UserLoginInfo+s.Name, e.UserLoginInfo+newName)
	//修改mysql
	db.Model(&model.User{}).Where("name = ?", s.Name).Update("name", newName)
	return dto.Success(message.Success, "修改成功")
}

// 修改密码
func (s *UserService) UpdatePassword() *dto.Result {
	//获取redis客户端
	redisClient := config.NewRedisClient()
	db := database.GetDB()
	//检验密码格式
	if isTrue := utils.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(message.WrongPasswordFormat, nil)
	}
	var email string
	db.Select("email").Find(&model.User{}).Where("name = ?", s.Name).Scan(&email)
	log.Println(email)

	if code := redisClient.Get(message.VerificationCodeKey + email).Val(); code != s.Code || code == "" {
		return dto.Fail(message.WrongCode, nil)
	}
	//密码加密
	hashPassword, _ := utils.Encrypt(s.Password)
	db.Model(&model.User{}).Where("name = ?", s.Name).Update("password", hashPassword)
	return dto.Success(message.Success, "修改成功")
}
