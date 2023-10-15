package api

import (
	"github.com/gin-gonic/gin"
	"go-v1/dto"
	"go-v1/service"
	"net/http"
)

func SendMoment(c *gin.Context) {
	momentService := service.NewMomentService()
	if err := c.ShouldBind(momentService); err == nil {
		c.JSON(http.StatusOK, momentService.SendMoment(c))
	} else {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}
