package auth

import (
	"gin-blog/plugins/configs"

	"github.com/gin-gonic/gin"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tk = c.Request.Header.Get("Authorization")
		claims, _, _, err := configs.ProccessToken(tk)
		if err != nil {
			configs.ErrorMessage(err, c)
			c.Abort()
			return
		} else {
			c.Set("claims", claims)
			c.Next()
		}
	}
}
