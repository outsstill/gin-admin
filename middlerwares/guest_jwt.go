package middlewares

import (
	"github.com/gin-gonic/gin"
	gokit "github.com/outsstill/go-kit"
	"github.com/outsstill/go-kit/response"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功了
			_, err := gokit.JWT().ParserTokenGin(c)
			if err == nil {
				response.Fail(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
