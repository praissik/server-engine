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
		viper.AddConfigPath("./config/")
	case "local":
		viper.SetConfigName("config.local")
		viper.AddConfigPath("./config/")
	case "test":
		viper.SetConfigName("config.test")
		viper.AddConfigPath("../config/")
	default:
		log.Fatal("LAUNCH_MODE not declared correctly")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
