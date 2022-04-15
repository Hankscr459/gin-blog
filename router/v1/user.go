package router

import (
	"gin-blog/middleware/validDto"
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var User = configs.User()
var Error = configs.Error()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.POST("/create", validDto.SignupValidator(), func(ctx *gin.Context) {
		value, _ := ctx.Get("user")
		body := value.(dto.SignupUser)
		Id, _, err := User.Signup(body)
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := User.FindById(ctx.Param("id"))
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})
}
