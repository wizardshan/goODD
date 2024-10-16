package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goODD/chapter8/controller"
	"goODD/chapter8/repository"
	"goODD/chapter8/repository/ent"
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
	engine.POST("/users1", controller.Wrapper(ctrUser.Many1))
	engine.POST("/users2", controller.Wrapper(ctrUser.Many2))
	engine.POST("/users3", controller.Wrapper(ctrUser.Many3))

	engine.Run()
}
