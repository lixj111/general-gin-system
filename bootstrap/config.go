package bootstrap

import (
	"fmt"

	"example.com/first_gin_attempt/global"

	"github.com/spf13/viper"
)

func InitializeConfig(config string) *viper.Viper {
	// 初始化viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
