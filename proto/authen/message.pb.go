// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authen/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type CustomerRequest struct {
	CustomerNumber       string   `protobuf:"bytes,1,opt,name=customerNumber,proto3" json:"customerNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerRequest) Reset()         { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()    {}
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{1}
}

func (m *CustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerRequest.Unmarshal(m, b)
}
func (m *CustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerRequest.Marshal(b, m, deterministic)
}
func (m *CustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerRequest.Merge(m, src)
}
func (m *CustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CustomerRequest.Size(m)
}
func (m *CustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerRequest proto.InternalMessageInfo

func (m *CustomerRequest) GetCustomerNumber() string {
	if m != nil {
		return m.CustomerNumber
	}
	return ""
}

type CustomerByUserNameRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerByUserNameRequest) Reset()         { *m = CustomerByUserNameRequest{} }
func (m *CustomerByUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerByUserNameRequest) ProtoMessage()    {}
func (*CustomerByUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{2}
}

func (m *CustomerByUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerByUserNameRequest.Unmarshal(m, b)
}
func (m *CustomerByUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerByUserNameRequest.Marshal(b, m, deterministic)
}
func (m *CustomerByUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerByUserNameRequest.Merge(m, src)
}
func (m *CustomerByUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_CustomerByUserNameRequest.Size(m)
}
func (m *CustomerByUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerByUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerByUserNameRequest proto.InternalMessageInfo

func (m *CustomerByUserNameRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type CustomerResponse struct {
	CustomerNumber       string           `protobuf:"bytes,1,opt,name=customerNumber,proto3" json:"customerNumber,omitempty"`
	Alias                string           `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Mobile               string           `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	FirstName            string           `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string           `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Address1             string           `protobuf:"bytes,6,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2             string           `protobuf:"bytes,7,opt,name=address2,proto3" json:"address2,omitempty"`
	Email                string           `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Setting              *CustomerSetting `protobuf:"bytes,9,opt,name=setting,proto3" json:"setting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomerResponse) Reset()         { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()    {}
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{3}
}

func (m *CustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerResponse.Unmarshal(m, b)
}
func (m *CustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerResponse.Marshal(b, m, deterministic)
}
func (m *CustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerResponse.Merge(m, src)
}
func (m *CustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CustomerResponse.Size(m)
}
func (m *CustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerResponse proto.InternalMessageInfo

func (m *CustomerResponse) GetCustomerNumber() string {
	if m != nil {
		return m.CustomerNumber
	}
	return ""
}

func (m *CustomerResponse) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *CustomerResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *CustomerResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CustomerResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CustomerResponse) GetAddress1() string {
	if m != nil {
		return m.Address1
	}
	return ""
}

func (m *CustomerResponse) GetAddress2() string {
	if m != nil {
		return m.Address2
	}
	return ""
}

func (m *CustomerResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerResponse) GetSetting() *CustomerSetting {
	if m != nil {
		return m.Setting
	}
	return nil
}

