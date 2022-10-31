package dto

import (
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupUserInput struct {
	Name     string `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email    string `binding:"required,email" label:"電子郵件" json:"email" bson:"email"`
	Password string `binding:"required" label:"密碼" json:"password" bson:"password"`
	Rank     string `label:"rank" json:"rank" bson:"rank,omitempty"`
}

type ReadUser struct {
	ID    string `bson:"_id,omitempty" json:"_id" label:"_id"`
	Name  string `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email string `binding:"required" label:"電子郵件" json:"email" bson:"email"`
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

type Read3 struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Friends []struct {
		Desc     string `json:"desc"`
		ReadUser `bson:",inline"`
	} `json:"friends,omitempty" bson:"friends,omitempty" ref:"users" type:"arrayOfObjId"`
}

type Read4 struct {
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Friends []ReadUser `json:"friends,omitempty" bson:"friends,omitempty" ref:"users"`
}

type Read5 struct {
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Friend ReadUser `json:"friend,omitempty" bson:"friend,omitempty" ref:"users" type:"obj"`
}

type UserPageParamsInput struct {
	Filter string `json:"filter"`
	Limit  string `json:"limit"`
	Page   string `json:"page"`
}

type ReadUserPage struct {
	Data       []ReadUser              `json:"data"`
	Pagination paginate.PaginationData `json:"pagination"`
}
