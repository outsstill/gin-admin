package adminPermission

import (
	"github.com/outsstill/gin-admin/model"
)

type AdminPermission struct {
	model.BaseModel
	Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
	Slug       string `json:"slug" gorm:"column:slug;type:varchar(255)"`
	HttpMethod string `json:"http_method" gorm:"column:http_method;type:varchar(255)"`
	HttpPath   string `json:"http_path" gorm:"column:http_path;type:varchar(255)"`
	Order      uint64 `json:"order" gorm:"column:order"`
	ParentId   uint64 `json:"parent_id" gorm:"column:parent_id"`
	model.CommonTimestampsField
}

func (model *AdminPermission) TableName() string {
	return "admin_permissions"
}
