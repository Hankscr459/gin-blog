package main

import (
	"gin-blog/plugins/configs"
	"gin-blog/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	if configs.CheckConnection() == 0 {
		log.Fatal("Fail to connect DB")
		return
	}
	basepath := server.Group("/v1")
	router.RegisterUserRoutes(basepath)
	server.Run(":1000")
}
