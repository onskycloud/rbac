// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: aws-sns/message.proto

package aws_sns

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
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

// Client API for AWSSNS service

type AWSSNSService interface {
	Publish(ctx context.Context, in *PublishInput, opts ...client.CallOption) (*PublishOutput, error)
	Endpoint(ctx context.Context, in *CreateEndpointRequest, opts ...client.CallOption) (*CreateEndpointResponse, error)
	PublishAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type aWSSNSService struct {
	c    client.Client
	name string
}

func NewAWSSNSService(name string, c client.Client) AWSSNSService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "aws_sns"
	}
	return &aWSSNSService{
		c:    c,
		name: name,
	}
}

func (c *aWSSNSService) Publish(ctx context.Context, in *PublishInput, opts ...client.CallOption) (*PublishOutput, error) {
	req := c.c.NewRequest(c.name, "AWSSNS.Publish", in)
	out := new(PublishOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSSNSService) Endpoint(ctx context.Context, in *CreateEndpointRequest, opts ...client.CallOption) (*CreateEndpointResponse, error) {
	req := c.c.NewRequest(c.name, "AWSSNS.Endpoint", in)
	out := new(CreateEndpointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSSNSService) PublishAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "AWSSNS.PublishAPI", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AWSSNS service

type AWSSNSHandler interface {
	Publish(context.Context, *PublishInput, *PublishOutput) error
	Endpoint(context.Context, *CreateEndpointRequest, *CreateEndpointResponse) error
	PublishAPI(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterAWSSNSHandler(s server.Server, hdlr AWSSNSHandler, opts ...server.HandlerOption) error {
	type aWSSNS interface {
		Publish(ctx context.Context, in *PublishInput, out *PublishOutput) error
		Endpoint(ctx context.Context, in *CreateEndpointRequest, out *CreateEndpointResponse) error
		PublishAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type AWSSNS struct {
		aWSSNS
	}
	h := &aWSSNSHandler{hdlr}
	return s.Handle(s.NewHandler(&AWSSNS{h}, opts...))
}

type aWSSNSHandler struct {
	AWSSNSHandler
}

func (h *aWSSNSHandler) Publish(ctx context.Context, in *PublishInput, out *PublishOutput) error {
	return h.AWSSNSHandler.Publish(ctx, in, out)
}

func (h *aWSSNSHandler) Endpoint(ctx context.Context, in *CreateEndpointRequest, out *CreateEndpointResponse) error {
	return h.AWSSNSHandler.Endpoint(ctx, in, out)
}

func (h *aWSSNSHandler) PublishAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.AWSSNSHandler.PublishAPI(ctx, in, out)
}
