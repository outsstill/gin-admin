package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/global"
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

	global.Logger.Info("AdminIndexController Index")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func (ic *AdminIndexController) Version(c *gin.Context) {
	fmt.Printf("%v \n", global.Config)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": global.Config.App.Version,
	})
}
