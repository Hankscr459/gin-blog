package router

import (
	"gin-blog/middleware/auth"
	"gin-blog/middleware/validDto"
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/jinzhu/copier"

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

	userRoute.POST("/signin", validDto.SigninValidator(), func(ctx *gin.Context) {
		value, _ := ctx.Get("SigninUser")
		body := value.(dto.SigninUser)
		user, err := configs.CheckUser(body.Email, body.Password)
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		readUser := dto.ReadUser{}
		copier.Copy(&readUser, user)
		tk, err := configs.GenerJWT(readUser)
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		data := gin.H{"token": tk}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.GET("/user/:id", auth.User(), func(ctx *gin.Context) {
		user, err := User.FindById(ctx.Param("id"))
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/list", func(ctx *gin.Context) {
		user, err := User.Find()
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.PUT("/:id", auth.User(), func(ctx *gin.Context) {
		//
	})
}
