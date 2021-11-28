package gengo

import (
	"strings"

	"github.com/golang/glog"
	"github.com/x64fun/terra/cmd/protoc-gen-go-terra/internal/version"
	"github.com/x64fun/terra/internal/tool"
	"github.com/x64fun/terra/pkg/protofile/kit"
	"github.com/x64fun/terra/pkg/protofile/sqlx"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

// SupportedFeatures reports the set of supported protobuf language features.
var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

const deprecationComment = "// Deprecated: Do not use."

type Generator struct {
	*protogen.Plugin
	requireUnimplemented    bool
	generateConvertProtobuf bool
	l                       glog.Verbose
}

func NewGenerator(p *protogen.Plugin, requireUnimplemented bool, generateConvertProtobuf bool) *Generator {
	return &Generator{Plugin: p, requireUnimplemented: requireUnimplemented, generateConvertProtobuf: generateConvertProtobuf, l: glog.V(1)}
}

// Standard library dependencies.
const (
	contextPackage = protogen.GoImportPath("context")
	errorsPackage  = protogen.GoImportPath("errors")
	netHTTPPackage = protogen.GoImportPath("net/http")
	syncPackage    = protogen.GoImportPath("sync")
	jsonPackage    = protogen.GoImportPath("encoding/json")
	xmlPackage     = protogen.GoImportPath("encoding/xml")
	ioPackage      = protogen.GoImportPath("io")
	ioutilPackage  = protogen.GoImportPath("io/ioutil")
	stringsPackage = protogen.GoImportPath("strings")
	timePackage    = protogen.GoImportPath("time")
	sqlPackage     = protogen.GoImportPath("database/sql")
	strconvPackage = protogen.GoImportPath("strconv")
)

// go-kit library dependencies.
const (
	kitLogPackage      = protogen.GoImportPath("github.com/go-kit/log")
	kitEndpointPackage = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	kitServerPackage   = protogen.GoImportPath("github.com/go-kit/kit/server")
	kitClientPackage   = protogen.GoImportPath("github.com/go-kit/kit/client")

	kitTransportHTTPPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	kitTransportGRPCPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
)

// sqlx library dependencies.
const (
	sqlxPackage = protogen.GoImportPath("github.com/jmoiron/sqlx")
)

// uuid library dependencies.
const (
	uuidPackage = protogen.GoImportPath("github.com/google/uuid")
)

// library dependencies.
const (
	utilPackage = protogen.GoImportPath("github.com/x64fun/terra/pkg/util")
)

// 获取当前文件的引用路径 | 区别于protobuf默认路径
func getCurrentImportPath(fileImportPath, filenamePrefix, importPath string) protogen.GoImportPath {

	fileImportPath = strings.Trim(fileImportPath, `"`)
	importPath = strings.Trim(importPath, `"`)
	fileImportPathSplit := strings.Split(fileImportPath, "/")
	importPathIndex := 1
	for {
		if strings.HasPrefix(importPath, strings.Join(fileImportPathSplit[:len(fileImportPathSplit)-importPathIndex], "/")) {
			break
		}
		importPathIndex++
	}
	if strings.Join(fileImportPathSplit[:len(fileImportPathSplit)-importPathIndex], "/") == "github.com" {
		return protogen.GoImportPath(importPath)
	}

	prefixSplit := strings.Split(filenamePrefix, "/")
	index := 1
	prefix := strings.Join(prefixSplit[:len(prefixSplit)-index], "/")
	for {
		if strings.HasSuffix(importPath, prefix) {
			break
		} else {
			index++
			prefix = strings.Join(prefixSplit[:len(prefixSplit)-index], "/")
		}
	}

	if prefix == "" {
		return protogen.GoImportPath(importPath + "/terra")
	} else {
		return protogen.GoImportPath(strings.Replace(importPath, prefix, "terra/"+prefix, 1))
	}
}

// 小驼峰命名
func camel(s string) string {
	currentIndex := 0
	for i := 0; i < len(s); i++ {
		tmp := string(s[i])
		if strings.ToUpper(tmp) == tmp {
			continue
		} else {
			s = s[0:currentIndex] + strings.ToLower(s[currentIndex:i]) + s[i:]
			currentIndex = i
			break
		}
	}
	if currentIndex == 0 {
		s = strings.ToLower(s[currentIndex:])
	}
	return s
}

// 生成文件头部
func generateFileHeader(g *protogen.GeneratedFile, plugin *protogen.Plugin, file *protogen.File) {
	g.P("// Code generated by protoc-gen-go-terra. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-terra ", version.String())
	g.P("// - protoc              ", tool.ProtocVersion(plugin))
	g.P("package ", file.GoPackageName)
	g.P()
}

func getMessageOptions(m *protogen.Message) protoreflect.Message {
	return m.Desc.Options().ProtoReflect()
}
func getMessageOptionSqlxDBTable(m *protogen.Message) *sqlx.DBTable {
	option := getMessageOptions(m).Get(sqlx.E_DbTable.TypeDescriptor())
	return sqlx.E_DbTable.InterfaceOf(option).(*sqlx.DBTable)
}

func getFieldOptions(f *protogen.Field) protoreflect.Message {
	return f.Desc.Options().ProtoReflect()
}

func getFieldOptionSqlxDBField(f *protogen.Field) *sqlx.DBField {
	option := getFieldOptions(f).Get(sqlx.E_DbField.TypeDescriptor())
	return sqlx.E_DbField.InterfaceOf(option).(*sqlx.DBField)
}
func getFieldOptionKitFieldType(f *protogen.Field) *kit.FieldType {
	option := getFieldOptions(f).Get(kit.E_Type.TypeDescriptor())
	return kit.E_Type.InterfaceOf(option).(*kit.FieldType)
}
func getFieldOptionKitTag(f *protogen.Field) *kit.StructTag {
	option := getFieldOptions(f).Get(kit.E_Tag.TypeDescriptor())
	return kit.E_Tag.InterfaceOf(option).(*kit.StructTag)
}
