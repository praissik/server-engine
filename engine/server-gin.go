package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/praissik/web-app-engine/engine/initializer"
	"github.com/spf13/viper"
	"log"
)

func PrepareGinServer() (*gin.Engine, string) {
	gin.SetMode(viper.GetString("gin.mode"))

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	initializer.Gin(r)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

	return r, host + ":" + port
}

func RunGinServer(r *gin.Engine, addr string) {
	err := r.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
	return
}
