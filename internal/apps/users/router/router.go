package router

import (
	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/internal/apps/users/views"
	"example.com/first_gin_attempt/middleware"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes 设置用户相关路由
func Routers(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.DELETE("/:id", views.DeleteUser)
		userGroup.PUT("/:id", views.UpdateUser)
		userGroup.GET("/sso/callback", views.SSOCallback, global.App.JWT.LoginHandler) // TODO：待实现
		userGroup.POST("/login/password", middleware.OTPMiddleware(), global.App.JWT.LoginHandler)
		userGroup.POST("/login/logout", global.App.JWT.LogoutHandler)        // 登出，但token依然有效
		userGroup.GET("/login/refresh_token", global.App.JWT.RefreshHandler) // 更新token，但是时间上未过期的tokrn依然有效
	}
	userGroup.Use(global.App.JWT.MiddlewareFunc())
	{
		userGroup.GET("/:name", views.GetUserByName)
		userGroup.POST("/", views.CreateUsersView)
		userGroup.GET("/", views.GetUsers)
	}
}
