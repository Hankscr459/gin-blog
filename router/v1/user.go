package router

import (
	"fmt"
	"gin-blog/middleware/validDto"
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var User = configs.User()
var Error = configs.Error()

func Valid(user dto.SignupUser, ctx *gin.Context) {
	fmt.Println("user: ", user)
	if user.Email == "12ff32@gmail.com" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false})
		panic("Email is required")
	}
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.POST("/create", func(ctx *gin.Context) {
		var user dto.SignupUser
		if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			Error.DtoError(err, ctx, &user)
		}
		Valid(user, ctx)
		encodePassword, err := configs.EncriptPassword(user.Password)
		user.Password = encodePassword
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		Id, err := Coll("users", user).Insert(user)
		if err != err {
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

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := Coll("users", dto.ReadUser{}).FindById(ctx.Param("id"))
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

	userRoute.PUT("/:id", func(ctx *gin.Context) {
		var f map[string]interface{}
		errBody := ctx.ShouldBindJSON(&f)
		if errBody != nil {
			Error.ErrorMessage(errBody, ctx)
			return
		}
		err := User.FindByIdAndUpdate(ctx.Param("id"), f)
		if err != nil {
			Error.ErrorMessage(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
