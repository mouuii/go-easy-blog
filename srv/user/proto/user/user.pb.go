// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	LoginReq
	LoginResp
*/
package user

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

type LoginReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name" xml:"name" form:"name"`
	Pwd  string `protobuf:"bytes,2,opt,name=pwd" json:"pwd" xml:"pwd" form:"pwd"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type LoginResp struct {
}

func (m *LoginResp) Reset()                    { *m = LoginResp{} }
func (m *LoginResp) String() string            { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()               {}
func (*LoginResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*LoginResp)(nil), "LoginResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginReq, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "userservice"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginReq, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Login", in)
	out := new(LoginReq)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Login(context.Context, *LoginReq, *LoginReq) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Login(ctx context.Context, in *LoginReq, out *LoginReq) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe0, 0xe2, 0xf0, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b,
	0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x0b, 0xca, 0x53, 0x24, 0x98, 0xc0, 0x42, 0x20, 0xa6,
	0x12, 0x37, 0x17, 0x27, 0x54, 0x47, 0x71, 0x81, 0x91, 0x1e, 0x17, 0x77, 0x68, 0x71, 0x6a, 0x51,
	0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3c, 0x17, 0x2b, 0x58, 0x4e, 0x88, 0x53, 0x0f,
	0x66, 0xaa, 0x14, 0x82, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x89, 0xfb, 0x52, 0xf2, 0x83, 0x00, 0x00, 0x00,
}
