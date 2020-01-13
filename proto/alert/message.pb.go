// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alert/message.proto

package message

import (
	fmt "fmt"
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

type AlertRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	CompareValue         string   `protobuf:"bytes,2,opt,name=compareValue,proto3" json:"compareValue,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertRequest) Reset()         { *m = AlertRequest{} }
func (m *AlertRequest) String() string { return proto.CompactTextString(m) }
func (*AlertRequest) ProtoMessage()    {}
func (*AlertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa71fe8c44cc448, []int{0}
}

func (m *AlertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertRequest.Unmarshal(m, b)
}
func (m *AlertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertRequest.Marshal(b, m, deterministic)
}
func (m *AlertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertRequest.Merge(m, src)
}
func (m *AlertRequest) XXX_Size() int {
	return xxx_messageInfo_AlertRequest.Size(m)
}
func (m *AlertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertRequest proto.InternalMessageInfo

func (m *AlertRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlertRequest) GetCompareValue() string {
	if m != nil {
		return m.CompareValue
	}
	return ""
}

func (m *AlertRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AlertResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertResponse) Reset()         { *m = AlertResponse{} }
func (m *AlertResponse) String() string { return proto.CompactTextString(m) }
func (*AlertResponse) ProtoMessage()    {}
func (*AlertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa71fe8c44cc448, []int{1}
}

func (m *AlertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertResponse.Unmarshal(m, b)
}
func (m *AlertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertResponse.Marshal(b, m, deterministic)
}
func (m *AlertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertResponse.Merge(m, src)
}
func (m *AlertResponse) XXX_Size() int {
	return xxx_messageInfo_AlertResponse.Size(m)
}
func (m *AlertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlertResponse proto.InternalMessageInfo

func (m *AlertResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

//--
type AlertTypeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertTypeRequest) Reset()         { *m = AlertTypeRequest{} }
func (m *AlertTypeRequest) String() string { return proto.CompactTextString(m) }
func (*AlertTypeRequest) ProtoMessage()    {}
func (*AlertTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa71fe8c44cc448, []int{2}
}

func (m *AlertTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertTypeRequest.Unmarshal(m, b)
}
func (m *AlertTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertTypeRequest.Marshal(b, m, deterministic)
}
func (m *AlertTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertTypeRequest.Merge(m, src)
}
func (m *AlertTypeRequest) XXX_Size() int {
	return xxx_messageInfo_AlertTypeRequest.Size(m)
}
func (m *AlertTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertTypeRequest proto.InternalMessageInfo

type AlertTypeResponse struct {
	Items                []*AlertType `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AlertTypeResponse) Reset()         { *m = AlertTypeResponse{} }
func (m *AlertTypeResponse) String() string { return proto.CompactTextString(m) }
func (*AlertTypeResponse) ProtoMessage()    {}
func (*AlertTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa71fe8c44cc448, []int{3}
}

func (m *AlertTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertTypeResponse.Unmarshal(m, b)
}
func (m *AlertTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertTypeResponse.Marshal(b, m, deterministic)
}
func (m *AlertTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertTypeResponse.Merge(m, src)
}
func (m *AlertTypeResponse) XXX_Size() int {
	return xxx_messageInfo_AlertTypeResponse.Size(m)
}
func (m *AlertTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlertTypeResponse proto.InternalMessageInfo

func (m *AlertTypeResponse) GetItems() []*AlertType {
	if m != nil {
		return m.Items
	}
	return nil
}

//--
type AlertType struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertType) Reset()         { *m = AlertType{} }
func (m *AlertType) String() string { return proto.CompactTextString(m) }
func (*AlertType) ProtoMessage()    {}
func (*AlertType) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa71fe8c44cc448, []int{4}
}

func (m *AlertType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertType.Unmarshal(m, b)
}
func (m *AlertType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertType.Marshal(b, m, deterministic)
}
func (m *AlertType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertType.Merge(m, src)
}
func (m *AlertType) XXX_Size() int {
	return xxx_messageInfo_AlertType.Size(m)
}
func (m *AlertType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertType.DiscardUnknown(m)
}

