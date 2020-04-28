// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Customer struct {
	CustomerID           int32                `protobuf:"varint,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	InternalCustomerID   string               `protobuf:"bytes,3,opt,name=InternalCustomerID,proto3" json:"InternalCustomerID,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	RowVersion           string               `protobuf:"bytes,6,opt,name=RowVersion,proto3" json:"RowVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetInternalCustomerID() string {
	if m != nil {
		return m.InternalCustomerID
	}
	return ""
}

func (m *Customer) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Customer) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Customer) GetRowVersion() string {
	if m != nil {
		return m.RowVersion
	}
	return ""
}

type CustomerCredentials struct {
	ClientAPIKey         string               `protobuf:"bytes,1,opt,name=ClientAPIKey,proto3" json:"ClientAPIKey,omitempty"`
	Secret               string               `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
	CustomerID           int32                `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerCredentials) Reset()         { *m = CustomerCredentials{} }
func (m *CustomerCredentials) String() string { return proto.CompactTextString(m) }
func (*CustomerCredentials) ProtoMessage()    {}
func (*CustomerCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *CustomerCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerCredentials.Unmarshal(m, b)
}
func (m *CustomerCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerCredentials.Marshal(b, m, deterministic)
}
func (m *CustomerCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerCredentials.Merge(m, src)
}
func (m *CustomerCredentials) XXX_Size() int {
	return xxx_messageInfo_CustomerCredentials.Size(m)
}
func (m *CustomerCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerCredentials proto.InternalMessageInfo

func (m *CustomerCredentials) GetClientAPIKey() string {
	if m != nil {
		return m.ClientAPIKey
	}
	return ""
}

func (m *CustomerCredentials) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *CustomerCredentials) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *CustomerCredentials) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CustomerCredentials) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type CustomerIP struct {
	CustomerIP           string               `protobuf:"bytes,1,opt,name=CustomerIP,proto3" json:"CustomerIP,omitempty"`
	CustomerID           int32                `protobuf:"varint,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerIP) Reset()         { *m = CustomerIP{} }
func (m *CustomerIP) String() string { return proto.CompactTextString(m) }
func (*CustomerIP) ProtoMessage()    {}
func (*CustomerIP) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *CustomerIP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerIP.Unmarshal(m, b)
}
func (m *CustomerIP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerIP.Marshal(b, m, deterministic)
}
func (m *CustomerIP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerIP.Merge(m, src)
}
func (m *CustomerIP) XXX_Size() int {
	return xxx_messageInfo_CustomerIP.Size(m)
}
func (m *CustomerIP) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerIP.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerIP proto.InternalMessageInfo

func (m *CustomerIP) GetCustomerIP() string {
	if m != nil {
		return m.CustomerIP
	}
	return ""
}

func (m *CustomerIP) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *CustomerIP) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CustomerIP) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type CustomerProduct struct {
	CustomerProductID    int32    `protobuf:"varint,1,opt,name=CustomerProductID,proto3" json:"CustomerProductID,omitempty"`
	CustomerProduct      string   `protobuf:"bytes,2,opt,name=CustomerProduct,proto3" json:"CustomerProduct,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerProduct) Reset()         { *m = CustomerProduct{} }
func (m *CustomerProduct) String() string { return proto.CompactTextString(m) }
func (*CustomerProduct) ProtoMessage()    {}
func (*CustomerProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *CustomerProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerProduct.Unmarshal(m, b)
}
func (m *CustomerProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerProduct.Marshal(b, m, deterministic)
}
func (m *CustomerProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerProduct.Merge(m, src)
}
func (m *CustomerProduct) XXX_Size() int {
	return xxx_messageInfo_CustomerProduct.Size(m)
}
func (m *CustomerProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerProduct.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerProduct proto.InternalMessageInfo

func (m *CustomerProduct) GetCustomerProductID() int32 {
	if m != nil {
		return m.CustomerProductID
	}
	return 0
}

func (m *CustomerProduct) GetCustomerProduct() string {
	if m != nil {
		return m.CustomerProduct
	}
	return ""
}

type CustomerServiceOfferings struct {
	CustomerID                int32                `protobuf:"varint,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	CustomerServiceOfferingID int32                `protobuf:"varint,2,opt,name=CustomerServiceOfferingID,proto3" json:"CustomerServiceOfferingID,omitempty"`
	CustomerClientProductID   int32                `protobuf:"varint,3,opt,name=CustomerClientProductID,proto3" json:"CustomerClientProductID,omitempty"`
	StartDate                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate                   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	BillingCycle              string               `protobuf:"bytes,6,opt,name=BillingCycle,proto3" json:"BillingCycle,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}             `json:"-"`
	XXX_unrecognized          []byte               `json:"-"`
	XXX_sizecache             int32                `json:"-"`
}

func (m *CustomerServiceOfferings) Reset()         { *m = CustomerServiceOfferings{} }
func (m *CustomerServiceOfferings) String() string { return proto.CompactTextString(m) }
func (*CustomerServiceOfferings) ProtoMessage()    {}
func (*CustomerServiceOfferings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4}
}

func (m *CustomerServiceOfferings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerServiceOfferings.Unmarshal(m, b)
}
func (m *CustomerServiceOfferings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerServiceOfferings.Marshal(b, m, deterministic)
}
func (m *CustomerServiceOfferings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerServiceOfferings.Merge(m, src)
}
func (m *CustomerServiceOfferings) XXX_Size() int {
	return xxx_messageInfo_CustomerServiceOfferings.Size(m)
}
func (m *CustomerServiceOfferings) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerServiceOfferings.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerServiceOfferings proto.InternalMessageInfo

func (m *CustomerServiceOfferings) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *CustomerServiceOfferings) GetCustomerServiceOfferingID() int32 {
	if m != nil {
		return m.CustomerServiceOfferingID
	}
	return 0
}

