// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api.proto

package pb

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

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9c,
	0x02, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x3a, 0x01, 0x2a,
	0x92, 0x41, 0x26, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0c, 0x53,
	0x65, 0x6e, 0x64, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x53, 0x4d, 0x53, 0x1a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x53, 0x4d, 0x53, 0x12, 0x81, 0x01, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4d, 0x53, 0x92, 0x41, 0x20, 0x0a, 0x08, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x53,
	0x4d, 0x53, 0x1a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x53, 0x4d, 0x53, 0x42, 0xa3, 0x01,
	0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x97, 0x01, 0x12, 0x5d, 0x0a, 0x0c,
	0x54, 0x65, 0x73, 0x74, 0x20, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x10,
	0x4e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x79, 0x20, 0x56, 0x61, 0x72, 0x6c, 0x61, 0x6d, 0x6f, 0x76,
	0x12, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x31, 0x34, 0x34, 0x1a, 0x16,
	0x76, 0x61, 0x72, 0x6c, 0x61, 0x6d, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6c, 0x40, 0x79, 0x61, 0x68,
	0x6f, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x33, 0x32, 0x30, 0x31, 0x2a, 0x02, 0x01, 0x02, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []interface{}{
	(*Messages_SendRequest)(nil),  // 0: api.Messages.SendRequest
	(*Messages_ListRequest)(nil),  // 1: api.Messages.ListRequest
	(*Messages_SendResponse)(nil), // 2: api.Messages.SendResponse
	(*Messages_ListResponse)(nil), // 3: api.Messages.ListResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.ApiService.SendMessage:input_type -> api.Messages.SendRequest
	1, // 1: api.ApiService.ListMessages:input_type -> api.Messages.ListRequest
	2, // 2: api.ApiService.SendMessage:output_type -> api.Messages.SendResponse
	3, // 3: api.ApiService.ListMessages:output_type -> api.Messages.ListResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_messages_proto_init()
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
