package main

import (
	"FashOJ_Backend/config"
	"FashOJ_Backend/global"
	"FashOJ_Backend/router"
	"FashOJ_Backend/utils"
)

func main(){
	config.InitConfig()
	utils.InitLogger()

	//Migrate all models at start
	utils.AutoMigrate()

	defer global.Logger.Sync()
	global.Logger.Info("InitConfig and Logger")	
	
	utils.SetJwtKey()

	r:=router.SetupRouter()
	r.Run(":"+config.FashOJConfig.FashOJApp.Port)
}