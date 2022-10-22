package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignupUserInput struct {
	Name     string `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email    string `binding:"required,email" label:"電子郵件" json:"email" bson:"email"`
	Password string `binding:"required" label:"密碼" json:"password" bson:"password"`
	Rank     string `label:"rank" json:"rank" bson:"rank,omitempty"`
}

type ReadUser struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id" label:"_id"`
	Name  string             `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email string             `binding:"required" label:"電子郵件" json:"email" bson:"email"`
}

type SigninUserInput struct {
	Email    string `binding:"required" label:"電子郵件" json:"email" bson:"email"`
	Password string `binding:"required" label:"密碼" json:"password" bson:"password"`
}

type ReadUserWithPassword struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id" label:"_id"`
	Name     string             `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email    string             `binding:"required" label:"電子郵件" json:"email" bson:"email"`
	Password string             `binding:"required" label:"密碼" json:"password" bson:"password"`
}

type UpdateUserInput struct {
	Name     string `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email    string `binding:"required,email" label:"電子郵件" json:"email" bson:"email"`
	Password string `label:"密碼" json:"password" bson:"password,omitempty"`
}
