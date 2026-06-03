package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/model/adminRole"
	"github.com/outsstill/gin-admin/model/adminUser"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
	"github.com/spf13/cast"
)

type AdminUserService struct {
	app *core.App
}

func NewAdminUserService(app *core.App) *AdminUserService {
	return &AdminUserService{
		app: app,
	}
}

func (service *AdminUserService) GetUserPermissions(userID uint64) ([]adminPermission.AdminPermission, error) {
	var user adminUser.AdminUser
	if err := service.app.DB.
		Preload("Roles.Permissions").
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		return nil, err
	}

	if user.IsSuperAdmin() {
		return NewAdminPermissionService(service.app).All(), nil
	}

	permissionMap := make(map[uint64]adminPermission.AdminPermission)
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			permissionMap[perm.ID] = perm
		}
	}

	var permissions []adminPermission.AdminPermission
	for _, perm := range permissionMap {
		permissions = append(permissions, perm)
	}

	return permissions, nil
}

func (service *AdminUserService) GetUserMenus(userID uint64) ([]adminMenu.AdminMenu, error) {
	var user adminUser.AdminUser
	if err := service.app.DB.
		Preload("Roles.Menus").
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		return nil, err
	}

	var menus []adminMenu.AdminMenu

	if user.IsSuperAdmin() {
		return NewAdminMenuService(service.app).All(), nil
	}

	menuMap := make(map[uint64]adminMenu.AdminMenu)
	for _, role := range user.Roles {
		for _, menu := range role.Menus {
			menuMap[menu.ID] = menu
		}
	}

	for _, perm := range menuMap {
		menus = append(menus, perm)
	}

	return menus, nil
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (service *AdminUserService) Create(c *gin.Context, request *requests.AdminUserStoreRequest) {
	tx := service.app.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			response.Abort500(c, "系统错误")
		}
	}()

	// 查询角色
	var roles []adminRole.AdminRole

	if len(request.RoleIDs) > 0 {
		if err := tx.Where("id IN ?", request.RoleIDs).Find(&roles).Error; err != nil {
			response.BadRequest(c, err, "角色查询失败")
			return
		}
	}

	model := adminUser.AdminUser{
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
		Roles:    roles,
	}

	if request.AvatarId > 0 {
		model.AvatarId = helpers.Uint64Ptr(cast.ToUint64(request.AvatarId))
	}

	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "创建账号失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminUserService) Update(c *gin.Context, request *requests.AdminUserUpdateRequest, userModel *adminUser.AdminUser) {

	tx := service.app.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		}
	}()

	// 查询角色
	var roles []adminRole.AdminRole

	if len(request.RoleIDs) > 0 {
		if err := tx.Where("id IN ?", request.RoleIDs).Find(&roles).Error; err != nil {
			response.BadRequest(c, err, "角色查询失败")
			return
		}
	}

	// 替换关联角色
	if err := tx.Model(&userModel).Association("Roles").Replace(roles); err != nil {
		tx.Rollback()
		return
	}

	if !helpers.Empty(request.Username) {
		userModel.Username = request.Username
	}

	if !helpers.Empty(request.Password) {
		userModel.Password = request.Password
	}

	if !helpers.Empty(request.Name) {
		userModel.Name = request.Name
	}

	if request.AvatarId > 0 {
		userModel.AvatarId = helpers.Uint64Ptr(cast.ToUint64(request.AvatarId))
	}

	//fmt.Printf("%T", userModel)

	if err := tx.Save(&userModel).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "创建账号失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminUserService) Save(model *adminUser.AdminUser) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *AdminUserService) Delete(model *adminUser.AdminUser) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminUserService) Get(idstr string) (model adminUser.AdminUser) {
	service.app.DB.Where("id", idstr).Preload("Roles").Preload("AvatarFile").First(&model)
	return
}

// Paginate 分页内容
func (service *AdminUserService) Paginate(c *gin.Context, perPage int) (users []adminUser.AdminUser, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.app.DB.Model(adminUser.AdminUser{}),
		&users,
		global.Config.VADMINURL(model.TableName(&adminUser.AdminUser{})),
		perPage,
	)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func (service *AdminUserService) GetByMulti(loginID string) (model adminUser.AdminUser) {
	service.app.DB.
		Where("username = ?", loginID).
		First(&model)
	return
}
