// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: checker.proto

package pb_hub

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 检查存活 请求参数
type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"` // nil-检查线程和文件 1-检查线程 2-检查文件
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_checker_proto_rawDescGZIP(), []int{0}
}

func (x *HealthRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

// 检查存活 响应参数
type HealthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAlive bool `protobuf:"varint,1,opt,name=isAlive,proto3" json:"isAlive,omitempty"`
}

func (x *HealthReply) Reset() {
	*x = HealthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthReply) ProtoMessage() {}

func (x *HealthReply) ProtoReflect() protoreflect.Message {
	mi := &file_checker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthReply.ProtoReflect.Descriptor instead.
func (*HealthReply) Descriptor() ([]byte, []int) {
	return file_checker_proto_rawDescGZIP(), []int{1}
}

func (x *HealthReply) GetIsAlive() bool {
	if x != nil {
		return x.IsAlive
	}
	return false
}

var File_checker_proto protoreflect.FileDescriptor

var file_checker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x23, 0x0a, 0x0d, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x27,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x32, 0x63, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x5a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x27, 0x2e, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x67, 0x69, 0x74, 0x2e, 0x77, 0x6f, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x6f, 0x74,
	0x61, 0x6c, 0x69, 0x75, 0x2f, 0x70, 0x62, 0x2d, 0x68, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_checker_proto_rawDescOnce sync.Once
	file_checker_proto_rawDescData = file_checker_proto_rawDesc
)

func file_checker_proto_rawDescGZIP() []byte {
	file_checker_proto_rawDescOnce.Do(func() {
		file_checker_proto_rawDescData = protoimpl.X.CompressGZIP(file_checker_proto_rawDescData)
	})
	return file_checker_proto_rawDescData
}

var file_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_checker_proto_goTypes = []interface{}{
	(*HealthRequest)(nil), // 0: trpc.checker.checkHealth.HealthRequest
	(*HealthReply)(nil),   // 1: trpc.checker.checkHealth.HealthReply
}
var file_checker_proto_depIdxs = []int32{
	0, // 0: trpc.checker.checkHealth.Check.Health:input_type -> trpc.checker.checkHealth.HealthRequest
	1, // 1: trpc.checker.checkHealth.Check.Health:output_type -> trpc.checker.checkHealth.HealthReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checker_proto_init() }
func file_checker_proto_init() {
	if File_checker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_checker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthReply); i {
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
			RawDescriptor: file_checker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checker_proto_goTypes,
		DependencyIndexes: file_checker_proto_depIdxs,
		MessageInfos:      file_checker_proto_msgTypes,
	}.Build()
	File_checker_proto = out.File
	file_checker_proto_rawDesc = nil
	file_checker_proto_goTypes = nil
	file_checker_proto_depIdxs = nil
}