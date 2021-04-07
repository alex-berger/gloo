// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	cluster "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"
	core1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"
	ec2 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/ec2"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/consul"
	kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"
	pipe "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/pipe"
	static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin. (plugins currently need to be compiled into Gloo)
type Upstream struct {
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,3,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	SslConfig         *UpstreamSslConfig `protobuf:"bytes,4,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Circuit breakers for this upstream. if not set, the defaults ones from the Gloo settings will be used.
	// if those are not set, [envoy's defaults](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
	// will be used.
	CircuitBreakers    *CircuitBreakerConfig     `protobuf:"bytes,5,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	LoadBalancerConfig *LoadBalancerConfig       `protobuf:"bytes,6,opt,name=load_balancer_config,json=loadBalancerConfig,proto3" json:"load_balancer_config,omitempty"`
	ConnectionConfig   *ConnectionConfig         `protobuf:"bytes,7,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	HealthChecks       []*core1.HealthCheck      `protobuf:"bytes,8,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	OutlierDetection   *cluster.OutlierDetection `protobuf:"bytes,9,opt,name=outlier_detection,json=outlierDetection,proto3" json:"outlier_detection,omitempty"`
	// Use http2 when communicating with this upstream
	// this field is evaluated `true` for upstreams
	// with a grpc service spec. otherwise defaults to `false`
	UseHttp2 *types.BoolValue `protobuf:"bytes,10,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
	// Note to developers: new Upstream plugins must be added to this oneof field
	// to be usable by Gloo. (plugins currently need to be compiled into Gloo)
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*Upstream_Kube
	//	*Upstream_Static
	//	*Upstream_Pipe
	//	*Upstream_Aws
	//	*Upstream_Azure
	//	*Upstream_Consul
	//	*Upstream_AwsEc2
	UpstreamType isUpstream_UpstreamType `protobuf_oneof:"upstream_type"`
	// Failover endpoints for this upstream. If omitted (the default) no failovers will be applied.
	Failover *Failover `protobuf:"bytes,18,opt,name=failover,proto3" json:"failover,omitempty"`
	// (UInt32Value) Initial stream-level flow-control window size.
	// Valid values range from 65535 (2^16 - 1, HTTP/2 default) to 2147483647 (2^31 - 1, HTTP/2 maximum)
	// and defaults to 268435456 (256 * 1024 * 1024).
	// NOTE: 65535 is the initial window size from HTTP/2 spec.
	// We only support increasing the default window size now, so it’s also the minimum.
	// This field also acts as a soft limit on the number of bytes Envoy will buffer per-stream
	// in the HTTP/2 codec buffers. Once the buffer reaches this pointer,
	// watermark callbacks will fire to stop the flow of data to the codec buffers.
	// Requires UseHttp2 to be true to be acknowledged.
	InitialStreamWindowSize *types.UInt32Value `protobuf:"bytes,19,opt,name=initial_stream_window_size,json=initialStreamWindowSize,proto3" json:"initial_stream_window_size,omitempty"`
	// (UInt32Value) Similar to initial_stream_window_size, but for connection-level flow-control window.
	// Currently, this has the same minimum/maximum/default as initial_stream_window_size.
	// Requires UseHttp2 to be true to be acknowledged.
	InitialConnectionWindowSize *types.UInt32Value `protobuf:"bytes,20,opt,name=initial_connection_window_size,json=initialConnectionWindowSize,proto3" json:"initial_connection_window_size,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}           `json:"-"`
	XXX_unrecognized            []byte             `json:"-"`
	XXX_sizecache               int32              `json:"-"`
}

