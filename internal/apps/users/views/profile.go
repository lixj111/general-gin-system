package views

import (
	"strconv"

	"example.com/first_gin_attempt/internal/apps/users/forms"
	"example.com/first_gin_attempt/internal/apps/users/services"
	"example.com/first_gin_attempt/internal/pkg/response"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserInfoView(c *gin.Context) {
	username := jwt.ExtractClaims(c)["id"].(string)
	service := services.GetUserInfoServies{C: c, Username: username}
	returnData, err := service.Run()
	if err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, returnData, "success")
	}
}

func UpdateUserInfoView(c *gin.Context) {
	var form *forms.UpdateUserInfoForm = &forms.UpdateUserInfoForm{}
	uid, _ := strconv.Atoi(c.Param("uid"))
	if err := c.ShouldBind(form); err == nil {
		service := services.UpdateUserInfoServices{
			C:                  c,
			UpdateUserInfoForm: form,
			UID:                uid,
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

func ChangeUserAvatarView(c *gin.Context) {
	username := jwt.ExtractClaims(c)["id"].(string)
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	service := services.ChangeUserAvatarService{
		C:        c,
		File:     file,
		Username: username,
	}
	err = service.Run()
	if err != nil {
		response.Fail(c, err.Error())
		return
	} else {
		response.Success(c, nil, "success")
	}
}

func ChangeUserPasswordView(c *gin.Context) {
	username := jwt.ExtractClaims(c)["id"].(string)
	var form *forms.ChangeUserPasswordForm = &forms.ChangeUserPasswordForm{}
	if err := c.ShouldBind(form); err == nil {
		service := services.ChangeUserPasswordService{C: c, Username: username, ChangeUserPasswordForm: form}
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
