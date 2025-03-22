package config

import (
	"FashOJ_Backend/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// initDb() initializes the database.
// It connects to the database and sets the global database.
// It also sets the max idle connections, max open connections and the max lifetime of a connection.
func initDb() {
	// DataSourceName is the connection string of database.
	DataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/main?charset=utf8mb4&parseTime=True&loc=Local",
		FashOJConfig.DataBase.UserName,
		FashOJConfig.DataBase.Password,
		FashOJConfig.DataBase.Host,
	)

	// Connect to the database.
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DataSourceName,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can't connect db,Error:%v \n", err)
	}

	// Get the Mysql object of database.
	sqldb, err := database.DB()
	if err != nil {
		log.Fatalf("Can't get db,Error:%v \n", err)
	}

	// Set the max idle connections , max open connections of database and the max lifetime of a connection.
	sqldb.SetMaxIdleConns(FashOJConfig.DataBase.MaxIdleConns)
	sqldb.SetMaxOpenConns(FashOJConfig.DataBase.MaxOpenConns)
	sqldb.SetConnMaxLifetime(time.Hour)

	// Set the global database.
	global.DB = database
}
