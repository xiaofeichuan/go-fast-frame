package global

import (
	"go-fast-frame/config"

	"gorm.io/gorm"
)

var (
	DB        *gorm.DB          //数据库
	AppConfig *config.AppConfig //配置
)
