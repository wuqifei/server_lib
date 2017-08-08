// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	HeartbeatReq
	HeartbeatAck
	ErrorAck
	TouristsReq
	TouristsAck
	TouristsLoginReq
	TouristsLoginAck
*/
package pb

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

// 心跳请求0
type HeartbeatReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartbeatReq) Reset()                    { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()               {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 心跳回复1
type HeartbeatAck struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartbeatAck) Reset()                    { *m = HeartbeatAck{} }
func (m *HeartbeatAck) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatAck) ProtoMessage()               {}
func (*HeartbeatAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 通用的错误回复2
type ErrorAck struct {
	Errid            *int32 `protobuf:"varint,1,req,name=errid" json:"errid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ErrorAck) Reset()                    { *m = ErrorAck{} }
func (m *ErrorAck) String() string            { return proto.CompactTextString(m) }
func (*ErrorAck) ProtoMessage()               {}
func (*ErrorAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ErrorAck) GetErrid() int32 {
	if m != nil && m.Errid != nil {
		return *m.Errid
	}
	return 0
}

// 请求生成游客账号1000
type TouristsReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TouristsReq) Reset()                    { *m = TouristsReq{} }
func (m *TouristsReq) String() string            { return proto.CompactTextString(m) }
func (*TouristsReq) ProtoMessage()               {}
func (*TouristsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 请求生成游客账号返回1001
type TouristsAck struct {
	Touristsid       *uint64 `protobuf:"varint,1,req,name=touristsid" json:"touristsid,omitempty"`
	Touristscot      *string `protobuf:"bytes,2,req,name=touristscot" json:"touristscot,omitempty"`
	Touristspwd      *string `protobuf:"bytes,3,req,name=touristspwd" json:"touristspwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristsAck) Reset()                    { *m = TouristsAck{} }
func (m *TouristsAck) String() string            { return proto.CompactTextString(m) }
func (*TouristsAck) ProtoMessage()               {}
func (*TouristsAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TouristsAck) GetTouristsid() uint64 {
	if m != nil && m.Touristsid != nil {
		return *m.Touristsid
	}
	return 0
}

func (m *TouristsAck) GetTouristscot() string {
	if m != nil && m.Touristscot != nil {
		return *m.Touristscot
	}
	return ""
}

func (m *TouristsAck) GetTouristspwd() string {
	if m != nil && m.Touristspwd != nil {
		return *m.Touristspwd
	}
	return ""
}

// 请求游客登录1002
type TouristsLoginReq struct {
	Touristscot      *string `protobuf:"bytes,1,req,name=touristscot" json:"touristscot,omitempty"`
	Touristspwd      *string `protobuf:"bytes,2,req,name=touristspwd" json:"touristspwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristsLoginReq) Reset()                    { *m = TouristsLoginReq{} }
func (m *TouristsLoginReq) String() string            { return proto.CompactTextString(m) }
func (*TouristsLoginReq) ProtoMessage()               {}
func (*TouristsLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TouristsLoginReq) GetTouristscot() string {
	if m != nil && m.Touristscot != nil {
		return *m.Touristscot
	}
	return ""
}

func (m *TouristsLoginReq) GetTouristspwd() string {
	if m != nil && m.Touristspwd != nil {
		return *m.Touristspwd
	}
	return ""
}

// 游客登录返回1003
type TouristsLoginAck struct {
	Str              *string `protobuf:"bytes,1,req,name=str" json:"str,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristsLoginAck) Reset()                    { *m = TouristsLoginAck{} }
func (m *TouristsLoginAck) String() string            { return proto.CompactTextString(m) }
func (*TouristsLoginAck) ProtoMessage()               {}
func (*TouristsLoginAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TouristsLoginAck) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartbeatReq)(nil), "protomsg.heartbeat_req")
	proto.RegisterType((*HeartbeatAck)(nil), "protomsg.heartbeat_ack")
	proto.RegisterType((*ErrorAck)(nil), "protomsg.error_ack")
	proto.RegisterType((*TouristsReq)(nil), "protomsg.tourists_req")
	proto.RegisterType((*TouristsAck)(nil), "protomsg.tourists_ack")
	proto.RegisterType((*TouristsLoginReq)(nil), "protomsg.tourists_login_req")
	proto.RegisterType((*TouristsLoginAck)(nil), "protomsg.tourists_login_ack")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xcb, 0xae, 0x82, 0x30,
	0x10, 0x86, 0x43, 0x39, 0x24, 0x87, 0xc1, 0x5b, 0x1a, 0x17, 0xac, 0x0c, 0x76, 0x61, 0x58, 0xf9,
	0x30, 0xac, 0xdc, 0x19, 0x84, 0x06, 0x89, 0x97, 0xc1, 0xe9, 0x18, 0x5e, 0xdf, 0x14, 0xc1, 0x14,
	0x59, 0xb8, 0x6a, 0xbf, 0x3f, 0x33, 0x5f, 0xfb, 0x43, 0x74, 0xc5, 0xaa, 0xbe, 0xef, 0x1b, 0x42,
	0x46, 0xf9, 0xdf, 0x1d, 0x37, 0x53, 0xa9, 0x25, 0xcc, 0xcf, 0x3a, 0x27, 0x3e, 0xe9, 0x9c, 0x8f,
	0xa4, 0x1f, 0xe3, 0x20, 0x2f, 0x2e, 0x6a, 0x0b, 0xa1, 0x26, 0x42, 0xb2, 0x20, 0xd7, 0x10, 0x68,
	0xa2, 0xba, 0x8c, 0xbd, 0x44, 0xa4, 0x41, 0xf6, 0x06, 0xb5, 0x80, 0x19, 0xe3, 0x93, 0x6a, 0xc3,
	0xa6, 0x73, 0x90, 0xc3, 0x76, 0x6b, 0x03, 0x30, 0x70, 0xbf, 0xfa, 0x97, 0x39, 0x89, 0x4c, 0x20,
	0x1a, 0xa8, 0x40, 0x8e, 0x45, 0x22, 0xd2, 0x30, 0x73, 0x23, 0x77, 0xa2, 0x69, 0xcb, 0xd8, 0x1f,
	0x4f, 0x34, 0x6d, 0xa9, 0x0e, 0x20, 0x3f, 0x6f, 0x76, 0x55, 0xed, 0x4f, 0xbe, 0xcd, 0xde, 0x4f,
	0xb3, 0x98, 0x9a, 0x77, 0x13, 0xb3, 0xed, 0xb4, 0x02, 0xdf, 0x30, 0xf5, 0x46, 0x7b, 0x7d, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x06, 0x59, 0x11, 0x28, 0x62, 0x01, 0x00, 0x00,
}