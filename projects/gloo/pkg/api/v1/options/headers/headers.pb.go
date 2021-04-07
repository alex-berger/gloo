// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/headers/headers.proto

package headers

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
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

// This plugin provides configuration options to append and remove headers from
// requests and responses
// HeaderManipulation can be specified on routes, virtual hosts, or weighted destinations
type HeaderManipulation struct {
	// Specifies a list of HTTP headers that should be added to each request
	// handled by this route or virtual host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,1,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// handled by this route or virtual host.
	RequestHeadersToRemove []string `protobuf:"bytes,2,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response
	// handled by this route or host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	ResponseHeadersToAdd []*HeaderValueOption `protobuf:"bytes,3,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// handled by this route or virtual host.
	ResponseHeadersToRemove []string `protobuf:"bytes,4,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *HeaderManipulation) Reset()         { *m = HeaderManipulation{} }
func (m *HeaderManipulation) String() string { return proto.CompactTextString(m) }
func (*HeaderManipulation) ProtoMessage()    {}
func (*HeaderManipulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0de64b70fd96e8, []int{0}
}
func (m *HeaderManipulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderManipulation.Unmarshal(m, b)
}
func (m *HeaderManipulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderManipulation.Marshal(b, m, deterministic)
}
func (m *HeaderManipulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderManipulation.Merge(m, src)
}
func (m *HeaderManipulation) XXX_Size() int {
	return xxx_messageInfo_HeaderManipulation.Size(m)
}
func (m *HeaderManipulation) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderManipulation.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderManipulation proto.InternalMessageInfo

func (m *HeaderManipulation) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HeaderManipulation) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HeaderManipulation) GetResponseHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *HeaderManipulation) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

// Header name/value pair plus option to control append behavior.
type HeaderValueOption struct {
	// Header name/value pair that this option applies to.
	Header *HeaderValue `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Should the value be appended? If true (default), the value is appended to
	// existing values.
	Append               *types.BoolValue `protobuf:"bytes,2,opt,name=append,proto3" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HeaderValueOption) Reset()         { *m = HeaderValueOption{} }
func (m *HeaderValueOption) String() string { return proto.CompactTextString(m) }
func (*HeaderValueOption) ProtoMessage()    {}
func (*HeaderValueOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0de64b70fd96e8, []int{1}
}
func (m *HeaderValueOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValueOption.Unmarshal(m, b)
}
func (m *HeaderValueOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValueOption.Marshal(b, m, deterministic)
}
func (m *HeaderValueOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValueOption.Merge(m, src)
}
func (m *HeaderValueOption) XXX_Size() int {
	return xxx_messageInfo_HeaderValueOption.Size(m)
}
func (m *HeaderValueOption) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValueOption.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValueOption proto.InternalMessageInfo

func (m *HeaderValueOption) GetHeader() *HeaderValue {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderValueOption) GetAppend() *types.BoolValue {
	if m != nil {
		return m.Append
	}
	return nil
}

// Header name/value pair.
type HeaderValue struct {
	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValue) Reset()         { *m = HeaderValue{} }
func (m *HeaderValue) String() string { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()    {}
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0de64b70fd96e8, []int{2}
}
func (m *HeaderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValue.Unmarshal(m, b)
}
func (m *HeaderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValue.Marshal(b, m, deterministic)
}
func (m *HeaderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValue.Merge(m, src)
}
func (m *HeaderValue) XXX_Size() int {
	return xxx_messageInfo_HeaderValue.Size(m)
}
func (m *HeaderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValue proto.InternalMessageInfo

func (m *HeaderValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeaderValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*HeaderManipulation)(nil), "headers.options.gloo.solo.io.HeaderManipulation")
	proto.RegisterType((*HeaderValueOption)(nil), "headers.options.gloo.solo.io.HeaderValueOption")
	proto.RegisterType((*HeaderValue)(nil), "headers.options.gloo.solo.io.HeaderValue")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/headers/headers.proto", fileDescriptor_fc0de64b70fd96e8)
}

var fileDescriptor_fc0de64b70fd96e8 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xe5, 0xb8, 0x44, 0xca, 0x64, 0x03, 0x43, 0xd4, 0x1a, 0xab, 0xaa, 0xa2, 0x88, 0x45,
	0x58, 0x30, 0x23, 0x8c, 0x58, 0x20, 0x36, 0xb4, 0x2b, 0x54, 0x09, 0x21, 0x59, 0x08, 0x09, 0x36,
	0xd1, 0x38, 0x7e, 0x75, 0x4d, 0x5d, 0xbf, 0x61, 0x66, 0x6c, 0xda, 0x2b, 0x70, 0x12, 0x8e, 0xc0,
	0x2d, 0xb8, 0x03, 0x77, 0x60, 0x8f, 0xe6, 0x4f, 0x22, 0x48, 0x42, 0x95, 0x55, 0xe6, 0xcd, 0x7b,
	0xdf, 0xf7, 0xfd, 0xe2, 0x37, 0xe4, 0xbc, 0xaa, 0xcd, 0x65, 0x57, 0xb0, 0x25, 0x5e, 0x73, 0x8d,
	0x0d, 0x3e, 0xad, 0x91, 0x57, 0x0d, 0x22, 0x97, 0x0a, 0x3f, 0xc3, 0xd2, 0x68, 0x5f, 0x09, 0x59,
	0xf3, 0xfe, 0x19, 0x47, 0x69, 0x6a, 0x6c, 0x35, 0xbf, 0x04, 0x51, 0x82, 0x5a, 0xff, 0x32, 0xa9,
	0xd0, 0x20, 0x3d, 0x5e, 0x95, 0x61, 0x8c, 0x59, 0x29, 0xb3, 0xae, 0xac, 0xc6, 0x74, 0x52, 0x61,
	0x85, 0x6e, 0x90, 0xdb, 0x93, 0xd7, 0xa4, 0x14, 0x6e, 0x8c, 0xbf, 0x84, 0x1b, 0x13, 0xee, 0x8e,
	0xa1, 0xed, 0xf1, 0xd6, 0x67, 0x66, 0x7c, 0x89, 0x0a, 0x78, 0x21, 0x34, 0x84, 0xee, 0x49, 0x85,
	0x58, 0x35, 0xc0, 0x5d, 0x55, 0x74, 0x17, 0xfc, 0xab, 0x12, 0x52, 0xae, 0x29, 0x66, 0x3f, 0x07,
	0x84, 0xbe, 0x71, 0x20, 0x6f, 0x45, 0x5b, 0xcb, 0xae, 0x11, 0x16, 0x86, 0x7e, 0x24, 0x87, 0x0a,
	0xbe, 0x74, 0xa0, 0xcd, 0x22, 0x60, 0x2e, 0x0c, 0x2e, 0x44, 0x59, 0x26, 0xd1, 0x34, 0x9e, 0x8f,
	0xb3, 0xc7, 0xcc, 0xa5, 0x32, 0x21, 0x6b, 0xd6, 0x67, 0xcc, 0xa6, 0x32, 0x6f, 0xf3, 0x41, 0x34,
	0x1d, 0xbc, 0x73, 0x7f, 0x29, 0x7f, 0x18, 0x3c, 0x7c, 0x47, 0xbf, 0xc7, 0xd3, 0xb2, 0xa4, 0x2f,
	0xc9, 0xa3, 0x1d, 0xd6, 0x0a, 0xae, 0xb1, 0x87, 0x64, 0x30, 0x8d, 0xe7, 0xa3, 0xfc, 0x70, 0x53,
	0x97, 0xbb, 0x2e, 0xbd, 0x20, 0x47, 0x0a, 0xb4, 0xc4, 0x56, 0xc3, 0x26, 0x56, 0xec, 0xb0, 0x38,
	0xbb, 0xeb, 0xa3, 0xee, 0x20, 0x9c, 0xac, 0xfc, 0xfe, 0x41, 0x7c, 0x45, 0xd2, 0x5d, 0x39, 0x81,
	0xf1, 0xc0, 0x31, 0x1e, 0x6d, 0x29, 0x3d, 0xe4, 0xec, 0x5b, 0x44, 0x1e, 0x6c, 0x05, 0xd1, 0x53,
	0x32, 0xf4, 0x4e, 0x49, 0x34, 0x8d, 0xe6, 0xe3, 0xec, 0xc9, 0xde, 0xa4, 0x79, 0x10, 0xd2, 0x8c,
	0x0c, 0xed, 0xea, 0xda, 0x32, 0x19, 0x38, 0x8b, 0x94, 0xf9, 0xdd, 0xb2, 0xd5, 0x6e, 0xd9, 0x19,
	0x62, 0x13, 0x34, 0x7e, 0x72, 0xf6, 0x82, 0x8c, 0xff, 0xb2, 0xa2, 0xf7, 0x49, 0x7c, 0x05, 0xb7,
	0x0e, 0x61, 0x94, 0xdb, 0x23, 0x9d, 0x90, 0x7b, 0xbd, 0x6d, 0x39, 0xcf, 0x51, 0xee, 0x8b, 0xb3,
	0xf3, 0x1f, 0xbf, 0x0f, 0xa2, 0xef, 0xbf, 0x4e, 0xa2, 0x4f, 0xaf, 0xf7, 0x7b, 0xf1, 0xf2, 0xaa,
	0xfa, 0xcf, 0xab, 0x2f, 0x86, 0x0e, 0xef, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xb5,
	0xdf, 0xce, 0x3c, 0x03, 0x00, 0x00,
}

func (this *HeaderManipulation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderManipulation)
	if !ok {
		that2, ok := that.(HeaderManipulation)
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
	if len(this.RequestHeadersToAdd) != len(that1.RequestHeadersToAdd) {
		return false
	}
	for i := range this.RequestHeadersToAdd {
		if !this.RequestHeadersToAdd[i].Equal(that1.RequestHeadersToAdd[i]) {
			return false
		}
	}
	if len(this.RequestHeadersToRemove) != len(that1.RequestHeadersToRemove) {
		return false
	}
	for i := range this.RequestHeadersToRemove {
		if this.RequestHeadersToRemove[i] != that1.RequestHeadersToRemove[i] {
			return false
		}
	}
	if len(this.ResponseHeadersToAdd) != len(that1.ResponseHeadersToAdd) {
		return false
	}
	for i := range this.ResponseHeadersToAdd {
		if !this.ResponseHeadersToAdd[i].Equal(that1.ResponseHeadersToAdd[i]) {
			return false
		}
	}
	if len(this.ResponseHeadersToRemove) != len(that1.ResponseHeadersToRemove) {
		return false
	}
	for i := range this.ResponseHeadersToRemove {
		if this.ResponseHeadersToRemove[i] != that1.ResponseHeadersToRemove[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderValueOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderValueOption)
	if !ok {
		that2, ok := that.(HeaderValueOption)
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
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !this.Append.Equal(that1.Append) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
