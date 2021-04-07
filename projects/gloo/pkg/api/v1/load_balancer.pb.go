// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/lbhash"
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

// LoadBalancerConfig is the settings for the load balancer used to send request to the Upstream
// endpoints.
type LoadBalancerConfig struct {
	// Configures envoy's panic threshold Percent between 0-100. Once the number of non health hosts
	// reaches this percentage, envoy disregards health information.
	// see more info [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/panic_threshold.html).
	HealthyPanicThreshold *types.DoubleValue `protobuf:"bytes,1,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// This allows batch updates of endpoints health/weight/metadata that happen during a time window.
	// this help lower cpu usage when endpoint change rate is high. defaults to 1 second.
	// Set to 0 to disable and have changes applied immediately.
	UpdateMergeWindow *time.Duration `protobuf:"bytes,2,opt,name=update_merge_window,json=updateMergeWindow,proto3,stdduration" json:"update_merge_window,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*LoadBalancerConfig_RoundRobin_
	//	*LoadBalancerConfig_LeastRequest_
	//	*LoadBalancerConfig_Random_
	//	*LoadBalancerConfig_RingHash_
	//	*LoadBalancerConfig_Maglev_
	Type                 isLoadBalancerConfig_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LoadBalancerConfig) Reset()         { *m = LoadBalancerConfig{} }
func (m *LoadBalancerConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig) ProtoMessage()    {}
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0}
}
func (m *LoadBalancerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig.Merge(m, src)
}
func (m *LoadBalancerConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig.Size(m)
}
func (m *LoadBalancerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig proto.InternalMessageInfo

type isLoadBalancerConfig_Type interface {
	isLoadBalancerConfig_Type()
	Equal(interface{}) bool
}

type LoadBalancerConfig_RoundRobin_ struct {
	RoundRobin *LoadBalancerConfig_RoundRobin `protobuf:"bytes,3,opt,name=round_robin,json=roundRobin,proto3,oneof" json:"round_robin,omitempty"`
}
type LoadBalancerConfig_LeastRequest_ struct {
	LeastRequest *LoadBalancerConfig_LeastRequest `protobuf:"bytes,4,opt,name=least_request,json=leastRequest,proto3,oneof" json:"least_request,omitempty"`
}
type LoadBalancerConfig_Random_ struct {
	Random *LoadBalancerConfig_Random `protobuf:"bytes,5,opt,name=random,proto3,oneof" json:"random,omitempty"`
}
type LoadBalancerConfig_RingHash_ struct {
	RingHash *LoadBalancerConfig_RingHash `protobuf:"bytes,6,opt,name=ring_hash,json=ringHash,proto3,oneof" json:"ring_hash,omitempty"`
}
type LoadBalancerConfig_Maglev_ struct {
	Maglev *LoadBalancerConfig_Maglev `protobuf:"bytes,7,opt,name=maglev,proto3,oneof" json:"maglev,omitempty"`
}

func (*LoadBalancerConfig_RoundRobin_) isLoadBalancerConfig_Type()   {}
func (*LoadBalancerConfig_LeastRequest_) isLoadBalancerConfig_Type() {}
func (*LoadBalancerConfig_Random_) isLoadBalancerConfig_Type()       {}
func (*LoadBalancerConfig_RingHash_) isLoadBalancerConfig_Type()     {}
func (*LoadBalancerConfig_Maglev_) isLoadBalancerConfig_Type()       {}

func (m *LoadBalancerConfig) GetType() isLoadBalancerConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *LoadBalancerConfig) GetHealthyPanicThreshold() *types.DoubleValue {
	if m != nil {
		return m.HealthyPanicThreshold
	}
	return nil
}

func (m *LoadBalancerConfig) GetUpdateMergeWindow() *time.Duration {
	if m != nil {
		return m.UpdateMergeWindow
	}
	return nil
}

func (m *LoadBalancerConfig) GetRoundRobin() *LoadBalancerConfig_RoundRobin {
	if x, ok := m.GetType().(*LoadBalancerConfig_RoundRobin_); ok {
		return x.RoundRobin
	}
	return nil
}

func (m *LoadBalancerConfig) GetLeastRequest() *LoadBalancerConfig_LeastRequest {
	if x, ok := m.GetType().(*LoadBalancerConfig_LeastRequest_); ok {
		return x.LeastRequest
	}
	return nil
}

func (m *LoadBalancerConfig) GetRandom() *LoadBalancerConfig_Random {
	if x, ok := m.GetType().(*LoadBalancerConfig_Random_); ok {
		return x.Random
	}
	return nil
}

func (m *LoadBalancerConfig) GetRingHash() *LoadBalancerConfig_RingHash {
	if x, ok := m.GetType().(*LoadBalancerConfig_RingHash_); ok {
		return x.RingHash
	}
	return nil
}

