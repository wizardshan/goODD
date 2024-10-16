package main

import (
	"github.com/gin-gonic/gin"
	"goODD/chapter5/controller"
)

func main() {

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.POST("/user/login", controller.Wrapper(ctrUser.Login))
	engine.POST("/user", controller.Wrapper(ctrUser.One))
	engine.POST("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
