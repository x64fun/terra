// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: validator/validator.proto

package validator

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var file_validator_validator_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         82295900,
		Name:          "go.terra.validator.rule",
		Tag:           "bytes,82295900,opt,name=rule",
		Filename:      "validator/validator.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional string rule = 82295900;
	E_Rule = &file_validator_validator_proto_extTypes[0]
)

var File_validator_validator_proto protoreflect.FileDescriptor

var file_validator_validator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3a, 0x34, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0xf8, 0x9e, 0x27, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x36, 0x34, 0x66, 0x75, 0x6e, 0x2f, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_validator_validator_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_validator_validator_proto_depIdxs = []int32{
	0, // 0: go.terra.validator.rule:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_validator_validator_proto_init() }
func file_validator_validator_proto_init() {
	if File_validator_validator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_validator_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_validator_validator_proto_goTypes,
		DependencyIndexes: file_validator_validator_proto_depIdxs,
		ExtensionInfos:    file_validator_validator_proto_extTypes,
	}.Build()
	File_validator_validator_proto = out.File
	file_validator_validator_proto_rawDesc = nil
	file_validator_validator_proto_goTypes = nil
	file_validator_validator_proto_depIdxs = nil
}