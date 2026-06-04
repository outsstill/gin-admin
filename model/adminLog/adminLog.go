package adminLog

import (
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminUser"
)

type AdminLog struct {
	model.BaseModel
	AdminUser adminUser.AdminUser `json:"admin_user" gorm:"foreignKey:UserId;references:ID"`
	UserId    uint64              `json:"user_id" gorm:"column:user_id"`
	Path      string              `json:"path" gorm:"column:path;type:varchar(255);index"`
	Url       string              `json:"url" gorm:"column:url;type:text"`
	Method    string              `json:"method" gorm:"column:method;type:varchar(255);index"`
	Ip        string              `json:"ip" gorm:"column:ip;type:varchar(255);index"`
	Input     string              `json:"input" gorm:"type:longtext;column:input"`
	model.CommonTimestampsField
}

func (model *AdminLog) TableName() string {
	return "admin_logs"
}
