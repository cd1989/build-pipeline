// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.golang.org/appengine/internal/channel/channel_service.proto

/*
Package channel is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/channel/channel_service.proto

It has these top-level messages:
	ChannelServiceError
	CreateChannelRequest
	CreateChannelResponse
	SendMessageRequest
*/
package channel

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

type ChannelServiceError_ErrorCode int32

const (
	ChannelServiceError_OK                             ChannelServiceError_ErrorCode = 0
	ChannelServiceError_INTERNAL_ERROR                 ChannelServiceError_ErrorCode = 1
	ChannelServiceError_INVALID_CHANNEL_KEY            ChannelServiceError_ErrorCode = 2
	ChannelServiceError_BAD_MESSAGE                    ChannelServiceError_ErrorCode = 3
	ChannelServiceError_INVALID_CHANNEL_TOKEN_DURATION ChannelServiceError_ErrorCode = 4
	ChannelServiceError_APPID_ALIAS_REQUIRED           ChannelServiceError_ErrorCode = 5
)

var ChannelServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "INVALID_CHANNEL_KEY",
	3: "BAD_MESSAGE",
	4: "INVALID_CHANNEL_TOKEN_DURATION",
	5: "APPID_ALIAS_REQUIRED",
}
var ChannelServiceError_ErrorCode_value = map[string]int32{
	"OK":                             0,
	"INTERNAL_ERROR":                 1,
	"INVALID_CHANNEL_KEY":            2,
	"BAD_MESSAGE":                    3,
	"INVALID_CHANNEL_TOKEN_DURATION": 4,
	"APPID_ALIAS_REQUIRED":           5,
}

