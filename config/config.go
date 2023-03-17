package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Jwt      JwtConfig      `yaml:"jwt"`
	Database DatabaseConfig `yaml:"database"`
}

type JwtConfig struct {
	Issuer    string `yaml:"issuer"`
	Signature string `yaml:"signature"`
}

type DatabaseConfig struct {
	Type    string `yaml:"type"`
	Default string `yaml:"default"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

// LoadAppConfig 加载配置
func LoadAppConfig() *AppConfig {
	//导入配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile("appconfig.yaml")

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置出现错误：" + err.Error())
	}

	//将配置文件读到结构体中
	var _config *AppConfig
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println("配置文件转实体出现错误" + err.Error())
	}
	return _config
}
