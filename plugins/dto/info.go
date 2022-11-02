package dto

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

type PageParamsInput struct {
	Ctx        *gin.Context `json:"ctx" bson:"Ctx"`
	SearchType []string     `json:"search_type"`
	Sort       string       `json:"sort"` // "asc" or "desc" 1, -1 exp: price,-1
}
