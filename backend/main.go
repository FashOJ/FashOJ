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

	defer global.Logger.Sync()
	global.Logger.Info("InitConfig and Logger")	
	var err error
	if global.JwtKey,err = utils.GenerateHMACKey(); err != nil{
		global.Logger.Errorf("Error at generate jwt key Error: %v",err)
		return
	}

	r:=router.SetupRouter()
	r.Run(":"+config.FashOJConfig.FashOJApp.Port)
}