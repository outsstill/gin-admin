package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/spf13/cast"
)

// BaseAPIController 基础控制器
type BaseAPIController struct {
	App *core.App
}

func (base *BaseAPIController) GetPerPage(c *gin.Context) int {
	key := global.Config.Paging.UrlQueryPerPage

	if len(key) == 0 {
		key = "per_page"
	}

	defaultPerPage := global.Config.Paging.PerPage

	if defaultPerPage <= 0 {
		defaultPerPage = 10
	}

	return cast.ToInt(c.DefaultQuery(key, cast.ToString(defaultPerPage)))
}
