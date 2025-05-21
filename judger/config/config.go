package config

import (
	"github.com/spf13/viper"
	"log"
)

// Structure of FashOJ configuration information.
type Config struct {
	FashOJApp struct { // FashOJApp is the configuration of FashOJ.
		Port string // port of FashOJ.
		LogPath  string
	}
	DataBase struct { // DataBase is the configuration of database.
		UserName     string // username of database.
		Password     string // password of database.
		Host         string // host of database.
		MaxIdleConns int    // max idle connections of database.
		MaxOpenConns int    // max open connections of database.
	}
}

// AppConfig is the global configuration of FashOJ.
var FashOJConfig *Config

// InitConfig() initializes the configuration of FashOJ.
// It reads the configuration file and unmarshals it to the global configuration.
// It also initializes the database.
func InitConfig() {
	viper.AddConfigPath("./config") // The path of the configuration file.
	viper.SetConfigName("config")   // The name of the configuration file.
	viper.SetConfigType("yaml")     // The type of the configuration file.

	// Read the configuration file.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read config file, Error: %v \n", err)
	}

	// Unmarshal the configuration file to the global configuration.
	FashOJConfig = &Config{}

	// Unmarshal the configuration file to the global configuration.
	if err := viper.Unmarshal(FashOJConfig); err != nil {
		log.Fatalf("Cat unmarshal the config file, Error: %v \n", err)
	}

	// Initialize the database.
	initDb()
}

