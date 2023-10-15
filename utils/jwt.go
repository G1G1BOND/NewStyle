package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go-v1/model"
	"strconv"
	"time"
)

// 设置 token 密钥
var jwtKey = []byte("token_test")

// 设置 token 请求体
type Claims struct {
	jwt.StandardClaims
}

// 发放token
// 参数：user对象
// 返回：token字符串 是否成功
func ReleaseToken(user model.User) (string, error) {
	// 设置超时时间
	expiresTime := time.Now().Add(60 * time.Minute)
	// 创建token请求体
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Name,          // 受众
			ExpiresAt: expiresTime.Unix(), // 失效时间
			// 这里要将 uint 转换为十进制 string
			Id:        strconv.FormatUint(uint64(user.ID), 10), // 编号
			IssuedAt:  time.Now().Unix(),                       // 签发时间
			Issuer:    "test02",                                // 签发人
			NotBefore: time.Now().Unix(),                       // 生效时间
			Subject:   "login",                                 // 主题
		},
	}
	// 使用 HS256 加密算法加密 token 请求体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 通过 token 密钥获取完整的 token 令牌
	if token, err := tokenClaims.SignedString(jwtKey); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// 解析token
// 参数：token字符串
// 返回：token请求体 错误信息
func ParseToken(tokenString string) (*Claims, error) {
	// 将 token 解析为 token 请求体
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// 错误处理
	if err != nil {
		return nil, err
	}
	// 有效验证
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
