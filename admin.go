package admin

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	middlewares "github.com/outsstill/gin-admin/middlerwares"
	"github.com/outsstill/gin-admin/pkg/cache"
	"github.com/outsstill/gin-admin/pkg/logger"
	redisClient "github.com/outsstill/gin-admin/pkg/redis"
	"github.com/outsstill/gin-admin/routes"
	"github.com/outsstill/gin-admin/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewApp(prefix string, config *setting.Setting, log *zap.Logger, db *gorm.DB, redis *redis.Client, cache *cache.CacheService) *core.App {

	if db == nil {
		panic("logger is required")
	}

	if redis == nil {
		panic("db is required")
	}

	if log == nil {
		panic("logger is required")
	}

	if config == nil {
		panic("config is required")
	}

	rClient := redisClient.NewClient(redis)

	if len(prefix) == 0 {
		prefix = "/admin"
	}

	app := &core.App{
		Prefix: prefix,
		DB:     db,
		Redis:  rClient,
		Cache:  cache,
	}

	// 全局变量初始化
	globalLogger := logger.New(log)
	global.Init(globalLogger, config)

	return app
}
func NewAppWithConfigFile(filepath string, prefix string, log *zap.Logger, db *gorm.DB, redis *redis.Client, cache *cache.CacheService) *core.App {

	config, err := core.LoadConfig(filepath)

	if err != nil {
		panic("配置文件错误" + err.Error())
	}

	if config == nil {
		panic("config is required")
	}

	if db == nil {
		panic("logger is required")
	}

	if redis == nil {
		panic("db is required")
	}

	if log == nil {
		panic("logger is required")
	}

	rClient := redisClient.NewClient(redis)

	if len(prefix) == 0 {
		prefix = "/admin"
	}

	app := &core.App{
		Prefix: prefix,
		DB:     db,
		Redis:  rClient,
		Cache:  cache,
	}

	// 全局变量初始化
	globalLogger := logger.New(log)
	global.Init(globalLogger, config)

	return app
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
	root.Use(middlewares.AuthAdminJWT(app))
	root.Use(middlewares.OperationLog(app))

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
		//gin.Logger(),
		middlewares.Logger(),
		middlewares.Recovery2(),
		//cors.Default(),
		//gin.Recovery(),
		//middlewares.ForceUA(),
		cors.New(cors.Config{
			AllowAllOrigins: true,
			//AllowOrigins:     []string{"http://localhost:4000"}, // 改成你的前端地址
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}),
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
