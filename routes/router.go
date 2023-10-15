package routes

import (
	"github.com/gin-gonic/gin"
	"go-v1/api"
	"go-v1/middleware"
	"go-v1/service"
	"go-v1/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.POST("/code", api.SendCode)
		router.POST("/register", api.Register)
		router.POST("/login/code", api.LoginByCode)
		router.POST("/login/password", api.LoginByPassword)

		router.PATCH("/password/:Name", middleware.Authorize(), api.UpdatePassword)
		router.PATCH("/edit/Name/:Name", middleware.Authorize(), api.UpdateName)

		//router.GET("/checkUsers", userService.CheckUsers)
		router.POST("/edit/avatar/:Name", middleware.Authorize(), api.UploadAvatar)

		router.POST("/send/moment", middleware.Authorize(), api.SendMoment)
		router.GET("/get/moment", middleware.Authorize(), service.GetMoment)
	}
	r.Run(utils.HttpPort)
}
