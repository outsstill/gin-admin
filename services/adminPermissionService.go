package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"github.com/outsstill/gin-admin/setting"
	"gorm.io/gorm"
)

type AdminPermissionService struct {
	DB *gorm.DB
}

func NewAdminPermissionService(db *gorm.DB) *AdminPermissionService {
	return &AdminPermissionService{
		DB: db,
	}
}

func (service *AdminPermissionService) Create(model *adminPermission.AdminPermission) {
	service.DB.Create(model)
}

func (service *AdminPermissionService) Save(model *adminPermission.AdminPermission) (rowsAffected int64) {
	result := service.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminPermissionService) Delete(model *adminPermission.AdminPermission) (rowsAffected int64) {
	result := service.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminPermissionService) Get(idstr string) (model *adminPermission.AdminPermission) {
	service.DB.Where("id", idstr).First(&model)
	return
}

func (service *AdminPermissionService) All() (models []adminPermission.AdminPermission) {
	service.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminPermissionService) Paginate(c *gin.Context, perPage int) (users []adminPermission.AdminPermission, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.DB.Model(adminPermission.AdminPermission{}),
		&users,
		setting.VADMINURL(model.TableName(&adminPermission.AdminPermission{})),
		perPage,
	)
	return
}
