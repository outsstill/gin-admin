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

	err = r.Run(":8789")
        if err != nil {
		panic(err)
	}

}
```

### module示例
```go
package topic

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/outsstill/gin-admin/setting"
)

type Module struct{}

func (m *Module) Name() string {
	return "topic"
}

func (m *Module) Prefix() string {
	return "/topic"
}

func (m *Module) Register(rg *gin.RouterGroup) {

	data, err := json.Marshal(setting.GlobalSetting)

	if err != nil {
		panic(err)
	}

	rg.GET("/index", func(c *gin.Context) {
		currentAdmin := auth.CurrentAdminUser(c)
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%v", string(data)),
			"msg":  currentAdmin,
		})
	})
}

```