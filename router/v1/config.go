package router

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-blog/plugins/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterConfigRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/config")

	configRoute.POST("/create", func(ctx *gin.Context) {
		var b map[string]interface{}
		err := ctx.ShouldBindJSON(&b)
		Id, err := CollW("configs").Create(b)
		ErrorMessage(err, ctx)
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})
	configRoute.GET("/create-index", func(ctx *gin.Context) {
		db := configs.MongoCN.Database("userdb")
		for _, f := range []string{"name", "email"} {
			_, err := db.Collection("users").Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys:    bson.D{{Key: f, Value: 1}},
					Options: options.Index().SetUnique(true),
				},
			)
			if err != nil {
				panic(err)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})

	configRoute.GET("/test/:id", func(ctx *gin.Context) {
		var b map[string]interface{}
		err := ctx.ShouldBindJSON(&b)
		config, err := CollR("configs", b).FindById(ctx.Param("id"))
		ErrorMessage(err, ctx)
		v, _ := json.Marshal(config)
		fmt.Println("string(yfile): ", string(v))
		value := gjson.Get(string(v), "data.myObj.objKey")
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": value.String()})
	})
}
