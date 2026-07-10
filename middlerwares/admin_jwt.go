package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/pkg/helpers"
	gokit "github.com/outsstill/go-kit"
	"github.com/outsstill/go-kit/jwt"
	"github.com/outsstill/go-kit/response"
)

func AuthAdminJWT(app *core.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		/**
		c.FullPath()   // "/user/:id"   路由模板
		c.Request.URL.Path  // "/user/123"  实际路径
		c.Param("id")  // "123"         路径参数
		c.Query("name") // "abc"        query参数
		*/
		path := c.FullPath()

		ignorePermissionPaths := []string{"/auth/logout", "/auth/refresh-token", "/roles/all", "/permissions/all", "/menus/all", "/auth/current", "/limit-test"}

		realIgnorePermissionPaths := make([]string, 0, len(ignorePermissionPaths))

		for _, itemPath := range ignorePermissionPaths {
			if !strings.HasPrefix(itemPath, "/") {
				itemPath = "/" + itemPath
			}
			realIgnorePermissionPaths = append(realIgnorePermissionPaths, fmt.Sprintf("%s%s", app.Prefix, itemPath))
		}

		ignorePaths := helpers.GetIgnorePaths(app.Prefix)
		//
		if !helpers.StringContains(ignorePaths, path) {
			// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
			claims, err := gokit.JWT().ParserTokenGin(c)

			// JWT 解析失败，有错误发生
			if err != nil {
				response.Abort401(c, "凭证过期，请重新登陆")
				// response.AuthFail(c, fmt.Sprintf("请查看相关的接口认证文档 path：%s", path))
				return
			}

			// 判断是否是 admin token
			if claims.Type != jwt.ADMIN_TOKEN_TYPE {
				response.Abort403(c, "非法操作!!!")
				return
			}

			// JWT 解析成功，设置用户信息
			userModel := app.GetAdminUserService().Get(claims.UserID)
			if userModel.ID == 0 {
				response.AuthFail(c, "找不到对应用户，用户可能已删除")
				return
			}
			//if !global.Config.IsProduction() {
			//	fmt.Printf("GetStringID 11 :%s \n", userModel.GetStringID())
			//}
			// 验证权限
			// 3. 比对请求的 Method + 路由模板
			isPass := false

			// 有些全局的操作也不做验证权限
			if helpers.StringContains(realIgnorePermissionPaths, path) {
				isPass = true
			} else {
				// 超级管理员
				if userModel.IsSuperAdmin() {
					isPass = true
				} else {
					perms, err := app.GetAdminUserService().GetUserPermissions(userModel.ID)
					if err != nil {
						response.BadRequest(c, err, "权限加载失败")
						return
					}

					reqMethod := c.Request.Method // GET, POST, PUT ...
					for _, perm := range perms {
						if perm.HttpMethod != "any" && strings.ToUpper(perm.HttpMethod) != strings.ToUpper(reqMethod) {
							continue
						}
						if helpers.IsPathAllowed(path, perm.HttpPath) {
							// 放行
							isPass = true
							break
						}
					}
				}
			}

			if isPass {
				// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
				c.Set("current_admin_user_id", userModel.GetStringID())
				//c.Set("current_user_name", userModel.Name)
				//c.Set("current_user", userModel)

				c.Next()
				return
			}

			// 4. 拒绝访问
			response.Abort403(c, "无权访问")
			return
		}

		c.Next()
	}
}
