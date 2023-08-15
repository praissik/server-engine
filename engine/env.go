package engine

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func InitViper() {
	switch os.Getenv("LAUNCH_MODE") {
	case "prod":
		viper.SetConfigName("config.prod")
	case "local":
		viper.SetConfigName("config.local")
	case "test":
		viper.SetConfigName("config.test")
	default:
		log.Fatal("LAUNCH_MODE not declared correctly")
	}

	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
