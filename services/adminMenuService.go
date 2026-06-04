package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/pkg/paginator"
)

type AdminMenuService struct {
	app *core.App
}

func NewAdminMenuService(app *core.App) *AdminMenuService {
	return &AdminMenuService{
		app: app,
	}
}

func (service *AdminMenuService) Create(model *adminMenu.AdminMenu) {
	service.app.DB.Create(model)
}

func (service *AdminMenuService) Save(model *adminMenu.AdminMenu) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminMenuService) Delete(model *adminMenu.AdminMenu) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminMenuService) Get(idstr string) (model adminMenu.AdminMenu) {
	service.app.DB.Where("id", idstr).Preload("Roles").Preload("AvatarFile").First(&model)
	return
}

func (service *AdminMenuService) All() (models []adminMenu.AdminMenu) {
	service.app.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminMenuService) Paginate(c *gin.Context, perPage int) (users []adminMenu.AdminMenu, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.app.DB.Model(adminMenu.AdminMenu{}),
		&users,
		global.Config.VADMINURL(model.TableName(&adminMenu.AdminMenu{})),
		perPage,
	)
	return
}
