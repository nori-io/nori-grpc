// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: nori.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_nori_proto protoreflect.FileDescriptor

var file_nori_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x72, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x08, 0x0a,
	0x04, 0x4e, 0x6f, 0x72, 0x69, 0x12, 0x42, 0x0a, 0x10, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x14, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x16, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x12, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x16, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_nori_proto_goTypes = []interface{}{
	(*PluginGetRequest)(nil),       // 0: plugin.PluginGetRequest
	(*PluginInstallRequest)(nil),   // 1: plugin.PluginInstallRequest
	(*PluginInterfaceRequest)(nil), // 2: plugin.PluginInterfaceRequest
	(*PluginListRequest)(nil),      // 3: plugin.PluginListRequest
	(*PluginMetaRequest)(nil),      // 4: plugin.PluginMetaRequest
	(*PluginPullRequest)(nil),      // 5: plugin.PluginPullRequest
	(*PluginRemoveRequest)(nil),    // 6: plugin.PluginRemoveRequest
	(*PluginStartRequest)(nil),     // 7: plugin.PluginStartRequest
	(*PluginStopRequest)(nil),      // 8: plugin.PluginStopRequest
	(*PluginUninstallRequest)(nil), // 9: plugin.PluginUninstallRequest
	(*PluginUploadRequest)(nil),    // 10: plugin.PluginUploadRequest
	(*ConfigGetRequest)(nil),       // 11: config.ConfigGetRequest
	(*ConfigSetRequest)(nil),       // 12: config.ConfigSetRequest
	(*ConfigUploadRequest)(nil),    // 13: config.ConfigUploadRequest
	(*ErrorReply)(nil),             // 14: common.ErrorReply
	(*PluginListReply)(nil),        // 15: plugin.PluginListReply
	(*PluginMetaReply)(nil),        // 16: plugin.PluginMetaReply
	(*ConfigGetReply)(nil),         // 17: config.ConfigGetReply
	(*ConfigSetReply)(nil),         // 18: config.ConfigSetReply
	(*ConfigUploadReply)(nil),      // 19: config.ConfigUploadReply
}
var file_nori_proto_depIdxs = []int32{
	0,  // 0: nori.Nori.PluginGetCommand:input_type -> plugin.PluginGetRequest
	1,  // 1: nori.Nori.PluginInstallCommand:input_type -> plugin.PluginInstallRequest
	2,  // 2: nori.Nori.PluginInterfaceCommand:input_type -> plugin.PluginInterfaceRequest
	3,  // 3: nori.Nori.PluginListCommand:input_type -> plugin.PluginListRequest
	4,  // 4: nori.Nori.PluginMetaCommand:input_type -> plugin.PluginMetaRequest
	5,  // 5: nori.Nori.PluginPullCommand:input_type -> plugin.PluginPullRequest
	6,  // 6: nori.Nori.PluginRemoveCommand:input_type -> plugin.PluginRemoveRequest
	7,  // 7: nori.Nori.PluginStartCommand:input_type -> plugin.PluginStartRequest
	8,  // 8: nori.Nori.PluginStopCommand:input_type -> plugin.PluginStopRequest
	9,  // 9: nori.Nori.PluginUninstallCommand:input_type -> plugin.PluginUninstallRequest
	10, // 10: nori.Nori.PluginUploadCommand:input_type -> plugin.PluginUploadRequest
	11, // 11: nori.Nori.ConfigGetCommand:input_type -> config.ConfigGetRequest
	12, // 12: nori.Nori.ConfigSetCommand:input_type -> config.ConfigSetRequest
	13, // 13: nori.Nori.ConfigUploadCommand:input_type -> config.ConfigUploadRequest
	14, // 14: nori.Nori.PluginGetCommand:output_type -> common.ErrorReply
	14, // 15: nori.Nori.PluginInstallCommand:output_type -> common.ErrorReply
	15, // 16: nori.Nori.PluginInterfaceCommand:output_type -> plugin.PluginListReply
	15, // 17: nori.Nori.PluginListCommand:output_type -> plugin.PluginListReply
	16, // 18: nori.Nori.PluginMetaCommand:output_type -> plugin.PluginMetaReply
	14, // 19: nori.Nori.PluginPullCommand:output_type -> common.ErrorReply
	14, // 20: nori.Nori.PluginRemoveCommand:output_type -> common.ErrorReply
	14, // 21: nori.Nori.PluginStartCommand:output_type -> common.ErrorReply
	14, // 22: nori.Nori.PluginStopCommand:output_type -> common.ErrorReply
	14, // 23: nori.Nori.PluginUninstallCommand:output_type -> common.ErrorReply
	14, // 24: nori.Nori.PluginUploadCommand:output_type -> common.ErrorReply
	17, // 25: nori.Nori.ConfigGetCommand:output_type -> config.ConfigGetReply
	18, // 26: nori.Nori.ConfigSetCommand:output_type -> config.ConfigSetReply
	19, // 27: nori.Nori.ConfigUploadCommand:output_type -> config.ConfigUploadReply
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_nori_proto_init() }
func file_nori_proto_init() {
	if File_nori_proto != nil {
		return
	}
	file_common_proto_init()
	file_config_proto_init()
	file_plugin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nori_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nori_proto_goTypes,
		DependencyIndexes: file_nori_proto_depIdxs,
	}.Build()
	File_nori_proto = out.File
	file_nori_proto_rawDesc = nil
	file_nori_proto_goTypes = nil
	file_nori_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NoriClient is the client API for Nori service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoriClient interface {
	//plugin
	PluginGetCommand(ctx context.Context, in *PluginGetRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginInstallCommand(ctx context.Context, in *PluginInstallRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginInterfaceCommand(ctx context.Context, in *PluginInterfaceRequest, opts ...grpc.CallOption) (*PluginListReply, error)
	PluginListCommand(ctx context.Context, in *PluginListRequest, opts ...grpc.CallOption) (*PluginListReply, error)
	PluginMetaCommand(ctx context.Context, in *PluginMetaRequest, opts ...grpc.CallOption) (*PluginMetaReply, error)
	PluginPullCommand(ctx context.Context, in *PluginPullRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginRemoveCommand(ctx context.Context, in *PluginRemoveRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginStartCommand(ctx context.Context, in *PluginStartRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginStopCommand(ctx context.Context, in *PluginStopRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginUninstallCommand(ctx context.Context, in *PluginUninstallRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginUploadCommand(ctx context.Context, in *PluginUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	//config
	ConfigGetCommand(ctx context.Context, in *ConfigGetRequest, opts ...grpc.CallOption) (*ConfigGetReply, error)
	ConfigSetCommand(ctx context.Context, in *ConfigSetRequest, opts ...grpc.CallOption) (*ConfigSetReply, error)
	ConfigUploadCommand(ctx context.Context, in *ConfigUploadRequest, opts ...grpc.CallOption) (*ConfigUploadReply, error)
}

type noriClient struct {
	cc grpc.ClientConnInterface
}

func NewNoriClient(cc grpc.ClientConnInterface) NoriClient {
	return &noriClient{cc}
}

func (c *noriClient) PluginGetCommand(ctx context.Context, in *PluginGetRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginGetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginInstallCommand(ctx context.Context, in *PluginInstallRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginInstallCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginInterfaceCommand(ctx context.Context, in *PluginInterfaceRequest, opts ...grpc.CallOption) (*PluginListReply, error) {
	out := new(PluginListReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginInterfaceCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginListCommand(ctx context.Context, in *PluginListRequest, opts ...grpc.CallOption) (*PluginListReply, error) {
	out := new(PluginListReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginListCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginMetaCommand(ctx context.Context, in *PluginMetaRequest, opts ...grpc.CallOption) (*PluginMetaReply, error) {
	out := new(PluginMetaReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginMetaCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginPullCommand(ctx context.Context, in *PluginPullRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginPullCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginRemoveCommand(ctx context.Context, in *PluginRemoveRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginRemoveCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginStartCommand(ctx context.Context, in *PluginStartRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginStartCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginStopCommand(ctx context.Context, in *PluginStopRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginStopCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginUninstallCommand(ctx context.Context, in *PluginUninstallRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginUninstallCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) PluginUploadCommand(ctx context.Context, in *PluginUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/PluginUploadCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) ConfigGetCommand(ctx context.Context, in *ConfigGetRequest, opts ...grpc.CallOption) (*ConfigGetReply, error) {
	out := new(ConfigGetReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/ConfigGetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) ConfigSetCommand(ctx context.Context, in *ConfigSetRequest, opts ...grpc.CallOption) (*ConfigSetReply, error) {
	out := new(ConfigSetReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/ConfigSetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noriClient) ConfigUploadCommand(ctx context.Context, in *ConfigUploadRequest, opts ...grpc.CallOption) (*ConfigUploadReply, error) {
	out := new(ConfigUploadReply)
	err := c.cc.Invoke(ctx, "/nori.Nori/ConfigUploadCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoriServer is the server API for Nori service.
type NoriServer interface {
	//plugin
	PluginGetCommand(context.Context, *PluginGetRequest) (*ErrorReply, error)
	PluginInstallCommand(context.Context, *PluginInstallRequest) (*ErrorReply, error)
	PluginInterfaceCommand(context.Context, *PluginInterfaceRequest) (*PluginListReply, error)
	PluginListCommand(context.Context, *PluginListRequest) (*PluginListReply, error)
	PluginMetaCommand(context.Context, *PluginMetaRequest) (*PluginMetaReply, error)
	PluginPullCommand(context.Context, *PluginPullRequest) (*ErrorReply, error)
	PluginRemoveCommand(context.Context, *PluginRemoveRequest) (*ErrorReply, error)
	PluginStartCommand(context.Context, *PluginStartRequest) (*ErrorReply, error)
	PluginStopCommand(context.Context, *PluginStopRequest) (*ErrorReply, error)
	PluginUninstallCommand(context.Context, *PluginUninstallRequest) (*ErrorReply, error)
	PluginUploadCommand(context.Context, *PluginUploadRequest) (*ErrorReply, error)
	//config
	ConfigGetCommand(context.Context, *ConfigGetRequest) (*ConfigGetReply, error)
	ConfigSetCommand(context.Context, *ConfigSetRequest) (*ConfigSetReply, error)
	ConfigUploadCommand(context.Context, *ConfigUploadRequest) (*ConfigUploadReply, error)
}

// UnimplementedNoriServer can be embedded to have forward compatible implementations.
type UnimplementedNoriServer struct {
}

func (*UnimplementedNoriServer) PluginGetCommand(context.Context, *PluginGetRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginGetCommand not implemented")
}
func (*UnimplementedNoriServer) PluginInstallCommand(context.Context, *PluginInstallRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInstallCommand not implemented")
}
func (*UnimplementedNoriServer) PluginInterfaceCommand(context.Context, *PluginInterfaceRequest) (*PluginListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInterfaceCommand not implemented")
}
func (*UnimplementedNoriServer) PluginListCommand(context.Context, *PluginListRequest) (*PluginListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginListCommand not implemented")
}
func (*UnimplementedNoriServer) PluginMetaCommand(context.Context, *PluginMetaRequest) (*PluginMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginMetaCommand not implemented")
}
func (*UnimplementedNoriServer) PluginPullCommand(context.Context, *PluginPullRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginPullCommand not implemented")
}
func (*UnimplementedNoriServer) PluginRemoveCommand(context.Context, *PluginRemoveRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginRemoveCommand not implemented")
}
func (*UnimplementedNoriServer) PluginStartCommand(context.Context, *PluginStartRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginStartCommand not implemented")
}
func (*UnimplementedNoriServer) PluginStopCommand(context.Context, *PluginStopRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginStopCommand not implemented")
}
func (*UnimplementedNoriServer) PluginUninstallCommand(context.Context, *PluginUninstallRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginUninstallCommand not implemented")
}
func (*UnimplementedNoriServer) PluginUploadCommand(context.Context, *PluginUploadRequest) (*ErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginUploadCommand not implemented")
}
func (*UnimplementedNoriServer) ConfigGetCommand(context.Context, *ConfigGetRequest) (*ConfigGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigGetCommand not implemented")
}
func (*UnimplementedNoriServer) ConfigSetCommand(context.Context, *ConfigSetRequest) (*ConfigSetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigSetCommand not implemented")
}
func (*UnimplementedNoriServer) ConfigUploadCommand(context.Context, *ConfigUploadRequest) (*ConfigUploadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigUploadCommand not implemented")
}

func RegisterNoriServer(s *grpc.Server, srv NoriServer) {
	s.RegisterService(&_Nori_serviceDesc, srv)
}

func _Nori_PluginGetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginGetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginGetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginGetCommand(ctx, req.(*PluginGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginInstallCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginInstallCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginInstallCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginInstallCommand(ctx, req.(*PluginInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginInterfaceCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginInterfaceCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginInterfaceCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginInterfaceCommand(ctx, req.(*PluginInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginListCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginListCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginListCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginListCommand(ctx, req.(*PluginListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginMetaCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginMetaCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginMetaCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginMetaCommand(ctx, req.(*PluginMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginPullCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginPullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginPullCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginPullCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginPullCommand(ctx, req.(*PluginPullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginRemoveCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginRemoveCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginRemoveCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginRemoveCommand(ctx, req.(*PluginRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginStartCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginStartCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginStartCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginStartCommand(ctx, req.(*PluginStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginStopCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginStopCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginStopCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginStopCommand(ctx, req.(*PluginStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginUninstallCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginUninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginUninstallCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginUninstallCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginUninstallCommand(ctx, req.(*PluginUninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_PluginUploadCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).PluginUploadCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/PluginUploadCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).PluginUploadCommand(ctx, req.(*PluginUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_ConfigGetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).ConfigGetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/ConfigGetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).ConfigGetCommand(ctx, req.(*ConfigGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_ConfigSetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).ConfigSetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/ConfigSetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).ConfigSetCommand(ctx, req.(*ConfigSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nori_ConfigUploadCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoriServer).ConfigUploadCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nori.Nori/ConfigUploadCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoriServer).ConfigUploadCommand(ctx, req.(*ConfigUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nori_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nori.Nori",
	HandlerType: (*NoriServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginGetCommand",
			Handler:    _Nori_PluginGetCommand_Handler,
		},
		{
			MethodName: "PluginInstallCommand",
			Handler:    _Nori_PluginInstallCommand_Handler,
		},
		{
			MethodName: "PluginInterfaceCommand",
			Handler:    _Nori_PluginInterfaceCommand_Handler,
		},
		{
			MethodName: "PluginListCommand",
			Handler:    _Nori_PluginListCommand_Handler,
		},
		{
			MethodName: "PluginMetaCommand",
			Handler:    _Nori_PluginMetaCommand_Handler,
		},
		{
			MethodName: "PluginPullCommand",
			Handler:    _Nori_PluginPullCommand_Handler,
		},
		{
			MethodName: "PluginRemoveCommand",
			Handler:    _Nori_PluginRemoveCommand_Handler,
		},
		{
			MethodName: "PluginStartCommand",
			Handler:    _Nori_PluginStartCommand_Handler,
		},
		{
			MethodName: "PluginStopCommand",
			Handler:    _Nori_PluginStopCommand_Handler,
		},
		{
			MethodName: "PluginUninstallCommand",
			Handler:    _Nori_PluginUninstallCommand_Handler,
		},
		{
			MethodName: "PluginUploadCommand",
			Handler:    _Nori_PluginUploadCommand_Handler,
		},
		{
			MethodName: "ConfigGetCommand",
			Handler:    _Nori_ConfigGetCommand_Handler,
		},
		{
			MethodName: "ConfigSetCommand",
			Handler:    _Nori_ConfigSetCommand_Handler,
		},
		{
			MethodName: "ConfigUploadCommand",
			Handler:    _Nori_ConfigUploadCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nori.proto",
}
