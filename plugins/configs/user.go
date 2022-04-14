package configs

import (
	"context"
	"fmt"
	"gin-blog/plugins/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DB = MongoCN.Database("userdb")
var UserCol = DB.Collection("users")

type UserService interface {
	Signup(dto.SignupUser) (string, bool, error)
	FindByEmail(string) (*dto.ReadUser, error)
}
type user struct{}

func User() UserService {
	return &user{}
}

func (*user) Signup(u dto.SignupUser) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	u.Password, _ = EncriptPassword(u.Password)
	result, err := UserCol.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

func (*user) FindByEmail(email string) (*dto.ReadUser, error) {
	var user *dto.ReadUser
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := bson.M{"email": email}
	err := UserCol.FindOne(ctx, query).Decode(&user)
	fmt.Println("err: ", err)
	return user, err
}

// func InsertoRegister(u models.User) (string, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()
// 	db := MongoCN.Database("go_twitter")
// 	col := db.Collection("users")
// 	u.Password, _ = EncriptPassword(u.Password)

// 	result, err := col.InsertOne(ctx, u)
// 	if err != nil {
// 		return "", false, err
// 	}
// 	ObjID, _ := result.InsertedID.(primitive.ObjectID)
// 	return ObjID.String(), true, nil
// }
