package router

import (
	"go-fast-frame/app/api"
	"go-fast-frame/middleware"

	"github.com/gin-gonic/gin"
)

var (
	sysUserApi = &api.SysUserApi{}
	sysAuthApi = &api.SysAuthApi{}
)

// 路由规范：业务/模板/操作，比如：taobao/user/add

// 初始化路由
func InitRoutes() *gin.Engine {

	routers := gin.Default()
	// 允许修复当前请求路径，如/FOO和/..//Foo会被修复为/foo，并进行重定向，默认为 false。
	routers.RedirectFixedPath = true

	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	routers.Use(middleware.Recover)

	// 跨域
	routers.Use(middleware.Cors())

	// 初始化权限路由
	privateGroup := routers.Group("/")
	privateGroup.Use(middleware.JwtAuth())
	{
		InitPrivateRouter(privateGroup)
	}

	// 初始化后台路由
	publicGroup := routers.Group("/")
	{
		InitPublicRouter(publicGroup)
	}

	return routers
}

// 初始化权限路由
func InitPrivateRouter(routerGroup *gin.RouterGroup) {
	// 用户
	userRouter := routerGroup.Group("/system/user")
	{
		userRouter.GET("query", sysUserApi.Query)    // 用户分页查询
		userRouter.POST("add", sysUserApi.Add)       // 添加用户
		userRouter.POST("update", sysUserApi.Update) // 更新用户
		userRouter.POST("delete", sysUserApi.Delete) // 删除用户
		userRouter.GET("detail", sysUserApi.Detail)  // 用户详情
		userRouter.GET("list", sysUserApi.List)      // 用户列表
	}

}

// 初始化无权限路由
func InitPublicRouter(routerGroup *gin.RouterGroup) {
	// 用户
	userRouter := routerGroup.Group("/system/auth")
	{
		userRouter.POST("login", sysAuthApi.Login) // 登录
	}

}
