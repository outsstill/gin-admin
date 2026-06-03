package adminMenu

import (
	"github.com/outsstill/gin-admin/model"
)

type AdminMenu struct {
	model.BaseModel
	ParentId uint64 `json:"parent_id" gorm:"column:parent_id"`
	Order    int64  `json:"order" gorm:"column:order"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255)"`
	Icon     string `json:"icon" gorm:"column:icon;type:varchar(255)"`
	Path     string `json:"path" gorm:"column:path;type:varchar(255)"`
	Uri      string `json:"uri" gorm:"column:uri;type:varchar(255)"`
	model.CommonTimestampsField
}

func (model *AdminMenu) TableName() string {
	return "admin_menus"
}
