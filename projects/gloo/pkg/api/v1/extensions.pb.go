// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
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

type Extensions struct {
	Configs              map[string]*types.Struct `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Extensions) Reset()         { *m = Extensions{} }
func (m *Extensions) String() string { return proto.CompactTextString(m) }
func (*Extensions) ProtoMessage()    {}
func (*Extensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6aa7a8f802feeb, []int{0}
}
func (m *Extensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extensions.Unmarshal(m, b)
}
func (m *Extensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extensions.Marshal(b, m, deterministic)
}
func (m *Extensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extensions.Merge(m, src)
}
func (m *Extensions) XXX_Size() int {
	return xxx_messageInfo_Extensions.Size(m)
}
func (m *Extensions) XXX_DiscardUnknown() {
	xxx_messageInfo_Extensions.DiscardUnknown(m)
}

var xxx_messageInfo_Extensions proto.InternalMessageInfo

func (m *Extensions) GetConfigs() map[string]*types.Struct {
	if m != nil {
		return m.Configs
	}
	return nil
}

type Extension struct {
	Config               *types.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6aa7a8f802feeb, []int{1}
}
func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (m *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(m, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetConfig() *types.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Extensions)(nil), "gloo.solo.io.Extensions")
	proto.RegisterMapType((map[string]*types.Struct)(nil), "gloo.solo.io.Extensions.ConfigsEntry")
	proto.RegisterType((*Extension)(nil), "gloo.solo.io.Extension")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto", fileDescriptor_eb6aa7a8f802feeb)
}

var fileDescriptor_eb6aa7a8f802feeb = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x56, 0x2b, 0xdd, 0xf6, 0x20, 0x8b, 0x60, 0x08, 0x22, 0xa1, 0x20, 0xe4, 0xd2,
	0x19, 0xad, 0x17, 0x29, 0x8a, 0xa0, 0xf4, 0x05, 0xd2, 0x9b, 0xb7, 0x26, 0x6c, 0xd7, 0xb5, 0x31,
	0x13, 0xb2, 0x9b, 0x92, 0xbe, 0x8e, 0x27, 0x1f, 0xc1, 0xe7, 0xf1, 0x1d, 0xbc, 0x4b, 0xb2, 0x69,
	0xed, 0x49, 0x7a, 0x9b, 0xdd, 0xff, 0x9b, 0xff, 0xe7, 0x67, 0xf8, 0x83, 0xd2, 0xf6, 0xb5, 0x8c,
	0x21, 0xa1, 0x77, 0x34, 0x94, 0xd2, 0x58, 0x13, 0xaa, 0x94, 0x08, 0xf3, 0x82, 0xde, 0x64, 0x62,
	0x8d, 0x7b, 0x2d, 0x72, 0x8d, 0xeb, 0x1b, 0x94, 0x95, 0x95, 0x99, 0xd1, 0x94, 0x19, 0xc8, 0x0b,
	0xb2, 0x24, 0x86, 0xb5, 0x0a, 0xf5, 0x22, 0x68, 0xf2, 0x2f, 0x14, 0x91, 0x4a, 0x25, 0x36, 0x5a,
	0x5c, 0x2e, 0xd1, 0xd8, 0xa2, 0x4c, 0xac, 0x63, 0xfd, 0x33, 0x45, 0x8a, 0x9a, 0x11, 0xeb, 0xa9,
	0xfd, 0x15, 0xb2, 0xb2, 0xee, 0x53, 0x56, 0x2d, 0x39, 0xfa, 0x60, 0x9c, 0xcf, 0x76, 0x51, 0xe2,
	0x91, 0x9f, 0x24, 0x94, 0x2d, 0xb5, 0x32, 0x1e, 0x0b, 0xba, 0xe1, 0x60, 0x72, 0x05, 0xfb, 0xb1,
	0xf0, 0x87, 0xc2, 0xb3, 0xe3, 0x66, 0x99, 0x2d, 0x36, 0xd1, 0x76, 0xcb, 0x9f, 0xf3, 0xe1, 0xbe,
	0x20, 0x4e, 0x79, 0x77, 0x25, 0x37, 0x1e, 0x0b, 0x58, 0xd8, 0x8f, 0xea, 0x51, 0x8c, 0xf9, 0xf1,
	0x7a, 0x91, 0x96, 0xd2, 0xeb, 0x04, 0x2c, 0x1c, 0x4c, 0xce, 0xc1, 0x35, 0x81, 0x6d, 0x13, 0x98,
	0x37, 0x4d, 0x22, 0x47, 0x4d, 0x3b, 0x77, 0x6c, 0x74, 0xcf, 0xfb, 0xbb, 0x60, 0x81, 0xbc, 0xe7,
	0xc2, 0x1a, 0xd3, 0x7f, 0x0c, 0x5a, 0xec, 0x69, 0xfa, 0xf5, 0x73, 0xc4, 0x3e, 0xbf, 0x2f, 0xd9,
	0xcb, 0xf5, 0x61, 0x17, 0xc8, 0x57, 0xaa, 0xbd, 0x42, 0xdc, 0x6b, 0x4c, 0x6f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x03, 0xec, 0x8d, 0xe3, 0xbc, 0x01, 0x00, 0x00,
}

func (this *Extensions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extensions)
	if !ok {
		that2, ok := that.(Extensions)
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
	if len(this.Configs) != len(that1.Configs) {
		return false
	}
	for i := range this.Configs {
		if !this.Configs[i].Equal(that1.Configs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Extension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extension)
	if !ok {
		that2, ok := that.(Extension)
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
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
