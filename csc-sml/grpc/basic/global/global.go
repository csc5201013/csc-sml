package global

import (
	"gorm.io/gorm"
	"grpc/basic/config"
)

var AppConf config.AppConfig
var DB *gorm.DB
