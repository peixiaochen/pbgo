// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/pbgo"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_acb411caf57e7eca, []int{0}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "hello_pb.String")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_acb411caf57e7eca) }

var fileDescriptor_hello_acb411caf57e7eca = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x73, 0xe2, 0x0b, 0x92, 0xa4, 0x94,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0x33, 0x12, 0x33, 0x8d,
	0x0c, 0x0c, 0x0d, 0xf4, 0x0b, 0x92, 0xd2, 0xf3, 0xc1, 0x04, 0x44, 0xb9, 0x92, 0x1c, 0x17, 0x5b,
	0x70, 0x49, 0x51, 0x66, 0x5e, 0xba, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x63, 0x14, 0xc5, 0xc5, 0xe3, 0x01, 0x32, 0x30, 0x38,
	0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x8b, 0x8b, 0x15, 0xcc, 0x17, 0x12, 0xd0, 0x83, 0x59,
	0xa4, 0x07, 0x31, 0x40, 0x0a, 0x43, 0x44, 0x49, 0xf6, 0xd6, 0xbb, 0xbf, 0x3e, 0x12, 0x52, 0x7c,
	0xfa, 0x60, 0x09, 0xfd, 0x6a, 0xb0, 0x91, 0xb5, 0x4a, 0x6c, 0x10, 0x7e, 0x12, 0x1b, 0xd8, 0x09,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x0f, 0x7b, 0x64, 0xc0, 0x00, 0x00, 0x00,
}
