package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"example.com/first_gin_attempt/global"
	usersModels "example.com/first_gin_attempt/internal/apps/users/models"
	"example.com/first_gin_attempt/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

/*
http://127.0.0.1:8181/apps/oauth2/authorize?client_id=tCiNuCV1cpdf5qiGLWQghf72dBekbhwpIFO4Kx9vx04s5q6OouWEeeogo59R7QhT&redirect_uri=http://127.0.0.1:8202/user/sso/callback&response_type=code
*/
func SSOCallback(c *gin.Context) {
	code := c.Query("code")
	ocrToken, err := getTokenByCode(code)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	accessToken := ocrToken.AccessToken
	ocrUser, err := getUserInfo(accessToken)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	if ocrUser.OCS.Meta.StatusCode != 200 {
		response.Fail(c, "OCS.Meta.StatusCode != 200")
		return
	}

	// TODO：验证用户信息
	err = userLogin(*ocrUser)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	var user usersModels.User
	global.App.DB.Table("user u").
		Where("u.username=?", ocrUser.OCS.Data.ID).
		Scan(&user)
	c.Set("loginUserName", user.Username)
	c.Set("loginOtpCode", user.OtpSecret)
	needsOTP := user.IsTwoFA
	if needsOTP {
		c.Set("loginNeedsOTP", "YES")
	} else {
		c.Set("loginNeedsOTP", "NO")
	}

	c.Next()
}

func getTokenByCode(code string) (*usersModels.OCSToken, error) {
	config := global.App.Config.Nextcloud
	token_api_url := config.TokenUrl

	formData := url.Values{}
	formData.Set("client_id", config.ClientID)
	formData.Set("client_secret", config.ClientSecret)
	formData.Set("grant_type", config.GrantType)
	formData.Set("code", code)
	formData.Set("redirect_uri", config.RedirectURI)

	resp, err := http.PostForm(token_api_url, formData)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var ocsToken usersModels.OCSToken
	err = json.Unmarshal(body, &ocsToken)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}

	return &ocsToken, nil
}

func getUserInfo(accessToken string) (*usersModels.OCSResponse, error) {
	user_info_api := global.App.Config.Nextcloud.UserInfoUrl

	req, err := http.NewRequest("GET", user_info_api, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var ocrUserInfo usersModels.OCSResponse
	err = json.Unmarshal([]byte(responseData), &ocrUserInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &ocrUserInfo, nil
}

func userLogin(ocrResp usersModels.OCSResponse) error {
	// TODO：验证用户信息、赋予token....
	var user usersModels.User
	if err := global.App.DB.Where("username = ?", ocrResp.OCS.Data.ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新用户
			err = createDefaultUser(ocrResp)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	return nil
}

func createDefaultUser(ocrResp usersModels.OCSResponse) error {
	var user usersModels.User
	user.Username = ocrResp.OCS.Data.ID
	user.Email = ocrResp.OCS.Data.Email
	user.Password = usersModels.BcryptPW("123456")
	user.NickName = ocrResp.OCS.Data.DisplayName
	user.AvatarFile = "/static/avatar2.jpg"

	return global.App.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&usersModels.User{}).Create(&user).Error; err != nil {
			mysqlErr := err.(*mysql.MySQLError)
			switch mysqlErr.Number {
			case 1062:
				return fmt.Errorf("用户`%s`已存在", user.Username)
			}
			global.App.Log.Error(err)
			return err
		}
		return nil
	})
}
