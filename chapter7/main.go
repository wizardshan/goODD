package main

import (
	"chapter7/controller"
	"chapter7/repository"
	"chapter7/repository/ent"
	"chapter7/rpc/rpcconnect"
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {

	interceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal(err)
	}
	interceptors := connect.WithInterceptors(interceptor)

	dsn := "root:123456@tcp(127.0.0.1:3306)/odd?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)

	mux := http.NewServeMux()
	mux.Handle(rpcconnect.NewUserHandler(ctrUser, interceptors))

	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(rpcconnect.UserName),
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(rpcconnect.UserName),
	))

	err = http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}
}
