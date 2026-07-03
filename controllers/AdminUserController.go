package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
)

type AdminUserController struct {
	*BaseAPIController
}

func NewAdminUserController(base *BaseAPIController) *AdminUserController {
	return &AdminUserController{
		BaseAPIController: base,
	}
}

func (uc *AdminUserController) Index(c *gin.Context) {

	data, pager := uc.App.GetAdminUserService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminUserController) Get(c *gin.Context) {

	user := uc.App.GetAdminUserService().Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminUserController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminUserStoreRequest{}
	if ok := requests.ValidateFunc(c, uc.App.DB, &request, requests.VerityAdminUserStore); !ok {
		return
	}

	uc.App.GetAdminUserService().Create(c, &request)

	response.Success(c)
}

func (uc *AdminUserController) Update(c *gin.Context) {

	id := c.Param("id")
	model := uc.App.GetAdminUserService().Get(id)
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	currentAdmin := uc.App.GetAuthService().CurrentAdminUser(c)

	if model.IsSuperAdmin() && !currentAdmin.IsSuperAdmin() {
		response.Fail(c, "非超级管理员不能修改超级管理员")
		return
	}

	// 验证
	request := requests.AdminUserUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, uc.App.DB, &request, requests.VerityAdminUserUpdate); !ok {
		return
	}

	uc.App.GetAdminUserService().Update(c, &request, model)

	response.Data(c, model)
}

func (uc *AdminUserController) Delete(c *gin.Context) {

	fmt.Printf("current_user_id :%v", uc.App.GetAuthService().CurrentAdminUser(c).ID)

	model := uc.App.GetAdminUserService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := uc.App.GetAdminUserService().Delete(model); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
