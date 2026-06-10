package service

import (
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model"
	configModel "github.com/outsstill/gin-admin/model/config"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"github.com/outsstill/gin-admin/setting"
	"gorm.io/gorm"
)

type ConfigService struct {
	DB *gorm.DB
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{
		DB: db,
	}
}

func (service *ConfigService) Create(model *configModel.Config) {
	service.DB.Create(model)
}

func (service *ConfigService) Save(model *configModel.Config) (rowsAffected int64) {
	result := service.DB.Save(model)
	return result.RowsAffected
}

func (service *ConfigService) Delete(model *configModel.Config) (rowsAffected int64) {
	result := service.DB.Delete(model)
	return result.RowsAffected
}

func (service *ConfigService) Get(idstr string) (model *configModel.Config) {
	service.DB.Where("id", idstr).First(&model)
	return
}

func (service *ConfigService) All() (models []configModel.Config) {
	service.DB.Find(&models)
	return
}

func (service *ConfigService) AllShow() (models []configModel.Config) {
	service.DB.Where("is_can_front = 1").Find(&models)
	return
}

// Paginate 分页内容
func (service *ConfigService) Paginate(c *gin.Context, perPage int) (users []configModel.Config, paging paginator.Paging) {

	db := service.DB.Model(configModel.Config{})

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
		setting.VADMINURL(model.TableName(&configModel.Config{})),
		perPage,
	)
	return
}
