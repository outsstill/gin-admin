package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminLog"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"github.com/outsstill/gin-admin/setting"
	"gorm.io/gorm"
)

type AdminLogService struct {
	DB *gorm.DB
}

func (service *AdminLogService) Name() string {
	return "AdminLogService"
}

func NewAdminLogService(db *gorm.DB) *AdminLogService {
	return &AdminLogService{
		DB: db,
	}
}

func (service *AdminLogService) Create(model *adminLog.AdminLog) {
	service.DB.Create(model)
}

func (service *AdminLogService) Save(model *adminLog.AdminLog) (rowsAffected int64) {
	result := service.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminLogService) Get(idstr string) (model *adminLog.AdminLog) {
	service.DB.Where("id", idstr).Preload("AdminUser").First(&model)
	return
}

func (service *AdminLogService) Delete(model *adminLog.AdminLog) (rowsAffected int64) {
	result := service.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminLogService) All() (models []adminLog.AdminLog) {
	service.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminLogService) Paginate(c *gin.Context, perPage int) (data []adminLog.AdminLog, paging paginator.Paging) {
	db := service.DB.Model(adminLog.AdminLog{})

	if c.Query("status") != "" {
		db = db.Where("status = ?", c.Query("status"))
	}

	if c.Query("path") != "" {
		db = db.Where("path LIKE ?", c.Query("path")+"%")
	}

	if c.Query("ip") != "" {
		db = db.Where("ip LIKE ?", c.Query("ip")+"%")
	}

	paging = paginator.Paginate(
		c,
		db,
		&data,
		setting.VADMINURL(model.TableName(&adminLog.AdminLog{})),
		perPage,
	)
	return
}
