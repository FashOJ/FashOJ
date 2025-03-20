package main

import (
	"FashOJ_Backend/config"
	"FashOJ_Backend/router"
)

func main(){
	config.InitConfig()
	r:=router.SetupRouter()
	r.Run(":"+config.AppConfig.App.Port)
}