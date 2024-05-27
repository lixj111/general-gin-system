package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"example.com/first_gin_attempt/bootstrap"
	"example.com/first_gin_attempt/global"
	"example.com/first_gin_attempt/middleware"
	"example.com/first_gin_attempt/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	categoryRouter "example.com/first_gin_attempt/internal/apps/categories/router"
	goodsDetailsRouter "example.com/first_gin_attempt/internal/apps/goods_details/router"
	userRouter "example.com/first_gin_attempt/internal/apps/users/router"
	/* _符号通常用作空白标识符。在导入包时，如果我们只需要使用该包中的一些功能，
	但不需要直接引用该包中的任何标识符（如变量、函数或结构体），可以使用空白标识符 _ 来避免编译器报未使用的警告	*/)

var version string

// 用于读取本地配置文件，等价于启动时使用 `-config config.yaml`参数指定配置文件路径
var ConfigFile = flag.String("config", "config.yaml", "config file")

// go:embed dist
var staticFS embed.FS

func RunServer() {
	// 生产环境模式
	if global.App.Config.App.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化认证中间件
	global.App.JWT, _ = middleware.InitAuthMiddleware()

	routers.Include(userRouter.Routers)
	routers.Include(categoryRouter.Routers)
	routers.Include(goodsDetailsRouter.Routers)

	r := routers.Init()

	// 嵌入的文件系统
	st, err := fs.Sub(staticFS, "dist")
	if err != nil {
		fmt.Println("Error accessing embedded filesystem:", err)
		return
	}
	r.StaticFS("/static", http.FS(st))

	// 提供其他非嵌入的文件系统
	// r.StaticFS("/media", http.Dir("media"))

	// 默认头像文件
	r.StaticFile("/avatar.png", "dist/avatar.png") // http://127.0.0.1:8080/avatar.png

	if err := r.Run(global.App.Config.App.ListenAddress); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 打印版本
	if version != "" {
		fmt.Println("goInsight Version：", version)
	}

	// 解析输入
	flag.Parse()
	configFilePath := *ConfigFile
	bootstrap.InitializeConfig(configFilePath)
	bootstrap.InitializeLog()
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	RunServer()
}
