package core

import (
	"github.com/outsstill/gin-admin/pkg/cache"
	redisClient "github.com/outsstill/gin-admin/pkg/redis"
	"gorm.io/gorm"
)

type App struct {
	Prefix string
	DB     *gorm.DB
	Redis  *redisClient.RedisClient
	Cache  *cache.CacheService
}
