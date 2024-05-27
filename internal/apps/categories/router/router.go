package routers

import (
	"example.com/first_gin_attempt/internal/apps/categories/views"

	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes 设置订单相关路由
func Routers(router *gin.Engine) {
	categoryGroup := router.Group("/category")
	{
		categoryGroup.GET("/:id", views.GetCategory())
		categoryGroup.POST("/", views.CreateCategory())
		categoryGroup.GET("/", views.GetCategories())
		categoryGroup.DELETE("/:id", views.DeleteCategory())
		categoryGroup.PUT("/:id", views.UpdateCategory())
	}
}
