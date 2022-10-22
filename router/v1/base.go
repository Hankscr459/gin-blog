package router

import (
	"gin-blog/plugins/configs"
	"os"
)

var db = configs.Database{}

// var Dto = validDto.Dto[any]()

func Coll[T any](collName string, dto T) *configs.Collection[T] {
	db.Connect(os.Getenv("MongoApplyURI"), "userdb")
	return configs.GetCollection[T](&db, collName)
}
