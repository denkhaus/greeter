// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package denkhaus_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	Request
	Response
*/
package denkhaus_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "denkhaus.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "denkhaus.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayClient struct {
	c           client.Client
	serviceName string
}

func NewSayClient(serviceName string, c client.Client) SayClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "denkhaus.micro.srv.greeter"
	}
	return &sayClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Say{hdlr}, opts...))
}

type Say struct {
	SayHandler
}

func (h *Say) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0x49, 0xcd, 0xcb, 0xce, 0x48, 0x2c,
	0x2d, 0xd6, 0xcb, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x4b, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0x52, 0x92, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95,
	0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x73,
	0x8b, 0xd3, 0xa1, 0xd2, 0x20, 0xa6, 0x51, 0x34, 0x17, 0x73, 0x70, 0x62, 0xa5, 0x50, 0x08, 0x17,
	0xab, 0x07, 0xc8, 0x3a, 0x21, 0x65, 0x3d, 0xdc, 0x36, 0xe9, 0x41, 0xad, 0x91, 0x52, 0xc1, 0xaf,
	0x08, 0x62, 0x99, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xf1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0x2b, 0xda, 0xc1, 0xcb, 0x00, 0x00, 0x00,
}
