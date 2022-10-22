package router

import (
	"gin-blog/middleware/auth"
	"gin-blog/middleware/valid"
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

	userRoute.POST("/create", valid.Dto[dto.SignupUser](), func(ctx *gin.Context) {
		body := configs.Body[dto.SignupUser](ctx)
		encodePassword, err := configs.EncriptPassword(body.Password)
		body.Password = encodePassword
		Error.ErrorMessage(err, ctx)
		Id, err := Coll("users", body).Insert(body)
		Error.ErrorMessage(err, ctx)
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.POST("/signin", valid.Dto[dto.SigninUser](), func(ctx *gin.Context) {
		body := configs.Body[dto.SigninUser](ctx)
		user, err := configs.CheckUser(body.Email, body.Password)
		Error.ErrorMessage(err, ctx)
		readUser := dto.ReadUser{}
		copier.Copy(&readUser, user)
		tk, err := configs.GenerJWT(readUser)
		Error.ErrorMessage(err, ctx)
		data := gin.H{"token": tk}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := Coll("users", dto.ReadUser{}).FindById(ctx.Param("id"))
		Error.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/list", func(ctx *gin.Context) {
		user, err := User.Find()
		Error.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.PUT("/:id", auth.User(), func(ctx *gin.Context) {
		var f map[string]interface{}
		err := ctx.ShouldBindJSON(&f)
		Error.ErrorMessage(err, ctx)
		err = User.FindByIdAndUpdate(ctx.Param("id"), f)
		Error.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
