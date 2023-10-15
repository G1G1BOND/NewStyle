package service

import (
	"github.com/gin-gonic/gin"
	"go-v1/database"
	"go-v1/dto"
	"go-v1/message"
	"go-v1/model"
	"go-v1/utils"
)

type MomentService struct {
	Name    string `form:"name"`
	Content string `form:"content"`
	Like    string `form:"like"`
	Avatar  string `form:"avatar"`
	//Picture string `form:"Picture"`
	//Code    string `form:"code"`
	//Token   string `form:"token"`
}

func NewMomentService() *MomentService {
	return &MomentService{}
}

func (s *MomentService) SendMoment(c *gin.Context) *dto.Result {
	//获取redis客户端
	//redisClient := config.NewRedisClient()

	//获取数据库指针
	db := database.MomentDB()

	//判断用户名是否存在
	if (db.Where("name=?", s.Name).Find(&model.User{}).RowsAffected == 0) {
		return dto.Fail(message.NotFoundName, nil)
	}

	//获取用户头像
	var avatar string
	db.Select("avatar").Where("name=?", s.Name).Find(&model.User{}).Scan(&avatar)

	//上传内容图片
	picture, fileHeader, _ := c.Request.FormFile("picture")

	fileSize := fileHeader.Size

	pictureUrl, code := utils.UpLoadFile(picture, fileSize)

	if code != 200 {
		return dto.Fail(message.WrongPictureFormat, nil)
	}

	//创建用户并持久化
	db.Create(&model.Moment{
		Name:    s.Name,
		Content: s.Content,
		Like:    s.Like,
		Avatar:  avatar,
		Picture: pictureUrl,
	})

	//返回存入信息
	return dto.Success(message.Success, "发布成功")
}

func GetMoment(c *gin.Context) {
	var moments []model.Moment
	db := database.MomentDB()
	_ = db.Find(&moments)
	c.JSON(200, gin.H{
		"status": 200,
		"data":   moments,
		"msg":    "OK",
	})

}
