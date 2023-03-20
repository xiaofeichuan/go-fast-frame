package main

import (
	"fmt"
	"go-fast-frame/common/tools"
	"go-fast-frame/config"
	docs "go-fast-frame/docs"
	"go-fast-frame/global"
	"go-fast-frame/router"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// 初始化配置
	global.AppConfig = config.LoadAppConfig()

	// 初始化数据库
	global.DB = tools.InitDB()

	// 初始化路由
	var routers = router.InitRoutes()

	// 初始化swagger文档
	InitSwagger(routers)

	//自定义 HTTP 服务器
	server := &http.Server{
		Addr:    ":9999",
		Handler: routers,
	}
	fmt.Printf("Swagger文档地址:http://127.0.0.1%s/swagger/index.html", server.Addr)
	server.ListenAndServe()
}

// 初始化swagger文档
func InitSwagger(routers *gin.Engine) {
	docs.SwaggerInfo.Title = "Go Fast Frame"
	docs.SwaggerInfo.Description = "Swagger Admin API"
	docs.SwaggerInfo.Version = "1.0"
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
