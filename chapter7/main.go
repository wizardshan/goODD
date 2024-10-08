package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goODD/chapter7/controller"
	"goODD/chapter7/repository"
	"goODD/chapter7/repository/ent"
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
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user", controller.Wrapper(ctrUser.One))
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
