package configs

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func timing(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(string)
	if ok && date != "" {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			return false
		}
	}
	return true
}

func Valid() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", timing)
		if err != nil {
			fmt.Println("success")
		}
	}
}
