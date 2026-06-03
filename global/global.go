package global

import (
	"github.com/outsstill/gin-admin/pkg/logger"
	"github.com/outsstill/gin-admin/setting"
)

var Config *setting.Setting

var Logger *logger.Logger

func Init(logger *logger.Logger, config *setting.Setting) {
	Config = config
	Logger = logger
}
