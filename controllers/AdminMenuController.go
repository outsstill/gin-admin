package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
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

	data, pager := uc.App.GetAdminMenuService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminMenuController) All(c *gin.Context) {

	menus := uc.App.GetAdminMenuService().All()

	response.Data(c, menus)
}

func (uc *AdminMenuController) Get(c *gin.Context) {

	user := uc.App.GetAdminMenuService().Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminMenuController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminMenuStoreRequest{}
	if ok := requests.ValidateFunc(c, uc.App.DB, &request, requests.VerityAdminMenuStore); !ok {
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

	uc.App.GetAdminMenuService().Create(u)

	response.Data(c, u)
}

func (uc *AdminMenuController) Update(c *gin.Context) {
	model := uc.App.GetAdminMenuService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminMenuUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, uc.App.DB, &request, requests.VerityAdminMenuUpdate); !ok {
		return
	}

	model.Uri = request.Uri
	model.Path = request.Path
	model.Icon = request.Icon
	model.Name = request.Name
	model.Order = request.Order
	model.ParentId = request.ParentId

	uc.App.GetAdminMenuService().Save(model)

	response.Data(c, model)
}

func (uc *AdminMenuController) Delete(c *gin.Context) {
	model := uc.App.GetAdminMenuService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	if res := uc.App.GetAdminMenuService().Delete(model); res > 0 {
		response.Success(c)
		return
	}

	response.Fail(c, "删除失败")

}
