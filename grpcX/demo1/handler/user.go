package handler

import (
	"connectrpc.com/connect"
	"context"
	"demo1/domain/vo"
	"demo1/rpc"
	"demo1/rpc/rpcconnect"
	"fmt"
)

type User struct {
	rpcconnect.UnimplementedUserHandler
}

func (s *User) Get(ctx context.Context, req *connect.Request[rpc.UserGetRequest]) (*connect.Response[rpc.UserGetResponse], error) {
	fmt.Println(req.Msg.GetID())
	res := connect.NewResponse(&rpc.UserGetResponse{
		ID: &vo.ID{Value: 122333},
	})
	return res, nil
}
