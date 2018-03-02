// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ipfix.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/ipfix.proto

It has these top-level messages:
	GetLocationRequest
	GetLocationResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetLocationRequest struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
}

func (m *GetLocationRequest) Reset()                    { *m = GetLocationRequest{} }
func (m *GetLocationRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetLocationRequest) ProtoMessage()               {}
func (*GetLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetLocationRequest) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type GetLocationResponse struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
}

func (m *GetLocationResponse) Reset()                    { *m = GetLocationResponse{} }
func (m *GetLocationResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetLocationResponse) ProtoMessage()               {}
func (*GetLocationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetLocationResponse) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func init() {
	proto1.RegisterType((*GetLocationRequest)(nil), "proto.GetLocationRequest")
	proto1.RegisterType((*GetLocationResponse)(nil), "proto.GetLocationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ipfix service

type IpfixClient interface {
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
}

type ipfixClient struct {
	cc *grpc.ClientConn
}

func NewIpfixClient(cc *grpc.ClientConn) IpfixClient {
	return &ipfixClient{cc}
}

func (c *ipfixClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := grpc.Invoke(ctx, "/proto.Ipfix/GetLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ipfix service

type IpfixServer interface {
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
}

func RegisterIpfixServer(s *grpc.Server, srv IpfixServer) {
	s.RegisterService(&_Ipfix_serviceDesc, srv)
}

func _Ipfix_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpfixServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ipfix/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpfixServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ipfix_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ipfix",
	HandlerType: (*IpfixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocation",
			Handler:    _Ipfix_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ipfix.proto",
}

func init() { proto1.RegisterFile("proto/ipfix.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x2c, 0x48, 0xcb, 0xac, 0xd0, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x31,
	0x97, 0x90, 0x7b, 0x6a, 0x89, 0x4f, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x90, 0x2c, 0x17, 0x57, 0x66, 0x41, 0x7c, 0x62, 0x4a, 0x4a, 0x51, 0x6a,
	0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x67, 0x66, 0x81, 0x23, 0x44, 0x40, 0xc9,
	0x84, 0x4b, 0x18, 0x45, 0x53, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x01, 0x5d, 0x46, 0xfe, 0x5c,
	0xac, 0x9e, 0x20, 0x07, 0x08, 0xb9, 0x71, 0x71, 0x23, 0x69, 0x17, 0x92, 0x84, 0xb8, 0x48, 0x0f,
	0xd3, 0x1d, 0x52, 0x52, 0xd8, 0xa4, 0x20, 0xb6, 0x29, 0x31, 0x38, 0xc9, 0x72, 0xf1, 0x66, 0xe6,
	0xeb, 0xa5, 0x17, 0x15, 0x24, 0xeb, 0x81, 0x7d, 0xe6, 0xc4, 0x05, 0x36, 0x3f, 0x00, 0xa4, 0x25,
	0x80, 0x31, 0x89, 0x0d, 0xac, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x0c, 0xc5, 0x91,
	0xfd, 0x00, 0x00, 0x00,
}