type CustomerSetting struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsAllowRegister      bool           `protobuf:"varint,2,opt,name=isAllowRegister,proto3" json:"isAllowRegister,omitempty"`
	Smtp                 *SMTPConfig    `protobuf:"bytes,3,opt,name=smtp,proto3" json:"smtp,omitempty"`
	Payment              *PaymentConfig `protobuf:"bytes,4,opt,name=payment,proto3" json:"payment,omitempty"`
	Other                string         `protobuf:"bytes,5,opt,name=other,proto3" json:"other,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CustomerSetting) Reset()         { *m = CustomerSetting{} }
func (m *CustomerSetting) String() string { return proto.CompactTextString(m) }
func (*CustomerSetting) ProtoMessage()    {}
func (*CustomerSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{4}
}

func (m *CustomerSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerSetting.Unmarshal(m, b)
}
func (m *CustomerSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerSetting.Marshal(b, m, deterministic)
}
func (m *CustomerSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerSetting.Merge(m, src)
}
func (m *CustomerSetting) XXX_Size() int {
	return xxx_messageInfo_CustomerSetting.Size(m)
}
func (m *CustomerSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerSetting proto.InternalMessageInfo

func (m *CustomerSetting) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerSetting) GetIsAllowRegister() bool {
	if m != nil {
		return m.IsAllowRegister
	}
	return false
}

func (m *CustomerSetting) GetSmtp() *SMTPConfig {
	if m != nil {
		return m.Smtp
	}
	return nil
}

func (m *CustomerSetting) GetPayment() *PaymentConfig {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *CustomerSetting) GetOther() string {
	if m != nil {
		return m.Other
	}
	return ""
}

type SMTPConfig struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Account              string   `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMTPConfig) Reset()         { *m = SMTPConfig{} }
func (m *SMTPConfig) String() string { return proto.CompactTextString(m) }
func (*SMTPConfig) ProtoMessage()    {}
func (*SMTPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{5}
}

func (m *SMTPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMTPConfig.Unmarshal(m, b)
}
func (m *SMTPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMTPConfig.Marshal(b, m, deterministic)
}
func (m *SMTPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMTPConfig.Merge(m, src)
}
func (m *SMTPConfig) XXX_Size() int {
	return xxx_messageInfo_SMTPConfig.Size(m)
}
func (m *SMTPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SMTPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SMTPConfig proto.InternalMessageInfo

func (m *SMTPConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SMTPConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SMTPConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SMTPConfig) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type PaymentConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentConfig) Reset()         { *m = PaymentConfig{} }
func (m *PaymentConfig) String() string { return proto.CompactTextString(m) }
func (*PaymentConfig) ProtoMessage()    {}
func (*PaymentConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{6}
}

