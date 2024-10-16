// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: domain/vo/id.proto

package vo

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	mi := &file_domain_vo_id_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_domain_vo_id_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_domain_vo_id_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_domain_vo_id_proto protoreflect.FileDescriptor

var file_domain_vo_id_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x6f, 0x2f, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x23, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x01, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x32, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_domain_vo_id_proto_rawDescOnce sync.Once
	file_domain_vo_id_proto_rawDescData = file_domain_vo_id_proto_rawDesc
)

func file_domain_vo_id_proto_rawDescGZIP() []byte {
	file_domain_vo_id_proto_rawDescOnce.Do(func() {
		file_domain_vo_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_vo_id_proto_rawDescData)
	})
	return file_domain_vo_id_proto_rawDescData
}

var file_domain_vo_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_vo_id_proto_goTypes = []any{
	(*ID)(nil), // 0: vo.ID
}
var file_domain_vo_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_vo_id_proto_init() }
func file_domain_vo_id_proto_init() {
	if File_domain_vo_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_vo_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_vo_id_proto_goTypes,
		DependencyIndexes: file_domain_vo_id_proto_depIdxs,
		MessageInfos:      file_domain_vo_id_proto_msgTypes,
	}.Build()
	File_domain_vo_id_proto = out.File
	file_domain_vo_id_proto_rawDesc = nil
	file_domain_vo_id_proto_goTypes = nil
	file_domain_vo_id_proto_depIdxs = nil
}
