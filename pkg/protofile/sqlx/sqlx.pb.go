// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: sqlx/sqlx.proto

package sqlx

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type DBField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db            string          `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`                                            // 字段映射
	Insert        bool            `protobuf:"varint,2,opt,name=insert,proto3" json:"insert,omitempty"`                                   // 插入时填入
	InsertDefault *DBFieldDefault `protobuf:"bytes,3,opt,name=insert_default,json=insertDefault,proto3" json:"insert_default,omitempty"` // 插入时默认值
	Update        []string        `protobuf:"bytes,4,rep,name=update,proto3" json:"update,omitempty"`                                    // 根据 Key 更新时填入
	UpdateDefault *DBFieldDefault `protobuf:"bytes,5,opt,name=update_default,json=updateDefault,proto3" json:"update_default,omitempty"` // 根据 Key 更新时默认值
	Delete        []string        `protobuf:"bytes,6,rep,name=delete,proto3" json:"delete,omitempty"`                                    // 根据 Key 删除时填入
	Replace       []string        `protobuf:"bytes,7,rep,name=replace,proto3" json:"replace,omitempty"`                                  // 根据 Key 替换时填入
	One           []string        `protobuf:"bytes,8,rep,name=one,proto3" json:"one,omitempty"`                                          // 根据 Key 查询单个时填入
	List          []string        `protobuf:"bytes,9,rep,name=list,proto3" json:"list,omitempty"`                                        // 根据 Key 查询列表时填入
}

func (x *DBField) Reset() {
	*x = DBField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sqlx_sqlx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBField) ProtoMessage() {}

func (x *DBField) ProtoReflect() protoreflect.Message {
	mi := &file_sqlx_sqlx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBField.ProtoReflect.Descriptor instead.
func (*DBField) Descriptor() ([]byte, []int) {
	return file_sqlx_sqlx_proto_rawDescGZIP(), []int{0}
}

func (x *DBField) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *DBField) GetInsert() bool {
	if x != nil {
		return x.Insert
	}
	return false
}

func (x *DBField) GetInsertDefault() *DBFieldDefault {
	if x != nil {
		return x.InsertDefault
	}
	return nil
}

func (x *DBField) GetUpdate() []string {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *DBField) GetUpdateDefault() *DBFieldDefault {
	if x != nil {
		return x.UpdateDefault
	}
	return nil
}

func (x *DBField) GetDelete() []string {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *DBField) GetReplace() []string {
	if x != nil {
		return x.Replace
	}
	return nil
}

func (x *DBField) GetOne() []string {
	if x != nil {
		return x.One
	}
	return nil
}

func (x *DBField) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type DBTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // 表名
	Update  []string `protobuf:"bytes,2,rep,name=update,proto3" json:"update,omitempty"`   // 根据 Key 更新时填入
	Delete  []string `protobuf:"bytes,3,rep,name=delete,proto3" json:"delete,omitempty"`   // 根据 Key 删除时填入
	Replace []string `protobuf:"bytes,4,rep,name=replace,proto3" json:"replace,omitempty"` // 替换
	One     []string `protobuf:"bytes,5,rep,name=one,proto3" json:"one,omitempty"`         // 根据 Key 查询单个时填入
	List    []string `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`       // 根据 Key 查询列表时填入
}

func (x *DBTable) Reset() {
	*x = DBTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sqlx_sqlx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTable) ProtoMessage() {}

func (x *DBTable) ProtoReflect() protoreflect.Message {
	mi := &file_sqlx_sqlx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTable.ProtoReflect.Descriptor instead.
func (*DBTable) Descriptor() ([]byte, []int) {
	return file_sqlx_sqlx_proto_rawDescGZIP(), []int{1}
}

func (x *DBTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DBTable) GetUpdate() []string {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *DBTable) GetDelete() []string {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *DBTable) GetReplace() []string {
	if x != nil {
		return x.Replace
	}
	return nil
}

func (x *DBTable) GetOne() []string {
	if x != nil {
		return x.One
	}
	return nil
}

func (x *DBTable) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type DBFieldDefault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoPkg   string `protobuf:"bytes,1,opt,name=go_pkg,json=goPkg,proto3" json:"go_pkg,omitempty"`
	GoValue string `protobuf:"bytes,2,opt,name=go_value,json=goValue,proto3" json:"go_value,omitempty"`
}

func (x *DBFieldDefault) Reset() {
	*x = DBFieldDefault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sqlx_sqlx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBFieldDefault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBFieldDefault) ProtoMessage() {}

