package router

import (
	"gin-blog/plugins/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var User = configs.User()
var Error = configs.Error()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := User.FindById(ctx.Param("id"))
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})
}
