package core

import (
	"github.com/outsstill/gin-admin/pkg/cache"
	redisClient "github.com/outsstill/gin-admin/pkg/redis"
	"github.com/outsstill/gin-admin/setting"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Prefix string
	DB     *gorm.DB
	Redis  *redisClient.RedisClient
	Cache  *cache.CacheService
}
