package service

import "github.com/outsstill/gin-admin/core"

type Services struct {
	AdminLog *AdminLogService
}

func NewServices(app *core.App) *Services {
	return &Services{
		AdminLog: NewAdminLogService(app),
	}
}
