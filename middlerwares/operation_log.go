// Package middlewares 存放系统中间件
package middlewares

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model/adminLog"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/spf13/cast"
)

func OperationLog(app *core.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		adminLog := &adminLog.AdminLog{}
		adminLog.UserId = cast.ToUint64(auth.CurrentAdminUID(c))

		// 设置开始时间
		c.Next()

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			if adminLog.UserId == 0 {
				adminLog.UserId = cast.ToUint64(auth.CurrentAdminUID(c))
			}

			if adminLog.UserId <= 0 {
				return
			}

			fullUrl := c.Request.URL.Path
			query := c.Request.URL.RawQuery

			if query != "" {
				fullUrl += "?" + query
			}

			adminLog.Path = c.FullPath()
			adminLog.Url = fullUrl
			adminLog.Method = c.Request.Method
			adminLog.Input = string(requestBody)
			adminLog.Ip = c.ClientIP()

			go func() {
				defer func() {
					if r := recover(); r != nil {
						data, _ := json.Marshal(adminLog)
						global.Logger.ErrorString("OperationLog", "记录失败", string(data))
					}
				}()
				app.GetAdminLogService().Create(adminLog)
			}()
		}
	}
}
