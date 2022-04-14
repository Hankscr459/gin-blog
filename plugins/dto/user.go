package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignupUser struct {
	Name     string `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email    string `binding:"required" label:"電子郵件" json:"email" bson:"email"`
	Password string `binding:"required" label:"密碼" json:"password" bson:"password"`
}

type ReadUser struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id" label:"_id"`
	Name  string             `binding:"required" label:"名稱" json:"name" bson:"name"`
	Email string             `binding:"required" label:"電子郵件" json:"email" bson:"email"`
}
