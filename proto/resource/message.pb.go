// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource/message.proto

package message

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ServiceByShortNameRequest struct {
	ShortName            string   `protobuf:"bytes,1,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceByShortNameRequest) Reset()         { *m = ServiceByShortNameRequest{} }
func (m *ServiceByShortNameRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceByShortNameRequest) ProtoMessage()    {}
func (*ServiceByShortNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3356f27430b4dff7, []int{0}
}

func (m *ServiceByShortNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceByShortNameRequest.Unmarshal(m, b)
}
func (m *ServiceByShortNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceByShortNameRequest.Marshal(b, m, deterministic)
}
func (m *ServiceByShortNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceByShortNameRequest.Merge(m, src)
}
func (m *ServiceByShortNameRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceByShortNameRequest.Size(m)
}
func (m *ServiceByShortNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceByShortNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceByShortNameRequest proto.InternalMessageInfo

func (m *ServiceByShortNameRequest) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

type RegisterInfoRequest struct {
	Name                 string          `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ShortName            string          `protobuf:"bytes,2,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Description          string          `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	IsCharge             bool            `protobuf:"varint,4,opt,name=IsCharge,proto3" json:"IsCharge,omitempty"`
	Data                 []*ResourceType `protobuf:"bytes,5,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterInfoRequest) Reset()         { *m = RegisterInfoRequest{} }
func (m *RegisterInfoRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterInfoRequest) ProtoMessage()    {}
func (*RegisterInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3356f27430b4dff7, []int{1}
}

func (m *RegisterInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterInfoRequest.Unmarshal(m, b)
}
func (m *RegisterInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterInfoRequest.Marshal(b, m, deterministic)
}
func (m *RegisterInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterInfoRequest.Merge(m, src)
}
func (m *RegisterInfoRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterInfoRequest.Size(m)
}
func (m *RegisterInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterInfoRequest proto.InternalMessageInfo

func (m *RegisterInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterInfoRequest) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *RegisterInfoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterInfoRequest) GetIsCharge() bool {
	if m != nil {
		return m.IsCharge
	}
	return false
}

func (m *RegisterInfoRequest) GetData() []*ResourceType {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResourceType struct {
	Name                 string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string    `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Data                 []*Action `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceType) Reset()         { *m = ResourceType{} }
func (m *ResourceType) String() string { return proto.CompactTextString(m) }
func (*ResourceType) ProtoMessage()    {}
func (*ResourceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3356f27430b4dff7, []int{2}
}

func (m *ResourceType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceType.Unmarshal(m, b)
}
func (m *ResourceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceType.Marshal(b, m, deterministic)
}
func (m *ResourceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceType.Merge(m, src)
}
func (m *ResourceType) XXX_Size() int {
	return xxx_messageInfo_ResourceType.Size(m)
}
func (m *ResourceType) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceType.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceType proto.InternalMessageInfo

func (m *ResourceType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ResourceType) GetData() []*Action {
	if m != nil {
		return m.Data
	}
	return nil
}

type Action struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_3356f27430b4dff7, []int{3}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Action) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ServiceInfoResponse struct {
	Id                   int32           `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CreatedAt            string          `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string          `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Name                 string          `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	ShortName            string          `protobuf:"bytes,5,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Version              string          `protobuf:"bytes,6,opt,name=Version,proto3" json:"Version,omitempty"`
	Description          string          `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	IsActive             bool            `protobuf:"varint,8,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	IsCharge             bool            `protobuf:"varint,9,opt,name=IsCharge,proto3" json:"IsCharge,omitempty"`
	Data                 []*ResourceType `protobuf:"bytes,10,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ServiceInfoResponse) Reset()         { *m = ServiceInfoResponse{} }
func (m *ServiceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceInfoResponse) ProtoMessage()    {}
func (*ServiceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3356f27430b4dff7, []int{4}
}

func (m *ServiceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfoResponse.Unmarshal(m, b)
}
func (m *ServiceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfoResponse.Marshal(b, m, deterministic)
}
func (m *ServiceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfoResponse.Merge(m, src)
}
func (m *ServiceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceInfoResponse.Size(m)
}
func (m *ServiceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfoResponse proto.InternalMessageInfo

func (m *ServiceInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServiceInfoResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ServiceInfoResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ServiceInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceInfoResponse) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *ServiceInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ServiceInfoResponse) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *ServiceInfoResponse) GetIsCharge() bool {
	if m != nil {
		return m.IsCharge
	}
	return false
}

func (m *ServiceInfoResponse) GetData() []*ResourceType {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceByShortNameRequest)(nil), "ServiceByShortNameRequest")
	proto.RegisterType((*RegisterInfoRequest)(nil), "RegisterInfoRequest")
	proto.RegisterType((*ResourceType)(nil), "ResourceType")
	proto.RegisterType((*Action)(nil), "Action")
	proto.RegisterType((*ServiceInfoResponse)(nil), "ServiceInfoResponse")
}

func init() { proto.RegisterFile("resource/message.proto", fileDescriptor_3356f27430b4dff7) }

var fileDescriptor_3356f27430b4dff7 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x1d, 0x3b, 0x71, 0x26, 0x14, 0xa4, 0x6d, 0x00, 0x13, 0x7a, 0x08, 0x96, 0x2a, 0xb5,
	0xaa, 0xe4, 0x4a, 0xe5, 0x8c, 0x50, 0xdd, 0x4a, 0x28, 0xe2, 0x47, 0x68, 0xc3, 0x8f, 0xc4, 0x6d,
	0xb1, 0x87, 0xd4, 0x12, 0xb1, 0xcd, 0xee, 0x36, 0x52, 0x9f, 0x81, 0x1b, 0x8f, 0xc1, 0xbb, 0xf0,
	0x02, 0x3c, 0x4a, 0x4e, 0xd8, 0xbb, 0xeb, 0x26, 0xce, 0xdf, 0x81, 0x9b, 0x77, 0xbe, 0xf9, 0xfb,
	0x66, 0xbe, 0x31, 0x3c, 0xe2, 0x28, 0xf2, 0x1b, 0x1e, 0xe3, 0xd9, 0x14, 0x85, 0x60, 0x13, 0x0c,
	0x0b, 0x9e, 0xcb, 0x7c, 0xf0, 0x78, 0xc6, 0xbe, 0xa7, 0x09, 0x93, 0x78, 0x56, 0x7f, 0x68, 0x20,
	0x88, 0xe0, 0xc9, 0x18, 0xf9, 0x2c, 0x8d, 0x31, 0xba, 0x1d, 0x5f, 0xe7, 0x5c, 0xbe, 0x63, 0x53,
	0xa4, 0xf8, 0xe3, 0x06, 0x85, 0x24, 0x47, 0xd0, 0xbd, 0xb3, 0xf9, 0xd6, 0xd0, 0x3a, 0xee, 0x46,
	0x9d, 0x79, 0xe4, 0x70, 0xfb, 0xcb, 0x1e, 0x5d, 0x20, 0xc1, 0x1f, 0x0b, 0x0e, 0x28, 0x4e, 0x52,
	0x21, 0x91, 0x8f, 0xb2, 0x6f, 0x79, 0x1d, 0xfe, 0x14, 0x9c, 0x4d, 0x91, 0xca, 0xd8, 0xcc, 0x6d,
	0x6f, 0xcb, 0x4d, 0x4e, 0xa0, 0x77, 0x85, 0x22, 0xe6, 0x69, 0x21, 0xd3, 0x3c, 0xf3, 0x5b, 0x4d,
	0xc7, 0x65, 0x8c, 0x0c, 0xc0, 0x1b, 0x89, 0xcb, 0x6b, 0xc6, 0x27, 0xe8, 0x3b, 0xa5, 0x9f, 0x47,
	0xef, 0xde, 0xe4, 0x14, 0x9c, 0x2b, 0x26, 0x99, 0xef, 0x0e, 0x5b, 0xc7, 0xbd, 0xf3, 0xfd, 0x90,
	0x9a, 0x31, 0x7d, 0xb8, 0x2d, 0x30, 0xf2, 0xe6, 0x91, 0xfb, 0xcb, 0xb2, 0x3d, 0x8b, 0x2a, 0xa7,
	0x40, 0xc2, 0xbd, 0x65, 0x7c, 0x37, 0x8f, 0x61, 0xb3, 0x41, 0xc5, 0xa4, 0xd9, 0xd7, 0x91, 0xa9,
	0xdd, 0x52, 0xb5, 0x3b, 0xe1, 0x45, 0x5c, 0x99, 0xd7, 0xaa, 0xfe, 0xb4, 0xa0, 0xad, 0xa1, 0xdd,
	0x05, 0x23, 0x70, 0xaa, 0xae, 0xcc, 0xcc, 0xc2, 0x79, 0x74, 0xca, 0x4f, 0xa8, 0x43, 0x91, 0x25,
	0xd4, 0xfd, 0xcc, 0x53, 0x89, 0xd4, 0x79, 0x53, 0xae, 0x82, 0xf6, 0xdf, 0x23, 0x9f, 0xa6, 0x42,
	0x94, 0x09, 0xdf, 0xb2, 0xac, 0x94, 0xc2, 0x14, 0x33, 0x49, 0x55, 0xec, 0x6a, 0xd3, 0xad, 0xb5,
	0xa6, 0x83, 0xdf, 0x36, 0x1c, 0x18, 0x61, 0xe8, 0x95, 0x8a, 0x22, 0xcf, 0x04, 0x92, 0xfb, 0x60,
	0x8f, 0x12, 0xd5, 0x98, 0x4b, 0xcb, 0x2f, 0x72, 0x08, 0xdd, 0x4b, 0x8e, 0xa5, 0x9e, 0x92, 0x0b,
	0x69, 0xc8, 0x2f, 0x0c, 0x15, 0xfa, 0xb1, 0x48, 0x0c, 0xaa, 0xab, 0x2c, 0x0c, 0x84, 0x18, 0x9a,
	0x8e, 0x02, 0x34, 0xbb, 0xc3, 0x65, 0x59, 0xb8, 0x3a, 0x62, 0xa1, 0x06, 0x1f, 0x3a, 0x9f, 0x90,
	0x57, 0x94, 0xfc, 0xb6, 0xc2, 0xea, 0xe7, 0x2a, 0xa3, 0xce, 0xfa, 0x1a, 0x94, 0x3c, 0xaa, 0x01,
	0xcf, 0xd0, 0xf7, 0x6a, 0x79, 0xe8, 0x77, 0x43, 0x3a, 0xdd, 0x15, 0xe9, 0x3c, 0x33, 0xeb, 0x83,
	0x0d, 0xd2, 0xd1, 0xab, 0x3b, 0xff, 0x6b, 0x41, 0xaf, 0x36, 0x8f, 0x67, 0x31, 0x79, 0x09, 0x0f,
	0xea, 0x7b, 0x30, 0x33, 0x24, 0xfd, 0x70, 0xc3, 0x85, 0x0c, 0xfa, 0xe1, 0x86, 0x19, 0x07, 0x7b,
	0xe4, 0x05, 0xec, 0xeb, 0x31, 0xfd, 0x5f, 0xf8, 0x6b, 0x78, 0xf8, 0x0a, 0xe5, 0xfa, 0x5d, 0x93,
	0x41, 0xb8, 0xf5, 0xd8, 0xb7, 0x25, 0xfb, 0xda, 0x56, 0x3f, 0x8a, 0xe7, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0x77, 0x53, 0x36, 0x5b, 0x04, 0x00, 0x00,
}
