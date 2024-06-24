package services

import (
	"fmt"
	"mime/multipart"
	"time"

	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/internal/apps/users/forms"
	"example.com/first_gin_attempt/internal/apps/users/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type GetUserInfoServies struct {
	C        *gin.Context
	Username string
}

func (s *GetUserInfoServies) Run() (responseData interface{}, err error) {
	var user map[string]interface{}
	// 查询数据库
	tx := global.App.DB.Table("user u").Where("u.username = ?", s.Username).Scan(&user)
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("用户 %s 不存在", s.Username)
	}
	return user, nil
}

type UpdateUserInfoServices struct {
	C *gin.Context
	*forms.UpdateUserInfoForm
	UID int
}

func (s *UpdateUserInfoServices) Run() error {
	tx := global.App.DB.Model(&models.User{}).Where("uid = ?", s.UID)
	data := make(map[string]interface{})
	if s.NickName != "" {
		data["nick_name"] = s.NickName
	}
	if s.Mobile != "" {
		data["mobile"] = s.Mobile
	}
	if s.Email != "" {
		data["email"] = s.Email
	}
	tx.Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

type ChangeUserAvatarService struct {
	C        *gin.Context
	Username string
	File     *multipart.FileHeader
}

func (s *ChangeUserAvatarService) Run() error {
	// 保存图片文件
	filename := fmt.Sprintf("%s_%d.jpg", s.Username, time.Now().Unix())
	err := s.C.SaveUploadedFile(s.File, "./media/avatar/"+filename)
	if err != nil {
		return err
	}

	// 更新用户头像路径
	tx := global.App.DB.Model(&models.User{}).Where("username = ?", s.Username).Update("avatar_file", "/media/avatar/"+filename)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

type ChangeUserPasswordService struct {
	C        *gin.Context
	Username string
	*forms.ChangeUserPasswordForm
}

func (s *ChangeUserPasswordService) Run() error {
	var user models.User
	tx := global.App.DB.Table("user u").Where("u.username = ?", s.Username).Scan(&user)
	if tx.RowsAffected == 0 {
		return fmt.Errorf("用户 %s 不存在", s.Username)
	}
	// 验证老密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.CurrentPassword))
	if err != nil {
		return fmt.Errorf("密码更改失败，旧密码输入不正确")
	}
	// 验证新老密码是否一致
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.NewPassword))
	if err == nil {
		return fmt.Errorf("密码更改失败，新密码不能与旧密码一致")
	}

	// 更新用户密码
	hashedPassword := models.BcryptPW(s.NewPassword)
	tx = global.App.DB.Model(&models.User{}).Where("username = ?", s.Username).Update("password", hashedPassword)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
