package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetupUsersMiddlewares 设置中间件
// Users路由指定中间件
func SetupUsersMiddlewares(r *gin.Engine) {
	// 在这里添加需要使用的中间件
	r.Use(TokenAuthMiddleware)
	r.Use(CORSMiddleware)
}

func SetupCategoryMiddlewares(r *gin.Engine) {
	// 在这里添加需要使用的中间件
	r.Use(TokenAuthMiddleware)
	r.Use(CORSMiddleware)
}

func SetupGoodsDetailsMiddlewares(r *gin.Engine) {
	// 在这里添加需要使用的中间件
	r.Use(TokenAuthMiddleware)
	r.Use(CORSMiddleware)
}

// TokenAuthMiddleware Token 验证中间件
func TokenAuthMiddleware(c *gin.Context) {
	// 从请求中获取 Token
	token := c.GetHeader("Authorization")

	// 验证 Token 的有效性，如果无效则返回未授权的错误
	if !isValidToken(token) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

// LoggerMiddleware 日志记录中间件
func LoggerMiddleware(c *gin.Context) {
	// 实现日志记录逻辑
	// fmt.Println("LoggerMiddleware")
	c.Next()
}

// CORSMiddleware 跨域请求中间件
func CORSMiddleware(c *gin.Context) {
	// 实现跨域请求处理逻辑
	// fmt.Println("CORSMiddleware")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Max-Age", "3600")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

// ValidationMiddleware 请求参数验证中间件
func ValidationMiddleware(c *gin.Context) {
	// 实现请求参数验证逻辑
	// fmt.Println("ValidationMiddleware")
	c.Next()
}

// 判断 Token 是否有效的函数
func isValidToken(token string) bool {
	// 在这里进行 Token 的验证，例如验证 Token 的签名、过期时间等
	return token == "valid_token"
}