var xxx_messageInfo_AlertType proto.InternalMessageInfo

func (m *AlertType) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AlertType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AlertRequest)(nil), "AlertRequest")
	proto.RegisterType((*AlertResponse)(nil), "AlertResponse")
	proto.RegisterType((*AlertTypeRequest)(nil), "AlertTypeRequest")
	proto.RegisterType((*AlertTypeResponse)(nil), "AlertTypeResponse")
	proto.RegisterType((*AlertType)(nil), "AlertType")
}

func init() { proto.RegisterFile("alert/message.proto", fileDescriptor_caa71fe8c44cc448) }

var fileDescriptor_caa71fe8c44cc448 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x50, 0x4d, 0x4f, 0xc2, 0x50,
	0x10, 0xe4, 0xc3, 0x12, 0x18, 0x29, 0x91, 0xd5, 0x18, 0xc2, 0x89, 0xbc, 0x8b, 0x5c, 0x7c, 0x24,
	0x18, 0x13, 0xaf, 0xc6, 0x83, 0xf7, 0x6a, 0x3c, 0x79, 0x79, 0xe2, 0x46, 0x89, 0x2d, 0xad, 0x6f,
	0x5f, 0x49, 0xfc, 0xf7, 0x96, 0xa5, 0x22, 0xe8, 0x6d, 0x76, 0x66, 0x77, 0x67, 0x67, 0x71, 0xea,
	0x52, 0xf6, 0x61, 0x96, 0xb1, 0x88, 0x7b, 0x63, 0x5b, 0xf8, 0x3c, 0xe4, 0xe6, 0x19, 0xfd, 0xdb,
	0x0d, 0x9d, 0xf0, 0x67, 0xc9, 0x12, 0x88, 0x70, 0x14, 0xbe, 0x0a, 0x1e, 0x35, 0x27, 0xcd, 0x69,
	0x2f, 0x51, 0x4c, 0x06, 0xfd, 0x45, 0x9e, 0x15, 0xce, 0xf3, 0x93, 0x4b, 0x4b, 0x1e, 0xb5, 0x54,
	0x3b, 0xe0, 0xe8, 0x0c, 0xd1, 0x5a, 0xc5, 0xb6, 0x8a, 0xdb, 0xc2, 0x5c, 0x20, 0xae, 0xb7, 0x4b,
	0x91, 0xaf, 0x84, 0xe9, 0x1c, 0x1d, 0xcf, 0x52, 0xa6, 0x41, 0x0d, 0xba, 0x49, 0x5d, 0x19, 0xc2,
	0x89, 0x36, 0x3e, 0x56, 0x7e, 0xf5, 0x29, 0xe6, 0x1a, 0xc3, 0x3d, 0xae, 0x5e, 0x30, 0x41, 0xb4,
	0x0c, 0x9c, 0x49, 0x35, 0xdf, 0x9e, 0x1e, 0xcf, 0x61, 0x7f, 0x5b, 0xb6, 0x82, 0x99, 0xa1, 0xb7,
	0xe3, 0x68, 0x80, 0xd6, 0xf2, 0x55, 0xbd, 0xa2, 0xa4, 0x42, 0x9b, 0x78, 0x2b, 0x97, 0xfd, 0x44,
	0x50, 0x3c, 0x17, 0x74, 0x75, 0xe0, 0xc1, 0xaf, 0xe9, 0x12, 0xb8, 0x7b, 0xe7, 0xc5, 0x87, 0x12,
	0x14, 0xdb, 0xfd, 0xdf, 0x8c, 0x07, 0xf6, 0x20, 0x8c, 0x69, 0xd0, 0x0d, 0xe2, 0x7b, 0x0e, 0x3b,
	0x3b, 0xa1, 0xa1, 0xfd, 0x1b, 0x63, 0x4c, 0xf6, 0x5f, 0x0a, 0xd3, 0x78, 0xe9, 0xe8, 0xfb, 0xaf,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x41, 0x83, 0x4a, 0x95, 0x01, 0x00, 0x00,
}
