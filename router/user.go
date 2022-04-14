package router

import (
	"gin-blog/plugins/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var User = configs.User()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.GET("/read", func(ctx *gin.Context) {
		email := ctx.Query("email")
		user, err := User.FindByEmail(email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})
}
