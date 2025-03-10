package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 测试程序
func main() {
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 运行服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
