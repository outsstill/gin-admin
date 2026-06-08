package auth

import "github.com/gin-gonic/gin"

func CurrentAdminUID(c *gin.Context) string {
	return c.GetString("current_admin_user_id")
}
