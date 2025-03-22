package main

import (
	"FashOJ_Backend/config"
	"FashOJ_Backend/global"
	"FashOJ_Backend/router"
	"FashOJ_Backend/utils"
	"fmt"
)

func main(){
	var err error
	if global.JwtKey,err = utils.GenerateHMACKey(); err != nil{
		fmt.Println("error when Gen JwtKey")
		return
	}
	config.InitConfig()
	r:=router.SetupRouter()
	r.Run(":"+config.FashOJConfig.FashOJApp.Port)
}