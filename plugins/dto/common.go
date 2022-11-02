package dto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cliam struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name  string             `json:"name"`
	jwt.StandardClaims
}

type PageParamsInput struct {
	Ctx        *gin.Context `json:"ctx" bson:"Ctx"`
	SearchType []string     `json:"search_type"`
	// "asc" or "desc" 1, -1 exp: price,-1
	DeSelect []string `json:"de_select"`
}
