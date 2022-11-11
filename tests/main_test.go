package tests

import (
	"fmt"
	"gin-blog/plugins/configs"
	"gin-blog/router/v1"
	"log"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupTsetServer() *gin.Engine {
	configs.Valid()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	if configs.CheckConnection() == 0 {
		log.Fatal("Fail to connect DB")
		return nil
	}
	basepath := server.Group("/v1")
	router.RegisterUserRoutes(basepath)
	router.RegisterConfigRoutes(basepath)
	router.RegisterOrgRoutes(basepath)
	return server
}

func Test(t *testing.T) {
	server := SetupTsetServer()

	UserTest(server, t)

	fmt.Println("env: ", os.Getenv("MongoApplyURI"))
	fmt.Println("DbName: ", os.Getenv("DbName"))
	fmt.Println("Mode: ", os.Getenv("Mode"))
}

// func init() {
// 	fmt.Println("db: ", os.Getenv("DbName"))
// }
