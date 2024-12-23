// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc/user.proto

package rpcconnect

import (
	rpc "chapter8/rpc"
	request "chapter8/rpc/request"
	response "chapter8/rpc/response"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserName is the fully-qualified name of the User service.
	UserName = "rpc.User"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserOneProcedure is the fully-qualified name of the User's One RPC.
	UserOneProcedure = "/rpc.User/One"
	// UserManyProcedure is the fully-qualified name of the User's Many RPC.
	UserManyProcedure = "/rpc.User/Many"
	// UserLoginProcedure is the fully-qualified name of the User's Login RPC.
	UserLoginProcedure = "/rpc.User/Login"
	// UserSmsRegisterProcedure is the fully-qualified name of the User's SmsRegister RPC.
	UserSmsRegisterProcedure = "/rpc.User/SmsRegister"
	// UserRegisterProcedure is the fully-qualified name of the User's Register RPC.
	UserRegisterProcedure = "/rpc.User/Register"
	// UserModifyProcedure is the fully-qualified name of the User's Modify RPC.
	UserModifyProcedure = "/rpc.User/Modify"
	// UserCashProcedure is the fully-qualified name of the User's Cash RPC.
	UserCashProcedure = "/rpc.User/Cash"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceDescriptor           = rpc.File_rpc_user_proto.Services().ByName("User")
	userOneMethodDescriptor         = userServiceDescriptor.Methods().ByName("One")
	userManyMethodDescriptor        = userServiceDescriptor.Methods().ByName("Many")
	userLoginMethodDescriptor       = userServiceDescriptor.Methods().ByName("Login")
	userSmsRegisterMethodDescriptor = userServiceDescriptor.Methods().ByName("SmsRegister")
	userRegisterMethodDescriptor    = userServiceDescriptor.Methods().ByName("Register")
	userModifyMethodDescriptor      = userServiceDescriptor.Methods().ByName("Modify")
	userCashMethodDescriptor        = userServiceDescriptor.Methods().ByName("Cash")
)

// UserClient is a client for the rpc.User service.
type UserClient interface {
	One(context.Context, *connect.Request[request.UserOne]) (*connect.Response[response.UserOne], error)
	Many(context.Context, *connect.Request[request.UserMany]) (*connect.Response[response.UserMany], error)
	Login(context.Context, *connect.Request[request.UserLogin]) (*connect.Response[response.UserLogin], error)
	SmsRegister(context.Context, *connect.Request[request.UserSmsRegister]) (*connect.Response[response.UserSmsRegister], error)
	Register(context.Context, *connect.Request[request.UserRegister]) (*connect.Response[response.UserRegister], error)
	Modify(context.Context, *connect.Request[request.UserModify]) (*connect.Response[response.UserModify], error)
	Cash(context.Context, *connect.Request[request.UserCash]) (*connect.Response[response.UserCash], error)
}

