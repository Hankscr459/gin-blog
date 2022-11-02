package router

import (
	"gin-blog/middleware/auth"
	"gin-blog/plugins/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOrgRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/org")

	configRoute.POST("/create", Valid[dto.CreateOrgInput](), auth.User(), func(ctx *gin.Context) {
		b := Body[dto.CreateOrgInput](ctx)
		user := Profile(ctx)
		b.Access = []dto.Access{{UserId: user.ID.Hex(), Role: "Owner"}}
		Id, err := CollW("orgs").Create(b)
		ErrorMessage(err, ctx)
		data := gin.H{"_id": Id}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	})
}