func (m *Upstream) Reset()         { *m = Upstream{} }
func (m *Upstream) String() string { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()    {}
func (*Upstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{0}
}
func (m *Upstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upstream.Unmarshal(m, b)
}
func (m *Upstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upstream.Marshal(b, m, deterministic)
}
func (m *Upstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upstream.Merge(m, src)
}
func (m *Upstream) XXX_Size() int {
	return xxx_messageInfo_Upstream.Size(m)
}
func (m *Upstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Upstream.DiscardUnknown(m)
}

var xxx_messageInfo_Upstream proto.InternalMessageInfo

type isUpstream_UpstreamType interface {
	isUpstream_UpstreamType()
	Equal(interface{}) bool
}

type Upstream_Kube struct {
	Kube *kubernetes.UpstreamSpec `protobuf:"bytes,11,opt,name=kube,proto3,oneof" json:"kube,omitempty"`
}
type Upstream_Static struct {
	Static *static.UpstreamSpec `protobuf:"bytes,12,opt,name=static,proto3,oneof" json:"static,omitempty"`
}
type Upstream_Pipe struct {
	Pipe *pipe.UpstreamSpec `protobuf:"bytes,13,opt,name=pipe,proto3,oneof" json:"pipe,omitempty"`
}
type Upstream_Aws struct {
	Aws *aws.UpstreamSpec `protobuf:"bytes,14,opt,name=aws,proto3,oneof" json:"aws,omitempty"`
}
type Upstream_Azure struct {
	Azure *azure.UpstreamSpec `protobuf:"bytes,15,opt,name=azure,proto3,oneof" json:"azure,omitempty"`
}
type Upstream_Consul struct {
	Consul *consul.UpstreamSpec `protobuf:"bytes,16,opt,name=consul,proto3,oneof" json:"consul,omitempty"`
}
type Upstream_AwsEc2 struct {
	AwsEc2 *ec2.UpstreamSpec `protobuf:"bytes,17,opt,name=aws_ec2,json=awsEc2,proto3,oneof" json:"aws_ec2,omitempty"`
}

func (*Upstream_Kube) isUpstream_UpstreamType()   {}
func (*Upstream_Static) isUpstream_UpstreamType() {}
func (*Upstream_Pipe) isUpstream_UpstreamType()   {}
func (*Upstream_Aws) isUpstream_UpstreamType()    {}
func (*Upstream_Azure) isUpstream_UpstreamType()  {}
func (*Upstream_Consul) isUpstream_UpstreamType() {}
func (*Upstream_AwsEc2) isUpstream_UpstreamType() {}

func (m *Upstream) GetUpstreamType() isUpstream_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *Upstream) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Upstream) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

func (m *Upstream) GetSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *Upstream) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *Upstream) GetLoadBalancerConfig() *LoadBalancerConfig {
	if m != nil {
		return m.LoadBalancerConfig
	}
	return nil
}

func (m *Upstream) GetConnectionConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnectionConfig
	}
	return nil
}

func (m *Upstream) GetHealthChecks() []*core1.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *Upstream) GetOutlierDetection() *cluster.OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *Upstream) GetUseHttp2() *types.BoolValue {
	if m != nil {
		return m.UseHttp2
	}
	return nil
}

func (m *Upstream) GetKube() *kubernetes.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *Upstream) GetStatic() *static.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Static); ok {
		return x.Static
	}
	return nil
}

func (m *Upstream) GetPipe() *pipe.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Pipe); ok {
		return x.Pipe
	}
	return nil
}

func (m *Upstream) GetAws() *aws.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Upstream) GetAzure() *azure.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Upstream) GetConsul() *consul.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Consul); ok {
		return x.Consul
	}
	return nil
}

func (m *Upstream) GetAwsEc2() *ec2.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_AwsEc2); ok {
		return x.AwsEc2
	}
	return nil
}

func (m *Upstream) GetFailover() *Failover {
	if m != nil {
		return m.Failover
	}
	return nil
}

func (m *Upstream) GetInitialStreamWindowSize() *types.UInt32Value {
	if m != nil {
		return m.InitialStreamWindowSize
	}
	return nil
}

func (m *Upstream) GetInitialConnectionWindowSize() *types.UInt32Value {
	if m != nil {
		return m.InitialConnectionWindowSize
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Upstream) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Upstream_Kube)(nil),
		(*Upstream_Static)(nil),
		(*Upstream_Pipe)(nil),
		(*Upstream_Aws)(nil),
		(*Upstream_Azure)(nil),
		(*Upstream_Consul)(nil),
		(*Upstream_AwsEc2)(nil),
	}
}

