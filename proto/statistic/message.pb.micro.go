// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: statistic/message.proto

package statistic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SystemSvc service

func NewSystemSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SystemSvc service

type SystemSvcService interface {
	GetCallingActivating(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	GetCountAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	GetMapDirectionAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	SendObserverEmailAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	GetChartData(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	HistoryBySerial(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	Offline(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	HeartRateStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	SnoringStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	SleepStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	SleepMonthStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	SleepWeekStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	BedSensorLog(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	TimeLine(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type systemSvcService struct {
	c    client.Client
	name string
}

func NewSystemSvcService(name string, c client.Client) SystemSvcService {
	return &systemSvcService{
		c:    c,
		name: name,
	}
}

func (c *systemSvcService) GetCallingActivating(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetCallingActivating", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) GetCountAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetCountAPI", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) GetMapDirectionAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetMapDirectionAPI", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) SendObserverEmailAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.SendObserverEmailAPI", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) GetChartData(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetChartData", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) HistoryBySerial(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.HistoryBySerial", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) Offline(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.Offline", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) HeartRateStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.HeartRateStatistic", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) SnoringStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.SnoringStatistic", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) SleepStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.SleepStatistic", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) SleepMonthStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.SleepMonthStatistic", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) SleepWeekStatistic(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.SleepWeekStatistic", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) BedSensorLog(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.BedSensorLog", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) TimeLine(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.TimeLine", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SystemSvc service

type SystemSvcHandler interface {
	GetCallingActivating(context.Context, *proto1.Request, *proto1.Response) error
	GetCountAPI(context.Context, *proto1.Request, *proto1.Response) error
	GetMapDirectionAPI(context.Context, *proto1.Request, *proto1.Response) error
	SendObserverEmailAPI(context.Context, *proto1.Request, *proto1.Response) error
	GetChartData(context.Context, *proto1.Request, *proto1.Response) error
	HistoryBySerial(context.Context, *proto1.Request, *proto1.Response) error
	Offline(context.Context, *proto1.Request, *proto1.Response) error
	HeartRateStatistic(context.Context, *proto1.Request, *proto1.Response) error
	SnoringStatistic(context.Context, *proto1.Request, *proto1.Response) error
	SleepStatistic(context.Context, *proto1.Request, *proto1.Response) error
	SleepMonthStatistic(context.Context, *proto1.Request, *proto1.Response) error
	SleepWeekStatistic(context.Context, *proto1.Request, *proto1.Response) error
	BedSensorLog(context.Context, *proto1.Request, *proto1.Response) error
	TimeLine(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterSystemSvcHandler(s server.Server, hdlr SystemSvcHandler, opts ...server.HandlerOption) error {
	type systemSvc interface {
		GetCallingActivating(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		GetCountAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		GetMapDirectionAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		SendObserverEmailAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		GetChartData(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		HistoryBySerial(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		Offline(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		HeartRateStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		SnoringStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		SleepStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		SleepMonthStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		SleepWeekStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		BedSensorLog(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		TimeLine(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type SystemSvc struct {
		systemSvc
	}
	h := &systemSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&SystemSvc{h}, opts...))
}

type systemSvcHandler struct {
	SystemSvcHandler
}

func (h *systemSvcHandler) GetCallingActivating(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.GetCallingActivating(ctx, in, out)
}

func (h *systemSvcHandler) GetCountAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.GetCountAPI(ctx, in, out)
}

func (h *systemSvcHandler) GetMapDirectionAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.GetMapDirectionAPI(ctx, in, out)
}

func (h *systemSvcHandler) SendObserverEmailAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.SendObserverEmailAPI(ctx, in, out)
}

func (h *systemSvcHandler) GetChartData(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.GetChartData(ctx, in, out)
}

func (h *systemSvcHandler) HistoryBySerial(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.HistoryBySerial(ctx, in, out)
}

func (h *systemSvcHandler) Offline(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.Offline(ctx, in, out)
}

func (h *systemSvcHandler) HeartRateStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.HeartRateStatistic(ctx, in, out)
}

func (h *systemSvcHandler) SnoringStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.SnoringStatistic(ctx, in, out)
}

func (h *systemSvcHandler) SleepStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.SleepStatistic(ctx, in, out)
}

func (h *systemSvcHandler) SleepMonthStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.SleepMonthStatistic(ctx, in, out)
}

func (h *systemSvcHandler) SleepWeekStatistic(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.SleepWeekStatistic(ctx, in, out)
}

func (h *systemSvcHandler) BedSensorLog(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.BedSensorLog(ctx, in, out)
}

func (h *systemSvcHandler) TimeLine(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SystemSvcHandler.TimeLine(ctx, in, out)
}
