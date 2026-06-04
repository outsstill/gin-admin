package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/pkg/paginator"
)

type AdminPermissionService struct {
	app *core.App
}

func NewAdminPermissionService(app *core.App) *AdminPermissionService {
	return &AdminPermissionService{
		app: app,
	}
}

func (service *AdminPermissionService) Create(model *adminPermission.AdminPermission) {
	service.app.DB.Create(model)
}

func (service *AdminPermissionService) Save(model *adminPermission.AdminPermission) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminPermissionService) Delete(model *adminPermission.AdminPermission) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminPermissionService) Get(idstr string) (model adminPermission.AdminPermission) {
	service.app.DB.Where("id", idstr).Preload("Roles").Preload("AvatarFile").First(&model)
	return
}

func (service *AdminPermissionService) All() (models []adminPermission.AdminPermission) {
	service.app.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminPermissionService) Paginate(c *gin.Context, perPage int) (users []adminPermission.AdminPermission, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.app.DB.Model(adminPermission.AdminPermission{}),
		&users,
		global.Config.VADMINURL(model.TableName(&adminPermission.AdminPermission{})),
		perPage,
	)
	return
}
