package views

import (
	"strconv"

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
func CreateUserView(c *gin.Context) {
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
func GetUsersView(c *gin.Context) {
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

// DeleteUserView 删除用户
func DeleteUserView(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	service := services.DeleteUsersService{
		C:   c,
		UID: uid,
	}
	err := service.Run()
	if err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil, "success")
	}
}

// UpdateUser 更新用户
func UpdateUsersView(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	var form *forms.UpdateUsersForm = &forms.UpdateUsersForm{}
	if err := c.ShouldBind(&form); err == nil {
		service := services.UpdateUsersService{
			C:               c,
			UpdateUsersForm: form,
			UID:             uid,
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
