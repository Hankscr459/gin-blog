package valid

import (
	"gin-blog/plugins/configs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Dto[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody T
		if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
			configs.DtoError(err, c, &reqBody)
		} else {
			c.Set("reqBody", reqBody)
			c.Next()
		}
	}
}
