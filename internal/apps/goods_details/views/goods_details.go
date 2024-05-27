package views

import (
	"fmt"

	"example.com/first_gin_attempt/global"
	goodsDetailsModels "example.com/first_gin_attempt/internal/apps/goods_details/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// CreateGoodsDetail 创建
func CreateGoodsDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var goodsDetail goodsDetailsModels.GoodsDetail
		c.BindJSON(&goodsDetail)

		global.App.DB.Create(&goodsDetail)
		c.JSON(200, goodsDetail)
	}
}

// GetGoodsDetails 获取所有
func GetGoodsDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var goodsDetails []goodsDetailsModels.GoodsDetail
		if err := global.App.DB.Find(&goodsDetails).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, goodsDetails)
		}
	}
}

// DeleteGoodsDetail 删除
func DeleteGoodsDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var goodsDetail goodsDetailsModels.GoodsDetail
		d := global.App.DB.Where("goods_id = ?", id).Delete(&goodsDetail)
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}

// UpdategoodsDetail 更新
func UpdateGoodsDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var goodsDetail goodsDetailsModels.GoodsDetail
		id := c.Params.ByName("id")

		if err := global.App.DB.Where("goods_id = ?", id).First(&goodsDetail).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		// 从请求的 JSON 数据中绑定新的用户名
		var newgoodsDetail goodsDetailsModels.GoodsDetail
		if err := c.BindJSON(&newgoodsDetail); err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
			return
		}

		// 修改用户名
		goodsDetail.GoodsName = newgoodsDetail.GoodsName

		// 保存修改后的用户对象到数据库
		if err := global.App.DB.Save(&goodsDetail).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Println(err)
			return
		}
		c.JSON(200, goodsDetail)
	}
}
