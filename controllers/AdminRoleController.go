package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/go-kit/response"
)

type AdminRoleController struct {
	*BaseAPIController
}

func NewAdminRoleController(base *BaseAPIController) *AdminRoleController {
	return &AdminRoleController{
		BaseAPIController: base,
	}
}

func (uc *AdminRoleController) Index(c *gin.Context) {

	data, pager := uc.App.GetAdminRoleService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminRoleController) All(c *gin.Context) {

	roles := uc.App.GetAdminRoleService().All()

	response.Data(c, roles)
}

func (uc *AdminRoleController) Get(c *gin.Context) {

	user := uc.App.GetAdminRoleService().Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminRoleController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminRoleStoreRequest{}
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminRoleStore); !ok {
		return
	}

	uc.App.GetAdminRoleService().Create(c, &request)
	response.Success(c)
}

func (uc *AdminRoleController) Update(c *gin.Context) {
	model := uc.App.GetAdminRoleService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminRoleUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminRoleUpdate); !ok {
		return
	}

	uc.App.GetAdminRoleService().Update(c, &request, model)

	response.Data(c, model)
}

func (uc *AdminRoleController) UpdateMenus(c *gin.Context) {
	model := uc.App.GetAdminRoleService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminRoleUpdateMenusRequest{}
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminRoleMenusUpdate); !ok {
		return
	}

	uc.App.GetAdminRoleService().UpdateMenus(c, &request, model)

	response.Data(c, model)
}

func (uc *AdminRoleController) UpdatePermissions(c *gin.Context) {
	model := uc.App.GetAdminRoleService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminRoleUpdatePermissionsRequest{}
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminRolePermissionsUpdate); !ok {
		return
	}

	uc.App.GetAdminRoleService().UpdatePermissions(c, &request, model)

	response.Data(c, model)
}

func (uc *AdminRoleController) Delete(c *gin.Context) {
	model := uc.App.GetAdminRoleService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := uc.App.GetAdminRoleService().Delete(model); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
