package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goODD/chapter7/controller/response"
	"net/http"
)

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

type Validator interface {
	Validate() error
}

type ctr struct{}

func (ctr *ctr) Bind(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		return err
	}
	if validator, ok := obj.(Validator); ok {
		if err := validator.Validate(); err != nil {
			return err
		}
	}
	return nil
}
