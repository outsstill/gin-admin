package controllers

import (
	"fmt"

	service "github.com/outsstill/gin-admin/services"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/auth"
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

	data, pager := service.NewAdminUserService(uc.App).Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminUserController) Get(c *gin.Context) {

	user := service.NewAdminUserService(uc.App).Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminUserController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminUserStoreRequest{}
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminUserStore); !ok {
		return
	}

	service.NewAdminUserService(uc.App).Create(c, &request)

	response.Success(c)
}

func (uc *AdminUserController) Update(c *gin.Context) {
	userModel := service.NewAdminUserService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminUserUpdateRequest{}
	request.ID = userModel.ID
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminUserUpdate); !ok {
		return
	}

	service.NewAdminUserService(uc.App).Update(c, &request, &userModel)

	response.Data(c, userModel)
}

func (uc *AdminUserController) Delete(c *gin.Context) {

	fmt.Printf("current_user_id :%v", auth.NewAuth(uc.App).CurrentAdminUser(c).ID)

	userModel := service.NewAdminUserService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := service.NewAdminUserService(uc.App).Delete(&userModel); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
