// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_json/grpc_json.proto

package grpc_json

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa94c5eabc74996, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	Equal(interface{}) bool
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof" json:"proto_descriptor,omitempty"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof" json:"proto_descriptor_bin,omitempty"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa94c5eabc74996, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "grpc_json.options.gloo.solo.io.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "grpc_json.options.gloo.solo.io.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_json/grpc_json.proto", fileDescriptor_caa94c5eabc74996)
}

var fileDescriptor_caa94c5eabc74996 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x4e, 0x14, 0x31,
	0x14, 0xc6, 0x19, 0xfe, 0x39, 0x94, 0x05, 0xb5, 0x22, 0x8c, 0x1b, 0xc4, 0xd5, 0x68, 0xb2, 0x89,
	0x71, 0x26, 0xe2, 0x8d, 0xd1, 0x0b, 0xc2, 0x44, 0x14, 0x8c, 0xe2, 0x3a, 0x6a, 0x4c, 0xbc, 0x69,
	0xca, 0xcc, 0x71, 0x28, 0xec, 0xb4, 0xa5, 0xed, 0x2c, 0xf0, 0x1a, 0xbe, 0x81, 0x77, 0x3e, 0x82,
	0xcf, 0xe3, 0x1b, 0x78, 0x61, 0x62, 0xbc, 0x32, 0x6d, 0x87, 0x65, 0x85, 0x68, 0xbc, 0xeb, 0x9c,
	0xef, 0xf7, 0x9d, 0x33, 0xe7, 0xf4, 0x14, 0x6d, 0x97, 0xcc, 0xec, 0xd6, 0x3b, 0x71, 0x2e, 0xaa,
	0x44, 0x8b, 0xbe, 0xb8, 0xc7, 0x44, 0x52, 0xf6, 0x85, 0x48, 0xa4, 0x12, 0x7b, 0x90, 0x1b, 0xed,
	0xbf, 0xa8, 0x64, 0xc9, 0xe0, 0x7e, 0x22, 0xa4, 0x61, 0x82, 0xeb, 0xa4, 0x54, 0x32, 0x27, 0x7b,
	0x5a, 0xf0, 0xd3, 0x53, 0x2c, 0x95, 0x30, 0x02, 0xaf, 0x9c, 0x06, 0x1a, 0x38, 0xb6, 0x09, 0x62,
	0x9b, 0x3b, 0x66, 0xa2, 0xbd, 0x50, 0x8a, 0x52, 0x38, 0x34, 0xb1, 0x27, 0xef, 0x6a, 0x63, 0x38,
	0x32, 0x3e, 0x08, 0x47, 0xa6, 0x89, 0x2d, 0x0d, 0x68, 0x9f, 0x15, 0xd4, 0x40, 0x72, 0x72, 0xf0,
	0xc2, 0xad, 0xcf, 0xd3, 0x08, 0x3f, 0x53, 0x32, 0x7f, 0xae, 0x05, 0x7f, 0xab, 0x28, 0xd7, 0xb9,
	0x28, 0x40, 0xe1, 0xbb, 0xe8, 0x92, 0xd3, 0x49, 0x01, 0x3a, 0x57, 0x4c, 0x1a, 0xa1, 0xa2, 0xa0,
	0x13, 0x74, 0x67, 0x36, 0xc7, 0xb2, 0x8b, 0x4e, 0x79, 0x32, 0x14, 0xf0, 0x2a, 0x5a, 0x38, 0x0b,
	0x93, 0x1d, 0xc6, 0xa3, 0xc9, 0x4e, 0xd0, 0x6d, 0x6d, 0x8e, 0x65, 0xf8, 0x8c, 0x21, 0x65, 0x1c,
	0xdf, 0x46, 0xa1, 0x06, 0x35, 0x60, 0x39, 0xe8, 0x68, 0xbc, 0x33, 0xd1, 0x9d, 0x49, 0xc3, 0x5f,
	0xe9, 0xd4, 0xa7, 0x60, 0x3c, 0x0c, 0xb2, 0xa1, 0x82, 0x0b, 0x34, 0x27, 0x15, 0xe3, 0x86, 0x34,
	0xed, 0x47, 0x13, 0x9d, 0xa0, 0x3b, 0xbb, 0xba, 0x16, 0xff, 0x7b, 0x30, 0xf1, 0xf9, 0x8e, 0xe2,
	0x9e, 0xcd, 0xf3, 0xca, 0xc3, 0x59, 0x4b, 0x8e, 0x7c, 0xe1, 0x35, 0xb4, 0x5c, 0x51, 0x93, 0xef,
	0x12, 0xc6, 0x73, 0x51, 0x31, 0x5e, 0x12, 0x05, 0x07, 0x35, 0x68, 0x43, 0x94, 0xa8, 0x0d, 0x44,
	0x53, 0x9d, 0xa0, 0x1b, 0x66, 0xd7, 0x1c, 0xb3, 0xd5, 0x20, 0x99, 0x27, 0x32, 0x0b, 0xe0, 0x87,
	0x28, 0x62, 0x25, 0x17, 0x0a, 0x0a, 0x72, 0x50, 0x83, 0x3a, 0x26, 0x92, 0x2a, 0x5a, 0x81, 0x01,
	0xa5, 0xa3, 0x69, 0xdb, 0x5c, 0xb6, 0xd8, 0xe8, 0xaf, 0xad, 0xdc, 0x1b, 0xaa, 0xf8, 0x26, 0x6a,
	0xd1, 0xda, 0x08, 0x52, 0x51, 0x29, 0x19, 0x2f, 0xa3, 0x0b, 0xae, 0xd4, 0xac, 0x8d, 0xbd, 0xf4,
	0x21, 0xbc, 0x81, 0x6e, 0x78, 0x33, 0xa9, 0xf9, 0x3e, 0x17, 0x87, 0xfc, 0x7c, 0x8d, 0xd0, 0xb9,
	0x96, 0x3d, 0xf6, 0xce, 0x53, 0x67, 0x2b, 0xc5, 0xe8, 0x4a, 0x2e, 0xf8, 0x00, 0x94, 0x21, 0x6e,
	0x78, 0xda, 0x50, 0x53, 0xeb, 0x68, 0xc6, 0x59, 0x2f, 0x37, 0x92, 0x9d, 0xdb, 0x1b, 0x27, 0xb4,
	0xbf, 0x07, 0xa8, 0x35, 0x3a, 0x33, 0x7c, 0x07, 0xcd, 0xd3, 0xa2, 0x20, 0x87, 0xbb, 0xcc, 0x80,
	0x96, 0x34, 0x07, 0xb7, 0x10, 0x61, 0x36, 0x47, 0x8b, 0xe2, 0xfd, 0x30, 0x88, 0xd7, 0xd1, 0x75,
	0xda, 0x3f, 0xa4, 0xc7, 0x9a, 0xf8, 0x9b, 0x93, 0x8a, 0x55, 0xcc, 0xb0, 0x01, 0x90, 0x8f, 0x0c,
	0xfa, 0x85, 0xbd, 0x6d, 0xeb, 0x6a, 0x7b, 0xc8, 0x55, 0xe8, 0x9d, 0x20, 0x4f, 0x1d, 0x81, 0x1f,
	0xa1, 0xf6, 0x1f, 0x29, 0x80, 0xd7, 0x95, 0x26, 0x54, 0x13, 0xc6, 0x8d, 0x5f, 0x81, 0x30, 0x5b,
	0x1c, 0xf1, 0x6f, 0x58, 0x7d, 0x5d, 0x6f, 0x71, 0xa3, 0xf1, 0x63, 0xd4, 0x96, 0x0a, 0xec, 0x02,
	0x01, 0xf1, 0x4b, 0xe9, 0xca, 0x12, 0x4e, 0x2b, 0xd0, 0x6e, 0x23, 0xc3, 0x6c, 0xe9, 0x84, 0xe8,
	0x59, 0xc0, 0x15, 0xdd, 0xb6, 0x72, 0x7a, 0x15, 0xcd, 0x8f, 0xac, 0xb0, 0x06, 0x83, 0x27, 0x7e,
	0xa6, 0x41, 0xfa, 0xe2, 0xeb, 0x8f, 0xc9, 0xe0, 0xcb, 0xb7, 0x95, 0xe0, 0x43, 0xfa, 0x7f, 0x0f,
	0x5c, 0xee, 0x97, 0x7f, 0x7d, 0xe4, 0x3b, 0xd3, 0xee, 0xb7, 0x1e, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x98, 0xbd, 0xc2, 0x89, 0x2d, 0x04, 0x00, 0x00,
}

