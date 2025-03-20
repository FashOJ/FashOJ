package config

import (
	"FashOJ_Backend/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDb(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/main?charset=utf8mb4&parseTime=True&loc=Local",AppConfig.DataBase.UserName,AppConfig.DataBase.Password,AppConfig.DataBase.Host)
	db,err := gorm.Open(mysql.New(mysql.Config{
		DSN:dsn,
	}),&gorm.Config{})

	if err != nil{
		log.Fatalln("error:can't connect db")
	}

	sqldb,err := db.DB()
	if err != nil{
		log.Fatalf("err:%v",err)
	}

	sqldb.SetMaxIdleConns(AppConfig.DataBase.MaxIdleConns)
	sqldb.SetMaxOpenConns(AppConfig.DataBase.MaxOpenConns)
	sqldb.SetConnMaxLifetime(time.Hour)

	global.DB = db
}