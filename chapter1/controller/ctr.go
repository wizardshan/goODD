package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"goODD/chapter1/controller/response"
	"net/http"
	"regexp"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {
	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(fl.Field().String())
	})
}

type HandlerFunc func(c *gin.Context) (response.Data, error)

// Wrapper 4xx client error you messed up 5xx server error I messed up
func Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusOK, response.Resp{Code: http.StatusInternalServerError, Message: fmt.Sprintf("%v", err)})
			}
		}()

		if data, err := handlerFunc(c); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, response.Resp{Code: http.StatusBadRequest, Message: err.Error()})
		} else {
			c.JSON(http.StatusOK, response.Resp{Code: http.StatusOK, Message: http.StatusText(http.StatusOK), Success: true, Data: data})
		}
	}
}