func (x ChannelServiceError_ErrorCode) Enum() *ChannelServiceError_ErrorCode {
	p := new(ChannelServiceError_ErrorCode)
	*p = x
	return p
}
func (x ChannelServiceError_ErrorCode) String() string {
	return proto.EnumName(ChannelServiceError_ErrorCode_name, int32(x))
}
func (x *ChannelServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChannelServiceError_ErrorCode_value, data, "ChannelServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ChannelServiceError_ErrorCode(value)
	return nil
}
func (ChannelServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ChannelServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChannelServiceError) Reset()                    { *m = ChannelServiceError{} }
func (m *ChannelServiceError) String() string            { return proto.CompactTextString(m) }
func (*ChannelServiceError) ProtoMessage()               {}
func (*ChannelServiceError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateChannelRequest struct {
	ApplicationKey   *string `protobuf:"bytes,1,req,name=application_key,json=applicationKey" json:"application_key,omitempty"`
	DurationMinutes  *int32  `protobuf:"varint,2,opt,name=duration_minutes,json=durationMinutes" json:"duration_minutes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateChannelRequest) Reset()                    { *m = CreateChannelRequest{} }
func (m *CreateChannelRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelRequest) ProtoMessage()               {}
func (*CreateChannelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateChannelRequest) GetApplicationKey() string {
	if m != nil && m.ApplicationKey != nil {
		return *m.ApplicationKey
	}
	return ""
}

func (m *CreateChannelRequest) GetDurationMinutes() int32 {
	if m != nil && m.DurationMinutes != nil {
		return *m.DurationMinutes
	}
	return 0
}

type CreateChannelResponse struct {
	Token            *string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	DurationMinutes  *int32  `protobuf:"varint,3,opt,name=duration_minutes,json=durationMinutes" json:"duration_minutes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateChannelResponse) Reset()                    { *m = CreateChannelResponse{} }
func (m *CreateChannelResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelResponse) ProtoMessage()               {}
func (*CreateChannelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateChannelResponse) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *CreateChannelResponse) GetDurationMinutes() int32 {
	if m != nil && m.DurationMinutes != nil {
		return *m.DurationMinutes
	}
	return 0
}

type SendMessageRequest struct {
	ApplicationKey   *string `protobuf:"bytes,1,req,name=application_key,json=applicationKey" json:"application_key,omitempty"`
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SendMessageRequest) Reset()                    { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()               {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SendMessageRequest) GetApplicationKey() string {
	if m != nil && m.ApplicationKey != nil {
		return *m.ApplicationKey
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelServiceError)(nil), "appengine.ChannelServiceError")
	proto.RegisterType((*CreateChannelRequest)(nil), "appengine.CreateChannelRequest")
	proto.RegisterType((*CreateChannelResponse)(nil), "appengine.CreateChannelResponse")
	proto.RegisterType((*SendMessageRequest)(nil), "appengine.SendMessageRequest")
}

func init() {
	proto.RegisterFile("google.golang.org/appengine/internal/channel/channel_service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0xee, 0xd2, 0x40,
	0x14, 0xc5, 0x6d, 0xff, 0x22, 0xe9, 0x35, 0x81, 0x66, 0xc0, 0xd8, 0x95, 0x21, 0xdd, 0x88, 0x1b,
	0x78, 0x86, 0xa1, 0x9d, 0x68, 0xd3, 0xd2, 0xe2, 0x14, 0xfc, 0xda, 0x4c, 0x26, 0x70, 0x53, 0x2b,
	0x65, 0xa6, 0x4e, 0x8b, 0x09, 0x4f, 0xe1, 0x63, 0xf8, 0x9a, 0x26, 0x14, 0x88, 0x21, 0x6c, 0x5c,
	0xcd, 0x9c, 0x93, 0xdf, 0x39, 0x33, 0x37, 0x17, 0x16, 0x85, 0xd6, 0x45, 0x85, 0xb3, 0x42, 0x57,
	0x52, 0x15, 0x33, 0x6d, 0x8a, 0xb9, 0xac, 0x6b, 0x54, 0x45, 0xa9, 0x70, 0x5e, 0xaa, 0x16, 0x8d,
	0x92, 0xd5, 0x7c, 0xfb, 0x5d, 0x2a, 0x85, 0xb7, 0x53, 0x34, 0x68, 0x7e, 0x95, 0x5b, 0x9c, 0xd5,
	0x46, 0xb7, 0x9a, 0x38, 0xb7, 0x84, 0xff, 0xc7, 0x82, 0x51, 0xd0, 0x41, 0x79, 0xc7, 0x30, 0x63,
	0xb4, 0xf1, 0x7f, 0x5b, 0xe0, 0x9c, 0x6f, 0x81, 0xde, 0x21, 0x79, 0x01, 0x76, 0x16, 0xbb, 0xcf,
	0x08, 0x81, 0x41, 0x94, 0xae, 0x19, 0x4f, 0x69, 0x22, 0x18, 0xe7, 0x19, 0x77, 0x2d, 0xf2, 0x1a,
	0x46, 0x51, 0xfa, 0x89, 0x26, 0x51, 0x28, 0x82, 0x0f, 0x34, 0x4d, 0x59, 0x22, 0x62, 0xf6, 0xd5,
	0xb5, 0xc9, 0x10, 0x5e, 0x2e, 0x68, 0x28, 0x96, 0x2c, 0xcf, 0xe9, 0x7b, 0xe6, 0x3e, 0x11, 0x1f,
	0xde, 0xdc, 0x93, 0xeb, 0x2c, 0x66, 0xa9, 0x08, 0x37, 0x9c, 0xae, 0xa3, 0x2c, 0x75, 0x9f, 0x13,
	0x0f, 0xc6, 0x74, 0xb5, 0x8a, 0x42, 0x41, 0x93, 0x88, 0xe6, 0x82, 0xb3, 0x8f, 0x9b, 0x88, 0xb3,
	0xd0, 0xed, 0xf9, 0x3f, 0x60, 0x1c, 0x18, 0x94, 0x2d, 0x5e, 0xbe, 0xcb, 0xf1, 0xe7, 0x11, 0x9b,
	0x96, 0xbc, 0x85, 0xa1, 0xac, 0xeb, 0xaa, 0xdc, 0xca, 0xb6, 0xd4, 0x4a, 0xec, 0xf1, 0xe4, 0x59,
	0x13, 0x7b, 0xea, 0xf0, 0xc1, 0x3f, 0x76, 0x8c, 0x27, 0xf2, 0x0e, 0xdc, 0xdd, 0xd1, 0x74, 0xd4,
	0xa1, 0x54, 0xc7, 0x16, 0x1b, 0xcf, 0x9e, 0x58, 0xd3, 0x1e, 0x1f, 0x5e, 0xfd, 0x65, 0x67, 0xfb,
	0x5f, 0xe0, 0xd5, 0xdd, 0x5b, 0x4d, 0xad, 0x55, 0x83, 0x64, 0x0c, 0xbd, 0x56, 0xef, 0x51, 0x9d,
	0x83, 0x0e, 0xef, 0xc4, 0xc3, 0xe6, 0xa7, 0xc7, 0xcd, 0x9f, 0x81, 0xe4, 0xa8, 0x76, 0x4b, 0x6c,
	0x1a, 0x59, 0xe0, 0x7f, 0xcf, 0xe0, 0x41, 0xff, 0xd0, 0x45, 0x3d, 0xfb, 0x0c, 0x5c, 0xe5, 0xc2,
	0xf9, 0xd6, 0xbf, 0x2c, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x0a, 0x77, 0x06, 0x23,
	0x02, 0x00, 0x00,
}