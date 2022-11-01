package router

import (
	"gin-blog/middleware/auth"
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
)

var User = configs.User()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.POST("/create", Valid[dto.SignupUserInput](), func(ctx *gin.Context) {
		body := Body[dto.SignupUserInput](ctx)
		encodePassword, err := configs.EncriptPassword(body.Password)
		body.Password = encodePassword
		ErrorMessage(err, ctx)
		Id, err := CollW("users").Create(body)
		ErrorMessage(err, ctx)
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.POST("/signin", Valid[dto.SigninUserInput](), func(ctx *gin.Context) {
		body := Body[dto.SigninUserInput](ctx)
		user, err := configs.CheckUser(body.Email, body.Password)
		ErrorMessage(err, ctx)
		readUser := dto.ReadUser{}
		copier.Copy(&readUser, user)
		tk, err := configs.GenerJWT(readUser)
		ErrorMessage(err, ctx)
		data := gin.H{"token": tk}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})

	userRoute.GET("/:id", func(ctx *gin.Context) {
		user, err := CollR("users", dto.ReadUser{}).FindById(ctx.Param("id"))
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/t2/:id", func(ctx *gin.Context) { // 634271c62b69193cb1ba2e56 array of objectId & ...obj
		user, err := CollR("users", dto.Read3{}).FindById(ctx.Param("id"), "friends.userId")
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/t3/:id", func(ctx *gin.Context) { // 630ec8817392454303408b33 array of only objectId
		user, err := CollR("users", dto.Read4{}).FindById(ctx.Param("id"), "friends")
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/t4/:id", func(ctx *gin.Context) { // 630ec2d209e44eb70b2f75b4 objectId
		user, err := CollR("users", dto.Read5{}).FindById(ctx.Param("id"), "friend")
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/list", func(ctx *gin.Context) {
		user, err := User.Find()
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.GET("/list/test", func(ctx *gin.Context) {
		query := dto.PageParamsInput{
			Filter:     ctx.Query("search"),
			SearchType: []string{"name", "email"},
			Limit:      ctx.Query("limit"),
			Page:       ctx.Query("page"),
			Sort:       ctx.Query("sortBy"),
		}
		user, err := CollR("users", dto.ReadUser{}).Paginate(query)
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	})

	userRoute.PUT("/:id", auth.User(), Valid[dto.UpdateUserInput](), func(ctx *gin.Context) {
		u := Body[dto.UpdateUserInput](ctx)
		if u.Password != "" {
			encodePassword, matchErr := configs.EncriptPassword(u.Password)
			u.Password = encodePassword
			ErrorMessage(matchErr, ctx)
		}
		err := CollW("users").FindByIdAndUpdate(ctx.Param("id"), u)
		ErrorMessage(err, ctx)
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