func (x *DBFieldDefault) ProtoReflect() protoreflect.Message {
	mi := &file_sqlx_sqlx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBFieldDefault.ProtoReflect.Descriptor instead.
func (*DBFieldDefault) Descriptor() ([]byte, []int) {
	return file_sqlx_sqlx_proto_rawDescGZIP(), []int{2}
}

func (x *DBFieldDefault) GetGoPkg() string {
	if x != nil {
		return x.GoPkg
	}
	return ""
}

func (x *DBFieldDefault) GetGoValue() string {
	if x != nil {
		return x.GoValue
	}
	return ""
}

var file_sqlx_sqlx_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*DBField)(nil),
		Field:         82295600,
		Name:          "go.terra.sqlx.db_field",
		Tag:           "bytes,82295600,opt,name=db_field",
		Filename:      "sqlx/sqlx.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*DBTable)(nil),
		Field:         82295601,
		Name:          "go.terra.sqlx.db_table",
		Tag:           "bytes,82295601,opt,name=db_table",
		Filename:      "sqlx/sqlx.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional go.terra.sqlx.DBField db_field = 82295600;
	E_DbField = &file_sqlx_sqlx_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional go.terra.sqlx.DBTable db_table = 82295601;
	E_DbTable = &file_sqlx_sqlx_proto_extTypes[1]
)

var File_sqlx_sqlx_proto protoreflect.FileDescriptor

var file_sqlx_sqlx_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x71, 0x6c, 0x78, 0x2f, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x73, 0x71, 0x6c, 0x78,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x07, 0x44, 0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44,
	0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x69,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x6f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x42, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x44, 0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x6f, 0x5f, 0x70, 0x6b, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f, 0x50, 0x6b, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x53, 0x0a, 0x08, 0x64, 0x62, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xb0, 0xf6, 0x9e, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x42, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x07, 0x64, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x55, 0x0a, 0x08, 0x64,
	0x62, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb1, 0xf6, 0x9e, 0x27, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x2e, 0x73, 0x71, 0x6c,
	0x78, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x64, 0x62, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x36, 0x34, 0x66, 0x75, 0x6e, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x71, 0x6c, 0x78,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sqlx_sqlx_proto_rawDescOnce sync.Once
	file_sqlx_sqlx_proto_rawDescData = file_sqlx_sqlx_proto_rawDesc
)

func file_sqlx_sqlx_proto_rawDescGZIP() []byte {
	file_sqlx_sqlx_proto_rawDescOnce.Do(func() {
		file_sqlx_sqlx_proto_rawDescData = protoimpl.X.CompressGZIP(file_sqlx_sqlx_proto_rawDescData)
	})
	return file_sqlx_sqlx_proto_rawDescData
}

var file_sqlx_sqlx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sqlx_sqlx_proto_goTypes = []interface{}{
	(*DBField)(nil),                   // 0: go.terra.sqlx.DBField
	(*DBTable)(nil),                   // 1: go.terra.sqlx.DBTable
	(*DBFieldDefault)(nil),            // 2: go.terra.sqlx.DBFieldDefault
	(*descriptor.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
}
var file_sqlx_sqlx_proto_depIdxs = []int32{
	2, // 0: go.terra.sqlx.DBField.insert_default:type_name -> go.terra.sqlx.DBFieldDefault
	2, // 1: go.terra.sqlx.DBField.update_default:type_name -> go.terra.sqlx.DBFieldDefault
	3, // 2: go.terra.sqlx.db_field:extendee -> google.protobuf.FieldOptions
	4, // 3: go.terra.sqlx.db_table:extendee -> google.protobuf.MessageOptions
	0, // 4: go.terra.sqlx.db_field:type_name -> go.terra.sqlx.DBField
	1, // 5: go.terra.sqlx.db_table:type_name -> go.terra.sqlx.DBTable
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	4, // [4:6] is the sub-list for extension type_name
	2, // [2:4] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sqlx_sqlx_proto_init() }
func file_sqlx_sqlx_proto_init() {
	if File_sqlx_sqlx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sqlx_sqlx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBField); i {
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
		file_sqlx_sqlx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBTable); i {
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
		file_sqlx_sqlx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBFieldDefault); i {
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
			RawDescriptor: file_sqlx_sqlx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_sqlx_sqlx_proto_goTypes,
		DependencyIndexes: file_sqlx_sqlx_proto_depIdxs,
		MessageInfos:      file_sqlx_sqlx_proto_msgTypes,
		ExtensionInfos:    file_sqlx_sqlx_proto_extTypes,
	}.Build()
	File_sqlx_sqlx_proto = out.File
	file_sqlx_sqlx_proto_rawDesc = nil
	file_sqlx_sqlx_proto_goTypes = nil
	file_sqlx_sqlx_proto_depIdxs = nil
}
