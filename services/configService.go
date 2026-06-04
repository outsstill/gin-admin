package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	configModel "github.com/outsstill/gin-admin/model/config"
	"github.com/outsstill/gin-admin/pkg/paginator"
)

type ConfigService struct {
	app *core.App
}

func NewConfigService(app *core.App) *ConfigService {
	return &ConfigService{
		app: app,
	}
}

func (service *ConfigService) Create(model *configModel.Config) {
	service.app.DB.Create(model)
}

func (service *ConfigService) Save(model *configModel.Config) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *ConfigService) Delete(model *configModel.Config) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *ConfigService) Get(idstr string) (model configModel.Config) {
	service.app.DB.Where("id", idstr).First(&model)
	return
}

func (service *ConfigService) All() (models []configModel.Config) {
	service.app.DB.Find(&models)
	return
}

func (service *ConfigService) AllShow() (models []configModel.Config) {
	service.app.DB.Where("is_can_front = 1").Find(&models)
	return
}

// Paginate 分页内容
func (service *ConfigService) Paginate(c *gin.Context, perPage int) (users []configModel.Config, paging paginator.Paging) {

	db := service.app.DB.Model(configModel.Config{})

	if c.Query("config_key") != "" {
		db = db.Where("config_key LIKE ?", c.Query("config_key")+"%")
	}

	if c.Query("config_label") != "" {
		db = db.Where("config_label LIKE ?", c.Query("config_label")+"%")
	}

	paging = paginator.Paginate(
		c,
		db,
		&users,
		global.Config.VADMINURL(model.TableName(&configModel.Config{})),
		perPage,
	)
	return
}