func (m *LoadBalancerConfig) GetMaglev() *LoadBalancerConfig_Maglev {
	if x, ok := m.GetType().(*LoadBalancerConfig_Maglev_); ok {
		return x.Maglev
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoadBalancerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoadBalancerConfig_RoundRobin_)(nil),
		(*LoadBalancerConfig_LeastRequest_)(nil),
		(*LoadBalancerConfig_Random_)(nil),
		(*LoadBalancerConfig_RingHash_)(nil),
		(*LoadBalancerConfig_Maglev_)(nil),
	}
}

type LoadBalancerConfig_RoundRobin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_RoundRobin) Reset()         { *m = LoadBalancerConfig_RoundRobin{} }
func (m *LoadBalancerConfig_RoundRobin) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RoundRobin) ProtoMessage()    {}
func (*LoadBalancerConfig_RoundRobin) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 0}
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.Merge(m, src)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Size(m)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RoundRobin proto.InternalMessageInfo

type LoadBalancerConfig_LeastRequest struct {
	// How many choices to take into account. defaults to 2.
	ChoiceCount          uint32   `protobuf:"varint,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_LeastRequest) Reset()         { *m = LoadBalancerConfig_LeastRequest{} }
func (m *LoadBalancerConfig_LeastRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_LeastRequest) ProtoMessage()    {}
func (*LoadBalancerConfig_LeastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 1}
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.Merge(m, src)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Size(m)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_LeastRequest proto.InternalMessageInfo

func (m *LoadBalancerConfig_LeastRequest) GetChoiceCount() uint32 {
	if m != nil {
		return m.ChoiceCount
	}
	return 0
}

type LoadBalancerConfig_Random struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_Random) Reset()         { *m = LoadBalancerConfig_Random{} }
func (m *LoadBalancerConfig_Random) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Random) ProtoMessage()    {}
func (*LoadBalancerConfig_Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 2}
}
func (m *LoadBalancerConfig_Random) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Random.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Random.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Random.Merge(m, src)
}
func (m *LoadBalancerConfig_Random) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Random.Size(m)
}
func (m *LoadBalancerConfig_Random) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Random.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Random proto.InternalMessageInfo

// Customizes the parameters used in the hashing algorithm to refine performance or resource usage.
type LoadBalancerConfig_RingHashConfig struct {
	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host)
	// the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited
	// to 8M entries.
	MinimumRingSize uint64 `protobuf:"varint,1,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further
	// constrain resource use.
	MaximumRingSize      uint64   `protobuf:"varint,2,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_RingHashConfig) Reset()         { *m = LoadBalancerConfig_RingHashConfig{} }
func (m *LoadBalancerConfig_RingHashConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHashConfig) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHashConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 3}
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Size(m)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHashConfig proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHashConfig) GetMinimumRingSize() uint64 {
	if m != nil {
		return m.MinimumRingSize
	}
	return 0
}

func (m *LoadBalancerConfig_RingHashConfig) GetMaximumRingSize() uint64 {
	if m != nil {
		return m.MaximumRingSize
	}
	return 0
}

type LoadBalancerConfig_RingHash struct {
	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig       *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *LoadBalancerConfig_RingHash) Reset()         { *m = LoadBalancerConfig_RingHash{} }
func (m *LoadBalancerConfig_RingHash) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHash) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 4}
}
func (m *LoadBalancerConfig_RingHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHash.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHash) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Size(m)
}
func (m *LoadBalancerConfig_RingHash) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHash.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHash proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHash) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if m != nil {
		return m.RingHashConfig
	}
	return nil
}

type LoadBalancerConfig_Maglev struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_Maglev) Reset()         { *m = LoadBalancerConfig_Maglev{} }
func (m *LoadBalancerConfig_Maglev) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Maglev) ProtoMessage()    {}
func (*LoadBalancerConfig_Maglev) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 5}
}
func (m *LoadBalancerConfig_Maglev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Maglev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Maglev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Maglev.Merge(m, src)
}
func (m *LoadBalancerConfig_Maglev) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Size(m)
}
func (m *LoadBalancerConfig_Maglev) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Maglev.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Maglev proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoadBalancerConfig)(nil), "gloo.solo.io.LoadBalancerConfig")
	proto.RegisterType((*LoadBalancerConfig_RoundRobin)(nil), "gloo.solo.io.LoadBalancerConfig.RoundRobin")
	proto.RegisterType((*LoadBalancerConfig_LeastRequest)(nil), "gloo.solo.io.LoadBalancerConfig.LeastRequest")
	proto.RegisterType((*LoadBalancerConfig_Random)(nil), "gloo.solo.io.LoadBalancerConfig.Random")
	proto.RegisterType((*LoadBalancerConfig_RingHashConfig)(nil), "gloo.solo.io.LoadBalancerConfig.RingHashConfig")
	proto.RegisterType((*LoadBalancerConfig_RingHash)(nil), "gloo.solo.io.LoadBalancerConfig.RingHash")
	proto.RegisterType((*LoadBalancerConfig_Maglev)(nil), "gloo.solo.io.LoadBalancerConfig.Maglev")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto", fileDescriptor_aaa1c019b03e4b0f)
}

var fileDescriptor_aaa1c019b03e4b0f = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcb, 0x6e, 0xd4, 0x3c,
	0x14, 0xc7, 0x67, 0xfa, 0xcd, 0x17, 0x8a, 0x3b, 0x2d, 0x34, 0x80, 0x08, 0x11, 0x2a, 0x97, 0x0d,
	0x37, 0x35, 0xa1, 0x20, 0x36, 0xac, 0x60, 0xca, 0x22, 0x8b, 0x16, 0x90, 0xa9, 0x40, 0xb0, 0x89,
	0x9c, 0xc4, 0x75, 0x0c, 0x8e, 0x4f, 0x70, 0x9c, 0xde, 0x9e, 0x84, 0x47, 0xe0, 0x11, 0x78, 0x9b,
	0x4a, 0xbc, 0x03, 0x7b, 0xe4, 0xcb, 0x94, 0x42, 0x55, 0xcd, 0xac, 0xc6, 0xc7, 0xe7, 0xfc, 0xfe,
	0xe7, 0x32, 0x27, 0x46, 0x2f, 0x18, 0xd7, 0x75, 0x5f, 0x24, 0x25, 0x34, 0x69, 0x07, 0x02, 0xd6,
	0x39, 0xa4, 0x4c, 0x00, 0xa4, 0xad, 0x82, 0xcf, 0xb4, 0xd4, 0x9d, 0xb3, 0x48, 0xcb, 0xd3, 0xbd,
	0x8d, 0x54, 0x00, 0xa9, 0xf2, 0x82, 0x08, 0x22, 0x4b, 0xaa, 0x92, 0x56, 0x81, 0x86, 0x70, 0x6c,
	0x02, 0x12, 0xc3, 0x26, 0x1c, 0xe2, 0x67, 0xe7, 0xc3, 0xd0, 0x6a, 0x0e, 0xb2, 0x4b, 0x45, 0x51,
	0x93, 0xae, 0xf6, 0x3f, 0x4e, 0x24, 0xbe, 0xca, 0x80, 0x81, 0x3d, 0xa6, 0xe6, 0xe4, 0x6f, 0xd7,
	0x18, 0x00, 0x13, 0x34, 0xb5, 0x56, 0xd1, 0xef, 0xa6, 0x55, 0xaf, 0x88, 0x11, 0x39, 0xcf, 0xbf,
	0xaf, 0x48, 0xdb, 0x52, 0xd5, 0x79, 0x7f, 0x48, 0x0f, 0xb4, 0x13, 0xa5, 0x07, 0xda, 0xdd, 0xdd,
	0x3d, 0x0e, 0x50, 0xb8, 0x05, 0xa4, 0x9a, 0xf8, 0x2e, 0x36, 0x41, 0xee, 0x72, 0x16, 0xee, 0xa0,
	0xeb, 0x35, 0x25, 0x42, 0xd7, 0x87, 0x79, 0x4b, 0x24, 0x2f, 0x73, 0x5d, 0x2b, 0xda, 0xd5, 0x20,
	0xaa, 0x68, 0x78, 0x7b, 0x78, 0x7f, 0xe9, 0xc9, 0xcd, 0xc4, 0x25, 0x4b, 0xa6, 0xc9, 0x92, 0x57,
	0xd0, 0x17, 0x82, 0xbe, 0x27, 0xa2, 0xa7, 0xf8, 0x9a, 0x87, 0xdf, 0x1a, 0x76, 0x67, 0x8a, 0x86,
	0x6f, 0xd0, 0x95, 0xbe, 0xad, 0x88, 0xa6, 0x79, 0x43, 0x15, 0xa3, 0xf9, 0x3e, 0x97, 0x15, 0xec,
	0x47, 0x0b, 0x56, 0xf1, 0xc6, 0x59, 0x45, 0xdf, 0xde, 0x64, 0xf4, 0xed, 0xf8, 0xd6, 0x10, 0xaf,
	0x3a, 0x76, 0xdb, 0xa0, 0x1f, 0x2c, 0x19, 0xbe, 0x46, 0x4b, 0x0a, 0x7a, 0x59, 0xe5, 0x0a, 0x0a,
	0x2e, 0xa3, 0xff, 0xac, 0xd0, 0xa3, 0xe4, 0xf4, 0x5f, 0x90, 0x9c, 0xed, 0x2e, 0xc1, 0x86, 0xc1,
	0x06, 0xc9, 0x06, 0x18, 0xa9, 0x13, 0x2b, 0xdc, 0x41, 0xcb, 0x82, 0x92, 0x4e, 0xe7, 0x8a, 0x7e,
	0xed, 0x69, 0xa7, 0xa3, 0x91, 0x55, 0x5c, 0x9f, 0xa9, 0xb8, 0x65, 0x28, 0xec, 0xa0, 0x6c, 0x80,
	0xc7, 0xe2, 0x94, 0x1d, 0xbe, 0x44, 0x81, 0x22, 0xb2, 0x82, 0x26, 0xfa, 0xdf, 0xca, 0xdd, 0x9b,
	0x5d, 0xa0, 0x0d, 0xcf, 0x06, 0xd8, 0x83, 0x61, 0x86, 0x2e, 0x2a, 0x2e, 0x59, 0x6e, 0x76, 0x24,
	0x0a, 0xac, 0xca, 0x83, 0xd9, 0x2a, 0x5c, 0xb2, 0x8c, 0x74, 0x75, 0x36, 0xc0, 0x8b, 0xca, 0x9f,
	0x4d, 0x31, 0x0d, 0x61, 0x82, 0xee, 0x45, 0x17, 0xe6, 0x2c, 0x66, 0xdb, 0x86, 0x9b, 0x62, 0x1c,
	0x18, 0x8f, 0x11, 0xfa, 0x33, 0xc1, 0x78, 0x03, 0x8d, 0x4f, 0x77, 0x1f, 0xde, 0x41, 0xe3, 0xb2,
	0x06, 0x5e, 0xd2, 0xbc, 0x84, 0x5e, 0x6a, 0xbb, 0x2f, 0xcb, 0x78, 0xc9, 0xdd, 0x6d, 0x9a, 0xab,
	0x78, 0x11, 0x05, 0xae, 0xc3, 0xb8, 0x46, 0x2b, 0xd3, 0x2a, 0xfd, 0xe6, 0x3d, 0x44, 0xab, 0x0d,
	0x97, 0xbc, 0xe9, 0x9b, 0xdc, 0x76, 0xdc, 0xf1, 0x23, 0x6a, 0x35, 0x46, 0xf8, 0x92, 0x77, 0x18,
	0xe2, 0x1d, 0x3f, 0xa2, 0x36, 0x96, 0x1c, 0xfc, 0x13, 0xbb, 0xe0, 0x63, 0x9d, 0x63, 0x1a, 0x1b,
	0x53, 0xb4, 0x38, 0xcd, 0x14, 0x7e, 0x44, 0x97, 0x4f, 0xa6, 0x99, 0x97, 0x36, 0xaf, 0x5f, 0xeb,
	0x74, 0xee, 0xa1, 0x3a, 0x13, 0xaf, 0xa8, 0xbf, 0x6c, 0xd3, 0x9a, 0x9b, 0xd7, 0x24, 0x40, 0x23,
	0x7d, 0xd8, 0xd2, 0xc9, 0xf3, 0x1f, 0xbf, 0x46, 0xc3, 0xef, 0x3f, 0xd7, 0x86, 0x9f, 0x1e, 0xcf,
	0xf7, 0xb8, 0xb4, 0x5f, 0x98, 0x7f, 0x23, 0x8a, 0xc0, 0x7e, 0x0b, 0x4f, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0xc7, 0x09, 0xc0, 0x97, 0x04, 0x00, 0x00,
}

func (this *LoadBalancerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig)
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
	if !this.HealthyPanicThreshold.Equal(that1.HealthyPanicThreshold) {
		return false
	}
	if this.UpdateMergeWindow != nil && that1.UpdateMergeWindow != nil {
		if *this.UpdateMergeWindow != *that1.UpdateMergeWindow {
			return false
		}
	} else if this.UpdateMergeWindow != nil {
		return false
	} else if that1.UpdateMergeWindow != nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin_)
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
	if !this.RoundRobin.Equal(that1.RoundRobin) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_LeastRequest_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest_)
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
	if !this.LeastRequest.Equal(that1.LeastRequest) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random_)
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
	if !this.Random.Equal(that1.Random) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash_)
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
	if !this.RingHash.Equal(that1.RingHash) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev_)
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
	if !this.Maglev.Equal(that1.Maglev) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_LeastRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest)
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
	if this.ChoiceCount != that1.ChoiceCount {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHashConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHashConfig)
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
	if this.MinimumRingSize != that1.MinimumRingSize {
		return false
	}
	if this.MaximumRingSize != that1.MaximumRingSize {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash)
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
	if !this.RingHashConfig.Equal(that1.RingHashConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
