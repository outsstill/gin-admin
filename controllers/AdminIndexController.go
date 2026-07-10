package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/logger"
	"github.com/outsstill/gin-admin/setting"
)

type AdminIndexController struct {
	*BaseAPIController
}

func NewAdminIndexController(base *BaseAPIController) *AdminIndexController {
	return &AdminIndexController{
		BaseAPIController: base,
	}
}

func (ic *AdminIndexController) Index(c *gin.Context) {

	logger.Info("AdminIndexController Index")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func (ic *AdminIndexController) Version(c *gin.Context) {
	fmt.Printf("%v \n", setting.Get())
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": setting.App().Version,
	})
}

func (ic *AdminIndexController) LimitTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "LimitTest",
	})
}
