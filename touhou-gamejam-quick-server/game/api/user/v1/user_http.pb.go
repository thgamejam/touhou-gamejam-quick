// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.PUT("/api/v1/user", _User_CreateUser0_HTTP_Handler(srv))
	r.GET("/api/v1/user/{id}", _User_GetUser0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in.Body); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/CreateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/api/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.User/CreateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Body, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.v1.User/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
