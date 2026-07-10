package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	gokit "github.com/outsstill/go-kit"
	"github.com/spf13/cast"
)

// BaseAPIController 基础控制器
type BaseAPIController struct {
	App *core.App
}

func (base *BaseAPIController) GetPerPage(c *gin.Context) int {
	key := gokit.Config().Paging.UrlQueryPerPage

	if len(key) == 0 {
		key = "per_page"
	}

	defaultPerPage := gokit.Config().Paging.PerPage

	if defaultPerPage <= 0 {
		defaultPerPage = 10
	}

	return cast.ToInt(c.DefaultQuery(key, cast.ToString(defaultPerPage)))
}
