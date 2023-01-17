package initializer

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Gin(r *gin.Engine) {
	setCors(r)
	setProxies(r)
}

func setCors(e *gin.Engine) {
	e.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("gin.cors"),
		AllowMethods:     viper.GetStringSlice("gin.allow_methods"),
		AllowHeaders:     viper.GetStringSlice("gin.allow_headers"),
		ExposeHeaders:    viper.GetStringSlice("gin.expose_headers"),
		MaxAge:           5 * time.Minute,
		AllowCredentials: true,
	}))
	return
}

func setProxies(e *gin.Engine) {
	err := e.SetTrustedProxies([]string{viper.GetString("server.host")})
	if err != nil {
		log.Fatal(err)
	}
	return
}
