package router

import (
	"gin-blog/middleware/valid"
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
)

var User = configs.User()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.POST("/create", valid.Dto[dto.SignupUserInput](), func(ctx *gin.Context) {
		body := configs.Body[dto.SignupUserInput](ctx)
		encodePassword, err := configs.EncriptPassword(body.Password)
		body.Password = encodePassword
		configs.ErrorMessage(err, ctx)
		Id, err := Coll("users", body).Insert(body)
		configs.ErrorMessage(err, ctx)
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.POST("/signin", valid.Dto[dto.SigninUserInput](), func(ctx *gin.Context) {
		body := configs.Body[dto.SigninUserInput](ctx)
		user, err := configs.CheckUser(body.Email, body.Password)
		configs.ErrorMessage(err, ctx)
		readUser := dto.ReadUser{}
		copier.Copy(&readUser, user)
		tk, err := configs.GenerJWT(readUser)
		configs.ErrorMessage(err, ctx)
		data := gin.H{"token": tk}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := Coll("users", dto.ReadUser{}).FindById(ctx.Param("id"))
		configs.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/list", func(ctx *gin.Context) {
		user, err := User.Find()
		configs.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.PUT("/:id", valid.Dto[dto.UpdateUserInput](), func(ctx *gin.Context) {
		u := configs.Body[dto.UpdateUserInput](ctx)
		if u.Password != "" {
			encodePassword, matchErr := configs.EncriptPassword(u.Password)
			u.Password = encodePassword
			configs.ErrorMessage(matchErr, ctx)
		}
		err := Coll("users", u).FindByIdAndUpdate(ctx.Param("id"), u)
		configs.ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
