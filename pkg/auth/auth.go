// Package auth 授权相关逻辑
package auth

import (
	"errors"

	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model/adminUser"
	service "github.com/outsstill/gin-admin/services"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	app *core.App
}

func NewAuth(app *core.App) *Auth {
	return &Auth{
		app: app,
	}
}

// AttemptAdmin 尝试登录
func (a *Auth) AttemptAdmin(email string, password string) (adminUser.AdminUser, error) {
	userService := service.NewAdminUserService(a.app)
	userModel := userService.GetByMulti(email)
	if userModel.ID == 0 {
		return adminUser.AdminUser{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return adminUser.AdminUser{}, errors.New("密码错误")
	}

	return userModel, nil
}

// CurrentAdminUser 从 gin.context 中获取当前登录用户
func (a *Auth) CurrentAdminUser(c *gin.Context) *adminUser.AdminUser {
	userService := service.NewAdminUserService(a.app)
	model := userService.Get(cast.ToString(c.MustGet("current_admin_user_id")))

	if model.ID <= 0 {
		global.Logger.LogIf(errors.New("无法获取用户"))
		//response.Fail(c, "没有找到")
		return &adminUser.AdminUser{}
	}

	if c.GetInt("menus_on") > 0 || cast.ToInt(c.Query("menus_on")) > 0 {
		// 获取账号的显示菜单
		menus, errs := userService.GetUserMenus(model.ID)

		if errs != nil {
			global.Logger.LogIf(errs)

			return model
		}

		model.Menus = menus
		model.ChildrenMenus = adminUser.BuildVbenRoutes(menus) // 一些前端框架需要带 children 的侧边路由
	}
	// db is now a *DB value
	return model
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID

func CurrentAdminUID(c *gin.Context) string {
	return c.GetString("current_admin_user_id")
}
