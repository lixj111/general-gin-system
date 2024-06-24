package router

import (
	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/internal/apps/users/views"
	"example.com/first_gin_attempt/middleware"
	"github.com/gin-gonic/gin"
)

// 管理员相关路由表
func AdminRoutes(userGroup *gin.RouterGroup) {
	admin := userGroup.Group("/admin")
	admin.Use(middleware.HasAdminPermission())

	// 用户
	admin.GET("/users", views.GetUsersView)
	admin.POST("/user", views.CreateUserView)
	admin.GET("/:name", views.GetUserByName)
	admin.PUT("/users/:uid", views.UpdateUsersView)
	admin.DELETE("/users/:uid", views.DeleteUserView)
	// admin.POST("/users/change/password", views.ChangeUsersPasswordView)

	// // 角色
	// admin.GET("/roles", views.GetRolesView)
	// admin.POST("/roles", views.CreateRolesView)
	// admin.PUT("/roles/:id", views.UpdateRolesView)
	// admin.DELETE("/roles/:id", views.DeleteRolesView)
	// // 组织
	// admin.GET("/organizations", views.GetOrganizationsView)
	// admin.PUT("/organizations", views.UpdateOrganizationsView)
	// admin.DELETE("/organizations", views.DeleteOrganizationsView)
	// admin.POST("/organizations/root-node", views.CreateRootOrganizationsView)
	// admin.POST("/organizations/child-node", views.CreateChildOrganizationsView)
	// admin.GET("/organizations/users", views.GetOrganizationsUsersView)
	// admin.POST("/organizations/users", views.BindOrganizationsUsersView)
	// admin.DELETE("/organizations/users", views.DeleteOrganizationsUsersView)
}

// SetupUserRoutes 设置用户相关路由
func Routers(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/sso/callback", views.SSOCallback, global.App.JWT.LoginHandler)
		userGroup.POST("/login/password", middleware.OTPMiddleware(), global.App.JWT.LoginHandler)
		userGroup.POST("/login/logout", global.App.JWT.LogoutHandler)        // 登出，但token依然有效
		userGroup.GET("/login/refresh_token", global.App.JWT.RefreshHandler) // 更新token，但是时间上未过期的token依然有效
	}
	userGroup.Use(global.App.JWT.MiddlewareFunc())
	{
		// 个人用户
		userGroup.GET("/user", views.GetUserInfoView)
		userGroup.PUT("/user/:uid", views.UpdateUserInfoView)
		userGroup.POST("/user/change/avatar", views.ChangeUserAvatarView)
		userGroup.POST("/user/change/password", views.ChangeUserPasswordView)

		// 管理员
		AdminRoutes(userGroup)
	}
}
