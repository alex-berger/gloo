// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto

package faultinjection

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RouteAbort struct {
	// Percentage of requests that should be aborted, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// This should be a standard HTTP status, i.e. 503. Defaults to 0.
	HttpStatus           uint32   `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAbort) Reset()         { *m = RouteAbort{} }
func (m *RouteAbort) String() string { return proto.CompactTextString(m) }
func (*RouteAbort) ProtoMessage()    {}
func (*RouteAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{0}
}
func (m *RouteAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAbort.Unmarshal(m, b)
}
func (m *RouteAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAbort.Marshal(b, m, deterministic)
}
func (m *RouteAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAbort.Merge(m, src)
}
func (m *RouteAbort) XXX_Size() int {
	return xxx_messageInfo_RouteAbort.Size(m)
}
func (m *RouteAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAbort.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAbort proto.InternalMessageInfo

func (m *RouteAbort) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteAbort) GetHttpStatus() uint32 {
	if m != nil {
		return m.HttpStatus
	}
	return 0
}

type RouteDelay struct {
	// Percentage of requests that should be delayed, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Fixed delay, defaulting to 0.
	FixedDelay           *time.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,proto3,stdduration" json:"fixed_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RouteDelay) Reset()         { *m = RouteDelay{} }
func (m *RouteDelay) String() string { return proto.CompactTextString(m) }
func (*RouteDelay) ProtoMessage()    {}
func (*RouteDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{1}
}
func (m *RouteDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDelay.Unmarshal(m, b)
}
func (m *RouteDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDelay.Marshal(b, m, deterministic)
}
func (m *RouteDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDelay.Merge(m, src)
}
func (m *RouteDelay) XXX_Size() int {
	return xxx_messageInfo_RouteDelay.Size(m)
}
func (m *RouteDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDelay.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDelay proto.InternalMessageInfo

func (m *RouteDelay) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteDelay) GetFixedDelay() *time.Duration {
	if m != nil {
		return m.FixedDelay
	}
	return nil
}

type RouteFaults struct {
	Abort                *RouteAbort `protobuf:"bytes,1,opt,name=abort,proto3" json:"abort,omitempty"`
	Delay                *RouteDelay `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RouteFaults) Reset()         { *m = RouteFaults{} }
func (m *RouteFaults) String() string { return proto.CompactTextString(m) }
func (*RouteFaults) ProtoMessage()    {}
func (*RouteFaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{2}
}
func (m *RouteFaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFaults.Unmarshal(m, b)
}
func (m *RouteFaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFaults.Marshal(b, m, deterministic)
}
func (m *RouteFaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFaults.Merge(m, src)
}
func (m *RouteFaults) XXX_Size() int {
	return xxx_messageInfo_RouteFaults.Size(m)
}
func (m *RouteFaults) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFaults.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFaults proto.InternalMessageInfo

func (m *RouteFaults) GetAbort() *RouteAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *RouteFaults) GetDelay() *RouteDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteAbort)(nil), "fault.options.gloo.solo.io.RouteAbort")
	proto.RegisterType((*RouteDelay)(nil), "fault.options.gloo.solo.io.RouteDelay")
	proto.RegisterType((*RouteFaults)(nil), "fault.options.gloo.solo.io.RouteFaults")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto", fileDescriptor_2594ab72fbfb9a06)
}

