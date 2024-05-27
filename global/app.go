package global

import (
	configs "example.com/first_gin_attempt/configs"
	"gorm.io/gorm"

	// "github.com/redis/go-redis/v9"
	// "github.com/robfig/cron/v3"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/sirupsen/logrus"
	// "github.com/spf13/viper"
	// "gorm.io/gorm"
)

type Application struct {
	// ConfigViper *viper.Viper
	Config configs.Configuration
	JWT    *jwt.GinJWTMiddleware
	Log    *logrus.Logger
	DB     *gorm.DB
	// Redis       *redis.Client
	// Cron        *cron.Cron
}

var App = new(Application)
