package tiedie

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_UNKNOWN_TYPE        Type = 0
	Type_CONNECT_REQUEST     Type = 1
	Type_DISCONNECT_REQUEST  Type = 2
	Type_CONNECT_RESPONSE    Type = 3
	Type_DISCONNECT_RESPONSE Type = 4
)

var Type_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	1: "CONNECT_REQUEST",
	2: "DISCONNECT_REQUEST",
	3: "CONNECT_RESPONSE",
	4: "DISCONNECT_RESPONSE",
}

var Type_value = map[string]int32{
	"UNKNOWN_TYPE":        0,
	"CONNECT_REQUEST":     1,
	"DISCONNECT_REQUEST":  2,
	"CONNECT_RESPONSE":    3,
	"DISCONNECT_RESPONSE": 4,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{0}
}

type Request struct {
	Type                 Type               `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	ConnectRequest       *ConnectRequest    `protobuf:"bytes,2,opt,name=connect_request,json=connectRequest,proto3" json:"connect_request,omitempty"`
	DisconnectRequest    *DisconnectRequest `protobuf:"bytes,3,opt,name=disconnect_request,json=disconnectRequest,proto3" json:"disconnect_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`

	XXX_sizecache int32 `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Type {
	if m != nil {
		return m.Type
	}

	return Type_UNKNOWN_TYPE
}

func (m *Request) GetConnectRequest() *ConnectRequest {
	if m != nil {
		return m.ConnectRequest
	}
	return nil
}

func (m *Request) GetDisconnectRequest() *DisconnectRequest {
	if m != nil {
		return m.DisconnectRequest
	}
	return nil
}

type Response struct {
	Type                 Type                `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	ConnectResponse      *ConnectResponse    `protobuf:"bytes,2,opt,name=connect_response,json=connectResponse,proto3" json:"connect_response,omitempty"`
	DisconnectResponse   *DisconnectResponse `protobuf:"bytes,3,opt,name=disconnect_response,json=disconnectResponse,proto3" json:"disconnect_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_UNKNOWN_TYPE
}
func (m *Response) GetConnectResponse() *ConnectResponse {
	if m != nil {
		return m.ConnectResponse
	}
	return nil
}

func (m *Response) GetDisconnectResponse() *DisconnectResponse {
	if m != nil {
		return m.DisconnectResponse
	}
	return nil
}

type ConnectRequest struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	RequestConnect       string   `protobuf:"bytes,2,opt,name=request_connect,json=requestConnect,proto3" json:"request_connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{2}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {

	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *ConnectRequest) GetRequestConnect() string {
	if m != nil {
		return m.RequestConnect
	}
	return ""
}

type DisconnectRequest struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	RequestDisconnect    string   `protobuf:"bytes,2,opt,name=request_disconnect,json=requestDisconnect,proto3" json:"request_disconnect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{3}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DisconnectRequest) GetRequestDisconnect() string {
	if m != nil {
		return m.RequestDisconnect
	}
	return ""
}

type ConnectResponse struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	AcknowledgeConnect   string   `protobuf:"bytes,2,opt,name=acknowledge_connect,json=acknowledgeConnect,proto3" json:"acknowledge_connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`

	XXX_sizecache int32 `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{4}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}

	return ""
}

func (m *ConnectResponse) GetAcknowledgeConnect() string {
	if m != nil {
		return m.AcknowledgeConnect
	}
	return ""
}

type DisconnectResponse struct {
	DeviceId              string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	AcknowledgeDisconnect string   `protobuf:"bytes,2,opt,name=acknowledge_disconnect,json=acknowledgeDisconnect,proto3" json:"acknowledge_disconnect,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4e5577304d102ae, []int{5}
}

func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (m *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(m, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DisconnectResponse) GetAcknowledgeDisconnect() string {
	if m != nil {
		return m.AcknowledgeDisconnect
	}
	return ""
}

func init() {
	proto.RegisterEnum("Type", Type_name, Type_value)
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*ConnectRequest)(nil), "ConnectRequest")
	proto.RegisterType((*DisconnectRequest)(nil), "DisconnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "ConnectResponse")
	proto.RegisterType((*DisconnectResponse)(nil), "DisconnectResponse")
}

func init() { proto.RegisterFile("tiedie.proto", fileDescriptor_b4e5577304d102ae) }

