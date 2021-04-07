// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/buffer/v3/buffer.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *types.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9fce05f80b94cf, []int{0}
}
func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9fce05f80b94cf, []int{1}
}
func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferPerRoute.Unmarshal(m, b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return xxx_messageInfo_BufferPerRoute.Size(m)
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	Equal(interface{}) bool
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof" json:"disabled,omitempty"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof" json:"buffer,omitempty"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.extensions.filters.http.buffer.v3.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.extensions.filters.http.buffer.v3.BufferPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/buffer/v3/buffer.proto", fileDescriptor_7e9fce05f80b94cf)
}

var fileDescriptor_7e9fce05f80b94cf = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x8a, 0x14, 0x31,
	0x10, 0x9d, 0xf4, 0x8e, 0x63, 0x93, 0x05, 0x67, 0x6d, 0x04, 0x97, 0x45, 0x96, 0x65, 0x41, 0x14,
	0xc1, 0x04, 0x76, 0xf0, 0x07, 0x72, 0x5a, 0x3d, 0x2d, 0x2d, 0x0a, 0x7a, 0x70, 0x49, 0xef, 0x54,
	0x67, 0x33, 0x66, 0xba, 0x62, 0x92, 0x6e, 0x7b, 0x7e, 0xc1, 0x7f, 0xf0, 0xee, 0xc9, 0xb3, 0xdf,
	0xe3, 0x3f, 0x08, 0x32, 0x27, 0xe9, 0xa4, 0x47, 0xaf, 0x83, 0xb7, 0x97, 0x57, 0xf5, 0xea, 0xbd,
	0xae, 0x2e, 0xaa, 0x94, 0x0e, 0xb7, 0x6d, 0xc5, 0x6e, 0x70, 0xcd, 0x3d, 0x1a, 0x7c, 0xae, 0x91,
	0x2b, 0x83, 0xc8, 0xad, 0xc3, 0x15, 0xdc, 0x04, 0x9f, 0x5e, 0xd2, 0x6a, 0x0e, 0x7d, 0x00, 0xd7,
	0x48, 0xc3, 0xa1, 0xe9, 0x70, 0x13, 0x9f, 0x8d, 0xd7, 0xd8, 0x78, 0x5e, 0x6b, 0x13, 0xc0, 0x79,
	0x7e, 0x1b, 0x82, 0xe5, 0x55, 0x5b, 0xd7, 0xe0, 0x78, 0xb7, 0x18, 0x11, 0xb3, 0x0e, 0x03, 0x16,
	0x4f, 0xa2, 0x8a, 0xfd, 0x53, 0xb1, 0x51, 0xc5, 0x06, 0x15, 0x1b, 0x7b, 0xbb, 0xc5, 0xc9, 0xa9,
	0x42, 0x54, 0x06, 0x78, 0x94, 0x55, 0x6d, 0xcd, 0x3f, 0x3b, 0x69, 0xed, 0xd0, 0x18, 0x99, 0x93,
	0x87, 0x9d, 0x34, 0x7a, 0x29, 0x03, 0xf0, 0x1d, 0x18, 0x0b, 0x0f, 0x14, 0x2a, 0x8c, 0x90, 0x0f,
	0x68, 0x64, 0x0b, 0xe8, 0x43, 0x22, 0xa1, 0x0f, 0x89, 0x3b, 0xd7, 0x74, 0x26, 0xa2, 0x5f, 0xf1,
	0x8e, 0xde, 0x5f, 0xcb, 0xfe, 0xda, 0xc1, 0xa7, 0x16, 0x7c, 0xb8, 0xae, 0x36, 0x01, 0xfc, 0x31,
	0x39, 0x23, 0x4f, 0x0f, 0x2f, 0x1e, 0xb1, 0x14, 0x84, 0xed, 0x82, 0xb0, 0x37, 0x2f, 0x9b, 0xb0,
	0xb8, 0x78, 0x2b, 0x4d, 0x0b, 0x62, 0xbe, 0x15, 0xd3, 0x67, 0xd9, 0xd9, 0x64, 0x2b, 0xee, 0x7c,
	0x21, 0xd9, 0x11, 0x29, 0xe7, 0x6b, 0xd9, 0x97, 0x69, 0x8c, 0x18, 0xa6, 0xbc, 0x9a, 0xe6, 0xd9,
	0xd1, 0xc1, 0xf9, 0x57, 0x42, 0xef, 0x25, 0xaf, 0x2b, 0x70, 0x25, 0xb6, 0x01, 0x8a, 0xc7, 0x34,
	0x5f, 0x6a, 0x2f, 0x2b, 0x03, 0xcb, 0x68, 0x95, 0x8b, 0xbb, 0x5b, 0x31, 0x5d, 0x65, 0x39, 0xb9,
	0x9c, 0x94, 0x7f, 0x4b, 0xc5, 0x6b, 0x3a, 0x4b, 0x4b, 0x39, 0xce, 0x62, 0x1e, 0xce, 0xf6, 0xdc,
	0x20, 0x4b, 0x7e, 0x22, 0xdf, 0x65, 0xbb, 0x9c, 0x94, 0xe3, 0x28, 0x31, 0xa7, 0x39, 0x76, 0xe0,
	0x9c, 0x5e, 0x42, 0x71, 0xf0, 0x5b, 0x10, 0xf1, 0x9d, 0xfc, 0xf8, 0x35, 0x25, 0xdf, 0x7e, 0x9e,
	0x12, 0xfa, 0x42, 0x63, 0xb2, 0xb0, 0x0e, 0xfb, 0xcd, 0xbe, 0x6e, 0xe2, 0x70, 0xfc, 0xbc, 0x61,
	0x4b, 0x57, 0xe4, 0xfd, 0x87, 0xfd, 0x0e, 0xca, 0x7e, 0x54, 0xff, 0x75, 0x54, 0xd5, 0x2c, 0xfe,
	0x8e, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xbc, 0x05, 0xdb, 0xb9, 0x02, 0x00, 0x00,
}

func (this *Buffer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Buffer)
	if !ok {
		that2, ok := that.(Buffer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MaxRequestBytes.Equal(that1.MaxRequestBytes) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BufferPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute)
	if !ok {
		that2, ok := that.(BufferPerRoute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Override == nil {
		if this.Override != nil {
			return false
		}
	} else if this.Override == nil {
		return false
	} else if !this.Override.Equal(that1.Override) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BufferPerRoute_Disabled) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute_Disabled)
	if !ok {
		that2, ok := that.(BufferPerRoute_Disabled)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	return true
}
func (this *BufferPerRoute_Buffer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute_Buffer)
	if !ok {
		that2, ok := that.(BufferPerRoute_Buffer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Buffer.Equal(that1.Buffer) {
		return false
	}
	return true
}
