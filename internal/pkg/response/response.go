package response

import (
	"fmt"
	"net/http"

	"example.com/first_gin_attempt/global"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// code 字段是表示请求状态
// code 的值为0000，表示请求成功；code 的值为0001，表示请求失败
type Response struct {
	RequestID string      `json:"request_id"`
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func writeResponse(c *gin.Context, code string, data interface{}, msg string) {
	claims := jwt.ExtractClaims(c)
	requestID := requestid.Get(c)
	username := claims["id"].(string)

	if code == "0000" {
		global.App.Log.WithField("request_id", requestID).WithField("username", username).Info(msg)
	} else {
		global.App.Log.WithField("request_id", requestID).WithField("username", username).Error(msg)
	}

	c.JSON(http.StatusOK, Response{requestID, code, data, msg})
}

func Fail(c *gin.Context, msg string) {
	writeResponse(c, "0001", nil, msg)
}

func Success(c *gin.Context, data interface{}, msg string) {
	fmt.Println(data)
	writeResponse(c, "0000", data, msg)
}

func ValidateFail(c *gin.Context, msg string) {
	writeResponse(c, "0001", nil, msg)
}
