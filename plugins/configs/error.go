package configs

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DtoError(err error, c *gin.Context, model interface{}) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			var errorMessage string = ""
			field, ok := reflect.TypeOf(model).Elem().FieldByName(fe.Field())
			if !ok {
				panic("Field not found")
			}
			label := string(field.Tag.Get("label"))
			switch fe.Tag() {
			case "required":
				errorMessage = label + " 欄位必填"
			case "email":
				errorMessage = "Email 格式不對"
			default:
				errorMessage = fe.Error()
			}
			out[i] = errorMessage
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  400,
			"message": "Dto驗證錯誤",
			"error": gin.H{
				"list": out,
			},
		})
		c.Abort()
	} else {
		ErrorMessage(err, c)
		c.Abort()
	}
}

func ErrorMessage(err error, c *gin.Context) {
	if err != nil {
		errorMessage := err.Error()
		if errorMessage == "mongo: no documents in result" {
			errorMessage = "找不到該資料"
		} else if errorMessage == "the provided hex string is not a valid ObjectID" {
			errorMessage = "Id格式不對"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  400,
			"message": errorMessage,
			"error":   err,
		})
		panic(err.Error())
	}
}
