package main

import (
	"gin-blog/plugins/configs"
	"gin-blog/router/v1"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	configs.Valid()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	if configs.CheckConnection() == 0 {
		log.Fatal("Fail to connect DB")
		return
	}
	basepath := server.Group("/v1")
	router.RegisterUserRoutes(basepath)
	router.RegisterConfigRoutes(basepath)
	router.RegisterOrgRoutes(basepath)
	server.Run(os.Getenv("PORT"))
}
