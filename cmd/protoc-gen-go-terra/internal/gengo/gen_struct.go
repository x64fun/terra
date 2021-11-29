package gengo

import (
	"path/filepath"
	"strings"

	"github.com/x64fun/terra/pkg/protofile/kit"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (gen Generator) GenerateStruct(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+".go")
	importPath := getCurrentImportPath(file.GoImportPath.String(), file.GeneratedFilenamePrefix, file.GoImportPath.String())
	file.GoImportPath = protogen.GoImportPath(importPath)
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(file.GoImportPath))
	generateFileHeader(g, gen.Plugin, file)
	g.P()

	for _, message := range file.Messages {
		gen.generateMessage(g, file, message)
	}
	for _, enum := range file.Enums {
		gen.generateEnum(g, file, enum)
	}
	return g
}

func (gen Generator) generateMessage(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	g.P("type ", m.GoIdent.GoName, " struct {")
	for _, field := range m.Fields {
		_type := getFieldOptionKitFieldType(field)
		_tag := getFieldOptionKitTag(field)
		_dbField := getFieldOptionSqlxDBField(field)
		_validatorRule := getFieldOptionValidator(field)
		var printArgs []interface{}
		printArgs = append(printArgs, field.GoName, " ")
		if _type != nil {
			if _type.GetGoPkg() != "" {
				tempImport := protogen.GoImportPath(_type.GetGoPkg())
				printArgs = append(printArgs, tempImport.Ident(_type.GetGoType()))
			} else {
				printArgs = append(printArgs, _type.GetGoType())
			}
		} else {
			if field.Desc.IsList() {
				printArgs = append(printArgs, "[]")
			}
			switch field.Desc.Kind() {
			case protoreflect.BoolKind:
				printArgs = append(printArgs, "bool")
			case protoreflect.EnumKind:
				printArgs = append(printArgs, getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Enum.GoIdent.GoImportPath.String()).Ident(field.Enum.GoIdent.GoName))
			case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
				printArgs = append(printArgs, "int32")
			case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
				printArgs = append(printArgs, "uint32")
			case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
				printArgs = append(printArgs, "int64")
			case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
				printArgs = append(printArgs, "uint64")
			case protoreflect.FloatKind:
				printArgs = append(printArgs, "float32")
			case protoreflect.DoubleKind:
				printArgs = append(printArgs, "float64")
			case protoreflect.StringKind:
				printArgs = append(printArgs, "string")
			case protoreflect.BytesKind:
				printArgs = append(printArgs, "[]byte")
			case protoreflect.MessageKind:
				printArgs = append(printArgs, "*", getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Message.GoIdent.GoImportPath.String()).Ident(field.Message.GoIdent.GoName))
			case protoreflect.GroupKind:
			}
		}
		printTag := []string{}
		if _dbField != nil {
			printTag = append(printTag, `db:"`+_dbField.GetDb()+`"`)
		}
		if _tag != nil {
			if _tag.GetOmitempty() {
				printTag = append(printTag, `json:"`+_tag.GetName()+`,omitempty"`)
				printTag = append(printTag, `xml:"`+_tag.GetName()+`,omitempty"`)
			} else {
				printTag = append(printTag, `json:"`+_tag.GetName()+`"`)
				printTag = append(printTag, `xml:"`+_tag.GetName()+`"`)
			}
		} else {
			printTag = append(printTag, `json:"`+field.Desc.JSONName()+`,omitempty"`)
			printTag = append(printTag, `xml:"`+field.Desc.JSONName()+`,omitempty"`)
		}
		if _validatorRule != "" {
			printTag = append(printTag, `validate:"`+_validatorRule+`"`)
		}
		printArgs = append(printArgs, "`")
		printArgs = append(printArgs, strings.Join(printTag, " "))
		printArgs = append(printArgs, "`")
		g.P(printArgs...)
	}
	g.P("}")
	g.P()
	if gen.generateConvertProtobuf {
		g.P("func (x *", m.GoIdent.GoName, ") ConvertToProtoMessage() (*", m.GoIdent, ", error) {")
		g.P("var err error")
		g.P("m := &", m.GoIdent, "{}")
		for _, field := range m.Fields {
			typeOption := field.Desc.Options().ProtoReflect().Get(kit.E_Type.TypeDescriptor())
			_type := kit.E_Type.InterfaceOf(typeOption).(*kit.FieldType)
			var printArgs []interface{}
			printArgs = append(printArgs, "m.", field.GoName, " = ")
			if _type != nil {
				switch _type.GetGoPkg() {
				case "github.com/google/uuid":
					switch _type.GetGoType() {
					case "UUID":
						g.P(append(printArgs, utilPackage.Ident("ConvertUUIDToString"), "(x.", field.GoName, ")")...)
					case "NullUUID":
						g.P(append(printArgs, utilPackage.Ident("ConvertNullUUIDToString"), "(x.", field.GoName, ")")...)
					}
				case "github.com/jmoiron/sqlx/types":
					switch _type.GetGoType() {
					case "JSONText":
						g.P("err = x.", field.GoName, ".Unmarshal(&m.", field.GoName, ")")
						g.P("if err != nil {")
						g.P("return nil, err")
						g.P("}")
					case "NullJSONText":
						g.P("err = x.", field.GoName, ".Unmarshal(&m.", field.GoName, ")")
						g.P("if err != nil {")
						g.P("return nil, err")
						g.P("}")
					}
				case "database/sql":
					switch _type.GetGoType() {
					case "NullString":
						g.P(append(printArgs, utilPackage.Ident("ConvertNullStringToString"), "(x.", field.GoName, ")")...)
					case "NullInt64":
						g.P(append(printArgs, utilPackage.Ident("ConvertNullInt64ToInt64"), "(x.", field.GoName, ")")...)
					}
				default:
					switch _type.GetGoType() {
					case "interface{}",
						"map[string]interface{}":
						g.P("m.", field.GoName, ", err = ", utilPackage.Ident("ConvertInterfaceToStruct"), "(x.", field.GoName, ")")
						g.P("if err != nil {")
						g.P("return nil, err")
						g.P("}")
					case "uint8":
						g.P(append(printArgs, utilPackage.Ident("ConvertUint8ToUInt32"), "(x.", field.GoName, ")")...)
					}
				}
			} else {
				switch field.Desc.Kind() {
				case protoreflect.BoolKind,
					protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
					protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
					protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
					protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
					protoreflect.FloatKind,
					protoreflect.DoubleKind,
					protoreflect.StringKind,
					protoreflect.BytesKind:
					g.P(append(printArgs, "x.", field.GoName)...)
				case protoreflect.EnumKind:

				case protoreflect.MessageKind:
					if field.Desc.IsList() {
						g.P("for _, item := range x.", field.GoName, " {")
						g.P("tmp, err := item.ConvertToProtoMessage()")
						g.P("if err != nil {")
						g.P("return nil, err")
						g.P("}")
						g.P("m.", field.GoName, " = append(m.", field.GoName, ", tmp)")
						g.P("}")
					} else {
						g.P("m.", field.GoName, ", err = ", "x.", field.GoName, ".ConvertToProtoMessage()")
						g.P("if err != nil {")
						g.P("return nil, err")
						g.P("}")
					}
				case protoreflect.GroupKind:
				}
			}
		}
		g.P("return m, err")
		g.P("}")
	}
}
func (gen Generator) generateEnum(g *protogen.GeneratedFile, f *protogen.File, e *protogen.Enum) {
	g.P("type ", e.GoIdent.GoName, " string")
	g.P("const (")
	for _, v := range e.Values {
		g.P(e.GoIdent.GoName, "_", v.Desc.Name(), " ", e.GoIdent.GoName, ` = "`, v.Desc.Name(), `"`)
	}
	g.P(")")
}
