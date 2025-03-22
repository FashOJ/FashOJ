package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	App struct{
		Port string
	}
	DataBase struct {
		UserName string
		Password string
		Host     string
		MaxIdleConns int
		MaxOpenConns int

	}
}

var AppConfig *Config

func InitConfig() {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/main?charset=utf8mb4&parseTime=True&loc=Local")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("cannot read config file")
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("cat unmarshal the config file %v", err)
	}

	initDb()
}
