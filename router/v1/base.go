package router

import (
	"gin-blog/middleware/valid"
	"gin-blog/plugins/configs"

	"github.com/gin-gonic/gin"
)

var db = configs.Database{}

func CollR[T any](collName string, dto T) *configs.Collection[T] {
	coll := configs.CollR(collName, dto)
	return coll
}

func CollW(collName string) *configs.MyColl {
	coll := configs.CollW(collName)
	return coll
}

func Valid[T any]() gin.HandlerFunc {
	return valid.Dto[T]()
}

func ErrorMessage(err error, ctx *gin.Context) {
	configs.ErrorMessage(err, ctx)
}

func Body[T any](ctx *gin.Context) T {
	return configs.Body[T](ctx)
}
