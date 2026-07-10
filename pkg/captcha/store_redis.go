package captcha

import (
	"errors"
	"time"

	gokit "github.com/outsstill/go-kit"
	"github.com/outsstill/go-kit/redis"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {

	ExpireTime := time.Minute * time.Duration(gokit.Config().Captcha.Expiration)

	// 方便本地开发调试
	if gokit.Config().App.Debug {
		ExpireTime = time.Minute * time.Duration(gokit.Config().Captcha.DebugExpireTime)
	}

	if err := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); err != nil {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val, _ := s.RedisClient.Get(key)
	if clear {
		_ = s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
