package validDto

import (
	"gin-blog/plugins/configs"
	"gin-blog/plugins/dto"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var Error = configs.Error()

func SignupValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.SignupUser
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			Error.DtoError(err, c, &user)
		} else {
			c.Set("user", user)
			c.Next()
		}
	}
}
