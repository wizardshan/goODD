package main

import (
	"github.com/gin-gonic/gin"
	"goODD/chapter1/controller"
)

func main() {

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.GET("/user/login/v1", controller.Wrapper(ctrUser.LoginV1))
	engine.GET("/user/login/v2", controller.Wrapper(ctrUser.LoginV2))
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user", controller.Wrapper(ctrUser.One))
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha/v1", controller.Wrapper(ctrSms.CaptchaV1))
	engine.GET("/sms/captcha/v2", controller.Wrapper(ctrSms.CaptchaV2))
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
