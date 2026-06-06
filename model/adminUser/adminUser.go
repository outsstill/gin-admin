package adminUser

import (
	"fmt"
	"strings"

	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/model/adminRole"
	"github.com/outsstill/gin-admin/model/file"
	"github.com/outsstill/gin-admin/pkg/hash"
	"gorm.io/gorm"
)

type AdminUser struct {
	model.BaseModel
	Username      string                            `json:"username" gorm:"column:username;type:varchar(255)"`
	Password      string                            `json:"-" gorm:"column:password;type:varchar(255)"`
	Name          string                            `json:"name" gorm:"column:name;type:varchar(255)"`
	AvatarFile    *file.File                        `json:"avatar" gorm:"foreignKey:AvatarId;references:ID"`
	AvatarId      *uint64                           `json:"avatar_id" gorm:"column:avatar_id"`
	Roles         []adminRole.AdminRole             `json:"roles" gorm:"many2many:admin_role_users;foreignKey:ID;joinForeignKey:UserID;references:ID;joinReferences:RoleID"`
	Permissions   []adminPermission.AdminPermission `json:"permissions" gorm:"-"`
	Menus         []adminMenu.AdminMenu             `json:"menus" gorm:"-"`
	ChildrenMenus []*RouteDTO                       `gorm:"-" json:"childrenMenus"`
	RealName      string                            `json:"realName" gorm:"-"`
	IsSuper       bool                              `json:"is_super" gorm:"column:is_super"`
	model.CommonTimestampsField
}

func (model *AdminUser) TableName() string {
	return "admin_users"
}

func (model *AdminUser) IsSuperAdmin() bool {
	return model.IsSuper
}

// ComparePassword 密码是否正确
func (model *AdminUser) ComparePassword(_password string) bool {
	return hash.BcryptCheckIn(_password, model.Password)
}

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (model *AdminUser) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(model.Password) {
		model.Password = hash.BcryptHash(model.Password)
	}
	return
}

type RouteDTO struct {
	Name      string      `json:"name"`      // 前端菜单名（替代 name）
	Path      string      `json:"path"`      // 路由路径
	Component string      `json:"component"` // 组件路径
	Icon      string      `json:"icon"`      // 图标
	Order     uint64      `json:"order"`     // 排序（前端可能用）
	Meta      RouteMeta   `json:"meta"`
	Children  []*RouteDTO `json:"children,omitempty"`
}

type RouteMeta struct {
	Title         string `json:"title"`
	Order         int    `json:"order"`
	Icon          string `json:"icon"` // 图标
	AffixTab      bool   `json:"affixTab,omitempty"`
	NoBasicLayout bool   `json:"noBasicLayout,omitempty"`
}

func BuildVbenRoutes(list []adminMenu.AdminMenu) []*RouteDTO {

	menuMap := make(map[uint64]*RouteDTO)
	var roots []*RouteDTO

	// 1. 转 DTO
	for _, item := range list {

		path := item.Path
		if !strings.HasPrefix("/", item.Path) {
			path = fmt.Sprintf("/%s", item.Path)
		}

		uri := item.Uri
		if !strings.HasPrefix("/", item.Uri) {
			uri = fmt.Sprintf("/%s", item.Uri)
		}

		menuMap[item.ID] = &RouteDTO{
			Name:      item.Name,
			Path:      path,
			Component: uri,
			Meta: RouteMeta{
				Title: item.Name,
				Order: int(item.Order),
				Icon:  item.Icon,
			},
			Children: []*RouteDTO{},
		}
	}

	// 2. 构建树
	for _, item := range list {
		node := menuMap[item.ID]

		if item.ParentId == 0 {
			roots = append(roots, node)
		} else {
			if parent, ok := menuMap[item.ParentId]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return roots
}
