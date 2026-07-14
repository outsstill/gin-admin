package core

import (
	service "github.com/outsstill/gin-admin/services"
)

type ServiceInterface interface {
	Name() string
}

type ServiceRegistrar interface {
	Register(name string, svc any)
}

type App struct {
	Prefix   string
	services map[string]any
}

func (a *App) Register(name string, svc any) {
	if a.services == nil {
		a.services = make(map[string]any)
	}
	a.services[name] = svc
}

func (a *App) GetAdminLogService() *service.AdminLogService {
	return a.services["admin_log"].(*service.AdminLogService)
}

func (a *App) GetAdminUserService() *service.AdminUserService {
	return a.services["admin_user"].(*service.AdminUserService)
}

func (a *App) GetAdminRoleService() *service.AdminRoleService {
	return a.services["admin_role"].(*service.AdminRoleService)
}

func (a *App) GetAdminMenuService() *service.AdminMenuService {
	return a.services["admin_menu"].(*service.AdminMenuService)
}

func (a *App) GetAdminPermissionService() *service.AdminPermissionService {
	return a.services["admin_permission"].(*service.AdminPermissionService)
}

func (a *App) GetConfigService() *service.ConfigService {
	return a.services["config"].(*service.ConfigService)
}

func (a *App) GetFileService() *service.FileService {
	return a.services["file"].(*service.FileService)
}

func (a *App) GetAuthService() *service.AuthService {
	return a.services["auth"].(*service.AuthService)
}
