package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"demo1/handler"
	"demo1/rpc/rpcconnect"
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

	mux := http.NewServeMux()
	mux.Handle(rpcconnect.NewUserHandler(&handler.User{}, interceptors))

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
