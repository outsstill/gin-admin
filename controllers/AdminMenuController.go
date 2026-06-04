package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
	service "github.com/outsstill/gin-admin/services"
)

type AdminMenuController struct {
	*BaseAPIController
}

func NewAdminMenuController(base *BaseAPIController) *AdminMenuController {
	return &AdminMenuController{
		BaseAPIController: base,
	}
}

func (uc *AdminMenuController) Index(c *gin.Context) {

	data, pager := service.NewAdminMenuService(uc.App).Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminMenuController) All(c *gin.Context) {

	menus := service.NewAdminMenuService(uc.App).All()

	response.Data(c, menus)
}

func (uc *AdminMenuController) Get(c *gin.Context) {

	user := service.NewAdminMenuService(uc.App).Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminMenuController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminMenuStoreRequest{}
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminMenuStore); !ok {
		return
	}

	u := &adminMenu.AdminMenu{
		Name:     request.Name,
		Icon:     request.Icon,
		Path:     request.Path,
		Uri:      request.Uri,
		Order:    request.Order,
		ParentId: request.ParentId,
	}

	service.NewAdminMenuService(uc.App).Create(u)

	response.Data(c, u)
}

func (uc *AdminMenuController) Update(c *gin.Context) {
	userModel := service.NewAdminMenuService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminMenuUpdateRequest{}
	request.ID = userModel.ID
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminMenuUpdate); !ok {
		return
	}

	userModel.Uri = request.Uri
	userModel.Path = request.Path
	userModel.Icon = request.Icon
	userModel.Name = request.Name
	userModel.Order = request.Order
	userModel.ParentId = request.ParentId

	service.NewAdminMenuService(uc.App).Save(&userModel)

	response.Data(c, userModel)
}

func (uc *AdminMenuController) Delete(c *gin.Context) {
	userModel := service.NewAdminMenuService(uc.App).Get(c.Param("id"))
	if userModel.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := service.NewAdminMenuService(uc.App).Delete(&userModel); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
