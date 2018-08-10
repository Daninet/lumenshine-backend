// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SendMailRequest struct {
	Base                 *BaseRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	From                 string       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	ToMultiple           []string     `protobuf:"bytes,4,rep,name=to_multiple,json=toMultiple,proto3" json:"to_multiple,omitempty"`
	Subject              string       `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string       `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	IsHtml               bool         `protobuf:"varint,7,opt,name=is_html,json=isHtml,proto3" json:"is_html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendMailRequest) Reset()         { *m = SendMailRequest{} }
func (m *SendMailRequest) String() string { return proto.CompactTextString(m) }
func (*SendMailRequest) ProtoMessage()    {}
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_mail_d63ba60ec7fd5a75, []int{0}
}
func (m *SendMailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMailRequest.Unmarshal(m, b)
}
func (m *SendMailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMailRequest.Marshal(b, m, deterministic)
}
func (dst *SendMailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMailRequest.Merge(dst, src)
}
func (m *SendMailRequest) XXX_Size() int {
	return xxx_messageInfo_SendMailRequest.Size(m)
}
func (m *SendMailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMailRequest proto.InternalMessageInfo

func (m *SendMailRequest) GetBase() *BaseRequest {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *SendMailRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendMailRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendMailRequest) GetToMultiple() []string {
	if m != nil {
		return m.ToMultiple
	}
	return nil
}

func (m *SendMailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendMailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendMailRequest) GetIsHtml() bool {
	if m != nil {
		return m.IsHtml
	}
	return false
}

func init() {
	proto.RegisterType((*SendMailRequest)(nil), "pb.SendMailRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MailServiceClient is the client API for MailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MailServiceClient interface {
	SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mailServiceClient struct {
	cc *grpc.ClientConn
}

func NewMailServiceClient(cc *grpc.ClientConn) MailServiceClient {
	return &mailServiceClient{cc}
}

func (c *mailServiceClient) SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.MailService/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServiceServer is the server API for MailService service.
type MailServiceServer interface {
	SendMail(context.Context, *SendMailRequest) (*Empty, error)
}

func RegisterMailServiceServer(s *grpc.Server, srv MailServiceServer) {
	s.RegisterService(&_MailService_serviceDesc, srv)
}

func _MailService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MailService/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendMail(ctx, req.(*SendMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MailService",
	HandlerType: (*MailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _MailService_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail.proto",
}

func init() { proto.RegisterFile("mail.proto", fileDescriptor_mail_d63ba60ec7fd5a75) }

var fileDescriptor_mail_d63ba60ec7fd5a75 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0x49, 0x1a, 0x92, 0xe6, 0x05, 0x51, 0xe9, 0x31, 0x60, 0x75, 0x21, 0x2a, 0x4b, 0x06,
	0x94, 0xa1, 0x8c, 0x6c, 0x48, 0x48, 0x2c, 0x5d, 0xdc, 0x1f, 0x50, 0xd9, 0xad, 0x01, 0x23, 0x1b,
	0x9b, 0xf8, 0x05, 0xa9, 0x7f, 0x8f, 0x5f, 0x86, 0x6c, 0xd2, 0xa5, 0xdb, 0xf9, 0x3b, 0xf9, 0x74,
	0xf7, 0x00, 0xac, 0xd0, 0xa6, 0xf7, 0x83, 0x23, 0x87, 0xb9, 0x97, 0xcb, 0xab, 0x77, 0xe3, 0xa4,
	0x98, 0xc8, 0xea, 0x37, 0x83, 0xc5, 0x56, 0x7d, 0x1d, 0x36, 0x42, 0x1b, 0xae, 0xbe, 0x47, 0x15,
	0x08, 0xef, 0xa1, 0x90, 0x22, 0x28, 0x96, 0xb5, 0x59, 0xd7, 0xac, 0x17, 0xbd, 0x97, 0xfd, 0xb3,
	0x08, 0x6a, 0xb2, 0x79, 0x32, 0x11, 0xa1, 0x78, 0x1b, 0x9c, 0x65, 0x79, 0x9b, 0x75, 0x35, 0x4f,
	0x1a, 0xaf, 0x21, 0x27, 0xc7, 0x66, 0x89, 0xe4, 0xe4, 0xf0, 0x0e, 0x1a, 0x72, 0x3b, 0x3b, 0x1a,
	0xd2, 0xde, 0x28, 0x56, 0xb4, 0xb3, 0xae, 0xe6, 0x40, 0x6e, 0x33, 0x11, 0x64, 0x50, 0x85, 0x51,
	0x7e, 0xaa, 0x3d, 0xb1, 0xcb, 0xf4, 0xeb, 0xf4, 0x8c, 0xf1, 0xd2, 0x1d, 0x8e, 0xac, 0xfc, 0x8f,
	0x8f, 0x1a, 0x6f, 0xa1, 0xd2, 0x61, 0xf7, 0x41, 0xd6, 0xb0, 0xaa, 0xcd, 0xba, 0x39, 0x2f, 0x75,
	0x78, 0x25, 0x6b, 0xd6, 0x4f, 0xd0, 0xc4, 0xfe, 0x5b, 0x35, 0xfc, 0xe8, 0xbd, 0xc2, 0x07, 0x98,
	0x9f, 0x26, 0xe1, 0x4d, 0x6c, 0x7f, 0x36, 0x70, 0x59, 0x47, 0xf8, 0x62, 0x3d, 0x1d, 0x57, 0x17,
	0xb2, 0x4c, 0x87, 0x78, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xec, 0x0e, 0x1c, 0x28, 0x01,
	0x00, 0x00,
}