// created by discovery services
type DiscoveryMetadata struct {
	// Labels inherited from the original upstream (e.g. Kubernetes labels)
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func (m *DiscoveryMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.DiscoveryMetadata.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto", fileDescriptor_b74df493149f644d)
}

var fileDescriptor_b74df493149f644d = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x1c, 0xc5, 0xeb, 0xc6, 0x49, 0x63, 0x26, 0x59, 0x62, 0x2e, 0x58, 0x85, 0xac, 0x73, 0x02, 0x0f,
	0x58, 0xb3, 0x0e, 0xa1, 0x56, 0x07, 0x43, 0x3b, 0x0f, 0x1d, 0x06, 0x3b, 0x19, 0x32, 0x34, 0xdd,
	0x00, 0x19, 0xdd, 0xb0, 0xdd, 0x08, 0x34, 0x4d, 0xdb, 0x9c, 0x19, 0x51, 0x10, 0x29, 0x3b, 0xce,
	0x65, 0x5f, 0x61, 0x2f, 0xb1, 0x47, 0xd8, 0x23, 0xec, 0x7e, 0xf7, 0xbd, 0xd8, 0x1b, 0x74, 0xc0,
	0xee, 0x07, 0x7e, 0xc8, 0xf1, 0x47, 0x5d, 0x6b, 0x17, 0xb6, 0x44, 0xf2, 0x9c, 0x9f, 0x28, 0xea,
	0xcf, 0x23, 0x81, 0xaf, 0x7a, 0x4c, 0xf5, 0xd3, 0x36, 0x22, 0xe2, 0xca, 0x97, 0x82, 0x8b, 0x13,
	0x26, 0xfc, 0x1e, 0x17, 0xc2, 0x8f, 0x13, 0xf1, 0x2b, 0x25, 0x4a, 0xda, 0x16, 0x8e, 0x99, 0x3f,
	0x7c, 0xec, 0xa7, 0xb1, 0x54, 0x09, 0xc5, 0x57, 0x28, 0x4e, 0x84, 0x12, 0x70, 0x5b, 0x8f, 0x21,
	0x6d, 0x43, 0x4c, 0x1c, 0xec, 0xf7, 0x44, 0x4f, 0x98, 0x01, 0x5f, 0x9f, 0x59, 0xcd, 0x01, 0xa4,
	0xd7, 0xca, 0x76, 0xd2, 0x6b, 0xe5, 0xfa, 0x2a, 0xe6, 0x4a, 0x03, 0xa6, 0x32, 0xee, 0x15, 0x55,
	0xb8, 0x83, 0x15, 0x76, 0xe3, 0x1f, 0x2f, 0x9f, 0x81, 0x94, 0xdc, 0x89, 0xde, 0x31, 0x4d, 0xc2,
	0x12, 0x92, 0x32, 0x15, 0xb6, 0x13, 0x8a, 0x07, 0x34, 0x71, 0x86, 0x93, 0xe5, 0x06, 0x2e, 0x70,
	0x27, 0x6c, 0x63, 0x8e, 0x23, 0x32, 0x91, 0x3f, 0x7a, 0x07, 0x5f, 0x44, 0x11, 0x25, 0x8a, 0x89,
	0xc8, 0x69, 0xcf, 0x96, 0x68, 0xe9, 0xb5, 0xa2, 0x49, 0x84, 0xb9, 0x4f, 0xa3, 0xa1, 0x18, 0x5b,
	0x7b, 0xcd, 0x27, 0x22, 0xa1, 0x7e, 0x9f, 0x62, 0xae, 0xfa, 0x21, 0xe9, 0x53, 0x32, 0x70, 0x94,
	0x07, 0xf3, 0xcb, 0x22, 0x15, 0x56, 0xa9, 0x74, 0xa3, 0x97, 0xff, 0xef, 0x1a, 0x3c, 0x95, 0x8a,
	0x26, 0xbe, 0x48, 0x15, 0x67, 0x34, 0x09, 0x3b, 0x54, 0xcd, 0xcc, 0x78, 0xe1, 0x11, 0x64, 0x6d,
	0x37, 0xfe, 0xc5, 0xf2, 0xbb, 0x17, 0xb1, 0xe6, 0x48, 0x33, 0x3b, 0x46, 0xdc, 0xc1, 0xd9, 0x1e,
	0xaf, 0xb6, 0xc5, 0x2c, 0xa6, 0xe6, 0xcf, 0x59, 0x9e, 0xad, 0xb6, 0x0c, 0xd2, 0x36, 0x4d, 0x22,
	0xaa, 0xe8, 0xf4, 0xe9, 0xea, 0x32, 0xc8, 0xec, 0x78, 0x64, 0x7e, 0xce, 0x70, 0x9a, 0xc3, 0x70,
	0x93, 0x26, 0xd4, 0xfe, 0xe7, 0x5f, 0x0e, 0x22, 0x22, 0x99, 0x72, 0x77, 0x70, 0xb6, 0x27, 0xf9,
	0x26, 0x47, 0x49, 0x4d, 0x1f, 0x43, 0x4a, 0x6a, 0xce, 0xf8, 0x70, 0xa5, 0xd1, 0x09, 0x8f, 0x97,
	0x0b, 0xbb, 0x98, 0x71, 0x31, 0x9c, 0xd4, 0x73, 0xa5, 0x27, 0x44, 0x8f, 0x53, 0xdf, 0xb4, 0xda,
	0x69, 0xd7, 0x1f, 0x25, 0x38, 0x8e, 0x69, 0xe2, 0x48, 0xd5, 0xbf, 0xb6, 0xc0, 0xe6, 0x4b, 0xb7,
	0xbf, 0xe1, 0x73, 0xb0, 0x61, 0x8b, 0xcf, 0x2b, 0x1c, 0x15, 0x8e, 0xb7, 0x6a, 0xfb, 0x48, 0x17,
	0x6d, 0xb6, 0xd5, 0x51, 0xcb, 0x8c, 0x35, 0x3e, 0xfa, 0xe3, 0xdf, 0x62, 0xe1, 0xcf, 0xd7, 0x87,
	0x77, 0xfe, 0x79, 0x7d, 0x58, 0x56, 0x54, 0xaa, 0x0e, 0xeb, 0x76, 0xeb, 0x55, 0xd6, 0x8b, 0x44,
	0x42, 0xab, 0x81, 0x43, 0xc0, 0xa7, 0x60, 0x33, 0xdb, 0xe0, 0xde, 0x5d, 0x83, 0xfb, 0x60, 0x16,
	0xf7, 0xc2, 0x8d, 0x36, 0x8a, 0x1a, 0x16, 0x4c, 0xd4, 0xf0, 0x7b, 0x00, 0x3b, 0x4c, 0x12, 0x7d,
	0x17, 0xe3, 0x70, 0xc2, 0x58, 0x33, 0x8c, 0x43, 0x34, 0x9d, 0x3e, 0xe8, 0x2c, 0xd3, 0x65, 0xb0,
	0xa0, 0xdc, 0x99, 0xef, 0x82, 0x5f, 0x03, 0x20, 0x25, 0x0f, 0x89, 0x88, 0xba, 0xac, 0xe7, 0x15,
	0xdf, 0xc6, 0xc9, 0x96, 0xa0, 0x25, 0x79, 0xd3, 0xc8, 0x82, 0x92, 0xcc, 0x4e, 0xe1, 0x0b, 0xb0,
	0x37, 0x97, 0x2d, 0xd2, 0x5b, 0x37, 0x94, 0xea, 0x2c, 0xa5, 0x69, 0x55, 0x0d, 0x2b, 0x72, 0xa0,
	0x5d, 0x32, 0xd3, 0x2b, 0x61, 0x00, 0xf6, 0x67, 0x92, 0x27, 0x9b, 0xd8, 0x86, 0x41, 0x1e, 0xcd,
	0x22, 0x2f, 0x05, 0xee, 0x34, 0x9c, 0xd0, 0x01, 0x21, 0x5f, 0xe8, 0x83, 0xcf, 0x41, 0xf9, 0x36,
	0x9e, 0x32, 0xe0, 0x3d, 0x03, 0xac, 0xcc, 0xcd, 0x71, 0x22, 0x73, 0xb8, 0x3d, 0x32, 0xd7, 0x03,
	0x9b, 0x60, 0x67, 0x3a, 0xa7, 0xa4, 0xb7, 0x79, 0xb4, 0x66, 0x40, 0x26, 0x6b, 0x10, 0x8e, 0x19,
	0x1a, 0xd6, 0xec, 0xb3, 0xbc, 0x30, 0xba, 0xa6, 0x96, 0x05, 0xdb, 0xfd, 0xdb, 0x86, 0x84, 0x2d,
	0x50, 0x5e, 0x48, 0x21, 0xaf, 0x64, 0x66, 0xf4, 0xc9, 0x1c, 0xc8, 0x86, 0x16, 0xfa, 0xc1, 0xca,
	0xcf, 0x32, 0x75, 0xb0, 0x27, 0xe6, 0x7a, 0xe0, 0x13, 0x50, 0x4a, 0x25, 0x0d, 0xfb, 0x4a, 0xc5,
	0x35, 0x0f, 0x18, 0xd8, 0x01, 0xb2, 0x15, 0x8e, 0xb2, 0x0a, 0x47, 0x0d, 0x21, 0xf8, 0x8f, 0x98,
	0xa7, 0x34, 0xd8, 0x4c, 0x25, 0xbd, 0xd0, 0x5a, 0xd8, 0x04, 0x45, 0x9d, 0x21, 0xde, 0x96, 0xf1,
	0x9c, 0xa0, 0xa9, 0x40, 0xc9, 0x76, 0xd6, 0xdb, 0xeb, 0x21, 0xa6, 0xe4, 0xe2, 0x4e, 0x60, 0xcc,
	0xb0, 0x69, 0xb7, 0x07, 0x23, 0xde, 0xb6, 0xc1, 0x7c, 0x8a, 0x5c, 0x0a, 0xe6, 0x41, 0x38, 0x2b,
	0x7c, 0x06, 0x8a, 0x3a, 0x06, 0xbd, 0x1d, 0x83, 0x78, 0x88, 0x4c, 0x26, 0xe6, 0x9a, 0x83, 0x56,
	0xc2, 0x3a, 0x58, 0xc3, 0x23, 0xe9, 0xbd, 0xe7, 0x16, 0x52, 0x07, 0x5c, 0x1e, 0xb3, 0x36, 0xc1,
	0x6f, 0xc0, 0xba, 0x49, 0x37, 0x6f, 0xd7, 0xb8, 0x8f, 0x91, 0xcd, 0xba, 0x3c, 0x7e, 0x6b, 0xd4,
	0x2b, 0x60, 0x93, 0xce, 0xdb, 0x73, 0x2b, 0xe0, 0x82, 0x2f, 0xd7, 0x0a, 0x58, 0x2d, 0x3c, 0x07,
	0xf7, 0x5c, 0xec, 0x79, 0x65, 0x43, 0x79, 0x84, 0xb2, 0x18, 0xcc, 0x85, 0xc1, 0x23, 0x79, 0x4e,
	0x6a, 0xb0, 0x06, 0x36, 0xb3, 0xac, 0xf3, 0xa0, 0xcb, 0x97, 0x19, 0xdf, 0xb7, 0x6e, 0x34, 0x98,
	0xe8, 0xe0, 0xcf, 0xe0, 0x80, 0x45, 0x4c, 0x31, 0xcc, 0x43, 0xcb, 0x0c, 0x47, 0x2c, 0xea, 0x88,
	0x51, 0x28, 0xd9, 0x0d, 0xf5, 0xde, 0x37, 0x94, 0x07, 0x0b, 0x05, 0xf5, 0xf2, 0xbb, 0x48, 0x9d,
	0xd6, 0x6c, 0x49, 0xdd, 0x77, 0xfe, 0x96, 0xb1, 0xff, 0x64, 0xdc, 0x2d, 0x76, 0x43, 0x21, 0x06,
	0x95, 0x0c, 0x3d, 0xb5, 0x13, 0xa7, 0xf1, 0xfb, 0x39, 0xf0, 0x1f, 0x3a, 0xc6, 0xed, 0x2e, 0xbd,
	0xbd, 0x44, 0xfd, 0xfe, 0xab, 0x37, 0xc5, 0x22, 0xb8, 0x9b, 0xca, 0x57, 0x6f, 0x8a, 0x5b, 0xb0,
	0x94, 0x7d, 0x96, 0xc9, 0xc6, 0x2e, 0xd8, 0xc9, 0x1a, 0xa1, 0x1a, 0xc7, 0xb4, 0xfa, 0x5b, 0x01,
	0x94, 0x17, 0xa2, 0x51, 0x3f, 0x3d, 0x8e, 0xdb, 0x94, 0xeb, 0x78, 0xd7, 0x1b, 0xfa, 0xb3, 0x15,
	0x59, 0x8a, 0x2e, 0x8d, 0xfa, 0x3c, 0x52, 0xc9, 0x38, 0x70, 0xd6, 0x83, 0x2f, 0xc1, 0xd6, 0x54,
	0x37, 0xdc, 0x03, 0x6b, 0x03, 0x3a, 0x36, 0xef, 0x8b, 0x52, 0xa0, 0x4f, 0xe1, 0x3e, 0x58, 0x1f,
	0xea, 0x7b, 0x31, 0xa1, 0x5f, 0x0a, 0x6c, 0xa3, 0x7e, 0xf7, 0x69, 0xa1, 0x51, 0xd7, 0x2f, 0x8e,
	0xdf, 0xff, 0xae, 0x14, 0x7e, 0xf9, 0x3c, 0xdf, 0xf7, 0x67, 0x3c, 0xe8, 0xb9, 0xd7, 0x5a, 0x7b,
	0xc3, 0x2c, 0xd7, 0xe9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x17, 0x11, 0x0e, 0x9f, 0xba, 0x0a,
	0x00, 0x00,
}

