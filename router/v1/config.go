package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterConfigRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/config")

	configRoute.POST("/create", func(ctx *gin.Context) {
		var b map[string]interface{}
		err := ctx.ShouldBindJSON(&b)
		Id, err := Coll("configs", b).Insert(b)
		if err != err {
			Error.ErrorMessage(err, ctx)
			return
		}
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})
}
