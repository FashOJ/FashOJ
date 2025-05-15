package main

import (
	"FashOJ_Backend/config"
	"FashOJ_Backend/internal/global"
	"FashOJ_Backend/internal/router"
	"FashOJ_Backend/pkg/utils"
)

func main() {
	config.InitConfig()
	utils.InitLogger()

	//Migrate all models at start
	utils.AutoMigrate()

	defer global.Logger.Sync()
	global.Logger.Info("InitConfig and Logger")

	utils.SetJwtKey()

	// 使用适配器
	r := router.SetupRouter()
	r.Run(":" + config.FashOJConfig.FashOJApp.Port)
}
