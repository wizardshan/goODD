// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: rpc/user.proto

package rpc

import (
	request "chapter8/rpc/request"
	response "chapter8/rpc/response"
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

var File_rpc_user_proto protoreflect.FileDescriptor

var file_rpc_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x84, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x04, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x13, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x6d, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6d, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12,
	0x13, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04,
	0x43, 0x61, 0x73, 0x68, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x61, 0x73, 0x68, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x73, 0x68, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x31, 0x30, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_rpc_user_proto_goTypes = []any{
	(*request.UserOne)(nil),          // 0: request.UserOne
	(*request.UserMany)(nil),         // 1: request.UserMany
	(*request.UserLogin)(nil),        // 2: request.UserLogin
	(*request.UserSmsRegister)(nil),  // 3: request.UserSmsRegister
	(*request.UserRegister)(nil),     // 4: request.UserRegister
	(*request.UserModify)(nil),       // 5: request.UserModify
	(*request.UserCash)(nil),         // 6: request.UserCash
	(*response.UserOne)(nil),         // 7: response.UserOne
	(*response.UserMany)(nil),        // 8: response.UserMany
	(*response.UserLogin)(nil),       // 9: response.UserLogin
	(*response.UserSmsRegister)(nil), // 10: response.UserSmsRegister
	(*response.UserRegister)(nil),    // 11: response.UserRegister
	(*response.UserModify)(nil),      // 12: response.UserModify
	(*response.UserCash)(nil),        // 13: response.UserCash
}
var file_rpc_user_proto_depIdxs = []int32{
	0,  // 0: rpc.User.One:input_type -> request.UserOne
	1,  // 1: rpc.User.Many:input_type -> request.UserMany
	2,  // 2: rpc.User.Login:input_type -> request.UserLogin
	3,  // 3: rpc.User.SmsRegister:input_type -> request.UserSmsRegister
	4,  // 4: rpc.User.Register:input_type -> request.UserRegister
	5,  // 5: rpc.User.Modify:input_type -> request.UserModify
	6,  // 6: rpc.User.Cash:input_type -> request.UserCash
	7,  // 7: rpc.User.One:output_type -> response.UserOne
	8,  // 8: rpc.User.Many:output_type -> response.UserMany
	9,  // 9: rpc.User.Login:output_type -> response.UserLogin
	10, // 10: rpc.User.SmsRegister:output_type -> response.UserSmsRegister
	11, // 11: rpc.User.Register:output_type -> response.UserRegister
	12, // 12: rpc.User.Modify:output_type -> response.UserModify
	13, // 13: rpc.User.Cash:output_type -> response.UserCash
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_user_proto_init() }
func file_rpc_user_proto_init() {
	if File_rpc_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_user_proto_goTypes,
		DependencyIndexes: file_rpc_user_proto_depIdxs,
	}.Build()
	File_rpc_user_proto = out.File
	file_rpc_user_proto_rawDesc = nil
	file_rpc_user_proto_goTypes = nil
	file_rpc_user_proto_depIdxs = nil
}
