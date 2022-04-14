package main

import (
	"gin-blog/plugins/configs"
	"gin-blog/router/v1"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := gin.Default()
	if configs.CheckConnection() == 0 {
		log.Fatal("Fail to connect DB")
		return
	}
	basepath := server.Group("/v1")
	router.RegisterUserRoutes(basepath)
	server.Run(os.Getenv("PORT"))
}
