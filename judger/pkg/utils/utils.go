package utils


import (
	"FashOJ/Judger/internal/global"
	"FashOJ/Judger/internal/models"
)

func AutoMigrate() {
	if err := global.DB.AutoMigrate(
		&models.File{},
		&models.Problem{},
		&models.Testcase{},
	); err != nil {
		panic(err)
	}
}