func (this *GrpcJsonTranscoder) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder)
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
	if that1.DescriptorSet == nil {
		if this.DescriptorSet != nil {
			return false
		}
	} else if this.DescriptorSet == nil {
		return false
	} else if !this.DescriptorSet.Equal(that1.DescriptorSet) {
		return false
	}
	if len(this.Services) != len(that1.Services) {
		return false
	}
	for i := range this.Services {
		if this.Services[i] != that1.Services[i] {
			return false
		}
	}
	if !this.PrintOptions.Equal(that1.PrintOptions) {
		return false
	}
	if this.MatchIncomingRequestRoute != that1.MatchIncomingRequestRoute {
		return false
	}
	if len(this.IgnoredQueryParameters) != len(that1.IgnoredQueryParameters) {
		return false
	}
	for i := range this.IgnoredQueryParameters {
		if this.IgnoredQueryParameters[i] != that1.IgnoredQueryParameters[i] {
			return false
		}
	}
	if this.AutoMapping != that1.AutoMapping {
		return false
	}
	if this.IgnoreUnknownQueryParameters != that1.IgnoreUnknownQueryParameters {
		return false
	}
	if this.ConvertGrpcStatus != that1.ConvertGrpcStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_ProtoDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_ProtoDescriptor)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_ProtoDescriptor)
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
	if this.ProtoDescriptor != that1.ProtoDescriptor {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_ProtoDescriptorBin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_ProtoDescriptorBin)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_ProtoDescriptorBin)
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
	if !bytes.Equal(this.ProtoDescriptorBin, that1.ProtoDescriptorBin) {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_PrintOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_PrintOptions)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_PrintOptions)
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
	if this.AddWhitespace != that1.AddWhitespace {
		return false
	}
	if this.AlwaysPrintPrimitiveFields != that1.AlwaysPrintPrimitiveFields {
		return false
	}
	if this.AlwaysPrintEnumsAsInts != that1.AlwaysPrintEnumsAsInts {
		return false
	}
	if this.PreserveProtoFieldNames != that1.PreserveProtoFieldNames {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
