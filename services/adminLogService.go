package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminLog"
	"github.com/outsstill/gin-admin/pkg/paginator"
)

type AdminLogService struct {
	app *core.App
}

func NewAdminLogService(app *core.App) *AdminLogService {
	return &AdminLogService{
		app: app,
	}
}

func (service *AdminLogService) Create(model *adminLog.AdminLog) {
	service.app.DB.Create(model)
}

func (service *AdminLogService) Save(model *adminLog.AdminLog) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminLogService) Delete(model *adminLog.AdminLog) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminLogService) Get(idstr string) (model adminLog.AdminLog) {
	service.app.DB.Where("id", idstr).Preload("AdminUser").First(&model)
	return
}

func (service *AdminLogService) All() (models []adminLog.AdminLog) {
	service.app.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminLogService) Paginate(c *gin.Context, perPage int) (data []adminLog.AdminLog, paging paginator.Paging) {
	db := service.app.DB.Model(adminLog.AdminLog{})

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
		global.Config.VADMINURL(model.TableName(&adminLog.AdminLog{})),
		perPage,
	)
	return
}
