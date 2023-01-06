package engine

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .engine file found")
	}
}

func InitViper() {
	if os.Getenv("LAUNCH_MODE") == "prod" {
		viper.SetConfigName("config.prod")
	} else if os.Getenv("LAUNCH_MODE") == "local" {
		viper.SetConfigName("config.local")
	} else {
		log.Fatal("LAUNCH_MODE not declared correctly")
	}

	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
