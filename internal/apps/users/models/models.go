package models

import (
	"example.com/first_gin_attempt/internal/apps/common/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uid         uint64           `gorm:"type:bigint;primaryKey;autoIncrement;comment:用户ID" json:"uid"`
	Username    string           `gorm:"type:varchar(32);not null;uniqueIndex:uniq_username;comment:用户名" json:"username"`
	Password    string           `gorm:"type:varchar(128);not null;comment:密码" json:"password"`
	Email       string           `gorm:"type:varchar(254);not null;comment:email" json:"email"`
	NickName    string           `gorm:"type:varchar(32);not null;comment:显示名" json:"nick_name"`
	Mobile      string           `gorm:"type:varchar(11);not null;comment:手机号" json:"mobile"`
	AvatarFile  string           `gorm:"type:varchar(254);not null;comment:头像文件地址" json:"avatar_file"`
	RoleID      uint64           `gorm:"type:bigint;null;comment:角色ID" json:"role_id"`
	IsSuperuser bool             `gorm:"type:boolean;default:false;comment:是否为超级管理员" json:"is_superuser"`
	IsActive    bool             `gorm:"type:boolean;default:true;comment:是否激活" json:"is_active"`
	IsStaff     bool             `gorm:"type:boolean;default:false;comment:是否为员工" json:"is_staff"`
	IsTwoFA     bool             `gorm:"type:boolean;default:false;comment:是否启用2FA认证" json:"is_two_fa"`
	OtpSecret   string           `gorm:"type:varchar(128);not null;comment:otp_secret" json:"otp_secret"`
	LastLogin   models.LocalTime `gorm:"autoUpdateTime;comment:最后一次登录" json:"last_login"`                    // 更新该条记录时更新
	DateJoined  models.LocalTime `gorm:"index:date_joined;autoCreateTime;comment:加入时间" json:"date_joined"`   // 第一次创建时更新，后续不变
	UpdatedAt   models.LocalTime `gorm:"index:idx_updated_at;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新该条记录时更新
}

func (User) TableName() string {
	return "user"
}

type OCSToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}

// 定义一个结构体来表示 JSON 数据的结构
type OCSResponse struct {
	OCS struct {
		Meta struct {
			Status     string `json:"status"`
			StatusCode int    `json:"statuscode"`
			Message    string `json:"message"`
		} `json:"meta"`
		Data struct {
			// Enabled     bool     `json:"enabled"`
			ID          string   `json:"id"`
			Email       string   `json:"email"`
			DisplayName string   `json:"displayname"`
			Role        string   `json:"role"`
			Groups      []string `json:"groups"`
			// 其他字段省略...
		} `json:"data"`
	} `json:"ocs"`
}

func BcryptPW(password string) string {
	cost := 10
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword)
}
