package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goODD/chapter6/controller"
	"goODD/chapter6/repository"
	"goODD/chapter6/repository/ent"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/odd?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()

	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.POST("/user/:id", controller.Wrapper(ctrUser.FetchOne))
	engine.POST("/user", controller.Wrapper(ctrUser.One))
	engine.POST("/users", controller.Wrapper(ctrUser.Many))
	engine.POST("/user/register", controller.Wrapper(ctrUser.Register))
	engine.POST("/user/sms/register", controller.Wrapper(ctrUser.SmsRegister))
	engine.POST("/user/login", controller.Wrapper(ctrUser.Login))
	engine.POST("/user/modify", controller.Wrapper(ctrUser.Modify))
	engine.POST("/user/cash", controller.Wrapper(ctrUser.Cash))

	engine.Run()
}
