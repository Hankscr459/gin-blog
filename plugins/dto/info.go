package dto

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

type PageParamsInput struct {
	Ctx        *gin.Context `json:"ctx" bson:"Ctx"`
	SearchType []string     `json:"search_type"`
	// "asc" or "desc" 1, -1 exp: price,-1
	DeSelect []string `json:"de_select"`
}
