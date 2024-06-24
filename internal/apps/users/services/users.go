package services

import (
	"fmt"

	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/internal/apps/users/forms"
	"example.com/first_gin_attempt/internal/apps/users/models"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CreateUsersService struct {
	*forms.CreateUsersForm
	C *gin.Context
}

func (s *CreateUsersService) Run() error {
	// 加密密码
	hashedPassword := models.BcryptPW(s.Password)
	user := models.User{
		Username:    s.Username,
		Password:    hashedPassword,
		Email:       s.Email,
		NickName:    s.NickName,
		Mobile:      s.Mobile,
		RoleID:      s.RoleID,
		IsTwoFA:     s.IsTwoFA,
		IsSuperuser: s.IsSuperuser,
		IsActive:    s.IsActive,
	}
	return global.App.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Create(&user).Error; err != nil {
			mysqlErr := err.(*mysql.MySQLError)
			switch mysqlErr.Number {
			case 1062:
				return fmt.Errorf("用户`%s`已存在", s.Username)
			}
			global.App.Log.Error(err)
			return err
		}
		return nil
	})
}

type UpdateUsersService struct {
	*forms.UpdateUsersForm
	C   *gin.Context
	UID int
}

func (s *UpdateUsersService) Run() error {
	return global.App.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("uid = ?", s.UID).Updates(map[string]interface{}{
			"username":     s.Username,
			"email":        s.Email,
			"nick_name":    s.NickName,
			"mobile":       s.Mobile,
			"role_id":      s.RoleID,
			"is_two_fa":    s.IsTwoFA,
			"is_superuser": s.IsSuperuser,
			"is_active":    s.IsActive,
		}).Error; err != nil {
			mysqlErr := err.(*mysql.MySQLError)
			switch mysqlErr.Number {
			case 1062:
				return fmt.Errorf("用户`%s`已存在", s.Username)
			}
			global.App.Log.Error(err)
			return err
		}
		return nil
	})
}

type DeleteUsersService struct {
	C   *gin.Context
	UID int
}

func (s *DeleteUsersService) Run() error {
	tx := global.App.DB.Where("uid = ?", s.UID).Delete(&models.User{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
