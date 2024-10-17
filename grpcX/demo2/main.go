package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"context"
	"demo2/handler"
	"demo2/rpc/rpcconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type Validator interface {
	Validate() error
}

func main() {

	interceptor := connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
				if r, ok := request.Any().(Validator); ok {
					if err := r.Validate(); err != nil {
						return nil, status.Error(codes.InvalidArgument, err.Error())
					}
				}
				return next(ctx, request)
			}
		},
	)
	interceptors := connect.WithInterceptors(interceptor)

	mux := http.NewServeMux()
	mux.Handle(rpcconnect.NewUserHandler(&handler.User{}, interceptors))

	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(rpcconnect.UserName),
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(rpcconnect.UserName),
	))

	err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}
}
