package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"goODD/chapter6/controller/response"
	"goODD/chapter6/domain/vo"
	"net/http"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {
	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		mobile := vo.Mobile{Value: fl.Field().String()}
		if err := mobile.Validate(); err != nil {
			return false
		}
		return true
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
