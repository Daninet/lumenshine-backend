// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global.proto

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

type EmailContentType int32

const (
	EmailContentType_text EmailContentType = 0
	EmailContentType_html EmailContentType = 1
)

var EmailContentType_name = map[int32]string{
	0: "text",
	1: "html",
}
var EmailContentType_value = map[string]int32{
	"text": 0,
	"html": 1,
}

func (x EmailContentType) String() string {
	return proto.EnumName(EmailContentType_name, int32(x))
}
func (EmailContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{0}
}

type NotificationType int32

const (
	NotificationType_ios     NotificationType = 0
	NotificationType_android NotificationType = 1
	NotificationType_mail    NotificationType = 2
)

var NotificationType_name = map[int32]string{
	0: "ios",
	1: "android",
	2: "mail",
}
var NotificationType_value = map[string]int32{
	"ios":     0,
	"android": 1,
	"mail":    2,
}

func (x NotificationType) String() string {
	return proto.EnumName(NotificationType_name, int32(x))
}
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{1}
}

type NotificationStatusCode int32

const (
	NotificationStatusCode_success NotificationStatusCode = 0
	NotificationStatusCode_error   NotificationStatusCode = 1
)

var NotificationStatusCode_name = map[int32]string{
	0: "success",
	1: "error",
}
var NotificationStatusCode_value = map[string]int32{
	"success": 0,
	"error":   1,
}

func (x NotificationStatusCode) String() string {
	return proto.EnumName(NotificationStatusCode_name, int32(x))
}
func (NotificationStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{2}
}

type DeviceType int32

const (
	DeviceType_apple  DeviceType = 0
	DeviceType_google DeviceType = 1
)

var DeviceType_name = map[int32]string{
	0: "apple",
	1: "google",
}
var DeviceType_value = map[string]int32{
	"apple":  0,
	"google": 1,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{3}
}

type DocumentType int32

const (
	DocumentType_passport           DocumentType = 0
	DocumentType_drivers_license    DocumentType = 1
	DocumentType_id_card            DocumentType = 2
	DocumentType_proof_of_residence DocumentType = 3
)

var DocumentType_name = map[int32]string{
	0: "passport",
	1: "drivers_license",
	2: "id_card",
	3: "proof_of_residence",
}
var DocumentType_value = map[string]int32{
	"passport":           0,
	"drivers_license":    1,
	"id_card":            2,
	"proof_of_residence": 3,
}

func (x DocumentType) String() string {
	return proto.EnumName(DocumentType_name, int32(x))
}
func (DocumentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{4}
}

type DocumentFormat int32

const (
	DocumentFormat_png  DocumentFormat = 0
	DocumentFormat_pdf  DocumentFormat = 1
	DocumentFormat_jpg  DocumentFormat = 2
	DocumentFormat_jpeg DocumentFormat = 3
)

var DocumentFormat_name = map[int32]string{
	0: "png",
	1: "pdf",
	2: "jpg",
	3: "jpeg",
}
var DocumentFormat_value = map[string]int32{
	"png":  0,
	"pdf":  1,
	"jpg":  2,
	"jpeg": 3,
}

func (x DocumentFormat) String() string {
	return proto.EnumName(DocumentFormat_name, int32(x))
}
func (DocumentFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{5}
}

type DocumentSide int32

const (
	DocumentSide_front DocumentSide = 0
	DocumentSide_back  DocumentSide = 1
)

var DocumentSide_name = map[int32]string{
	0: "front",
	1: "back",
}
var DocumentSide_value = map[string]int32{
	"front": 0,
	"back":  1,
}

func (x DocumentSide) String() string {
	return proto.EnumName(DocumentSide_name, int32(x))
}
func (DocumentSide) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{6}
}

type BaseRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	UpdateBy             string   `protobuf:"bytes,2,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{0}
}
func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (dst *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(dst, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type IDRequest struct {
	Base                 *BaseRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Id                   int64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{1}
}
func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (dst *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(dst, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetBase() *BaseRequest {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *IDRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type IDResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDResponse) Reset()         { *m = IDResponse{} }
func (m *IDResponse) String() string { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()    {}
func (*IDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{2}
}
func (m *IDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDResponse.Unmarshal(m, b)
}
func (m *IDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDResponse.Marshal(b, m, deterministic)
}
func (dst *IDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDResponse.Merge(dst, src)
}
func (m *IDResponse) XXX_Size() int {
	return xxx_messageInfo_IDResponse.Size(m)
}
func (m *IDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDResponse proto.InternalMessageInfo

func (m *IDResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type KeyRequest struct {
	Base                 *BaseRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{3}
}
func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRequest.Unmarshal(m, b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
}
func (dst *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(dst, src)
}
func (m *KeyRequest) XXX_Size() int {
	return xxx_messageInfo_KeyRequest.Size(m)
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetBase() *BaseRequest {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *KeyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type StringResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringResponse) Reset()         { *m = StringResponse{} }
func (m *StringResponse) String() string { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()    {}
func (*StringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{4}
}
func (m *StringResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringResponse.Unmarshal(m, b)
}
func (m *StringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringResponse.Marshal(b, m, deterministic)
}
func (dst *StringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringResponse.Merge(dst, src)
}
func (m *StringResponse) XXX_Size() int {
	return xxx_messageInfo_StringResponse.Size(m)
}
func (m *StringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringResponse proto.InternalMessageInfo

func (m *StringResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BoolResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolResponse) Reset()         { *m = BoolResponse{} }
func (m *BoolResponse) String() string { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()    {}
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{5}
}
func (m *BoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolResponse.Unmarshal(m, b)
}
func (m *BoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolResponse.Marshal(b, m, deterministic)
}
func (dst *BoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolResponse.Merge(dst, src)
}
func (m *BoolResponse) XXX_Size() int {
	return xxx_messageInfo_BoolResponse.Size(m)
}
func (m *BoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BoolResponse proto.InternalMessageInfo

func (m *BoolResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type IntResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntResponse) Reset()         { *m = IntResponse{} }
func (m *IntResponse) String() string { return proto.CompactTextString(m) }
func (*IntResponse) ProtoMessage()    {}
func (*IntResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{6}
}
func (m *IntResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntResponse.Unmarshal(m, b)
}
func (m *IntResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntResponse.Marshal(b, m, deterministic)
}
func (dst *IntResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntResponse.Merge(dst, src)
}
func (m *IntResponse) XXX_Size() int {
	return xxx_messageInfo_IntResponse.Size(m)
}
func (m *IntResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntResponse proto.InternalMessageInfo

func (m *IntResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Empty struct {
	Base                 *BaseRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{7}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetBase() *BaseRequest {
	if m != nil {
		return m.Base
	}
	return nil
}

type KeyValueRequest struct {
	Base                 *BaseRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KeyValueRequest) Reset()         { *m = KeyValueRequest{} }
func (m *KeyValueRequest) String() string { return proto.CompactTextString(m) }
func (*KeyValueRequest) ProtoMessage()    {}
func (*KeyValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_32df9260b6044f96, []int{8}
}
func (m *KeyValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValueRequest.Unmarshal(m, b)
}
func (m *KeyValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValueRequest.Marshal(b, m, deterministic)
}
func (dst *KeyValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueRequest.Merge(dst, src)
}
func (m *KeyValueRequest) XXX_Size() int {
	return xxx_messageInfo_KeyValueRequest.Size(m)
}
func (m *KeyValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueRequest proto.InternalMessageInfo

func (m *KeyValueRequest) GetBase() *BaseRequest {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *KeyValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValueRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "pb.BaseRequest")
	proto.RegisterType((*IDRequest)(nil), "pb.IDRequest")
	proto.RegisterType((*IDResponse)(nil), "pb.IDResponse")
	proto.RegisterType((*KeyRequest)(nil), "pb.KeyRequest")
	proto.RegisterType((*StringResponse)(nil), "pb.StringResponse")
	proto.RegisterType((*BoolResponse)(nil), "pb.BoolResponse")
	proto.RegisterType((*IntResponse)(nil), "pb.IntResponse")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*KeyValueRequest)(nil), "pb.KeyValueRequest")
	proto.RegisterEnum("pb.EmailContentType", EmailContentType_name, EmailContentType_value)
	proto.RegisterEnum("pb.NotificationType", NotificationType_name, NotificationType_value)
	proto.RegisterEnum("pb.NotificationStatusCode", NotificationStatusCode_name, NotificationStatusCode_value)
	proto.RegisterEnum("pb.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterEnum("pb.DocumentType", DocumentType_name, DocumentType_value)
	proto.RegisterEnum("pb.DocumentFormat", DocumentFormat_name, DocumentFormat_value)
	proto.RegisterEnum("pb.DocumentSide", DocumentSide_name, DocumentSide_value)
}

func init() { proto.RegisterFile("global.proto", fileDescriptor_global_32df9260b6044f96) }

var fileDescriptor_global_32df9260b6044f96 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0xa4, 0x38, 0xb1, 0xc6, 0x26, 0x59, 0xd4, 0x12, 0x02, 0x6d, 0xa1, 0xc8, 0x25, 0x14,
	0x51, 0x4c, 0x49, 0xe9, 0xbd, 0xd8, 0x4e, 0xc1, 0x04, 0x7a, 0xb0, 0x4b, 0x8e, 0x15, 0x2b, 0xed,
	0x58, 0xdd, 0x44, 0xd6, 0x6c, 0x77, 0x57, 0xa6, 0xfa, 0xfb, 0xb2, 0xb2, 0xe5, 0xf8, 0x90, 0x43,
	0x4b, 0x6f, 0x4f, 0x33, 0xef, 0xbd, 0x79, 0x4f, 0xb0, 0x30, 0x2a, 0x4a, 0xca, 0x78, 0x39, 0x51,
	0x9a, 0x2c, 0x45, 0xbe, 0xca, 0xe2, 0x05, 0x0c, 0xa7, 0xdc, 0xe0, 0x12, 0x7f, 0xd5, 0x68, 0x6c,
	0xf4, 0x06, 0x40, 0xef, 0x60, 0x2a, 0xc5, 0x95, 0xf7, 0xd6, 0x7b, 0x1f, 0x2e, 0xc3, 0xfd, 0x64,
	0x21, 0xa2, 0x57, 0x10, 0xd6, 0x4a, 0x70, 0x8b, 0x69, 0xd6, 0x5c, 0xf9, 0xed, 0x76, 0xb0, 0x1b,
	0x4c, 0x9b, 0xf8, 0x0b, 0x84, 0x8b, 0x79, 0x67, 0x34, 0x86, 0x93, 0x8c, 0x1b, 0x6c, 0x2d, 0x86,
	0x37, 0x17, 0x13, 0x95, 0x4d, 0x8e, 0xee, 0x2c, 0xdb, 0x65, 0x74, 0x0e, 0xbe, 0x14, 0xad, 0x4f,
	0xb0, 0xf4, 0xa5, 0x88, 0x5f, 0x03, 0x38, 0x07, 0xa3, 0xa8, 0x3a, 0x6c, 0xbd, 0xc3, 0x76, 0x06,
	0x70, 0x87, 0xcd, 0x3f, 0x1d, 0x60, 0x10, 0x3c, 0x62, 0x97, 0xd4, 0xc1, 0xf8, 0x1a, 0xce, 0x57,
	0x56, 0xcb, 0xaa, 0x38, 0x9c, 0x79, 0x09, 0xfd, 0x2d, 0x2f, 0x6b, 0xdc, 0xb7, 0xdd, 0x7d, 0xc4,
	0xef, 0x60, 0x34, 0x25, 0x2a, 0x9f, 0x67, 0x0d, 0x3a, 0xd6, 0x18, 0x86, 0x8b, 0xca, 0x3e, 0x4f,
	0x0a, 0x3a, 0xd2, 0x07, 0xe8, 0xdf, 0x6e, 0x94, 0x6d, 0xfe, 0x2a, 0x72, 0xfc, 0x03, 0x2e, 0xee,
	0xb0, 0xb9, 0x77, 0xca, 0xff, 0xab, 0xfa, 0x94, 0x26, 0x38, 0x2a, 0x96, 0x5c, 0x03, 0xbb, 0xdd,
	0x70, 0x59, 0xce, 0xa8, 0xb2, 0x58, 0xd9, 0xef, 0x8d, 0xc2, 0x68, 0x00, 0x27, 0x16, 0x7f, 0x5b,
	0xd6, 0x73, 0xe8, 0xa7, 0xdd, 0x94, 0xcc, 0x4b, 0x6e, 0x80, 0x7d, 0x23, 0x2b, 0xd7, 0x32, 0xe7,
	0x56, 0x52, 0xd5, 0xf2, 0xce, 0x20, 0x90, 0x64, 0x58, 0x2f, 0x1a, 0xc2, 0x19, 0xaf, 0x84, 0x26,
	0x29, 0x98, 0xe7, 0x34, 0xce, 0x90, 0xf9, 0xc9, 0x47, 0xb8, 0x3c, 0xd6, 0xac, 0x2c, 0xb7, 0xb5,
	0x99, 0x91, 0x40, 0x27, 0x30, 0x75, 0x9e, 0xa3, 0x71, 0xea, 0x10, 0xfa, 0xa8, 0x35, 0x69, 0xe6,
	0x25, 0x63, 0x80, 0x39, 0x6e, 0x65, 0x8e, 0xad, 0x7f, 0x08, 0x7d, 0xae, 0x54, 0x89, 0xac, 0x17,
	0x01, 0x9c, 0x16, 0x44, 0x45, 0x89, 0xcc, 0x4b, 0xee, 0x61, 0x34, 0xa7, 0xbc, 0xde, 0x74, 0x71,
	0x47, 0x30, 0x50, 0xdc, 0x18, 0x45, 0xda, 0x45, 0x7e, 0x01, 0x17, 0x42, 0xcb, 0x2d, 0x6a, 0x93,
	0x96, 0x32, 0xc7, 0xca, 0x20, 0xf3, 0xdc, 0x3d, 0x29, 0xd2, 0x9c, 0x6b, 0xc1, 0xfc, 0xe8, 0x12,
	0x22, 0xa5, 0x89, 0xd6, 0x29, 0xad, 0x53, 0x8d, 0x46, 0x0a, 0xac, 0x72, 0x64, 0x41, 0xf2, 0x19,
	0xce, 0x3b, 0xdf, 0xaf, 0xa4, 0x37, 0xdc, 0xba, 0x82, 0xaa, 0x2a, 0x58, 0xaf, 0x05, 0x62, 0xcd,
	0x3c, 0x07, 0x1e, 0x54, 0xc1, 0x7c, 0xd7, 0xf2, 0x41, 0x61, 0xc1, 0x82, 0x64, 0xfc, 0x14, 0x67,
	0x25, 0x45, 0x9b, 0x7a, 0xad, 0xa9, 0xda, 0xff, 0xbe, 0x8c, 0xe7, 0x8f, 0xcc, 0xcb, 0x4e, 0xdb,
	0x27, 0xf6, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x91, 0x92, 0x3c, 0x72, 0x03, 0x00,
	0x00,
}
