package views

import (
	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/internal/apps/users/forms"
	usersModels "example.com/first_gin_attempt/internal/apps/users/models"
	"example.com/first_gin_attempt/internal/apps/users/services"
	"example.com/first_gin_attempt/internal/pkg/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/* crud */

// CreateUsers 创建用户
func CreateUsersView(c *gin.Context) {
	var form *forms.CreateUsersForm = &forms.CreateUsersForm{}
	if err := c.ShouldBind(&form); err == nil {
		service := services.CreateUsersService{
			CreateUsersForm: form,
			C:               c,
		}
		err := service.Run()
		if err != nil {
			response.Fail(c, err.Error())
		} else {
			response.Success(c, nil, "success")
		}
	} else {
		response.ValidateFail(c, err.Error())
	}
}

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {
	var users []usersModels.User
	if err := global.App.DB.Find(&users).Error; err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, users, "success")
	}
}

// GetUserByName 获取用户
func GetUserByName(c *gin.Context) {
	name := c.Params.ByName("name") // rest路由传参
	if name == "" {
		response.ValidateFail(c, "name不能为空")
		return
	}
	var user usersModels.User
	if err := global.App.DB.Where("username = ?", name).First(&user).Error; err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user, "success")
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		response.ValidateFail(c, "id不能为空")
		return
	}
	var user usersModels.User
	if err := global.App.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := global.App.DB.Delete(&user).Error; err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user, "success")
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		response.ValidateFail(c, "id不能为空")
		return
	}
	var user usersModels.User
	if err := global.App.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		response.Fail(c, err.Error())
		return
	}
	// 从请求的 JSON 数据中绑定新的用户名
	var newUser usersModels.User
	if err := c.BindJSON(&newUser); err != nil {
		response.ValidateFail(c, err.Error())
		return
	}
	// 修改用户名
	user.Username = newUser.Username
	// TODO：修改其他字段

	// 保存修改后的用户对象到数据库
	if err := global.App.DB.Save(&user).Error; err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user, "success")
}
