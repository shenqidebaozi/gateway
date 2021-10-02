// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gateway/config/v1/protocol.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Protocol int32

const (
	Protocol_UNKNOWN Protocol = 0
	Protocol_HTTP    Protocol = 1
	Protocol_GRPC    Protocol = 2
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "UNKNOWN",
		1: "HTTP",
		2: "GRPC",
	}
	Protocol_value = map[string]int32{
		"UNKNOWN": 0,
		"HTTP":    1,
		"GRPC":    2,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_config_v1_protocol_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_gateway_config_v1_protocol_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_gateway_config_v1_protocol_proto_rawDescGZIP(), []int{0}
}

var File_gateway_config_v1_protocol_proto protoreflect.FileDescriptor

var file_gateway_config_v1_protocol_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0x2b, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43,
	0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_config_v1_protocol_proto_rawDescOnce sync.Once
	file_gateway_config_v1_protocol_proto_rawDescData = file_gateway_config_v1_protocol_proto_rawDesc
)

func file_gateway_config_v1_protocol_proto_rawDescGZIP() []byte {
	file_gateway_config_v1_protocol_proto_rawDescOnce.Do(func() {
		file_gateway_config_v1_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_config_v1_protocol_proto_rawDescData)
	})
	return file_gateway_config_v1_protocol_proto_rawDescData
}

var file_gateway_config_v1_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_config_v1_protocol_proto_goTypes = []interface{}{
	(Protocol)(0), // 0: gateway.config.v1.Protocol
}
var file_gateway_config_v1_protocol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_config_v1_protocol_proto_init() }
func file_gateway_config_v1_protocol_proto_init() {
	if File_gateway_config_v1_protocol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_config_v1_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_config_v1_protocol_proto_goTypes,
		DependencyIndexes: file_gateway_config_v1_protocol_proto_depIdxs,
		EnumInfos:         file_gateway_config_v1_protocol_proto_enumTypes,
	}.Build()
	File_gateway_config_v1_protocol_proto = out.File
	file_gateway_config_v1_protocol_proto_rawDesc = nil
	file_gateway_config_v1_protocol_proto_goTypes = nil
	file_gateway_config_v1_protocol_proto_depIdxs = nil
}
