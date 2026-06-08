package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"gorm.io/gorm"
)

type AdminMenuService struct {
	DB *gorm.DB
}

func NewAdminMenuService(db *gorm.DB) *AdminMenuService {
	return &AdminMenuService{
		DB: db,
	}
}

func (service *AdminMenuService) Create(model *adminMenu.AdminMenu) {
	service.DB.Create(model)
}

func (service *AdminMenuService) Save(model *adminMenu.AdminMenu) (rowsAffected int64) {
	result := service.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminMenuService) Delete(model *adminMenu.AdminMenu) (rowsAffected int64) {
	result := service.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminMenuService) Get(idstr string) (model *adminMenu.AdminMenu) {
	service.DB.Where("id", idstr).First(&model)
	return
}

func (service *AdminMenuService) All() (models []adminMenu.AdminMenu) {
	service.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminMenuService) Paginate(c *gin.Context, perPage int) (users []adminMenu.AdminMenu, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.DB.Model(adminMenu.AdminMenu{}),
		&users,
		global.Config.VADMINURL(model.TableName(&adminMenu.AdminMenu{})),
		perPage,
	)
	return
}
