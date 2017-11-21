// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	ResponseData
	Product
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResponseData_ResponseStatusEnum int32

const (
	ResponseData_SUCCESS ResponseData_ResponseStatusEnum = 0
	ResponseData_FAILED  ResponseData_ResponseStatusEnum = 1
)

var ResponseData_ResponseStatusEnum_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}
var ResponseData_ResponseStatusEnum_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x ResponseData_ResponseStatusEnum) String() string {
	return proto.EnumName(ResponseData_ResponseStatusEnum_name, int32(x))
}
func (ResponseData_ResponseStatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type Product_ResetApprovePassword int32

const (
	Product_AVAILABLE Product_ResetApprovePassword = 0
	Product_BOOKED    Product_ResetApprovePassword = 1
	Product_ACCEPTED  Product_ResetApprovePassword = 2
)

var Product_ResetApprovePassword_name = map[int32]string{
	0: "AVAILABLE",
	1: "BOOKED",
	2: "ACCEPTED",
}
var Product_ResetApprovePassword_value = map[string]int32{
	"AVAILABLE": 0,
	"BOOKED":    1,
	"ACCEPTED":  2,
}

func (x Product_ResetApprovePassword) String() string {
	return proto.EnumName(Product_ResetApprovePassword_name, int32(x))
}
func (Product_ResetApprovePassword) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ResponseData struct {
	Message string                          `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Status  ResponseData_ResponseStatusEnum `protobuf:"varint,2,opt,name=status,enum=api.ResponseData_ResponseStatusEnum" json:"status,omitempty"`
}

func (m *ResponseData) Reset()                    { *m = ResponseData{} }
func (m *ResponseData) String() string            { return proto.CompactTextString(m) }
func (*ResponseData) ProtoMessage()               {}
func (*ResponseData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResponseData) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResponseData) GetStatus() ResponseData_ResponseStatusEnum {
	if m != nil {
		return m.Status
	}
	return ResponseData_SUCCESS
}

// Product - product structre
type Product struct {
	Id     int64                        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string                       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status Product_ResetApprovePassword `protobuf:"varint,3,opt,name=status,enum=api.Product_ResetApprovePassword" json:"status,omitempty"`
	Hash   string                       `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Product) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetStatus() Product_ResetApprovePassword {
	if m != nil {
		return m.Status
	}
	return Product_AVAILABLE
}

func (m *Product) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseData)(nil), "api.ResponseData")
	proto.RegisterType((*Product)(nil), "api.Product")
	proto.RegisterEnum("api.ResponseData_ResponseStatusEnum", ResponseData_ResponseStatusEnum_name, ResponseData_ResponseStatusEnum_value)
	proto.RegisterEnum("api.Product_ResetApprovePassword", Product_ResetApprovePassword_name, Product_ResetApprovePassword_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0x06, 0xf0, 0x37, 0xed, 0xd2, 0xbe, 0x9d, 0x5d, 0x97, 0x32, 0x78, 0xe8, 0x71, 0x2d, 0x1e,
	0xf6, 0x62, 0x0f, 0x7a, 0x12, 0x04, 0xc9, 0xb6, 0x11, 0x16, 0x0b, 0x5b, 0x5a, 0xf5, 0x1e, 0x6d,
	0x70, 0x0b, 0xb6, 0x09, 0x4d, 0xaa, 0x9f, 0xc4, 0x8f, 0xe3, 0x77, 0x93, 0x66, 0xff, 0x20, 0xe8,
	0x6d, 0x86, 0xe4, 0x79, 0xf8, 0x31, 0x30, 0x6d, 0x65, 0x2d, 0xde, 0x12, 0xd5, 0x4b, 0x23, 0xd1,
	0xe5, 0xaa, 0x89, 0x3f, 0x09, 0xcc, 0x4a, 0xa1, 0x95, 0xec, 0xb4, 0xc8, 0xb8, 0xe1, 0x18, 0x81,
	0xdf, 0x0a, 0xad, 0xf9, 0xab, 0x88, 0xc8, 0x82, 0x2c, 0x83, 0xf2, 0xb0, 0xe2, 0x0d, 0x78, 0xda,
	0x70, 0x33, 0xe8, 0xc8, 0x59, 0x90, 0xe5, 0xfc, 0xf2, 0x3c, 0xe1, 0xaa, 0x49, 0x7e, 0x86, 0x8f,
	0x4b, 0x65, 0xff, 0xb1, 0x6e, 0x68, 0xcb, 0x7d, 0x26, 0xbe, 0x00, 0xfc, 0xfd, 0x8a, 0x53, 0xf0,
	0xab, 0xc7, 0x34, 0x65, 0x55, 0x15, 0xfe, 0x43, 0x00, 0xef, 0x8e, 0xae, 0x73, 0x96, 0x85, 0x24,
	0xfe, 0x22, 0xe0, 0x17, 0xbd, 0xac, 0x87, 0x17, 0x83, 0x73, 0x70, 0x9a, 0xda, 0x6a, 0xdc, 0xd2,
	0x69, 0x6a, 0x44, 0x98, 0x74, 0xbc, 0x15, 0x96, 0x11, 0x94, 0x76, 0xc6, 0xeb, 0x23, 0xce, 0xb5,
	0xb8, 0x33, 0x8b, 0xdb, 0x37, 0x8c, 0x2e, 0x61, 0xa8, 0x52, 0xbd, 0x7c, 0x17, 0x05, 0xd7, 0xfa,
	0x43, 0xf6, 0xf5, 0x41, 0x36, 0xd6, 0x6d, 0xb9, 0xde, 0x46, 0x93, 0x5d, 0xdd, 0x38, 0xc7, 0xb7,
	0x70, 0xfa, 0x57, 0x06, 0x4f, 0x20, 0xa0, 0x4f, 0x74, 0x9d, 0xd3, 0x55, 0xce, 0x76, 0xe2, 0xd5,
	0x66, 0x73, 0x3f, 0x8a, 0x71, 0x06, 0xff, 0x69, 0x9a, 0xb2, 0xe2, 0x81, 0x65, 0xa1, 0xf3, 0xec,
	0xd9, 0x1b, 0x5f, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x35, 0x7e, 0x73, 0xe9, 0x72, 0x01, 0x00,
	0x00,
}
