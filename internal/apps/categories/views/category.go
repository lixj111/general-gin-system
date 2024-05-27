package views

import (
	"fmt"

	"example.com/first_gin_attempt/global"
	categoriesModels "example.com/first_gin_attempt/internal/apps/categories/models"
	goodsDetailsModels "example.com/first_gin_attempt/internal/apps/goods_details/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// CreateCategory 创建用户
func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category categoriesModels.Category
		c.BindJSON(&category)

		global.App.DB.Create(&category)
		c.JSON(200, category)
	}
}

// GetCategories 获取所有
func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []categoriesModels.Category
		if err := global.App.DB.Find(&categories).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, categories)
		}
	}
}

// DeleteCategory 删除
func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var category categoriesModels.Category
		d := global.App.DB.Where("cat_id = ?", id).Delete(&category)
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}

// UpdateCategory 更新
func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category categoriesModels.Category
		id := c.Params.ByName("id")

		if err := global.App.DB.Where("cat_id = ?", id).First(&category).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		// 从请求的 JSON 数据中绑定新的用户名
		var newCategory categoriesModels.Category
		if err := c.BindJSON(&newCategory); err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
			return
		}

		// 修改用户名
		category.Name = newCategory.Name

		// 保存修改后的用户对象到数据库
		if err := global.App.DB.Save(&category).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Println(err)
			return
		}
		c.JSON(200, category)
	}
}

func GetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var category categoriesModels.Category
		if err := global.App.DB.Where("cat_id = ?", id).First(&category).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}

		// 查询与当前 Category 相关联的 GoodsDetail 记录
		var goodsDetails []goodsDetailsModels.GoodsDetail
		if err := global.App.DB.Where("goods_cat = ?", category.Name).Find(&goodsDetails).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Println(err)
			return
		}

		type CategoryWithGoodsDetails struct {
			Category     categoriesModels.Category        `json:"category"`
			GoodsDetails []goodsDetailsModels.GoodsDetail `json:"goods_details"`
		}

		// 创建一个新的结构体，用于存储 Category 和其关联的 GoodsDetail 记录
		categoryWithGoodsDetails := CategoryWithGoodsDetails{
			Category:     category,
			GoodsDetails: goodsDetails,
		}

		c.JSON(200, categoryWithGoodsDetails)
	}
}
