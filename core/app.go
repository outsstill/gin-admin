package core

import (
	"github.com/outsstill/gin-admin/pkg/cache"
	redisClient "github.com/outsstill/gin-admin/pkg/redis"
	"github.com/outsstill/gin-admin/setting"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	Prefix string
	DB     *gorm.DB
	Redis  *redisClient.RedisClient
	Cache  *cache.CacheService
}

func LoadConfig(path string) (*setting.Setting, error) {
	v := viper.New()

	if len(path) > 0 {
		v.SetConfigType("yaml") // 类型
		v.AddConfigPath(".")    // 当前目录
		v.SetConfigFile(path)

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	cfg := setting.Setting{}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