var fileDescriptor_2594ab72fbfb9a06 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x4e, 0x2b, 0x31,
	0x10, 0xfd, 0x8e, 0x7e, 0x28, 0xbc, 0xd0, 0xac, 0x90, 0x08, 0x29, 0x92, 0x28, 0x05, 0x8a, 0x90,
	0xb0, 0x45, 0x68, 0x69, 0x88, 0x22, 0xa0, 0x41, 0x88, 0xa5, 0xa3, 0x89, 0xbc, 0x59, 0xc7, 0x31,
	0x2c, 0x3b, 0xd6, 0xee, 0x6c, 0xb4, 0x1c, 0x81, 0x1b, 0x70, 0x04, 0xc4, 0x09, 0xb8, 0x0d, 0x12,
	0x77, 0xa0, 0xa1, 0x42, 0xb6, 0x37, 0x22, 0x0d, 0x84, 0x6e, 0xe6, 0xf9, 0xbd, 0x79, 0x7e, 0xa3,
	0xa1, 0x97, 0x4a, 0xe3, 0xbc, 0x8c, 0xd9, 0x14, 0xee, 0x79, 0x01, 0x29, 0x1c, 0x68, 0xe0, 0x2a,
	0x05, 0xe0, 0x26, 0x87, 0x5b, 0x39, 0xc5, 0xc2, 0x77, 0xc2, 0x68, 0xbe, 0x38, 0xe4, 0x60, 0x50,
	0x43, 0x56, 0xf0, 0x99, 0x28, 0x53, 0xd4, 0x99, 0x25, 0x68, 0xc8, 0x7c, 0xcb, 0x4c, 0x0e, 0x08,
	0x61, 0xdb, 0x37, 0x35, 0x93, 0x59, 0x35, 0xb3, 0x83, 0x99, 0x86, 0x76, 0x47, 0x01, 0xa8, 0x54,
	0x72, 0xc7, 0x8c, 0xcb, 0x19, 0x4f, 0xca, 0x5c, 0x58, 0x9e, 0xd7, 0xb6, 0x77, 0x16, 0x22, 0xd5,
	0x89, 0x40, 0xc9, 0x97, 0x45, 0xfd, 0xb0, 0xad, 0x40, 0x81, 0x2b, 0xb9, 0xad, 0x6a, 0x34, 0x94,
	0x15, 0x7a, 0x50, 0x56, 0xb5, 0x7d, 0xff, 0x82, 0xd2, 0x08, 0x4a, 0x94, 0x27, 0x31, 0xe4, 0x18,
	0x76, 0x28, 0x35, 0x32, 0x9f, 0xca, 0x0c, 0x85, 0x92, 0x2d, 0xd2, 0x23, 0x83, 0x46, 0xb4, 0x82,
	0x84, 0x5d, 0x1a, 0xcc, 0x11, 0xcd, 0xa4, 0x40, 0x81, 0x65, 0xd1, 0x6a, 0xf4, 0xc8, 0x60, 0x2b,
	0xa2, 0x16, 0xba, 0x76, 0x48, 0x7f, 0x51, 0x8f, 0x1b, 0xcb, 0x54, 0x3c, 0xac, 0x1d, 0x77, 0x4e,
	0x83, 0x99, 0xae, 0x64, 0x32, 0x49, 0x2c, 0xdd, 0x8d, 0x0b, 0x86, 0xbb, 0xcc, 0xa7, 0x66, 0xcb,
	0xd4, 0x6c, 0x5c, 0xa7, 0x1e, 0x6d, 0x7e, 0x8e, 0x9a, 0x2f, 0xa4, 0xb1, 0xff, 0xef, 0xe9, 0xad,
	0x4b, 0x22, 0xea, 0xb4, 0xce, 0xa9, 0xff, 0x48, 0x68, 0xe0, 0x8c, 0x4f, 0xed, 0x36, 0x8b, 0xf0,
	0x98, 0x36, 0x85, 0x4d, 0xe4, 0x4c, 0x83, 0xe1, 0x1e, 0xfb, 0x79, 0xcb, 0xec, 0x3b, 0x7f, 0xe4,
	0x45, 0x56, 0xbd, 0xfa, 0xa3, 0xf5, 0x6a, 0xf7, 0x89, 0xc8, 0x8b, 0x46, 0x57, 0xaf, 0x1f, 0xff,
	0xc9, 0xf3, 0x7b, 0x87, 0xdc, 0x9c, 0xfd, 0xed, 0x58, 0xcc, 0x9d, 0xfa, 0xfd, 0x60, 0xe2, 0x0d,
	0xb7, 0x8b, 0xa3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xdc, 0xe3, 0x7a, 0x7e, 0x02, 0x00,
	0x00,
}

func (this *RouteAbort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteAbort)
	if !ok {
		that2, ok := that.(RouteAbort)
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
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.HttpStatus != that1.HttpStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteDelay) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteDelay)
	if !ok {
		that2, ok := that.(RouteDelay)
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
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.FixedDelay != nil && that1.FixedDelay != nil {
		if *this.FixedDelay != *that1.FixedDelay {
			return false
		}
	} else if this.FixedDelay != nil {
		return false
	} else if that1.FixedDelay != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteFaults) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteFaults)
	if !ok {
		that2, ok := that.(RouteFaults)
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
	if !this.Abort.Equal(that1.Abort) {
		return false
	}
	if !this.Delay.Equal(that1.Delay) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
