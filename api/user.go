package api

import (
	"github.com/gin-gonic/gin"
	"go-v1/dto"
	"go-v1/service"
	"net/http"
)

func SendCode(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		c.JSON(http.StatusOK, userService.SendCode())
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func Register(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		c.JSON(http.StatusOK, userService.Register(c))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func LoginByCode(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		c.JSON(http.StatusOK, userService.LoginByCode(c))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func LoginByPassword(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		c.JSON(http.StatusOK, userService.LoginByPassword(c))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UploadAvatar(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		userService.Name = c.Param("Name")
		c.JSON(http.StatusOK, userService.UploadAvatar(c))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UpdateName(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		userService.Name = c.Param("Name")
		newName := c.PostForm("newName")
		c.JSON(http.StatusOK, userService.UpdateName(newName))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UpdatePassword(c *gin.Context) {
	userService := service.NewUserService()
	if err := c.ShouldBind(userService); err == nil {
		userService.Name = c.Param("Name")
		c.JSON(http.StatusOK, userService.UpdatePassword())
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

//func SendMoment(c *gin.Context) {
//	userService := service.NewUserService()
//	//if err := c
//}
