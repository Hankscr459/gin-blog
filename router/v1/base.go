package router

import (
	"gin-blog/plugins/configs"
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
