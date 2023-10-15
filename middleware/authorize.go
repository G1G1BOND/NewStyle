package middleware

import (
	"github.com/gin-gonic/gin"
	"go-v1/database"
	"go-v1/model"
	"go-v1/utils"
	"net/http"
	"strings"
)

var userID int

// 身份认证中间件
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 gin 上下文的 Header 中获取身份验证信息
		tokenString := c.GetHeader("Authorization")
		// 判断是否有 token 信息或者 token 是否以 Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		// token切片 去除Bearer (此处有个空格)
		tokenString = tokenString[7:]
		// 解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		// 获取userid
		userId := claims.StandardClaims.Id
		//userID = userId
		// 查询数据
		db := database.GetDB()
		var user model.User
		// 此处直接查询主键 即ID
		if !(db.Find(&user, userId).RowsAffected > 0) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		// 将user传入gin上下文
		c.Set("user", user)
		// 执行下一步
		c.Next()
	}
}
