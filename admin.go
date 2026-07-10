package admin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	middlewares "github.com/outsstill/gin-admin/middlerwares"
	"github.com/outsstill/gin-admin/pkg/captcha"
	"github.com/outsstill/gin-admin/routes"
	service "github.com/outsstill/gin-admin/services"
	gokit "github.com/outsstill/go-kit"
)

func New(prefix string) (*core.App, error) {

	if len(prefix) == 0 {
		prefix = "/admin"
	}

	if !strings.HasPrefix(prefix, "/") {
		prefix = fmt.Sprintf("/%s", prefix)
	}

	app := &core.App{
		Prefix: prefix,
	}

	// 注册 services
	app.Register("admin_log", service.NewAdminLogService(gokit.DB().DB()))
	app.Register("admin_user", service.NewAdminUserService(gokit.DB().DB()))
	app.Register("admin_role", service.NewAdminRoleService(gokit.DB().DB()))
	app.Register("admin_menu", service.NewAdminMenuService(gokit.DB().DB()))
	app.Register("admin_permission", service.NewAdminPermissionService(gokit.DB().DB()))
	app.Register("file", service.NewFileService(gokit.DB().DB()))
	app.Register("config", service.NewConfigService(gokit.DB().DB()))
	app.Register("auth", service.NewAuthService(gokit.DB().DB()))
	app.Register("captcha", captcha.NewCaptcha(gokit.Redis()))

	return app, nil
}

var builtinModules = []core.Module{}

// 注册内置模块
func RegisterBuiltin(mods ...core.Module) {
	builtinModules = append(builtinModules, mods...)
}

func Register(r *gin.Engine, app *core.App, modules ...core.Module) {

	// 注册全局中间件
	registerGlobalMiddleWare(r)

	root := r.Group(app.Prefix)

	// 注册中间件
	//root.Use(middlewares.LimitIP("500-H"))
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	root.Use(middlewares.OperationLog(app))
	root.Use(middlewares.LimitIP(app, gokit.Config().Limit.Rate))
	root.Use(middlewares.AuthAdminJWT(app))

	// 1️⃣ 内置模块
	//for _, m := range builtinModules {
	//	registerModule(root, m)
	//}

	// 注册静态资源路由
	routes.RegisterStaticRoutes(r)

	routes.RegisterAdminRoutes(root, app)

	// 2️⃣ 业务模块
	for _, m := range modules {
		registerModule(app, root, m)
	}

	//  配置 404 路由
	routes.Setup404Handler(r)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		////gin.Logger(),
		//middlewares.Logger(),
		//middlewares.Recovery2(),
		////cors.Default(),
		////gin.Recovery(),
		////middlewares.ForceUA(),
		//cors.New(cors.Config{
		//	AllowAllOrigins: true,
		//	//AllowOrigins:     []string{"http://localhost:4000"}, // 改成你的前端地址
		//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "access-token", "x-access-token"},
		//	ExposeHeaders:    []string{"Content-Length", "Authorization"},
		//	AllowCredentials: false,
		//	MaxAge:           12 * time.Hour,
		//}),
	)
}

// 核心逻辑
func registerModule(app *core.App, root *gin.RouterGroup, m core.Module) {
	prefix := m.Prefix()

	// 自动补 /
	if prefix != "" && prefix[0] != '/' {
		prefix = "/" + prefix
	}

	group := root
	if prefix != "" {
		group = root.Group(prefix)
	}

	m.Register(app, group)
}
