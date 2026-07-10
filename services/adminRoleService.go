package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/model/adminMenu"
	"github.com/outsstill/gin-admin/model/adminPermission"
	"github.com/outsstill/gin-admin/model/adminRole"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/paginator"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/go-kit/response"
	"gorm.io/gorm"
)

type AdminRoleService struct {
	DB *gorm.DB
}

func NewAdminRoleService(db *gorm.DB) *AdminRoleService {
	return &AdminRoleService{
		DB: db,
	}
}

func (service *AdminRoleService) Create(c *gin.Context, request *requests.AdminRoleStoreRequest) {

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		}
	}()

	// 查询菜单和权限
	var permissions []adminPermission.AdminPermission
	var menus []adminMenu.AdminMenu

	if len(request.PermissionIDs) > 0 {
		if err := tx.Where("id IN ?", request.PermissionIDs).Find(&permissions).Error; err != nil {
			response.BadRequest(c, err, "权限查询失败")
			return
		}
	}

	if len(request.MenuIDs) > 0 {
		if err := tx.Where("id IN ?", request.MenuIDs).Find(&menus).Error; err != nil {
			response.BadRequest(c, err, "菜单查询失败")
			return
		}
	}

	role := &adminRole.AdminRole{
		Name:        request.Name,
		Slug:        request.Slug,
		Permissions: permissions,
		Menus:       menus,
	}

	if err := tx.Create(role).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "创建角色失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminRoleService) Update(c *gin.Context, request *requests.AdminRoleUpdateRequest, userModel *adminRole.AdminRole) {

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		}
	}()

	// 查询菜单和权限
	//var permissions []adminPermission.AdminPermission
	//var menus []adminMenu.AdminMenu
	//
	//if len(request.PermissionIDs) > 0 {
	//	if err := tx.Where("id IN ?", request.PermissionIDs).Find(&permissions).Error; err != nil {
	//		response.BadRequest(c, err, "权限查询失败")
	//		return
	//	}
	//}
	//
	//if len(request.MenuIDs) > 0 {
	//	if err := tx.Where("id IN ?", request.MenuIDs).Find(&menus).Error; err != nil {
	//		response.BadRequest(c, err, "菜单查询失败")
	//		return
	//	}
	//}
	//
	//// 替换关联权限和菜单
	//if err := tx.Model(&userModel).Association("Permissions").Replace(permissions); err != nil {
	//	tx.Rollback()
	//	return
	//}
	//if err := tx.Model(&userModel).Association("Menus").Replace(menus); err != nil {
	//	tx.Rollback()
	//	return
	//}

	//role := adminRole.AdminRole{
	//	Name:        request.Name,
	//	Slug:        request.Slug,
	//	Permissions: permissions,
	//	Menus:       menus,
	//}

	//fmt.Printf("%v \n", request)

	userModel.Name = request.Name
	userModel.Slug = request.Slug
	//userModel.Permissions = permissions
	//userModel.Menus = menus

	//fmt.Printf("%T", userModel)

	if err := tx.Save(&userModel).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "更新角色信息失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminRoleService) UpdateMenus(c *gin.Context, request *requests.AdminRoleUpdateMenusRequest, userModel *adminRole.AdminRole) {

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		}
	}()

	// 查询菜单和权限
	var menus []adminMenu.AdminMenu

	if len(request.MenuIDs) > 0 {
		if err := tx.Where("id IN ?", request.MenuIDs).Find(&menus).Error; err != nil {
			response.BadRequest(c, err, "菜单查询失败")
			return
		}
	}

	// 替换关联权限和菜单
	if err := tx.Model(&userModel).Association("Menus").Replace(menus); err != nil {
		tx.Rollback()
		return
	}
	userModel.Menus = menus

	if err := tx.Save(&userModel).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "更新角色菜单失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminRoleService) UpdatePermissions(c *gin.Context, request *requests.AdminRoleUpdatePermissionsRequest, userModel *adminRole.AdminRole) {

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		}
	}()

	// 查询菜单和权限
	var permissions []adminPermission.AdminPermission

	if len(request.PermissionIDs) > 0 {
		if err := tx.Where("id IN ?", request.PermissionIDs).Find(&permissions).Error; err != nil {
			response.BadRequest(c, err, "权限查询失败")
			return
		}
	}

	// 替换关联权限和菜单
	if err := tx.Model(&userModel).Association("Permissions").Replace(permissions); err != nil {
		tx.Rollback()
		return
	}

	userModel.Permissions = permissions

	//fmt.Printf("%T", userModel)

	if err := tx.Save(&userModel).Error; err != nil {
		tx.Rollback()
		response.BadRequest(c, err, "更新角色权限失败")
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.BadRequest(c, err, "提交事务失败")
		return
	}
}

func (service *AdminRoleService) Delete(model *adminRole.AdminRole) (rowsAffected int64) {
	result := service.DB.Delete(model)
	return result.RowsAffected
}

func (service *AdminRoleService) Get(idstr string) (model *adminRole.AdminRole) {
	service.DB.Where("id", idstr).Preload("Menus").Preload("Permissions").First(&model)
	return
}

func (service *AdminRoleService) All() (models []adminRole.AdminRole) {
	service.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *AdminRoleService) Paginate(c *gin.Context, perPage int) (users []adminRole.AdminRole, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		service.DB.Model(adminRole.AdminRole{}),
		&users,
		helpers.VADMINURL(model.TableName(&adminRole.AdminRole{})),
		perPage,
	)
	return
}
