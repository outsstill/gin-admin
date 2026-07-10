package controllers

import (
	"github.com/outsstill/gin-admin/model/adminLog"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/outsstill/go-kit/response"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

type AdminLogController struct {
	*BaseAPIController
}

func NewAdminLogController(base *BaseAPIController) *AdminLogController {
	return &AdminLogController{
		BaseAPIController: base,
	}
}

func (uc *AdminLogController) Index(c *gin.Context) {

	data, pager := uc.App.GetAdminLogService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminLogController) All(c *gin.Context) {
	response.Data(c, uc.App.GetAdminLogService().All())
}

func (uc *AdminLogController) Get(c *gin.Context) {
	response.Data(c, uc.App.GetAdminLogService().Get(c.Param("id")))
}

func (uc *AdminLogController) Store(c *gin.Context) {

	u := &adminLog.AdminLog{
		UserId: cast.ToUint64(auth.CurrentAdminUID(c)),
	}

	uc.App.GetAdminLogService().Create(u)

	response.Data(c, u)
}

func (uc *AdminLogController) Update(c *gin.Context) {
	response.Success(c)
}

func (uc *AdminLogController) Delete(c *gin.Context) {
	response.Fail(c, "删除失败")

}
