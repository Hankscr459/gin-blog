package configs

import (
	"errors"
	"gin-blog/plugins/dto"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DtoError(err error, c *gin.Context, model interface{}) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]dto.Error, len(ve))
		for i, fe := range ve {
			Error := dto.Error{}
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
			Error.Message = errorMessage
			out[i] = Error
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		c.Abort()
	} else {
		ErrorMessage(err, c)
		c.Abort()
	}
}

func ErrorMessage(err error, c *gin.Context) {
	if err != nil {
		errorMessage := err.Error()
		out := make([]dto.Error, 1)
		if errorMessage == "mongo: no documents in result" {
			out[0] = dto.Error{Message: "找不到該資料"}
		} else if errorMessage == "the provided hex string is not a valid ObjectID" {
			out[0] = dto.Error{Message: "Id格式不對"}
		} else {
			out[0] = dto.Error{Message: err.Error()}
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "errors": out})
		panic(err.Error())
	}
}
