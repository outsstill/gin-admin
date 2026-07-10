package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/go-kit/response"
)

type AdminPermissionController struct {
	*BaseAPIController
}

func NewAdminPermissionController(base *BaseAPIController) *AdminPermissionController {
	return &AdminPermissionController{
		BaseAPIController: base,
	}
}

func (uc *AdminPermissionController) Index(c *gin.Context) {

	data, pager := uc.App.GetAdminPermissionService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminPermissionController) All(c *gin.Context) {

	models := uc.App.GetAdminPermissionService().All()

	response.Data(c, models)
}

func (uc *AdminPermissionController) Get(c *gin.Context) {

	user := uc.App.GetAdminPermissionService().Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminPermissionController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminPermissionStoreRequest{}
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminPermissionStore); !ok {
		return
	}

	u := &adminPermission.AdminPermission{
		Name:       request.Name,
		Slug:       request.Slug,
		HttpMethod: strings.ToLower(request.HttpMethod),
		HttpPath:   request.HttpPath,
		Order:      request.Order,
		ParentId:   request.ParentId,
	}
	uc.App.GetAdminPermissionService().Create(u)

	response.Data(c, u)
}

func (uc *AdminPermissionController) Update(c *gin.Context) {
	model := uc.App.GetAdminPermissionService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminPermissionUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminPermissionUpdate); !ok {
		return
	}

	model.HttpMethod = strings.ToLower(request.HttpMethod)
	model.HttpPath = request.HttpPath
	model.Name = request.Name
	model.Slug = request.Slug
	model.Order = request.Order
	model.ParentId = request.ParentId

	uc.App.GetAdminPermissionService().Save(model)

	response.Data(c, model)
}

func (uc *AdminPermissionController) Delete(c *gin.Context) {
	model := uc.App.GetAdminPermissionService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := uc.App.GetAdminPermissionService().Delete(model); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
