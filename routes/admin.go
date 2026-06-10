// Package routes 注册路由
package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/controllers"
	"github.com/outsstill/gin-admin/core"
	middlewares "github.com/outsstill/gin-admin/middlerwares"
	"github.com/outsstill/gin-admin/setting"
)

// RegisterAdminRoutes 注册 admin 相关路由

func RegisterAdminRoutes(admin *gin.RouterGroup, app *core.App) {
	//var admin *gin.RouterGroup
	//
	//admin = r.Group("/admin")

	base := &controllers.BaseAPIController{
		App: app,
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	//admin.Use(middlewares.LimitIP("500-H"))
	//admin.Use(middlewares.AuthAdminJWT())
	{
		ic := controllers.NewAdminIndexController(base)
		cc := controllers.NewAdminConfigController(base)
		admin.GET("/index", ic.Index)
		admin.GET("/version", ic.Version)
		admin.GET("/setting-all", cc.AllShow)

		authGroup := admin.Group("/auth")
		// 登录
		lgc := controllers.NewAdminAuthController(base)
		authGroup.POST("/login", middlewares.GuestJWT(), lgc.Login)
		authGroup.POST("/refresh-token", lgc.RefreshToken)
		authGroup.GET("/current", lgc.Current)
		authGroup.POST("/profile", lgc.UpdateProfile)
		authGroup.POST("/profile-pass", lgc.UpdatePassword)
		authGroup.POST("/logout", lgc.Logout)
		authGroup.GET("/captcha", lgc.ShowCaptcha)

		auc := controllers.NewAdminUserController(base)
		// 账号
		admin.GET("/users", auc.Index)
		admin.GET("/user/:id", auc.Get)
		admin.POST("/user", auc.Store)
		admin.PUT("/user/:id", auc.Update)
		admin.DELETE("/user/:id", auc.Delete)

		// 角色
		arc := controllers.NewAdminRoleController(base)
		admin.GET("/roles", arc.Index)
		admin.GET("/roles/all", arc.All)
		admin.GET("/role/:id", arc.Get)
		admin.POST("/role", arc.Store)
		admin.PUT("/role/:id/menus", arc.UpdateMenus)
		admin.PUT("/role/:id/permissions", arc.UpdatePermissions)
		admin.PUT("/role/:id", arc.Update)
		admin.DELETE("/role/:id", arc.Delete)

		// 菜单
		amc := controllers.NewAdminMenuController(base)
		admin.GET("/menus", amc.Index)
		admin.GET("/menus/all", amc.All)
		admin.GET("/menu/:id", amc.Get)
		admin.POST("/menu", amc.Store)
		admin.PUT("/menu/:id", amc.Update)
		admin.DELETE("/menu/:id", amc.Delete)

		// 权限
		apc := controllers.NewAdminPermissionController(base)
		admin.GET("/permissions", apc.Index)
		admin.GET("/permissions/all", apc.All)
		admin.GET("/permission/:id", apc.Get)
		admin.POST("/permission", apc.Store)
		admin.PUT("/permission/:id", apc.Update)
		admin.DELETE("/permission/:id", apc.Delete)

		// 配置
		admin.GET("/configs", cc.Index)
		admin.GET("/configs/all", cc.All)
		admin.GET("/config/:id", cc.Get)
		admin.POST("/config", cc.Store)
		admin.PUT("/config/:id", cc.Update)
		admin.DELETE("/config/:id", cc.Delete)

		fc := controllers.NewAdminFileController(base)
		admin.POST("/upload", fc.Upload)
		admin.POST("/file", fc.Store)
		admin.GET("/files", fc.Index)
		admin.GET("/file/:id", fc.Get)
		admin.PUT("/file/:id", fc.Update)
		admin.DELETE("/file/:id", fc.Delete)

		olc := controllers.NewAdminLogController(base)
		admin.POST("/log", olc.Store)
		admin.GET("/logs", olc.Index)
		admin.GET("/log/:id", olc.Get)
		admin.PUT("/log/:id", olc.Update)
		admin.DELETE("/log/:id", olc.Delete)

	}

}
func RegisterStaticRoutes(r *gin.Engine) {

	staticPath := "static"

	if setting.Storage().Local.StaticPrefix != "" {
		staticPath = setting.Storage().Local.StaticPrefix
	}

	staticRoute := fmt.Sprintf("/%s", staticPath)

	if HasRoute(r, staticRoute) {
		// 本地文件
		r.Static(staticRoute, "storage/files")
	}

}

func Setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
func HasRoute(r *gin.Engine, path string, ops ...string) bool {
	for _, route := range r.Routes() {

		if len(ops) > 0 && route.Method == ops[0] && route.Path == path {
			return true
		} else {
			if route.Path == path {
				return true
			}
		}

	}
	return false
}
