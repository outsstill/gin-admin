package admin

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	middlewares "github.com/outsstill/gin-admin/middlerwares"
	"github.com/outsstill/gin-admin/pkg/cache"
	"github.com/outsstill/gin-admin/pkg/captcha"
	"github.com/outsstill/gin-admin/pkg/logger"
	redisClient "github.com/outsstill/gin-admin/pkg/redis"
	"github.com/outsstill/gin-admin/routes"
	service "github.com/outsstill/gin-admin/services"
	"github.com/outsstill/gin-admin/setting"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewApp(prefix string, opts ...Option) (*core.App, error) {

	cfg := &InitConfig{}

	// 2️⃣ 应用手动配置（覆盖）
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.Config == nil {
		return nil, errors.New("config is required")
	}

	if cfg.DB == nil {
		return nil, errors.New("db is required")
	}

	if cfg.Redis == nil {
		return nil, errors.New("redis is required")
	}

	if cfg.Logger == nil {
		return nil, errors.New("logger is required")
	}

	rClient := redisClient.NewClient(cfg.Redis)

	if len(prefix) == 0 {
		prefix = "/admin"
	}

	app := &core.App{
		Prefix: prefix,
		DB:     cfg.DB,
		Redis:  rClient,
		Cache:  cfg.Cache,
	}

	// 全局变量初始化
	globalLogger := logger.New(cfg.Logger, cfg.Config)
	global.Init(globalLogger, cfg.Config)

	return app, nil
}
func NewAppWithConfigFile(filepath string, prefix string, opts ...Option) (*core.App, error) {

	cfg, err := loadConfig(filepath)

	if err != nil {
		return nil, errors.New("配置文件错误")
	}

	if cfg.Config == nil {
		return nil, errors.New("config is nil")
	}

	// 2️⃣ 应用手动配置（覆盖）
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.DB == nil {
		return nil, errors.New("db is required")
	}

	if cfg.Redis == nil {
		return nil, errors.New("redis is required")
	}

	if cfg.Logger == nil {
		return nil, errors.New("logger is required")
	}

	rClient := redisClient.NewClient(cfg.Redis)

	if len(prefix) == 0 {
		prefix = "/admin"
	}

	if !strings.HasPrefix(prefix, "/") {
		prefix = fmt.Sprintf("/%s", prefix)
	}

	app := &core.App{
		Prefix: prefix,
		DB:     cfg.DB,
		Redis:  rClient,
		Cache:  cfg.Cache,
	}

	// 注册 services
	app.Register("admin_log", service.NewAdminLogService(app.DB))
	app.Register("admin_user", service.NewAdminUserService(app.DB))
	app.Register("admin_role", service.NewAdminRoleService(app.DB))
	app.Register("admin_menu", service.NewAdminMenuService(app.DB))
	app.Register("admin_permission", service.NewAdminPermissionService(app.DB))
	app.Register("file", service.NewFileService(app.DB))
	app.Register("config", service.NewConfigService(app.DB))
	app.Register("auth", service.NewAuthService(app.DB))
	app.Register("captcha", captcha.NewCaptcha(app.Redis))

	// 全局变量初始化
	globalLogger := logger.New(cfg.Logger, cfg.Config)
	global.Init(globalLogger, cfg.Config)

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

type InitConfig struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
	Config *setting.Setting
	Cache  *cache.CacheService
}

type Option func(config *InitConfig)

func loadConfig(path string) (*InitConfig, error) {
	v := viper.New()

	if len(path) > 0 {
		v.SetConfigType("yaml") // 类型
		v.AddConfigPath(".")    // 当前目录
		v.SetConfigFile(path)

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	cfg := InitConfig{}
	if err := v.Unmarshal(&cfg.Config); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func WithConfig(setting *setting.Setting) Option {
	return func(c *InitConfig) {
		c.Config = setting
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(c *InitConfig) {
		c.Logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(c *InitConfig) {
		c.DB = db
	}
}

func WithRedis(redis *redis.Client) Option {
	return func(c *InitConfig) {
		c.Redis = redis
	}
}

func WithCache(cache *cache.CacheService) Option {
	return func(c *InitConfig) {
		c.Cache = cache
	}
}
