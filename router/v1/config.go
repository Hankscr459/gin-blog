package router

import (
	"context"
	"gin-blog/plugins/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterConfigRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/config")

	configRoute.POST("/create", func(ctx *gin.Context) {
		var b map[string]interface{}
		err := ctx.ShouldBindJSON(&b)
		Id, err := Coll("configs", b).Insert(b)
		configs.ErrorMessage(err, ctx)
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
}
