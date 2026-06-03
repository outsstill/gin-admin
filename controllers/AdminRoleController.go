package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
	service "github.com/outsstill/gin-admin/services"
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

	data, pager := service.NewAdminRoleService(uc.App).Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminRoleController) All(c *gin.Context) {

	roles := service.NewAdminRoleService(uc.App).All()

	response.Data(c, roles)
}

func (uc *AdminRoleController) Get(c *gin.Context) {

	user := service.NewAdminRoleService(uc.App).Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminRoleController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminRoleStoreRequest{}
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminRoleStore); !ok {
		return
	}

	service.NewAdminRoleService(uc.App).Create(c, &request)
	response.Success(c)
}

func (uc *AdminRoleController) Update(c *gin.Context) {
	userModel := service.NewAdminRoleService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminRoleUpdateRequest{}
	request.ID = userModel.ID
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminRoleUpdate); !ok {
		return
	}

	service.NewAdminRoleService(uc.App).Update(c, &request, &userModel)

	response.Data(c, userModel)
}

func (uc *AdminRoleController) Delete(c *gin.Context) {
	userModel := service.NewAdminRoleService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := service.NewAdminRoleService(uc.App).Delete(&userModel); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
