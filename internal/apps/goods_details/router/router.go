package router

import (
	"example.com/first_gin_attempt/internal/apps/goods_details/views"
	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes 设置订单相关路由
func Routers(router *gin.Engine) {
	goodsDetailGroup := router.Group("/goods_detail")
	{
		goodsDetailGroup.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Get user by ID",
			})
		})
		goodsDetailGroup.POST("/", views.CreateGoodsDetail())
		goodsDetailGroup.GET("/", views.GetGoodsDetails())
		goodsDetailGroup.DELETE("/:id", views.DeleteGoodsDetail())
		goodsDetailGroup.PUT("/:id", views.UpdateGoodsDetail())
	}
}
