package cmd

import (
	"fmt"
	"gin-blog/plugins/configs"
	"gin-blog/router/v1"
	"log"
	"os"

	_ "gin-blog/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/joho/godotenv/autoload"
)

var port string
var runSwagger bool

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(c *cobra.Command, args []string) {
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
		// http://localhost:1000/swagger/index.html
		if runSwagger {
			url := ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:%s/swagger/doc.json", port))
			server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}
		server.Run(":" + port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	// --p YourPort
	serverCmd.Flags().StringVar(&port, "p", os.Getenv("PORT"), "PORT")
	// -d
	serverCmd.Flags().BoolVarP(&runSwagger, "docs", "d", false, "Api 文件明細")
}