var fileDescriptor_b4e5577304d102ae = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6b, 0xa3, 0x40,
	0x14, 0xc5, 0xd7, 0x24, 0xbb, 0x9b, 0xdc, 0x0d, 0x6a, 0xae, 0xbb, 0xd9, 0x2c, 0x0b, 0x4b, 0xf0,
	0x65, 0xc3, 0xc2, 0x2a, 0xa4, 0x14, 0x0a, 0x7d, 0x6a, 0x13, 0x1f, 0x42, 0xc1, 0xa4, 0xa3, 0x69,
	0x69, 0x5f, 0x24, 0x75, 0x86, 0xd6, 0x36, 0x51, 0x1b, 0xed, 0x9f, 0x7c, 0x9c, 0xbe, 0xf5, 0x63,
	0x96, 0xea, 0xa4, 0x46, 0x05, 0x1f, 0xe7, 0xdc, 0x7b, 0xcf, 0xfd, 0xcd, 0x61, 0x06, 0xda, 0xb1,

	0xc7, 0xa8, 0xc7, 0xb4, 0x70, 0x1d, 0xc4, 0x81, 0xfa, 0x22, 0xc0, 0x57, 0xc2, 0xee, 0x1f, 0x58,
	0x14, 0xe3, 0x2f, 0x68, 0xc4, 0x9b, 0x90, 0xf5, 0x84, 0xbe, 0x30, 0x10, 0x87, 0x9f, 0x35, 0x7b,
	0x13, 0x32, 0x92, 0x48, 0x78, 0x00, 0x92, 0x1b, 0xf8, 0x3e, 0x73, 0x63, 0x67, 0x9d, 0x76, 0xf7,
	0x6a, 0x7d, 0x61, 0xf0, 0x6d, 0x28, 0x69, 0xa3, 0x54, 0xe7, 0x26, 0x44, 0x74, 0x73, 0x67, 0x3c,
	0x02, 0xa4, 0x5e, 0x54, 0x1c, 0xae, 0x27, 0xc3, 0xa8, 0x8d, 0x3f, 0x4a, 0xdb, 0xf9, 0x0e, 0x2d,
	0x4a, 0xea, 0xab, 0x00, 0x4d, 0xc2, 0xa2, 0x30, 0xf0, 0x23, 0x56, 0x05, 0x79, 0x08, 0x72, 0xb6,
	0x27, 0x6d, 0xe7, 0x94, 0x72, 0x46, 0x99, 0xea, 0x44, 0x72, 0xf3, 0x02, 0x8e, 0x41, 0xc9, 0x71,
	0xf2, 0xf9, 0x14, 0x54, 0xc9, 0x81, 0x72, 0x0b, 0xa4, 0x25, 0x4d, 0x3d, 0x03, 0x31, 0x9f, 0x07,
	0xfe, 0x86, 0x16, 0x65, 0x8f, 0x9e, 0xcb, 0x1c, 0x8f, 0x26, 0xd0, 0x2d, 0xd2, 0x4c, 0x85, 0x09,
	0xc5, 0xbf, 0x20, 0xf1, 0x44, 0x1c, 0xee, 0x94, 0x00, 0xb7, 0x88, 0xc8, 0x65, 0x6e, 0xa6, 0x3a,
	0xd0, 0x29, 0x45, 0x55, 0x6d, 0xfd, 0x1f, 0x70, 0x6b, 0x9d, 0x71, 0x72, 0xf7, 0x0e, 0xaf, 0x64,
	0x96, 0xaa, 0x03, 0x52, 0x21, 0xa2, 0x6a, 0x7b, 0x1d, 0x94, 0x85, 0x7b, 0xe7, 0x07, 0x4f, 0x4b,
	0x46, 0xaf, 0x59, 0x81, 0x1e, 0x77, 0x4a, 0xdb, 0x1b, 0xdc, 0x00, 0x96, 0x33, 0xac, 0xde, 0xb1,
	0x0f, 0xdd, 0xdd, 0x1d, 0xa5, 0x6b, 0xfc, 0xd8, 0xa9, 0x66, 0xde, 0xff, 0x62, 0x68, 0xbc, 0x3f,
	0x0a, 0x94, 0xa1, 0x3d, 0x37, 0x4f, 0xcc, 0xe9, 0xb9, 0xe9, 0xd8, 0x17, 0x33, 0x43, 0xfe, 0x84,
	0x0a, 0x48, 0xa3, 0xa9, 0x69, 0x1a, 0x23, 0xdb, 0x21, 0xc6, 0xe9, 0xdc, 0xb0, 0x6c, 0x59, 0xc0,
	0x2e, 0xe0, 0x78, 0x62, 0x15, 0xf5, 0x1a, 0x7e, 0x07, 0x39, 0x13, 0xad, 0xd9, 0xd4, 0xb4, 0x0c,
	0xb9, 0x8e, 0x3f, 0x41, 0xc9, 0x75, 0xf3, 0x42, 0xe3, 0xb8, 0x7f, 0xf9, 0x87, 0x3d, 0x2f, 0x56,
	0xe1, 0x92, 0x69, 0x6e, 0xb0, 0xd2, 0xc3, 0x75, 0x70, 0xab, 0x27, 0x5f, 0x2c, 0xd2, 0xd3, 0x0f,
	0x77, 0xf5, 0x25, 0x39, 0xee, 0xbd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x6b, 0x8f, 0xaf, 0x81,
	0x03, 0x00, 0x00,
}