// NewUserClient constructs a client for the rpc.User service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userClient{
		one: connect.NewClient[request.UserOne, response.UserOne](
			httpClient,
			baseURL+UserOneProcedure,
			connect.WithSchema(userOneMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		many: connect.NewClient[request.UserMany, response.UserMany](
			httpClient,
			baseURL+UserManyProcedure,
			connect.WithSchema(userManyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[request.UserLogin, response.UserLogin](
			httpClient,
			baseURL+UserLoginProcedure,
			connect.WithSchema(userLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		smsRegister: connect.NewClient[request.UserSmsRegister, response.UserSmsRegister](
			httpClient,
			baseURL+UserSmsRegisterProcedure,
			connect.WithSchema(userSmsRegisterMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		register: connect.NewClient[request.UserRegister, response.UserRegister](
			httpClient,
			baseURL+UserRegisterProcedure,
			connect.WithSchema(userRegisterMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		modify: connect.NewClient[request.UserModify, response.UserModify](
			httpClient,
			baseURL+UserModifyProcedure,
			connect.WithSchema(userModifyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		cash: connect.NewClient[request.UserCash, response.UserCash](
			httpClient,
			baseURL+UserCashProcedure,
			connect.WithSchema(userCashMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userClient implements UserClient.
type userClient struct {
	one         *connect.Client[request.UserOne, response.UserOne]
	many        *connect.Client[request.UserMany, response.UserMany]
	login       *connect.Client[request.UserLogin, response.UserLogin]
	smsRegister *connect.Client[request.UserSmsRegister, response.UserSmsRegister]
	register    *connect.Client[request.UserRegister, response.UserRegister]
	modify      *connect.Client[request.UserModify, response.UserModify]
	cash        *connect.Client[request.UserCash, response.UserCash]
}

// One calls rpc.User.One.
func (c *userClient) One(ctx context.Context, req *connect.Request[request.UserOne]) (*connect.Response[response.UserOne], error) {
	return c.one.CallUnary(ctx, req)
}

// Many calls rpc.User.Many.
func (c *userClient) Many(ctx context.Context, req *connect.Request[request.UserMany]) (*connect.Response[response.UserMany], error) {
	return c.many.CallUnary(ctx, req)
}

// Login calls rpc.User.Login.
func (c *userClient) Login(ctx context.Context, req *connect.Request[request.UserLogin]) (*connect.Response[response.UserLogin], error) {
	return c.login.CallUnary(ctx, req)
}

// SmsRegister calls rpc.User.SmsRegister.
func (c *userClient) SmsRegister(ctx context.Context, req *connect.Request[request.UserSmsRegister]) (*connect.Response[response.UserSmsRegister], error) {
	return c.smsRegister.CallUnary(ctx, req)
}

// Register calls rpc.User.Register.
func (c *userClient) Register(ctx context.Context, req *connect.Request[request.UserRegister]) (*connect.Response[response.UserRegister], error) {
	return c.register.CallUnary(ctx, req)
}

// Modify calls rpc.User.Modify.
func (c *userClient) Modify(ctx context.Context, req *connect.Request[request.UserModify]) (*connect.Response[response.UserModify], error) {
	return c.modify.CallUnary(ctx, req)
}

// Cash calls rpc.User.Cash.
func (c *userClient) Cash(ctx context.Context, req *connect.Request[request.UserCash]) (*connect.Response[response.UserCash], error) {
	return c.cash.CallUnary(ctx, req)
}

// UserHandler is an implementation of the rpc.User service.
type UserHandler interface {
	One(context.Context, *connect.Request[request.UserOne]) (*connect.Response[response.UserOne], error)
	Many(context.Context, *connect.Request[request.UserMany]) (*connect.Response[response.UserMany], error)
	Login(context.Context, *connect.Request[request.UserLogin]) (*connect.Response[response.UserLogin], error)
	SmsRegister(context.Context, *connect.Request[request.UserSmsRegister]) (*connect.Response[response.UserSmsRegister], error)
	Register(context.Context, *connect.Request[request.UserRegister]) (*connect.Response[response.UserRegister], error)
	Modify(context.Context, *connect.Request[request.UserModify]) (*connect.Response[response.UserModify], error)
	Cash(context.Context, *connect.Request[request.UserCash]) (*connect.Response[response.UserCash], error)
}

// NewUserHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserHandler(svc UserHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userOneHandler := connect.NewUnaryHandler(
		UserOneProcedure,
		svc.One,
		connect.WithSchema(userOneMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userManyHandler := connect.NewUnaryHandler(
		UserManyProcedure,
		svc.Many,
		connect.WithSchema(userManyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userLoginHandler := connect.NewUnaryHandler(
		UserLoginProcedure,
		svc.Login,
		connect.WithSchema(userLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userSmsRegisterHandler := connect.NewUnaryHandler(
		UserSmsRegisterProcedure,
		svc.SmsRegister,
		connect.WithSchema(userSmsRegisterMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userRegisterHandler := connect.NewUnaryHandler(
		UserRegisterProcedure,
		svc.Register,
		connect.WithSchema(userRegisterMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userModifyHandler := connect.NewUnaryHandler(
		UserModifyProcedure,
		svc.Modify,
		connect.WithSchema(userModifyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userCashHandler := connect.NewUnaryHandler(
		UserCashProcedure,
		svc.Cash,
		connect.WithSchema(userCashMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/rpc.User/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserOneProcedure:
			userOneHandler.ServeHTTP(w, r)
		case UserManyProcedure:
			userManyHandler.ServeHTTP(w, r)
		case UserLoginProcedure:
			userLoginHandler.ServeHTTP(w, r)
		case UserSmsRegisterProcedure:
			userSmsRegisterHandler.ServeHTTP(w, r)
		case UserRegisterProcedure:
			userRegisterHandler.ServeHTTP(w, r)
		case UserModifyProcedure:
			userModifyHandler.ServeHTTP(w, r)
		case UserCashProcedure:
			userCashHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserHandler returns CodeUnimplemented from all methods.
type UnimplementedUserHandler struct{}

func (UnimplementedUserHandler) One(context.Context, *connect.Request[request.UserOne]) (*connect.Response[response.UserOne], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.One is not implemented"))
}

func (UnimplementedUserHandler) Many(context.Context, *connect.Request[request.UserMany]) (*connect.Response[response.UserMany], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.Many is not implemented"))
}

func (UnimplementedUserHandler) Login(context.Context, *connect.Request[request.UserLogin]) (*connect.Response[response.UserLogin], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.Login is not implemented"))
}

func (UnimplementedUserHandler) SmsRegister(context.Context, *connect.Request[request.UserSmsRegister]) (*connect.Response[response.UserSmsRegister], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.SmsRegister is not implemented"))
}

func (UnimplementedUserHandler) Register(context.Context, *connect.Request[request.UserRegister]) (*connect.Response[response.UserRegister], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.Register is not implemented"))
}

func (UnimplementedUserHandler) Modify(context.Context, *connect.Request[request.UserModify]) (*connect.Response[response.UserModify], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.Modify is not implemented"))
}

func (UnimplementedUserHandler) Cash(context.Context, *connect.Request[request.UserCash]) (*connect.Response[response.UserCash], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.User.Cash is not implemented"))
}
