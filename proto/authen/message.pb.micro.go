// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authen/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthenSvc service

type AuthenSvcService interface {
	GetCustomer(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	GetCustomerByEmail(ctx context.Context, in *CustomerByUserNameRequest, opts ...client.CallOption) (*CustomerResponse, error)
	VerifyPermission(ctx context.Context, in *VerifyPermissionRequest, opts ...client.CallOption) (*VerifyPermissionResponse, error)
	CountCustomer(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*wrappers.Int64Value, error)
}

type authenSvcService struct {
	c    client.Client
	name string
}

func NewAuthenSvcService(name string, c client.Client) AuthenSvcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "authensvc"
	}
	return &authenSvcService{
		c:    c,
		name: name,
	}
}

func (c *authenSvcService) GetCustomer(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthenSvc.GetCustomer", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenSvcService) GetCustomerByEmail(ctx context.Context, in *CustomerByUserNameRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthenSvc.GetCustomerByEmail", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenSvcService) VerifyPermission(ctx context.Context, in *VerifyPermissionRequest, opts ...client.CallOption) (*VerifyPermissionResponse, error) {
	req := c.c.NewRequest(c.name, "AuthenSvc.VerifyPermission", in)
	out := new(VerifyPermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenSvcService) CountCustomer(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*wrappers.Int64Value, error) {
	req := c.c.NewRequest(c.name, "AuthenSvc.CountCustomer", in)
	out := new(wrappers.Int64Value)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthenSvc service

type AuthenSvcHandler interface {
	GetCustomer(context.Context, *CustomerRequest, *CustomerResponse) error
	GetCustomerByEmail(context.Context, *CustomerByUserNameRequest, *CustomerResponse) error
	VerifyPermission(context.Context, *VerifyPermissionRequest, *VerifyPermissionResponse) error
	CountCustomer(context.Context, *empty.Empty, *wrappers.Int64Value) error
}

func RegisterAuthenSvcHandler(s server.Server, hdlr AuthenSvcHandler, opts ...server.HandlerOption) error {
	type authenSvc interface {
		GetCustomer(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error
		GetCustomerByEmail(ctx context.Context, in *CustomerByUserNameRequest, out *CustomerResponse) error
		VerifyPermission(ctx context.Context, in *VerifyPermissionRequest, out *VerifyPermissionResponse) error
		CountCustomer(ctx context.Context, in *empty.Empty, out *wrappers.Int64Value) error
	}
	type AuthenSvc struct {
		authenSvc
	}
	h := &authenSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthenSvc{h}, opts...))
}

type authenSvcHandler struct {
	AuthenSvcHandler
}

func (h *authenSvcHandler) GetCustomer(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.AuthenSvcHandler.GetCustomer(ctx, in, out)
}

func (h *authenSvcHandler) GetCustomerByEmail(ctx context.Context, in *CustomerByUserNameRequest, out *CustomerResponse) error {
	return h.AuthenSvcHandler.GetCustomerByEmail(ctx, in, out)
}

func (h *authenSvcHandler) VerifyPermission(ctx context.Context, in *VerifyPermissionRequest, out *VerifyPermissionResponse) error {
	return h.AuthenSvcHandler.VerifyPermission(ctx, in, out)
}

func (h *authenSvcHandler) CountCustomer(ctx context.Context, in *empty.Empty, out *wrappers.Int64Value) error {
	return h.AuthenSvcHandler.CountCustomer(ctx, in, out)
}
