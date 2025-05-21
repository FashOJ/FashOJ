package main

import (
	"FashOJ/Judger/config"
	"FashOJ/Judger/internal/router"
	"FashOJ/Judger/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// _, filename, _, _ := runtime.Caller(0)
	// baseDir := filepath.Dir(filepath.Dir(filename))
	// sourceCPP := filepath.Join(baseDir, "testTEMP", "sourceCode", "testCPP.cpp")
	// exec := filepath.Join(baseDir, "testTEMP", "execTemp", "test")
	// input := filepath.Join(baseDir, "testTEMP", "testCase", "input1.txt")
	// answer := filepath.Join(baseDir, "testTEMP", "outputCase", "output1.txt")
	// timeLimit := 1000 // 1s

	// result := judge.Judge(sourceCPP, exec, input, answer, timeLimit)
	// fmt.Println("评测结果:", result)

	config.InitConfig()
	utils.AutoMigrate()

	app := gin.Default()

	judger := app.Group("/")
	router.SetupRouter(judger)

	app.Run(":3000")
}