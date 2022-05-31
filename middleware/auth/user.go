package auth

import (
	"fmt"
	"gin-blog/plugins/configs"

	"github.com/gin-gonic/gin"
)

var Error = configs.Error()

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tk = c.Request.Header.Get("Authorization")
		claims, _, _, err := configs.ProccessToken(tk)
		fmt.Println(err)
		if err != nil {
			Error.ErrorMessage(err, c)
			c.Abort()
			return
		} else {
			c.Set("claims", claims)
			c.Next()
		}
	}
}