func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if !this.LoadBalancerConfig.Equal(that1.LoadBalancerConfig) {
		return false
	}
	if !this.ConnectionConfig.Equal(that1.ConnectionConfig) {
		return false
	}
	if len(this.HealthChecks) != len(that1.HealthChecks) {
		return false
	}
	for i := range this.HealthChecks {
		if !this.HealthChecks[i].Equal(that1.HealthChecks[i]) {
			return false
		}
	}
	if !this.OutlierDetection.Equal(that1.OutlierDetection) {
		return false
	}
	if !this.UseHttp2.Equal(that1.UseHttp2) {
		return false
	}
	if that1.UpstreamType == nil {
		if this.UpstreamType != nil {
			return false
		}
	} else if this.UpstreamType == nil {
		return false
	} else if !this.UpstreamType.Equal(that1.UpstreamType) {
		return false
	}
	if !this.Failover.Equal(that1.Failover) {
		return false
	}
	if !this.InitialStreamWindowSize.Equal(that1.InitialStreamWindowSize) {
		return false
	}
	if !this.InitialConnectionWindowSize.Equal(that1.InitialConnectionWindowSize) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Upstream_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Kube)
	if !ok {
		that2, ok := that.(Upstream_Kube)
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
	if !this.Kube.Equal(that1.Kube) {
		return false
	}
	return true
}
func (this *Upstream_Static) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Static)
	if !ok {
		that2, ok := that.(Upstream_Static)
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
	if !this.Static.Equal(that1.Static) {
		return false
	}
	return true
}
func (this *Upstream_Pipe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Pipe)
	if !ok {
		that2, ok := that.(Upstream_Pipe)
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
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	return true
}
func (this *Upstream_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Aws)
	if !ok {
		that2, ok := that.(Upstream_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Upstream_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Azure)
	if !ok {
		that2, ok := that.(Upstream_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Upstream_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Consul)
	if !ok {
		that2, ok := that.(Upstream_Consul)
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
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
func (this *Upstream_AwsEc2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_AwsEc2)
	if !ok {
		that2, ok := that.(Upstream_AwsEc2)
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
	if !this.AwsEc2.Equal(that1.AwsEc2) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
