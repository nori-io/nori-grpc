// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: config.proto

package proto

import (
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

// cmd: config get
type ConfigGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *ID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfigGetRequest) Reset() {
	*x = ConfigGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigGetRequest) ProtoMessage() {}

func (x *ConfigGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigGetRequest.ProtoReflect.Descriptor instead.
func (*ConfigGetRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigGetRequest) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

type ConfigGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[string]string `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigGetReply) Reset() {
	*x = ConfigGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigGetReply) ProtoMessage() {}

func (x *ConfigGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigGetReply.ProtoReflect.Descriptor instead.
func (*ConfigGetReply) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigGetReply) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

// cmd: config set
type ConfigSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *ID    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigSetRequest) Reset() {
	*x = ConfigSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSetRequest) ProtoMessage() {}

func (x *ConfigSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSetRequest.ProtoReflect.Descriptor instead.
func (*ConfigSetRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigSetRequest) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ConfigSetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ConfigSetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ConfigSetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ConfigSetReply) Reset() {
	*x = ConfigSetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSetReply) ProtoMessage() {}

func (x *ConfigSetReply) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSetReply.ProtoReflect.Descriptor instead.
func (*ConfigSetReply) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigSetReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// cmd: config upload
type ConfigUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConfigUploadRequest) Reset() {
	*x = ConfigUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigUploadRequest) ProtoMessage() {}

func (x *ConfigUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigUploadRequest.ProtoReflect.Descriptor instead.
func (*ConfigUploadRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigUploadRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type ConfigUploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[string]string `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigUploadReply) Reset() {
	*x = ConfigUploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigUploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigUploadReply) ProtoMessage() {}

func (x *ConfigUploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigUploadReply.ProtoReflect.Descriptor instead.
func (*ConfigUploadReply) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigUploadReply) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x7a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x55,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2d, 0x0a, 0x13, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x33, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_proto_goTypes = []interface{}{
	(*ConfigGetRequest)(nil),    // 0: proto.ConfigGetRequest
	(*ConfigGetReply)(nil),      // 1: proto.ConfigGetReply
	(*ConfigSetRequest)(nil),    // 2: proto.ConfigSetRequest
	(*ConfigSetReply)(nil),      // 3: proto.ConfigSetReply
	(*ConfigUploadRequest)(nil), // 4: proto.ConfigUploadRequest
	(*ConfigUploadReply)(nil),   // 5: proto.ConfigUploadReply
	nil,                         // 6: proto.ConfigGetReply.MapEntry
	nil,                         // 7: proto.ConfigUploadReply.MapEntry
	(*ID)(nil),                  // 8: proto.ID
	(*Error)(nil),               // 9: proto.Error
}
var file_config_proto_depIdxs = []int32{
	8, // 0: proto.ConfigGetRequest.id:type_name -> proto.ID
	6, // 1: proto.ConfigGetReply.map:type_name -> proto.ConfigGetReply.MapEntry
	8, // 2: proto.ConfigSetRequest.id:type_name -> proto.ID
	9, // 3: proto.ConfigSetReply.error:type_name -> proto.Error
	7, // 4: proto.ConfigUploadReply.map:type_name -> proto.ConfigUploadReply.MapEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigGetRequest); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigGetReply); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSetRequest); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSetReply); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigUploadRequest); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigUploadReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}