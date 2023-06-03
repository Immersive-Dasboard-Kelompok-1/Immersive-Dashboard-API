package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DB_USERNAME					string
	DB_PASS							string
	DB_HOSTNAME					string
	DB_PORT							int
	DB_NAME							string
	JWT_ACCESS_TOKEN		string 
}

func InitConfig() *AppConfig {
	appConfig := AppConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigName(".local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("error load config: ", err.Error())
		return nil
	}

	appConfig.DB_USERNAME = viper.Get("DB_USERNAME").(string)
	appConfig.DB_PASS = viper.Get("DB_PASS").(string)
	appConfig.DB_HOSTNAME = viper.Get("DB_HOSTNAME").(string)
	appConfig.DB_PORT, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
	appConfig.DB_NAME = viper.Get("DB_NAME").(string)
	appConfig.JWT_ACCESS_TOKEN = viper.Get("JWT_ACCESS_TOKEN").(string)

	return &appConfig
}