package configs

import (
	"context"
	"gin-blog/plugins/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DB      = MongoCN.Database("userdb")
	UserCol = DB.Collection("users")
)

type UserService interface {
	FindById(string) (*dto.ReadUser, error)
	FindByEmail(string) (*dto.ReadUserWithPassword, error)
	FindOne(bson.M) (*dto.ReadUser, error)
	Signup(dto.SignupUser) (string, bool, error)
}
type user struct{}

func User() UserService {
	return &user{}
}

func (*user) FindById(Id string) (*dto.ReadUser, error) {
	var user *dto.ReadUser
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, idErr := primitive.ObjectIDFromHex(Id)
	if idErr != nil {
		return nil, idErr
	}
	query := bson.M{"_id": objID}
	err := UserCol.FindOne(ctx, query).Decode(&user)
	return user, err
}

func (*user) FindByEmail(email string) (*dto.ReadUserWithPassword, error) {
	var user *dto.ReadUserWithPassword
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := bson.M{"email": email}
	err := UserCol.FindOne(ctx, query).Decode(&user)
	return user, err
}

func (*user) FindOne(query bson.M) (*dto.ReadUser, error) {
	var user *dto.ReadUser
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := UserCol.FindOne(ctx, query).Decode(&user)
	return user, err
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

func (*user) Signin(u dto.SigninUser) (string, bool, error) {
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