func (m *CustomerServiceOfferings) GetCustomerClientProductID() int32 {
	if m != nil {
		return m.CustomerClientProductID
	}
	return 0
}

func (m *CustomerServiceOfferings) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CustomerServiceOfferings) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *CustomerServiceOfferings) GetBillingCycle() string {
	if m != nil {
		return m.BillingCycle
	}
	return ""
}

func init() {
	proto.RegisterType((*Customer)(nil), "models.Customer")
	proto.RegisterType((*CustomerCredentials)(nil), "models.CustomerCredentials")
	proto.RegisterType((*CustomerIP)(nil), "models.CustomerIP")
	proto.RegisterType((*CustomerProduct)(nil), "models.CustomerProduct")
	proto.RegisterType((*CustomerServiceOfferings)(nil), "models.CustomerServiceOfferings")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xbf, 0x4e, 0xe3, 0x40,
	0x10, 0xc6, 0xe5, 0xfc, 0xf1, 0x5d, 0xe6, 0xa2, 0x3b, 0xdd, 0x9e, 0x74, 0xe7, 0x4b, 0x01, 0x91,
	0x2b, 0x17, 0xc8, 0x91, 0x80, 0x22, 0x05, 0x0d, 0x38, 0x14, 0x11, 0x12, 0x58, 0x0e, 0xa2, 0x77,
	0xec, 0x89, 0xb5, 0xd2, 0x7a, 0x37, 0x5a, 0x6f, 0x40, 0x79, 0x37, 0x6a, 0xde, 0x83, 0x17, 0x41,
	0x88, 0x8d, 0x9d, 0x38, 0x09, 0x26, 0x88, 0x22, 0xdd, 0xfa, 0x9b, 0x99, 0xf5, 0x6f, 0xbe, 0x9d,
	0x81, 0x9f, 0xd1, 0x2c, 0x53, 0x22, 0x45, 0xe9, 0x4e, 0xa5, 0x50, 0x82, 0x98, 0xa9, 0x88, 0x91,
	0x65, 0x9d, 0xc3, 0x44, 0x88, 0x84, 0x61, 0x4f, 0xab, 0xe3, 0xd9, 0xa4, 0xa7, 0x68, 0x8a, 0x99,
	0x0a, 0xd3, 0xe9, 0x22, 0xd1, 0x7e, 0x31, 0xe0, 0xbb, 0x97, 0xd7, 0x92, 0x03, 0x80, 0xe2, 0x3c,
	0x1c, 0x58, 0x46, 0xd7, 0x70, 0x9a, 0x41, 0x49, 0x21, 0x04, 0x1a, 0xd7, 0x61, 0x8a, 0x56, 0xad,
	0x6b, 0x38, 0xad, 0x40, 0x9f, 0x89, 0x0b, 0x64, 0xc8, 0x15, 0x4a, 0x1e, 0xb2, 0x52, 0x6d, 0x5d,
	0x67, 0xbc, 0x13, 0x21, 0x7d, 0x68, 0x8d, 0x54, 0x28, 0xd5, 0x20, 0x54, 0x68, 0x35, 0xba, 0x86,
	0xf3, 0xe3, 0xb8, 0xe3, 0x2e, 0x28, 0xdd, 0x82, 0xd2, 0xbd, 0x2d, 0x28, 0x83, 0x55, 0x32, 0x39,
	0x85, 0x6f, 0x97, 0x3c, 0xd6, 0x75, 0xcd, 0x9d, 0x75, 0x45, 0xea, 0x5b, 0x4f, 0x81, 0x78, 0xb8,
	0x43, 0x99, 0x51, 0xc1, 0x2d, 0x53, 0x73, 0x95, 0x14, 0xfb, 0xd9, 0x80, 0x3f, 0x05, 0x9e, 0x27,
	0x31, 0x46, 0xae, 0x68, 0xc8, 0x32, 0x62, 0x43, 0xdb, 0x63, 0x14, 0xb9, 0x3a, 0xf7, 0x87, 0x57,
	0x38, 0xd7, 0x6e, 0xb4, 0x82, 0x35, 0x8d, 0xfc, 0x05, 0x73, 0x84, 0x91, 0x44, 0x95, 0x3b, 0x92,
	0x7f, 0x6d, 0xf8, 0x58, 0xdf, 0xf2, 0x71, 0xcf, 0x1e, 0xd8, 0x8f, 0x46, 0x09, 0xc8, 0x5f, 0xc3,
	0xf3, 0xf3, 0xc6, 0xaa, 0xe2, 0x03, 0xdd, 0xda, 0x07, 0xf8, 0xf5, 0x2f, 0xe2, 0x37, 0x3e, 0x8f,
	0x4f, 0xe1, 0x57, 0xf1, 0x77, 0x5f, 0x8a, 0x78, 0x16, 0x29, 0x72, 0x04, 0xbf, 0x37, 0xa4, 0xe5,
	0xc0, 0x6e, 0x07, 0x88, 0xb3, 0x75, 0x41, 0xfe, 0x60, 0x9b, 0xb2, 0xfd, 0x54, 0x03, 0xab, 0xd0,
	0x46, 0x28, 0xef, 0x69, 0x84, 0x37, 0x93, 0x09, 0x4a, 0xca, 0x93, 0x6c, 0xe7, 0x7a, 0x9c, 0xc1,
	0xff, 0x8a, 0xda, 0xa5, 0x8d, 0xd5, 0x09, 0xa4, 0x0f, 0xff, 0x96, 0x73, 0xa8, 0x87, 0x6c, 0xd5,
	0xd8, 0x62, 0x82, 0xaa, 0xc2, 0x7b, 0x5f, 0x29, 0x1b, 0xda, 0x17, 0x94, 0x31, 0xca, 0x13, 0x6f,
	0x1e, 0x31, 0xcc, 0x97, 0x6a, 0x4d, 0x1b, 0x9b, 0xfa, 0x82, 0x93, 0xd7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0xe8, 0x83, 0x91, 0x99, 0x04, 0x00, 0x00,
}
