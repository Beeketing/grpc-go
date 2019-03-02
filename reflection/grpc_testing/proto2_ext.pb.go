// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2_ext.proto

package grpc_testing

import proto "github.com/Beeketing/protobuf/proto"
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

type Extension struct {
	Whatzit              *int32   `protobuf:"varint,1,opt,name=whatzit" json:"whatzit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto2_ext_4437118420d604f2, []int{0}
}
func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (dst *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(dst, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetWhatzit() int32 {
	if m != nil && m.Whatzit != nil {
		return *m.Whatzit
	}
	return 0
}

var E_Foo = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtended)(nil),
	ExtensionType: (*int32)(nil),
	Field:         13,
	Name:          "grpc.testing.foo",
	Tag:           "varint,13,opt,name=foo",
	Filename:      "proto2_ext.proto",
}

var E_Bar = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtended)(nil),
	ExtensionType: (*Extension)(nil),
	Field:         17,
	Name:          "grpc.testing.bar",
	Tag:           "bytes,17,opt,name=bar",
	Filename:      "proto2_ext.proto",
}

var E_Baz = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtended)(nil),
	ExtensionType: (*SearchRequest)(nil),
	Field:         19,
	Name:          "grpc.testing.baz",
	Tag:           "bytes,19,opt,name=baz",
	Filename:      "proto2_ext.proto",
}

func init() {
	proto.RegisterType((*Extension)(nil), "grpc.testing.Extension")
	proto.RegisterExtension(E_Foo)
	proto.RegisterExtension(E_Bar)
	proto.RegisterExtension(E_Baz)
}

func init() { proto.RegisterFile("proto2_ext.proto", fileDescriptor_proto2_ext_4437118420d604f2) }

var fileDescriptor_proto2_ext_4437118420d604f2 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x8a, 0x4f, 0xad, 0x28, 0xd1, 0x03, 0x33, 0x85, 0x78, 0xd2, 0x8b, 0x0a, 0x92, 0xf5,
	0x4a, 0x52, 0x8b, 0x4b, 0x32, 0xf3, 0xd2, 0xa5, 0x78, 0x20, 0xf2, 0x10, 0x39, 0x29, 0x2e, 0x90,
	0x30, 0x84, 0xad, 0xa4, 0xca, 0xc5, 0xe9, 0x5a, 0x51, 0x92, 0x9a, 0x57, 0x9c, 0x99, 0x9f, 0x27,
	0x24, 0xc1, 0xc5, 0x5e, 0x9e, 0x91, 0x58, 0x52, 0x95, 0x59, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x1a, 0x04, 0xe3, 0x5a, 0xe9, 0x70, 0x31, 0xa7, 0xe5, 0xe7, 0x0b, 0x49, 0xe9, 0x21, 0x1b, 0xab,
	0x17, 0x92, 0xef, 0x94, 0x0a, 0xd6, 0x9d, 0x92, 0x9a, 0x22, 0xc1, 0x0b, 0xd6, 0x01, 0x52, 0x66,
	0xe5, 0xca, 0xc5, 0x9c, 0x94, 0x58, 0x84, 0x57, 0xb5, 0xa0, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x38,
	0xaa, 0x0a, 0xb8, 0x4b, 0x82, 0x40, 0xfa, 0xad, 0x3c, 0x41, 0xc6, 0x54, 0xe1, 0x35, 0x46, 0x18,
	0x6c, 0x8c, 0x34, 0xaa, 0x8a, 0xe0, 0xd4, 0xc4, 0xa2, 0xe4, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x90, 0x51, 0x55, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x6b, 0x94, 0x9f, 0x21,
	0x01, 0x00, 0x00,
}
