// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: nori.proto

package proto

import (
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

var File_nori_proto protoreflect.FileDescriptor

var file_nori_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x07,
	0x0a, 0x04, 0x4e, 0x6f, 0x72, 0x69, 0x12, 0x34, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x47, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x75,
	0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_nori_proto_goTypes = []interface{}{
	(*PluginRequest)(nil),          // 0: proto.PluginRequest
	(*PluginInstallRequest)(nil),   // 1: proto.PluginInstallRequest
	(*PluginInterfaceRequest)(nil), // 2: proto.PluginInterfaceRequest
	(*PluginListRequest)(nil),      // 3: proto.PluginListRequest
	(*PluginMetaRequest)(nil),      // 4: proto.PluginMetaRequest
	(*PluginPullRequest)(nil),      // 5: proto.PluginPullRequest
	(*PluginStartRequest)(nil),     // 6: proto.PluginStartRequest
	(*PluginStopRequest)(nil),      // 7: proto.PluginStopRequest
	(*PluginUninstallRequest)(nil), // 8: proto.PluginUninstallRequest
	(*PluginUploadChunk)(nil),      // 9: proto.PluginUploadChunk
	(*ConfigGetRequest)(nil),       // 10: proto.ConfigGetRequest
	(*ConfigSetRequest)(nil),       // 11: proto.ConfigSetRequest
	(*ConfigUploadRequest)(nil),    // 12: proto.ConfigUploadRequest
	(*Reply)(nil),                  // 13: proto.Reply
	(*PluginListReply)(nil),        // 14: proto.PluginListReply
	(*PluginMetaReply)(nil),        // 15: proto.PluginMetaReply
	(*ConfigGetReply)(nil),         // 16: proto.ConfigGetReply
	(*ConfigSetReply)(nil),         // 17: proto.ConfigSetReply
	(*ConfigUploadReply)(nil),      // 18: proto.ConfigUploadReply
}
var file_nori_proto_depIdxs = []int32{
	0,  // 0: proto.Nori.PluginEnable:input_type -> proto.PluginRequest
	0,  // 1: proto.Nori.PluginDisable:input_type -> proto.PluginRequest
	0,  // 2: proto.Nori.PluginGet:input_type -> proto.PluginRequest
	1,  // 3: proto.Nori.PluginInstall:input_type -> proto.PluginInstallRequest
	2,  // 4: proto.Nori.PluginInterface:input_type -> proto.PluginInterfaceRequest
	3,  // 5: proto.Nori.PluginList:input_type -> proto.PluginListRequest
	4,  // 6: proto.Nori.PluginMeta:input_type -> proto.PluginMetaRequest
	5,  // 7: proto.Nori.PluginPull:input_type -> proto.PluginPullRequest
	0,  // 8: proto.Nori.PluginRemove:input_type -> proto.PluginRequest
	6,  // 9: proto.Nori.PluginStart:input_type -> proto.PluginStartRequest
	7,  // 10: proto.Nori.PluginStop:input_type -> proto.PluginStopRequest
	8,  // 11: proto.Nori.PluginUninstall:input_type -> proto.PluginUninstallRequest
	9,  // 12: proto.Nori.PluginUpload:input_type -> proto.PluginUploadChunk
	10, // 13: proto.Nori.ConfigGet:input_type -> proto.ConfigGetRequest
	11, // 14: proto.Nori.ConfigSet:input_type -> proto.ConfigSetRequest
	12, // 15: proto.Nori.ConfigUpload:input_type -> proto.ConfigUploadRequest
	13, // 16: proto.Nori.PluginEnable:output_type -> proto.Reply
	13, // 17: proto.Nori.PluginDisable:output_type -> proto.Reply
	13, // 18: proto.Nori.PluginGet:output_type -> proto.Reply
	13, // 19: proto.Nori.PluginInstall:output_type -> proto.Reply
	14, // 20: proto.Nori.PluginInterface:output_type -> proto.PluginListReply
	14, // 21: proto.Nori.PluginList:output_type -> proto.PluginListReply
	15, // 22: proto.Nori.PluginMeta:output_type -> proto.PluginMetaReply
	13, // 23: proto.Nori.PluginPull:output_type -> proto.Reply
	13, // 24: proto.Nori.PluginRemove:output_type -> proto.Reply
	13, // 25: proto.Nori.PluginStart:output_type -> proto.Reply
	13, // 26: proto.Nori.PluginStop:output_type -> proto.Reply
	13, // 27: proto.Nori.PluginUninstall:output_type -> proto.Reply
	13, // 28: proto.Nori.PluginUpload:output_type -> proto.Reply
	16, // 29: proto.Nori.ConfigGet:output_type -> proto.ConfigGetReply
	17, // 30: proto.Nori.ConfigSet:output_type -> proto.ConfigSetReply
	18, // 31: proto.Nori.ConfigUpload:output_type -> proto.ConfigUploadReply
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
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
