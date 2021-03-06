// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package reverse

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type DoReq struct {
	Header               []byte   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Route                string   `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	Round                string   `protobuf:"bytes,4,opt,name=round,proto3" json:"round,omitempty"`
	ReqID                string   `protobuf:"bytes,5,opt,name=reqID,proto3" json:"reqID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReq) Reset()         { *m = DoReq{} }
func (m *DoReq) String() string { return proto.CompactTextString(m) }
func (*DoReq) ProtoMessage()    {}
func (*DoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *DoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReq.Unmarshal(m, b)
}
func (m *DoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReq.Marshal(b, m, deterministic)
}
func (m *DoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReq.Merge(m, src)
}
func (m *DoReq) XXX_Size() int {
	return xxx_messageInfo_DoReq.Size(m)
}
func (m *DoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DoReq proto.InternalMessageInfo

func (m *DoReq) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DoReq) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *DoReq) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *DoReq) GetRound() string {
	if m != nil {
		return m.Round
	}
	return ""
}

func (m *DoReq) GetReqID() string {
	if m != nil {
		return m.ReqID
	}
	return ""
}

type DoResp struct {
	Header               []byte   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Route                string   `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	Round                string   `protobuf:"bytes,4,opt,name=round,proto3" json:"round,omitempty"`
	ReqID                string   `protobuf:"bytes,5,opt,name=reqID,proto3" json:"reqID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoResp) Reset()         { *m = DoResp{} }
func (m *DoResp) String() string { return proto.CompactTextString(m) }
func (*DoResp) ProtoMessage()    {}
func (*DoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *DoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoResp.Unmarshal(m, b)
}
func (m *DoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoResp.Marshal(b, m, deterministic)
}
func (m *DoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoResp.Merge(m, src)
}
func (m *DoResp) XXX_Size() int {
	return xxx_messageInfo_DoResp.Size(m)
}
func (m *DoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DoResp.DiscardUnknown(m)
}

var xxx_messageInfo_DoResp proto.InternalMessageInfo

func (m *DoResp) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DoResp) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *DoResp) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *DoResp) GetRound() string {
	if m != nil {
		return m.Round
	}
	return ""
}

func (m *DoResp) GetReqID() string {
	if m != nil {
		return m.ReqID
	}
	return ""
}

func init() {
	proto.RegisterType((*DoReq)(nil), "reverse.DoReq")
	proto.RegisterType((*DoResp)(nil), "reverse.DoResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Restful service

type RestfulClient interface {
	Echo(ctx context.Context, in *DoReq, opts ...grpc.CallOption) (*DoResp, error)
}

type restfulClient struct {
	cc *grpc.ClientConn
}

func NewRestfulClient(cc *grpc.ClientConn) RestfulClient {
	return &restfulClient{cc}
}

func (c *restfulClient) Echo(ctx context.Context, in *DoReq, opts ...grpc.CallOption) (*DoResp, error) {
	out := new(DoResp)
	err := grpc.Invoke(ctx, "/reverse.Restful/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Restful service

type RestfulServer interface {
	Echo(context.Context, *DoReq) (*DoResp, error)
}

func RegisterRestfulServer(s *grpc.Server, srv RestfulServer) {
	s.RegisterService(&_Restful_serviceDesc, srv)
}

func _Restful_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reverse.Restful/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Echo(ctx, req.(*DoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Restful_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reverse.Restful",
	HandlerType: (*RestfulServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Restful_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x90, 0x4d, 0x4a, 0xc5, 0x30,
	0x14, 0x46, 0x8d, 0xf6, 0x87, 0x5e, 0xc4, 0x42, 0x10, 0x09, 0x0e, 0xa4, 0x74, 0x54, 0x27, 0x19,
	0xa8, 0x2b, 0x28, 0x15, 0x71, 0x56, 0xb2, 0x03, 0x9b, 0x5c, 0x6d, 0x41, 0x9a, 0xf4, 0xa6, 0x2d,
	0xb8, 0x7b, 0x69, 0x0c, 0xc2, 0xdb, 0xc0, 0x9b, 0x7d, 0xe7, 0x4c, 0x3e, 0x38, 0x00, 0x5f, 0xe4,
	0xb4, 0x74, 0x64, 0x57, 0xcb, 0x73, 0xc2, 0x1d, 0xc9, 0x63, 0xbd, 0x41, 0xda, 0x59, 0x85, 0x0b,
	0xbf, 0x83, 0x6c, 0xc4, 0x0f, 0x83, 0x24, 0x58, 0xc5, 0x9a, 0x6b, 0x15, 0x89, 0x73, 0x48, 0x06,
	0x6b, 0x7e, 0xc4, 0x65, 0xb0, 0x61, 0xf3, 0x5b, 0x48, 0xc9, 0x6e, 0x2b, 0x8a, 0xab, 0x8a, 0x35,
	0x85, 0xfa, 0x83, 0x68, 0x67, 0x23, 0x92, 0x7f, 0x3b, 0x9b, 0x60, 0x71, 0x79, 0xef, 0x44, 0x1a,
	0xed, 0x01, 0xf5, 0x0e, 0xd9, 0x71, 0xeb, 0xdd, 0x79, 0x7f, 0x9f, 0x5e, 0x20, 0x57, 0xe8, 0xd7,
	0xcf, 0xed, 0x9b, 0x3f, 0x42, 0xf2, 0xaa, 0x47, 0xcb, 0x6f, 0x64, 0x6c, 0x21, 0x43, 0x88, 0xfb,
	0xf2, 0x84, 0xbd, 0xab, 0x2f, 0xda, 0x07, 0x28, 0x27, 0x2b, 0x43, 0x3e, 0x8f, 0xb4, 0x4f, 0x1a,
	0xdb, 0xe2, 0x8d, 0x9c, 0xee, 0x8f, 0x96, 0x3d, 0x1b, 0xb2, 0x10, 0xf5, 0xf9, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x71, 0xc8, 0x40, 0xb7, 0x62, 0x01, 0x00, 0x00,
}
