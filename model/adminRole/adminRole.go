package adminRole

import (
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/model/adminPermission"
)

type AdminRole struct {
	model.BaseModel
	Name        string                            `json:"name" gorm:"column:name;type:varchar(255)"`
	Slug        string                            `json:"slug" gorm:"column:slug;type:varchar(255)"`
	Permissions []adminPermission.AdminPermission `json:"permissions" gorm:"many2many:admin_role_permissions;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:PermissionID"`
	Menus       []adminMenu.AdminMenu             `json:"menus" gorm:"many2many:admin_role_menus;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:MenuID"`
	model.CommonTimestampsField
}

func (model *AdminRole) TableName() string {
	return "admin_roles"
}
