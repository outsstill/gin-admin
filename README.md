## 使用

### 初始化
```go
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db := getDB()
	redis := getRedis()
    logger := getLogger()
	cache := getCache()

	// 通过配置文件传入配置,并传入必要的实例
	app, err := admin.NewAppWithConfigFile("config.yaml", "/admin", admin.WithDB(database.DB), admin.WithRedis(redis.Redis.Client), admin.WithLogger(logger.Logger))

    if err != nil {
        panic(err)
    }
	// 初始化路由，注册外部module
	admin.Register(r, app, &topic.Module{})

	err = r.Run(":" + config.GetString("app.http_port"))
        if err != nil {
		panic(err)
	}

}
```

### module示例
```go
type Module struct {
}

func (m *Module) Name() string {
	return "product"
}

func (m *Module) Prefix() string {
	return "/product"
}

func (m *Module) Register(app *core.App, rg *gin.RouterGroup) {

	// /admin/product/list
	rg.GET("/list", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code":    200,
			"data":    []string{"1", "2"},
			"message": "success!!!",
		})
	})

	// /admin/product/all-setting
	rg.GET("/all-setting", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code":    200,
			"data":    configModel.AllShow(),
			"message": "success!!!",
		})
	})

	// 不走验证 ; 需要带上 当前module的 prefix , 初始化的前缀不需要
	helpers.AppendIgnorePaths([]string{"/product/all-setting"})
}
```