func (m *PaymentConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentConfig.Unmarshal(m, b)
}
func (m *PaymentConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentConfig.Marshal(b, m, deterministic)
}
func (m *PaymentConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentConfig.Merge(m, src)
}
func (m *PaymentConfig) XXX_Size() int {
	return xxx_messageInfo_PaymentConfig.Size(m)
}
func (m *PaymentConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentConfig proto.InternalMessageInfo

type VerifyPermissionRequest struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	CompareResources     []string `protobuf:"bytes,2,rep,name=compareResources,proto3" json:"compareResources,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyPermissionRequest) Reset()         { *m = VerifyPermissionRequest{} }
func (m *VerifyPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyPermissionRequest) ProtoMessage()    {}
func (*VerifyPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{7}
}

func (m *VerifyPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyPermissionRequest.Unmarshal(m, b)
}
func (m *VerifyPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyPermissionRequest.Marshal(b, m, deterministic)
}
func (m *VerifyPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyPermissionRequest.Merge(m, src)
}
func (m *VerifyPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyPermissionRequest.Size(m)
}
func (m *VerifyPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyPermissionRequest proto.InternalMessageInfo

func (m *VerifyPermissionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *VerifyPermissionRequest) GetCompareResources() []string {
	if m != nil {
		return m.CompareResources
	}
	return nil
}

type VerifyPermissionResponse struct {
	Resources            []string  `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	Success              bool      `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	AuthUser             *AuthUser `protobuf:"bytes,3,opt,name=authUser,proto3" json:"authUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VerifyPermissionResponse) Reset()         { *m = VerifyPermissionResponse{} }
func (m *VerifyPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyPermissionResponse) ProtoMessage()    {}
func (*VerifyPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{8}
}

func (m *VerifyPermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyPermissionResponse.Unmarshal(m, b)
}
func (m *VerifyPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyPermissionResponse.Marshal(b, m, deterministic)
}
func (m *VerifyPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyPermissionResponse.Merge(m, src)
}
func (m *VerifyPermissionResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyPermissionResponse.Size(m)
}
func (m *VerifyPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyPermissionResponse proto.InternalMessageInfo

func (m *VerifyPermissionResponse) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *VerifyPermissionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *VerifyPermissionResponse) GetAuthUser() *AuthUser {
	if m != nil {
		return m.AuthUser
	}
	return nil
}

type AuthUser struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerID           int32    `protobuf:"varint,2,opt,name=customerID,proto3" json:"customerID,omitempty"`
	CustomerNumber       string   `protobuf:"bytes,3,opt,name=customerNumber,proto3" json:"customerNumber,omitempty"`
	Clients              []string `protobuf:"bytes,4,rep,name=clients,proto3" json:"clients,omitempty"`
	Username             string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	UserUuid             string   `protobuf:"bytes,6,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUser) Reset()         { *m = AuthUser{} }
func (m *AuthUser) String() string { return proto.CompactTextString(m) }
func (*AuthUser) ProtoMessage()    {}
func (*AuthUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_48722578ab44b1f8, []int{9}
}

func (m *AuthUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUser.Unmarshal(m, b)
}
func (m *AuthUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUser.Marshal(b, m, deterministic)
}
func (m *AuthUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUser.Merge(m, src)
}
func (m *AuthUser) XXX_Size() int {
	return xxx_messageInfo_AuthUser.Size(m)
}
func (m *AuthUser) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUser.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUser proto.InternalMessageInfo

func (m *AuthUser) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuthUser) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *AuthUser) GetCustomerNumber() string {
	if m != nil {
		return m.CustomerNumber
	}
	return ""
}

func (m *AuthUser) GetClients() []string {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *AuthUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthUser) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*CustomerRequest)(nil), "CustomerRequest")
	proto.RegisterType((*CustomerByUserNameRequest)(nil), "CustomerByUserNameRequest")
	proto.RegisterType((*CustomerResponse)(nil), "CustomerResponse")
	proto.RegisterType((*CustomerSetting)(nil), "CustomerSetting")
	proto.RegisterType((*SMTPConfig)(nil), "SMTPConfig")
	proto.RegisterType((*PaymentConfig)(nil), "PaymentConfig")
	proto.RegisterType((*VerifyPermissionRequest)(nil), "VerifyPermissionRequest")
	proto.RegisterType((*VerifyPermissionResponse)(nil), "VerifyPermissionResponse")
	proto.RegisterType((*AuthUser)(nil), "AuthUser")
}

func init() { proto.RegisterFile("authen/message.proto", fileDescriptor_48722578ab44b1f8) }

var fileDescriptor_48722578ab44b1f8 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xa6, 0xdd, 0xfa, 0x77, 0xca, 0xba, 0x62, 0x4d, 0x23, 0xeb, 0xd0, 0x40, 0x91, 0x40, 0x13,
	0x17, 0x99, 0x28, 0x13, 0x88, 0xcb, 0xfd, 0x09, 0xed, 0x02, 0x98, 0x32, 0xb6, 0x3b, 0x2e, 0xdc,
	0xd4, 0xed, 0x8c, 0x92, 0x38, 0xd8, 0xce, 0xa6, 0xbc, 0x0b, 0x8f, 0xc0, 0x25, 0xef, 0xc1, 0x2b,
	0x61, 0x3b, 0x76, 0xd3, 0xb5, 0xab, 0xb4, 0x3b, 0x7f, 0xdf, 0x67, 0x1f, 0x1f, 0x9f, 0xf3, 0x1d,
	0xc3, 0x16, 0xce, 0xe5, 0x0d, 0x49, 0x0f, 0x12, 0x22, 0x04, 0x9e, 0x92, 0x20, 0xe3, 0x4c, 0xb2,
	0xc1, 0xee, 0x94, 0xb1, 0x69, 0x4c, 0x0e, 0x0c, 0x1a, 0xe5, 0x93, 0x03, 0x92, 0x64, 0xb2, 0xb0,
	0xe2, 0xde, 0xa2, 0x78, 0xc7, 0x71, 0x96, 0x11, 0x2e, 0x4a, 0xdd, 0xef, 0xc1, 0xd3, 0x33, 0xbd,
	0x3d, 0x24, 0xbf, 0x72, 0x22, 0xa4, 0xff, 0x09, 0x36, 0x4f, 0x72, 0x21, 0x59, 0x42, 0xb8, 0xa5,
	0xd0, 0x1b, 0xe8, 0x45, 0x96, 0xfa, 0x9a, 0x27, 0x23, 0xc2, 0xbd, 0xda, 0xab, 0xda, 0x7e, 0x27,
	0x5c, 0x60, 0xfd, 0x8f, 0xb0, 0xe3, 0x8e, 0x1e, 0x17, 0x57, 0x42, 0xf1, 0x38, 0x21, 0x2e, 0xc8,
	0x00, 0xda, 0xb9, 0xa5, 0xec, 0xf1, 0x19, 0xf6, 0x7f, 0xd7, 0xa1, 0x5f, 0x5d, 0x2a, 0x32, 0x96,
	0x0a, 0xf2, 0xd8, 0x5b, 0xd1, 0x16, 0x34, 0x70, 0x4c, 0xb1, 0xf0, 0xea, 0x46, 0x2e, 0x01, 0xda,
	0x86, 0x66, 0xc2, 0x46, 0x34, 0x26, 0xde, 0x9a, 0xa1, 0x2d, 0x42, 0x2f, 0xa0, 0x33, 0xa1, 0x5c,
	0x48, 0x93, 0xc7, 0xba, 0x91, 0x2a, 0x42, 0x27, 0x19, 0x63, 0x2b, 0x36, 0xca, 0x24, 0x1d, 0xd6,
	0x1a, 0x1e, 0x8f, 0xb9, 0xaa, 0xfc, 0x3b, 0xaf, 0x59, 0x6a, 0x0e, 0xcf, 0x69, 0x43, 0xaf, 0x75,
	0x4f, 0x1b, 0xea, 0xfc, 0x48, 0x82, 0x69, 0xec, 0xb5, 0xcb, 0xfc, 0x0c, 0x40, 0x6f, 0xa1, 0x25,
	0x88, 0x94, 0x34, 0x9d, 0x7a, 0x1d, 0xc5, 0x77, 0x87, 0xfd, 0xc0, 0x55, 0xe0, 0xb2, 0xe4, 0x43,
	0xb7, 0xc1, 0xff, 0x53, 0xab, 0x7a, 0x62, 0x45, 0xd4, 0x83, 0x3a, 0x1d, 0x9b, 0x8a, 0x34, 0x42,
	0xb5, 0x42, 0xfb, 0xb0, 0x49, 0xc5, 0x51, 0x1c, 0xb3, 0xbb, 0x90, 0x4c, 0xa9, 0x90, 0xaa, 0x5c,
	0xba, 0x1e, 0xed, 0x70, 0x91, 0x46, 0x2f, 0x61, 0x5d, 0x24, 0x32, 0x33, 0x75, 0xe9, 0x0e, 0xbb,
	0xc1, 0xe5, 0x97, 0xef, 0x17, 0x27, 0x2c, 0x9d, 0xd0, 0x69, 0x68, 0x04, 0x15, 0xaa, 0x95, 0xe1,
	0x22, 0x21, 0xa9, 0x34, 0x05, 0xea, 0x0e, 0x7b, 0xc1, 0x45, 0x89, 0xed, 0x36, 0x27, 0xeb, 0xa7,
	0x31, 0xe5, 0x47, 0x6e, 0x6b, 0x55, 0x02, 0xff, 0x27, 0x40, 0x15, 0x13, 0x21, 0x58, 0xbf, 0x61,
	0x42, 0xda, 0xe6, 0x99, 0xb5, 0xe6, 0x32, 0xc6, 0xa5, 0xc9, 0xb0, 0x11, 0x9a, 0xb5, 0x2e, 0x61,
	0x86, 0x85, 0xb8, 0x63, 0x7c, 0x6c, 0x5b, 0x36, 0xc3, 0xc8, 0x83, 0x16, 0x8e, 0x22, 0x96, 0xdb,
	0x8c, 0x3a, 0xa1, 0x83, 0xfe, 0x26, 0x6c, 0xdc, 0xcb, 0xcd, 0xff, 0x01, 0xcf, 0xaf, 0x09, 0xa7,
	0x93, 0xe2, 0x82, 0xf0, 0x84, 0x0a, 0x41, 0x59, 0xea, 0x1c, 0xa8, 0x2c, 0x81, 0x23, 0xa9, 0x08,
	0x9b, 0x8b, 0x45, 0xaa, 0x15, 0xfd, 0x88, 0x25, 0x19, 0xe6, 0xca, 0xab, 0x82, 0xe5, 0x3c, 0x22,
	0xda, 0x4b, 0x6b, 0x6a, 0xc7, 0x12, 0xef, 0x17, 0xe0, 0x2d, 0x87, 0xb7, 0x86, 0x55, 0xd6, 0xe2,
	0xb3, 0x00, 0x35, 0x13, 0xa0, 0x22, 0xf4, 0x1b, 0x44, 0x1e, 0xa9, 0x95, 0xb0, 0x8d, 0x71, 0x10,
	0xbd, 0x56, 0xe6, 0x51, 0x63, 0xad, 0x07, 0xc6, 0x36, 0xa5, 0x13, 0x1c, 0x59, 0x22, 0x9c, 0x49,
	0xfe, 0xdf, 0x1a, 0xb4, 0x1d, 0xbd, 0xd4, 0xfe, 0x3d, 0x00, 0x37, 0x16, 0xe7, 0xa7, 0xb6, 0xae,
	0x73, 0xcc, 0x03, 0xc3, 0xb4, 0xf6, 0xe0, 0x30, 0xa9, 0x2c, 0xa3, 0x98, 0xaa, 0x72, 0x0a, 0x55,
	0x69, 0xfd, 0x02, 0x07, 0xdd, 0xfc, 0xa6, 0x73, 0xa3, 0xe1, 0xb0, 0xd3, 0xae, 0x72, 0x95, 0x53,
	0xb3, 0xd2, 0x34, 0x1e, 0xfe, 0xab, 0x43, 0xe7, 0xc8, 0xfc, 0x5a, 0x97, 0xb7, 0x11, 0x3a, 0x84,
	0xee, 0x67, 0x22, 0x9d, 0x99, 0x51, 0x65, 0x7a, 0xdb, 0xa4, 0xc1, 0xb3, 0x60, 0xf1, 0x23, 0xf0,
	0x9f, 0xd8, 0x53, 0xdf, 0x46, 0x2a, 0xe6, 0xed, 0xe3, 0x4f, 0x9d, 0x01, 0x9a, 0xbb, 0xeb, 0xb8,
	0x38, 0x33, 0x83, 0x37, 0x08, 0x56, 0xfe, 0x51, 0x0f, 0x87, 0x39, 0x87, 0xfe, 0x62, 0xcb, 0x91,
	0x17, 0xac, 0x30, 0xd9, 0x60, 0x27, 0x58, 0xe5, 0x0f, 0x15, 0xea, 0x14, 0x36, 0x4e, 0xb4, 0x6d,
	0x67, 0xef, 0xdf, 0x0e, 0xca, 0xdf, 0x39, 0x70, 0xbf, 0x73, 0x60, 0xfe, 0xe2, 0xc1, 0xee, 0x12,
	0x7f, 0x9e, 0xca, 0x0f, 0x87, 0xd7, 0x38, 0xce, 0xc9, 0xa8, 0x69, 0xc8, 0xf7, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x9c, 0x79, 0xa2, 0xfc, 0x0d, 0x06, 0x00, 0x00,
}
