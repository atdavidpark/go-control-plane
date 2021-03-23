// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/extensions/access_loggers/stderror/v3/stderror.proto

package envoy_extensions_access_loggers_stderror_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Custom configuration for an :ref:`AccessLog <envoy_api_msg_config.accesslog.v3.AccessLog>`
// that writes log entries directly to the operating system's standard error.
type StdErrorAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccessLogFormat:
	//	*StdErrorAccessLog_LogFormat
	AccessLogFormat isStdErrorAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
}

func (x *StdErrorAccessLog) Reset() {
	*x = StdErrorAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdErrorAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdErrorAccessLog) ProtoMessage() {}

func (x *StdErrorAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdErrorAccessLog.ProtoReflect.Descriptor instead.
func (*StdErrorAccessLog) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescGZIP(), []int{0}
}

func (m *StdErrorAccessLog) GetAccessLogFormat() isStdErrorAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (x *StdErrorAccessLog) GetLogFormat() *v3.SubstitutionFormatString {
	if x, ok := x.GetAccessLogFormat().(*StdErrorAccessLog_LogFormat); ok {
		return x.LogFormat
	}
	return nil
}

type isStdErrorAccessLog_AccessLogFormat interface {
	isStdErrorAccessLog_AccessLogFormat()
}

type StdErrorAccessLog_LogFormat struct {
	// Configuration to form access log data and format.
	// If not specified, use :ref:`default format <config_access_log_default_format>`.
	LogFormat *v3.SubstitutionFormatString `protobuf:"bytes,1,opt,name=log_format,json=logFormat,proto3,oneof"`
}

func (*StdErrorAccessLog_LogFormat) isStdErrorAccessLog_AccessLogFormat() {}

var File_envoy_extensions_access_loggers_stderror_v3_stderror_proto protoreflect.FileDescriptor

var file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2f, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x64,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x59,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09,
	0x6c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x54,
	0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x53, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescOnce sync.Once
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescData = file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDesc
)

func file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescGZIP() []byte {
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescData)
	})
	return file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDescData
}

var file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_goTypes = []interface{}{
	(*StdErrorAccessLog)(nil),           // 0: envoy.extensions.access_loggers.stderror.v3.StdErrorAccessLog
	(*v3.SubstitutionFormatString)(nil), // 1: envoy.config.core.v3.SubstitutionFormatString
}
var file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.access_loggers.stderror.v3.StdErrorAccessLog.log_format:type_name -> envoy.config.core.v3.SubstitutionFormatString
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_init() }
func file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_init() {
	if File_envoy_extensions_access_loggers_stderror_v3_stderror_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdErrorAccessLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StdErrorAccessLog_LogFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_msgTypes,
	}.Build()
	File_envoy_extensions_access_loggers_stderror_v3_stderror_proto = out.File
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_rawDesc = nil
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_goTypes = nil
	file_envoy_extensions_access_loggers_stderror_v3_stderror_proto_depIdxs = nil
}
