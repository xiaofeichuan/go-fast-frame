package tools

import (
	"fmt"
	"go-fast-frame/global"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitDB 初始化数据库
func InitDB() *gorm.DB {

	var dialector gorm.Dialector
	switch global.AppConfig.Database.Type {
	case "mysql":
		dialector = mysql.New(mysql.Config{
			DSN:               global.AppConfig.Database.Default, // DSN data source name
			DefaultStringSize: 256,                               // string 类型字段的默认长度
		})
		break
	case "postgres":
		dialector = postgres.New(postgres.Config{
			DSN: "host=localhost user=postgres password=123456 dbname=go_fast_frame port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		})
		break
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("连接数据库出现错误：" + err.Error())
	}
	return db
}
