package gengo

import (
	"path/filepath"
	"strings"

	"github.com/x64fun/terra/pkg/protofile/kit"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (gen Generator) GenerateDAODefault(file *protogen.File, daoDefaultMap map[string]bool) {
	filename := filepath.Join("terra", strings.Join(strings.Split(file.GeneratedFilenamePrefix, "/")[0:len(strings.Split(file.GeneratedFilenamePrefix, "/"))-1], "/")+"/dao.go")
	if _, ok := daoDefaultMap[filename]; !ok {
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		generateFileHeader(g, gen.Plugin, file)
		g.P()
		g.P(`var DEFAULT_TABLE_PREFIX = "" // 全局表前缀`)
		daoDefaultMap[filename] = true
	}
}

func (gen Generator) GenerateDAOMySQL(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+"_dao_mysql.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(g, gen.Plugin, file)
	g.P()
	for _, message := range file.Messages {
		gen.generateDAOMySQLMessage(g, file, message)
	}
	return g
}
func (gen Generator) generateDAOMySQLMessage(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	_dbTable := getMessageOptionSqlxDBTable(m)
	if _dbTable != nil {
		g.P("func MySQL", m.GoIdent.GoName, "TableName(prefix string) string {")
		g.P(`return "`, "`", `" + prefix + "`, _dbTable.Name, "`", `"`)
		g.P("}")
	}
}

func (gen Generator) GenerateDAOPostgres(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+"_dao_postgres.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(g, gen.Plugin, file)
	g.P()
	for _, message := range file.Messages {
		gen.generateDAOPostgresMessage(g, file, message)
	}
	return g
}
func (gen Generator) generateDAOPostgresMessage(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	_dbTable := getMessageOptionSqlxDBTable(m)
	if _dbTable != nil {
		g.P("func Postgres", m.GoIdent.GoName, "TableName(schema string) string {")
		g.P("return schema + `.", `"`, _dbTable.Name, `"`, "`")
		g.P("}")
		g.P("const (")
		for _, field := range m.Fields {
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				g.P("postgres", m.GoIdent.GoName, "TableColumn", field.GoName, " = ", "`", `"`, _dbField.Db, `"`, "`")
			}
		}
		g.P(")")
		g.P("var postgres", m.GoIdent.GoName, "TableAllColumn = []string{")
		for _, field := range m.Fields {
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				g.P("postgres", m.GoIdent.GoName, "TableColumn", field.GoName, ",")
			}
		}
		g.P("}")
		for _, field := range m.Fields {
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				g.P("func Postgres", m.GoIdent.GoName, "TableColumn", field.GoName, "(schema string) string {")
				g.P("return Postgres", m.GoIdent.GoName, "TableName(schema) + `.` + postgres", m.GoIdent.GoName, "TableColumn", field.GoName)
				g.P("}")
			}
		}
		for _, replaceKey := range _dbTable.GetReplace() {
			g.P("func generatePostgresReplace", m.GoIdent.GoName, "By", replaceKey, "SQL(schema string) string {")
			g.P("conflictKeyList := []string{")
			for _, key := range strings.Split(replaceKey, "And") {
				g.P("postgres", m.GoIdent.GoName, "TableColumn", key, ",")
			}
			g.P("}")
			g.P("updateList := []string{")
			for _, field := range m.Fields {
				_dbField := getFieldOptionSqlxDBField(field)
				if _dbField != nil {
					for _, _replaceKey := range _dbField.GetReplace() {
						if replaceKey == _replaceKey {
							g.P("postgres", m.GoIdent.GoName, "TableColumn", field.GoName, ` + " = EXCLUDED." + postgres`, m.GoIdent.GoName, "TableColumn", field.GoName, ",")
						}
					}
				}
			}
			g.P("postgres", m.GoIdent.GoName, `TableColumnUpdatedAt + " = EXCLUDED." + postgres`, m.GoIdent.GoName, "TableColumnCreatedAt,")
			g.P("}")
			g.P(`query := "INSERT INTO " + Postgres`, m.GoIdent.GoName, `TableName(schema) + " (" +`)
			g.P(`postgres`, m.GoIdent.GoName, `TableColumnID + ", " +`)
			for _, field := range m.Fields {
				_dbField := getFieldOptionSqlxDBField(field)
				if _dbField != nil {
					for _, _replaceKey := range _dbField.GetReplace() {
						if replaceKey == _replaceKey {
							g.P(`postgres`, m.GoIdent.GoName, `TableColumn`, field.GoName, ` + ", " +`)
						}
					}
				}
			}
			g.P(`postgres`, m.GoIdent.GoName, `TableColumnCreatedAt +`)
			g.P(`") VALUES (" +`)
			g.P(m.GoIdent.GoName, `FieldIDNamedMapping + ", " +`)
			for _, field := range m.Fields {
				_dbField := getFieldOptionSqlxDBField(field)
				if _dbField != nil {
					for _, _replaceKey := range _dbField.GetReplace() {
						if replaceKey == _replaceKey {
							g.P(m.GoIdent.GoName, `Field`, field.GoName, `NamedMapping + ", " +`)
						}
					}
				}
			}
			g.P(m.GoIdent.GoName, `FieldCreatedAtNamedMapping +`)
			g.P(`") ON CONFLICT (" + `, stringsPackage.Ident(`Join(conflictKeyList, ", ")`), ` + `)
			g.P(`") DO UPDATE SET " + `, stringsPackage.Ident(`Join(updateList, ", ")`))
			g.P(`return query`)
			g.P("}")
		}
	}
	g.P()
}
func (gen Generator) GenerateDAO(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+"_dao.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(g, gen.Plugin, file)
	g.P()

	for _, message := range file.Messages {
		gen.generateDAOMessage(g, file, message)
	}
	return g
}

func (gen Generator) generateDAOMessage(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	_dbTable := getMessageOptionSqlxDBTable(m)
	if _dbTable != nil {
		gen.generateDAODBField(g, f, m)
		g.P()
		gen.generateDAOWhere(g, f, m)
		g.P()
		gen.generateDAOInsert(g, f, m)
		g.P()
		for _, updateKey := range _dbTable.GetUpdate() {
			gen.generateDAOUpdate(g, f, m, updateKey)
		}
		for _, replaceKey := range _dbTable.GetReplace() {
			gen.generateDAOReplace(g, f, m, replaceKey)
		}
		for _, deleteKey := range _dbTable.GetDelete() {
			gen.generateDAODelete(g, f, m, deleteKey)
		}

		for _, oneKey := range _dbTable.GetOne() {
			gen.generateDAOSelect(g, f, "one", false, m, oneKey, func(g *protogen.GeneratedFile, structName, key string) {
				g.P(`whereList := []string{}`)
				g.P(`whereDataList := []interface{}{}`)
				for _, k := range strings.Split(key, "And") {
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, k, `[tx.DriverName()] + " = ?")`)
					g.P(`whereDataList = append(whereDataList, `, camel(k), `)`)
				}
				g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, m.GoIdent.GoName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + strings.Join(whereList, " AND ") + " LIMIT 1"`)
				g.P(`query = tx.Rebind(query)`)
			})
			if len(strings.Split(oneKey, "And")) == 1 {
				gen.generateDAOSelect(g, f, "one", true, m, oneKey, func(g *protogen.GeneratedFile, structName, key string) {
					g.P(`step := 1000 // 每次最多查询1000个`)
					g.P(`max := len(`, camel(key), `List)`)
					g.P(`for index := 0; index < max; index += step {`)
					g.P(`start := index`)
					g.P(`end := index + step`)
					g.P(`if end > max { end = max }`)
					g.P(`var tmpList []*`, structName)
					g.P(`whereList := []string{}`)
					g.P(`whereDataList := []interface{}{}`)
					g.P(`{`)
					g.P(`var q string`)
					g.P(`var a []interface{}`)
					g.P(`q, a, err = sqlx.In(`, structName, `TableColumn`, key, `[tx.DriverName()] + " IN (?)", `, camel(key), `List[start:end])`)
					g.P(`if err != nil {`)
					g.P(`return`)
					g.P(`}`)
					g.P(`whereList = append(whereList, q)`)
					g.P(`whereDataList = append(whereDataList, a...)`)
					g.P(`}`)
					g.P(`whereList = append(whereList, `, structName, `TableColumnDeletedAt[tx.DriverName()] + " IS NULL")`)
					g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + strings.Join(whereList, " AND ")`)
					g.P(`query = tx.Rebind(query)`)
					g.P(`err = tx.SelectContext(ctx, &tmpList, query, whereDataList...)`)
					g.P(`if err != nil {`)
					g.P(`return`)
					g.P(`}`)
					g.P(`list = append(list, tmpList...)`)
					g.P(`}`)
				})
			}
		}
		for _, listKey := range _dbTable.GetList() {
			gen.generateDAOSelect(g, f, "list", false, m, listKey, func(g *protogen.GeneratedFile, structName, key string) {
				g.P(`whereList := []string{}`)
				g.P(`whereDataList := []interface{}{}`)
				for _, k := range strings.Split(key, "And") {
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, k, `[tx.DriverName()] + " = ?")`)
					g.P(`whereDataList = append(whereDataList, `, camel(k), `)`)
				}
				g.P(`whereList = append(whereList, `, structName, `TableColumnDeletedAt[tx.DriverName()] + " IS NULL")`)
				g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + strings.Join(whereList, " AND ")`)
				g.P(`query = tx.Rebind(query)`)
				g.P(`err = tx.SelectContext(ctx, &list, query, whereDataList...)`)
				g.P(`if err != nil {`)
				g.P(`return`)
				g.P(`}`)
			})
			if len(strings.Split(listKey, "And")) == 1 {
				gen.generateDAOSelect(g, f, "list", true, m, listKey, func(g *protogen.GeneratedFile, structName, key string) {
					g.P(`list = make([]*`, structName, `, 0)`)
					g.P(`step := 1000 // 每次最多查询1000个`)
					g.P(`max := len(`, camel(listKey), `List)`)
					g.P(`for index := 0; index < max; index += step {`)
					g.P(`start := index`)
					g.P(`end := index + step`)
					g.P(`if end > max { end = max }`)
					g.P(`var tmpList []*`, structName)
					g.P(`whereList := []string{}`)
					g.P(`whereDataList := []interface{}{}`)
					g.P(`{`)
					g.P(`var q string`)
					g.P(`var a []interface{}`)
					g.P(`q, a, err = sqlx.In(`, structName, `TableColumn`, listKey, `[tx.DriverName()] + " IN (?)", `, camel(listKey), `List[start:end])`)
					g.P(`if err != nil {`)
					g.P(`return`)
					g.P(`}`)
					g.P(`whereList = append(whereList, q)`)
					g.P(`whereDataList = append(whereDataList, a...)`)
					g.P(`}`)
					g.P(`whereList = append(whereList, `, structName, `TableColumnDeletedAt[tx.DriverName()] + " IS NULL")`)
					g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + strings.Join(whereList, " AND ")`)
					g.P(`query = tx.Rebind(query)`)
					g.P(`err = tx.SelectContext(ctx, &tmpList, query, whereDataList...)`)
					g.P(`if err != nil {`)
					g.P(`return`)
					g.P(`}`)
					g.P(`list = append(list, tmpList...)`)
					g.P(`}`)
				})
			}
		}
		g.P(`func Get`, m.GoIdent.GoName, `AllList(`)
		g.P(`ctx `, contextPackage.Ident("Context"), `,`)
		g.P(`tx *`, sqlxPackage.Ident("Tx"), `,`)
		g.P(`order *`, utilPackage.Ident("Order"), `,`)
		g.P(`columns ...string,`)
		g.P(`) (list []*`, m.GoIdent.GoName, `, err error) {`)
		g.P(`if len(columns) == 0 { return }`)
		g.P(`list = make([]*`, m.GoIdent.GoName, `, 0)`)
		g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, m.GoIdent.GoName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + `, m.GoIdent.GoName, `TableColumnDeletedAt[tx.DriverName()] + " IS NULL"`)
		g.P(`if order != nil {`)
		g.P(`query += " ORDER BY " + order.Field + " " + string(order.Direction)`)
		g.P(`}`)
		g.P(`query = tx.Rebind(query)`)
		g.P(`err = tx.SelectContext(ctx, &list, query)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`return`)
		g.P(`}`)

		g.P(`func Get`, m.GoIdent.GoName, `Count(`)
		g.P(`ctx `, contextPackage.Ident("Context"), `,`)
		g.P(`tx *`, sqlxPackage.Ident("Tx"), `,`)
		g.P(`where *`, m.GoIdent.GoName, `Where,`)
		g.P(`) (count uint64, err error) {`)
		g.P(`whereDataList := []interface{}{}`)
		g.P(`query := "SELECT COUNT(" + `, m.GoIdent.GoName, `TableColumnID[tx.DriverName()] + ") FROM " + `, m.GoIdent.GoName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX)`)
		g.P(`if where != nil {`)
		g.P(`var query string`)
		g.P(`var args []interface{}`)
		g.P(`query, args, err = Generate`, m.GoIdent.GoName, `WhereSQL(tx.DriverName(), where)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`query += " WHERE " + query`)
		g.P(`whereDataList = append(whereDataList, args...)`)
		g.P(`}`)
		g.P(`err = tx.GetContext(ctx, &count, tx.Rebind(query), whereDataList...)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`return`)
		g.P(`}`)

		g.P(`func Get`, m.GoIdent.GoName, `List(`)
		g.P(`ctx `, contextPackage.Ident("Context"), `,`)
		g.P(`tx *`, sqlxPackage.Ident("Tx"), `,`)
		g.P(`pageNum, pageSize int64,`)
		g.P(`order *`, utilPackage.Ident("Order"), `,`)
		g.P(`where *`, m.GoIdent.GoName, `Where,`)
		g.P(`columns ...string,`)
		g.P(`) (list []*`, m.GoIdent.GoName, `, total uint64, err error) {`)
		g.P(`if len(columns) == 0 { return }`)
		g.P(`list = make([]*`, m.GoIdent.GoName, `, 0)`)
		g.P(`whereDataList := []interface{}{}`)
		g.P(`query := "SELECT " + strings.Join(columns, ", ") + " FROM " + `, m.GoIdent.GoName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX)`)
		g.P(`if where != nil {`)
		g.P(`var query string`)
		g.P(`var args []interface{}`)
		g.P(`query, args, err = Generate`, m.GoIdent.GoName, `WhereSQL(tx.DriverName(), where)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`query += " WHERE " + query`)
		g.P(`whereDataList = append(whereDataList, args...)`)
		g.P(`}`)
		g.P(`if order != nil {`)
		g.P(`query += " ORDER BY " + order.Field + " " + string(order.Direction)`)
		g.P(`}`)
		g.P(`offset := (pageNum - 1) * pageSize`)
		g.P(`count := pageSize`)
		g.P(`if count != 0 {`)
		g.P(`query += " LIMIT "`)
		g.P(`if offset != 0 {`)
		g.P(`query += `, strconvPackage.Ident("FormatInt(offset, 10)"), ` + ", "`)
		g.P(`}`)
		g.P(`query += `, strconvPackage.Ident("FormatInt(count, 10)"))
		g.P(`}`)
		g.P(`total, err = Get`, m.GoIdent.GoName, `Count(ctx, tx, where)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`err = tx.SelectContext(ctx, &list, tx.Rebind(query), whereDataList...)`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`return`)
		g.P(`}`)
	}
}
func (gen Generator) generateDAODBField(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	structName := m.GoIdent.GoName
	g.P("const (")
	for _, field := range m.Fields {
		_dbField := getFieldOptionSqlxDBField(field)
		if _dbField != nil {
			g.P(structName, "Field", field.GoName, `NamedMapping = ":`, _dbField.GetDb(), `"`)
		}
	}
	g.P(")")
	g.P("var (")
	g.P(structName, "TableName = map[string]func(string) string {")
	g.P(`"postgres": Postgres`, structName, `TableName,`)
	g.P("}")
	for _, field := range m.Fields {
		_dbField := getFieldOptionSqlxDBField(field)
		if _dbField != nil {
			g.P(structName, "TableColumn", field.GoName, " = map[string]string {")
			g.P(`"postgres": postgres`, structName, `TableColumn`, field.GoName, `,`)
			g.P("}")
		}
	}
	g.P(structName, "TableAllColumn = map[string][]string{")
	g.P(`"postgres": postgres`, structName, `TableAllColumn,`)
	g.P("}")
	g.P(")")
}
func (gen Generator) generateDAOWhere(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	structName := m.GoIdent.GoName
	structWhereName := structName + "Where"
	g.P("type ", structWhereName, " struct {")
	g.P("Not *", structWhereName, " `json:", `"not,omitempty"`, "`")
	g.P("Or []*", structWhereName, " `json:", `"or,omitempty"`, "`")
	g.P("And []*", structWhereName, " `json:", `"and,omitempty"`, "`")
	for _, field := range m.Fields {
		_type := getFieldOptionKitFieldType(field)
		_tag := getFieldOptionKitTag(field)
		_dbField := getFieldOptionSqlxDBField(field)
		if _dbField != nil {
			printFlag := true
			var fieldType interface{}
			fieldCanBeNull := false
			if _type != nil {
				fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident(_type.GetGoType())
				switch _type.GetGoPkg() {
				case "github.com/google/uuid":
					goType := _type.GetGoType()
					goType = strings.TrimPrefix(goType, "Null")
					fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident(goType)
					switch _type.GetGoType() {
					case "NullUUID":
						fieldCanBeNull = true
					}
				case "github.com/jmoiron/sqlx/types":
					printFlag = false
				case "database/sql":
					switch _type.GetGoType() {
					case "NullString":
						fieldType = "string"
						fieldCanBeNull = true
					case "NullInt64":
						fieldType = "int64"
						fieldCanBeNull = true
					case "NullInt32":
						fieldType = "int32"
						fieldCanBeNull = true
					case "NullInt16":
						fieldType = "int16"
						fieldCanBeNull = true
					case "NullByte":
						fieldType = "byte"
						fieldCanBeNull = true
					case "NullFloat64":
						fieldType = "float64"
						fieldCanBeNull = true
					case "NullBool":
						fieldType = "bool"
						fieldCanBeNull = true
					case "NullTime":
						fieldType = protogen.GoImportPath("time").Ident("Time")
						fieldCanBeNull = true
					}
				default:
					fieldType = _type.GetGoType()
					switch _type.GetGoType() {
					}
				}
			} else {
				switch field.Desc.Kind() {
				case protoreflect.BoolKind:
					fieldType = "bool"
				case protoreflect.EnumKind:
					fieldType = getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Enum.GoIdent.GoImportPath.String()).Ident(field.Enum.GoIdent.GoName)
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
					fieldType = "int32"
				case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
					fieldType = "uint32"
				case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
					fieldType = "int64"
				case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
					fieldType = "uint64"
				case protoreflect.FloatKind:
					fieldType = "float32"
				case protoreflect.DoubleKind:
					fieldType = "float64"
				case protoreflect.StringKind:
					fieldType = "string"
				case protoreflect.BytesKind:
					fieldType = "[]byte"
				case protoreflect.MessageKind:
					fieldType = getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Message.GoIdent.GoImportPath.String()).Ident(field.Message.GoIdent.GoName)
				case protoreflect.GroupKind:
				}
			}
			var jsonTag string
			if _tag != nil {
				if _tag.GetName() != "" {
					jsonTag = _tag.GetName()
				}
			} else {
				jsonTag = field.Desc.JSONName()
			}
			if printFlag {
				g.P(field.GoName, " *", fieldType, " `json:", `"`, jsonTag, `,omitempty"`, "`")
				g.P(field.GoName, "NEQ *", fieldType, " `json:", `"`, jsonTag, `NEQ,omitempty"`, "`")
				if fieldType != "bool" {
					g.P(field.GoName, "In *", fieldType, " `json:", `"`, jsonTag, `In,omitempty"`, "`")
					g.P(field.GoName, "NotIn *", fieldType, " `json:", `"`, jsonTag, `NotIn,omitempty"`, "`")
				}
				if fieldType == "string" ||
					fieldType == "int8" || fieldType == "int16" || fieldType == "int" || fieldType == "int32" || fieldType == "int64" ||
					fieldType == "uint8" || fieldType == "uint16" || fieldType == "uint" || fieldType == "uint32" || fieldType == "uint64" ||
					fieldType == "float32" || fieldType == "float64" ||
					fieldType == "time.Time" {
					g.P(field.GoName, "GT *", fieldType, " `json:", `"`, jsonTag, `GT,omitempty"`, "`")
					g.P(field.GoName, "GTE *", fieldType, " `json:", `"`, jsonTag, `GTE,omitempty"`, "`")
					g.P(field.GoName, "LT *", fieldType, " `json:", `"`, jsonTag, `LT,omitempty"`, "`")
					g.P(field.GoName, "LTE *", fieldType, " `json:", `"`, jsonTag, `LTE,omitempty"`, "`")
				}
				if fieldType == "string" {
					g.P(field.GoName, "Contains *", fieldType, " `json:", `"`, jsonTag, `Contains,omitempty"`, "`")
					g.P(field.GoName, "HasPrefix *", fieldType, " `json:", `"`, jsonTag, `HasPrefix,omitempty"`, "`")
					g.P(field.GoName, "HasSuffix *", fieldType, " `json:", `"`, jsonTag, `HasSuffix,omitempty"`, "`")
					g.P(field.GoName, "EqualFold *", fieldType, " `json:", `"`, jsonTag, `EqualFold,omitempty"`, "`")
					g.P(field.GoName, "ContainsFold *", fieldType, " `json:", `"`, jsonTag, `ContainsFold,omitempty"`, "`")
				}
				if fieldCanBeNull {
					g.P(field.GoName, "IsNull bool `json:", `"`, jsonTag, `IsNull,omitempty"`, "`")
					g.P(field.GoName, "IsNotNull bool `json:", `"`, jsonTag, `IsNotNull,omitempty"`, "`")
				}
			}
		}
	}
	g.P("}")
	g.P(`func Generate`, structWhereName, `SQL(driverName string, w *`, m.GoIdent.GoName, `Where) (string, []interface{}, error) {`)
	g.P(`if w == nil {`)
	g.P(`return "", nil, nil`)
	g.P(`}`)
	g.P(`whereList := []string{}`)
	g.P(`whereDataList := []interface{}{}`)
	g.P(`if w.Not != nil {`)
	g.P(`notQuery, args, err := Generate`, structWhereName, `SQL(driverName, w.Not)`)
	g.P(`if err != nil {`)
	g.P(`return "", nil, err`)
	g.P(`}`)
	g.P(`whereList = append(whereList, "NOT (" + notQuery + ")")`)
	g.P(`whereDataList = append(whereDataList, args...)`)
	g.P("}")
	g.P(`if w.Or != nil {`)
	g.P(`orList := []string{}`)
	g.P(`for _, or := range w.Or {`)
	g.P(`orQuery, args, err := Generate`, structWhereName, `SQL(driverName, or)`)
	g.P(`if err != nil {`)
	g.P(`return "", nil, err`)
	g.P(`}`)
	g.P(`orList = append(orList, "(" + orQuery + ")")`)
	g.P(`whereDataList = append(whereDataList, args...)`)
	g.P(`}`)
	g.P(`whereList = append(whereList, "(" + `, stringsPackage.Ident(`Join(orList, " OR ")`), ` + ")")`)
	g.P(`}`)
	g.P(`if w.And != nil {`)
	g.P(`andList := []string{}`)
	g.P(`for _, and := range w.And {`)
	g.P(`andQuery, args, err := Generate`, structWhereName, `SQL(driverName, and)`)
	g.P(`if err != nil {`)
	g.P(`return "", nil, err`)
	g.P(`}`)
	g.P(`andList = append(andList, "(" + andQuery + ")")`)
	g.P(`whereDataList = append(whereDataList, args...)`)
	g.P(`}`)
	g.P(`whereList = append(whereList, "(" + `, stringsPackage.Ident(`Join(andList, " AND ")`), ` + ")")`)
	g.P(`}`)
	for _, field := range m.Fields {
		_dbField := getFieldOptionSqlxDBField(field)
		_type := getFieldOptionKitFieldType(field)
		fieldName := field.GoName

		if _dbField != nil {
			printFlag := true
			var fieldType interface{}
			fieldCanBeNull := false
			if _type != nil {
				fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident(_type.GetGoType())
				switch _type.GetGoPkg() {
				case "github.com/google/uuid":
					switch _type.GetGoType() {
					case "NullUUID":
						fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident("UUID")
						fieldCanBeNull = true
					}
				case "github.com/jmoiron/sqlx/types":
					printFlag = false
				case "database/sql":
					switch _type.GetGoType() {
					case "NullString":
						fieldType = "string"
						fieldCanBeNull = true
					case "NullInt64":
						fieldType = "int64"
						fieldCanBeNull = true
					case "NullInt32":
						fieldType = "int32"
						fieldCanBeNull = true
					case "NullInt16":
						fieldType = "int16"
						fieldCanBeNull = true
					case "NullByte":
						fieldType = "byte"
						fieldCanBeNull = true
					case "NullFloat64":
						fieldType = "float64"
						fieldCanBeNull = true
					case "NullBool":
						fieldType = "bool"
						fieldCanBeNull = true
					case "NullTime":
						fieldType = protogen.GoImportPath("time").Ident("Time")
						fieldCanBeNull = true
					}
				default:
					fieldType = _type.GetGoType()
					switch _type.GetGoType() {
					}
				}
			} else {
				switch field.Desc.Kind() {
				case protoreflect.BoolKind:
					fieldType = "bool"
				case protoreflect.EnumKind:
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
					fieldType = "int32"
				case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
					fieldType = "uint32"
				case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
					fieldType = "int64"
				case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
					fieldType = "uint64"
				case protoreflect.FloatKind:
					fieldType = "float32"
				case protoreflect.DoubleKind:
					fieldType = "float64"
				case protoreflect.StringKind:
					fieldType = "string"
				case protoreflect.BytesKind:
					fieldType = "[]byte"
				case protoreflect.MessageKind:
					fieldType = getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Message.GoIdent.GoImportPath.String()).Ident(field.Message.GoIdent.GoName)
				case protoreflect.GroupKind:
				}
			}
			if printFlag {
				g.P(`if w.`, fieldName, ` != nil {`)
				g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " = ?")`)
				g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `)`)
				g.P(`}`)
				if fieldType != "bool" {
					g.P(`if w.`, fieldName, `In != nil {`)
					g.P(`query, args, err := `, sqlxPackage.Ident("In"), `(`, structName, `TableColumn`, fieldName, `[driverName] + " IN (?)", *w.`, fieldName, `In)`)
					g.P(`if err != nil {`)
					g.P(`return "", nil, err`)
					g.P(`}`)
					g.P(`whereList = append(whereList, query)`)
					g.P(`whereDataList = append(whereDataList, args...)`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `NotIn != nil {`)
					g.P(`query, args, err := `, sqlxPackage.Ident("In"), `(`, structName, `TableColumn`, fieldName, `[driverName] + " NOT IN (?)", *w.`, fieldName, `NotIn)`)
					g.P(`if err != nil {`)
					g.P(`return "", nil, err`)
					g.P(`}`)
					g.P(`whereList = append(whereList, query)`)
					g.P(`whereDataList = append(whereDataList, args...)`)
					g.P(`}`)
				}
				if fieldType == "string" ||
					fieldType == "int8" || fieldType == "int16" || fieldType == "int" || fieldType == "int32" || fieldType == "int64" ||
					fieldType == "uint8" || fieldType == "uint16" || fieldType == "uint" || fieldType == "uint32" || fieldType == "uint64" ||
					fieldType == "float32" || fieldType == "float64" ||
					fieldType == "time.Time" {
					g.P(`if w.`, fieldName, `GT != nil {`)
					g.P(`whereList = append(whereList, `, m.GoIdent.GoName, `TableColumn`, fieldName, `[driverName] + " > ?")`)
					g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `GT)`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `GTE != nil {`)
					g.P(`whereList = append(whereList, `, m.GoIdent.GoName, `TableColumn`, fieldName, `[driverName] + " >= ?")`)
					g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `GTE)`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `LT != nil {`)
					g.P(`whereList = append(whereList, `, m.GoIdent.GoName, `TableColumn`, fieldName, `[driverName] + " < ?")`)
					g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `LT)`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `LTE != nil {`)
					g.P(`whereList = append(whereList, `, m.GoIdent.GoName, `TableColumn`, fieldName, `[driverName] + " <= ?")`)
					g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `LTE)`)
					g.P(`}`)
				}
				if fieldType == "string" {
					g.P(`if w.`, fieldName, `Contains != nil {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " LIKE ?")`)
					g.P(`whereDataList = append(whereDataList, "%"+*w.`, fieldName, `Contains+"%")`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `HasPrefix != nil {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " LIKE ?")`)
					g.P(`whereDataList = append(whereDataList, *w.`, fieldName, `HasPrefix+"%")`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `HasSuffix != nil {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " LIKE ?")`)
					g.P(`whereDataList = append(whereDataList, "%"+*w.`, fieldName, `HasSuffix)`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `EqualFold != nil {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " = ?")`)
					g.P(`whereDataList = append(whereDataList, strings.ToLower(*w.`, fieldName, `EqualFold))`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `ContainsFold != nil {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " LIKE ?")`)
					g.P(`whereDataList = append(whereDataList, "%"+strings.ToLower(*w.`, fieldName, `ContainsFold)+"%")`)
					g.P(`}`)
				}
				if fieldCanBeNull {
					g.P(`if w.`, fieldName, `IsNull {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " IS NULL")`)
					g.P(`}`)
					g.P(`if w.`, fieldName, `IsNotNull {`)
					g.P(`whereList = append(whereList, `, structName, `TableColumn`, fieldName, `[driverName] + " IS NOT NULL")`)
					g.P(`}`)
				}
			}
		}
	}
	g.P(`return `, stringsPackage.Ident(`Join(whereList, " AND ")`), `, whereDataList, nil`)
	g.P(`}`)
}
func (gen Generator) generateDAOInsert(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	gen.generateDAOOperate(g, m, "Insert", "", func(g *protogen.GeneratedFile, batch bool, structName, key string) {
		if batch {
			g.P(`if len(tmp) == 0 { return }`)
			g.P(`list = tmp`)
			g.P(`for index := range list {`)
			g.P(`list[index].ID = `, uuidPackage.Ident("New()"))
			g.P(`list[index].CreatedAt = `, timePackage.Ident("Now().Unix()"))
			g.P(`}`)
		} else {
			g.P(`if tmp == nil { return }`)
			g.P(`data = tmp`)
			g.P(`data.ID = `, uuidPackage.Ident("New()"))
			g.P(`data.CreatedAt = `, timePackage.Ident("Now().Unix()"))
		}
		g.P(`query := "INSERT INTO " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " (" +`)
		g.P(structName, `TableColumnID[tx.DriverName()] + ", " +`)
		for _, field := range m.Fields {
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				if _dbField.GetInsert() {
					g.P(structName, `TableColumn`, field.GoName, `[tx.DriverName()] + ", " +`)
				}
			}
		}
		g.P(structName, `TableColumnCreatedAt[tx.DriverName()] +`)
		g.P(`") VALUES (" +`)
		g.P(structName, `FieldIDNamedMapping + ", " +`)
		for _, field := range m.Fields {
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				if _dbField.GetInsert() {
					g.P(structName, `Field`, field.GoName, `NamedMapping + ", " +`)
				}
			}
		}
		g.P(structName, `FieldCreatedAtNamedMapping +`)
		g.P(`")"`)
	})
}
func (gen Generator) generateDAOUpdate(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message, key string) {
	gen.generateDAOOperate(g, m, "Update", key, func(g *protogen.GeneratedFile, batch bool, structName, key string) {
		if batch {
			g.P(`if len(tmp) == 0 { return }`)
			g.P(`list = tmp`)
			g.P(`for index := range list {`)
			g.P(`list[index].UpdatedAt = `, sqlPackage.Ident("NullInt64"), `{ Int64: `, timePackage.Ident("Now().Unix()"), `, Valid: true }`)
			g.P(`}`)
		} else {
			g.P(`if tmp == nil { return }`)
			g.P(`data = tmp`)
			g.P(`data.UpdatedAt = `, sqlPackage.Ident("NullInt64"), `{ Int64: `, timePackage.Ident("Now().Unix()"), `, Valid: true }`)
		}
		g.P(`updateList := []string{`)
		for _, field := range m.Fields {
			fieldName := field.GoName
			_dbField := getFieldOptionSqlxDBField(field)
			if _dbField != nil {
				for _, _updateKey := range _dbField.GetUpdate() {
					if key == _updateKey {
						g.P(structName, `TableColumn`, fieldName, `[tx.DriverName()] + " = " + `, structName, `Field`, fieldName, `NamedMapping,`)
					}
				}
			}
		}
		g.P(structName, `TableColumnUpdatedAt[tx.DriverName()] + " = " + `, structName, `FieldDeletedAtNamedMapping,`)
		g.P(`}`)
		g.P(`whereList := []string{`)
		for _, k := range strings.Split(key, "And") {
			g.P(structName, `TableColumn`, k, `[tx.DriverName()] + " = " + `, structName, `Field`, k, `NamedMapping,`)
		}
		g.P(`}`)
		g.P(`query := "UPDATE " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " SET " + strings.Join(updateList, ", ") + " WHERE " + strings.Join(whereList, " AND ")`)
	})
}
func (gen Generator) generateDAOReplace(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message, key string) {
	gen.generateDAOOperate(g, m, "Replace", key, func(g *protogen.GeneratedFile, batch bool, structName, key string) {
		if batch {
			g.P(`if len(tmp) == 0 { return }`)
			g.P(`list = tmp`)
			g.P(`for index := range list {`)
			g.P(`list[index].ID = `, uuidPackage.Ident("New()"))
			g.P(`list[index].CreatedAt = `, timePackage.Ident("Now().Unix()"))
			g.P(`}`)
		} else {
			g.P(`if tmp == nil { return }`)
			g.P(`data = tmp`)
			g.P(`data.ID = `, uuidPackage.Ident("New()"))
			g.P(`data.CreatedAt = `, timePackage.Ident("Now().Unix()"))
		}
		g.P(`query := ""`)
		g.P(`switch tx.DriverName() {`)
		g.P(`case "postgres":`)
		g.P(`query = generatePostgresReplace`, structName, `By`, key, `SQL(DEFAULT_TABLE_PREFIX)`)
		g.P(`case "mysql":`)
		g.P(`}`)
		g.P(`if query == "" { return }`)
	})
}
func (gen Generator) generateDAODelete(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message, key string) {
	gen.generateDAOOperate(g, m, "Delete", key, func(g *protogen.GeneratedFile, batch bool, structName, key string) {
		g.P(`whereList := []string{`)
		for _, k := range strings.Split(key, "And") {
			g.P(structName, `TableColumn`, k, `[tx.DriverName()] + " = " + `, structName, `Field`, k, `NamedMapping,`)
		}
		g.P(`}`)
		g.P(`query := "DELETE FROM " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " WHERE " + strings.Join(whereList, " AND ")`)
	})
	gen.generateDAOOperate(g, m, "SoftDelete", key, func(g *protogen.GeneratedFile, batch bool, structName, key string) {
		if batch {
			g.P(`if len(tmp) == 0 { return }`)
			g.P(`for index := range tmp {`)
			g.P(`tmp[index].UpdatedAt = `, sqlPackage.Ident("NullInt64"), `{ Int64: `, timePackage.Ident("Now().Unix()"), `, Valid: true }`)
			g.P(`}`)
		} else {
			g.P(`if tmp == nil { return }`)
			g.P(`tmp.DeletedAt = sql.NullInt64{Int64: `, timePackage.Ident("Now().Unix()"), `, Valid: true}`)
		}
		g.P(`whereList := []string{`)
		for _, k := range strings.Split(key, "And") {
			g.P(structName, `TableColumn`, k, `[tx.DriverName()] + " = " + `, structName, `Field`, k, `NamedMapping,`)
		}
		g.P(`}`)
		g.P(`query := "UPDATE " + `, structName, `TableName[tx.DriverName()](DEFAULT_TABLE_PREFIX) + " SET " + `, m.GoIdent.GoName, `TableColumnDeletedAt[tx.DriverName()] + " = " + `, m.GoIdent.GoName, `FieldDeletedAtNamedMapping + " WHERE " + strings.Join(whereList, " AND ")`)
	})
}
func (gen Generator) generateDAOSelect(
	g *protogen.GeneratedFile,
	f *protogen.File,
	selectType string,
	flag bool,
	m *protogen.Message,
	key string,
	generateFuncBody func(g *protogen.GeneratedFile, structName, key string),
) {
	var generateFuncHeaderParams = func(g *protogen.GeneratedFile, f *protogen.File, selectType string, flag bool, field *protogen.Field, key string) {
		var getFieldType = func(field *protogen.Field, _type *kit.FieldType) interface{} {
			var fieldType interface{}
			if _type != nil {
				switch _type.GetGoPkg() {
				case "github.com/google/uuid":
					goType := _type.GetGoType()
					if strings.HasPrefix(goType, "Null") {
						goType = strings.Replace(goType, "Null", "", 1)
					}
					fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident(goType)
				case "database/sql":
					switch _type.GetGoType() {
					case "NullString":
						fieldType = `string`
					case "NullInt64":
						fieldType = `int64`
					case "NullInt32":
						fieldType = `int32`
					case "NullInt16":
						fieldType = `int16`
					case "NullByte":
						fieldType = `byte`
					case "NullFloat64":
						fieldType = `float64`
					case "NullBool":
						fieldType = `bool`
					case "NullTime":
						fieldType = timePackage.Ident("Time")
					default:
						fieldType = protogen.GoImportPath(_type.GetGoPkg()).Ident(_type.GoType)
					}
				}
			} else {
				switch field.Desc.Kind() {
				case protoreflect.BoolKind:
					fieldType = `bool`
				case protoreflect.EnumKind:
					fieldType = getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Enum.GoIdent.GoImportPath.String()).Ident(field.Enum.GoIdent.GoName)
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
					fieldType = `int32`
				case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
					fieldType = `uint32`
				case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
					fieldType = `int64`
				case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
					fieldType = `uint64`
				case protoreflect.FloatKind:
					fieldType = `float32`
				case protoreflect.DoubleKind:
					fieldType = `float64`
				case protoreflect.StringKind:
					fieldType = `string`
				case protoreflect.BytesKind:
					fieldType = `[]byte`
				case protoreflect.MessageKind:
					fieldType = getCurrentImportPath(f.GoImportPath.String(), f.GeneratedFilenamePrefix, field.Message.GoIdent.GoImportPath.String()).Ident(field.Message.GoIdent.GoName)
				case protoreflect.GroupKind:
				}
			}
			return fieldType
		}

		fieldName := camel(field.GoName)
		_type := getFieldOptionKitFieldType(field)
		_dbField := getFieldOptionSqlxDBField(field)
		if _dbField != nil {
			var fieldType interface{}
			switch selectType {
			case "one":
				for _, _key := range _dbField.GetOne() {
					if key == _key {
						fieldType = getFieldType(field, _type)
					}
				}
			case "list":
				for _, _key := range _dbField.GetList() {
					if key == _key {

						fieldType = getFieldType(field, _type)
					}
				}
			}
			if fieldType != nil {
				if flag {
					g.P(fieldName, `List []`, fieldType, `,`)
				} else {
					g.P(fieldName, ` `, fieldType, `,`)
				}
			}
		}
	}

	var generateFuncHeader = func(g *protogen.GeneratedFile, selectType string, flag bool, structName string, key string) {
		funcName := "func Get" + structName
		if flag {
			funcName += "List"
		} else {
			switch selectType {
			case "list":
				funcName += "List"
			}
		}
		funcName += "By" + key
		if flag {
			funcName += "List"
		}
		out := "data "
		switch selectType {
		case "list":
			out = "list []"
		}
		if flag {
			out = "list []"
		}
		out += "*" + structName
		switch selectType {
		case "one":
			out += ", exist bool"
		case "list":

		}
		g.P(funcName, `(`)
		g.P(`ctx `, contextPackage.Ident("Context"), `,`)
		g.P(`tx *`, sqlxPackage.Ident("Tx"), `,`)
		for _, field := range m.Fields {
			generateFuncHeaderParams(g, f, selectType, flag, field, key)
		}
		g.P(`columns ...string,`)
		g.P(`) (`, out, `, err error) {`)
		g.P(`if len(columns) == 0 { return }`)

		if flag {
			g.P(`list = make([]*`, structName, `, 0)`)
		} else {
			switch selectType {
			case "one":
				g.P(`data = new(`, structName, `)`)
			case "list":
				g.P(`list = make([]*`, structName, `, 0)`)
			}
		}
	}
	var generateFuncFooter = func(g *protogen.GeneratedFile, selectType string, flag bool) {
		if !flag {
			switch selectType {
			case "one":
				g.P(`err = tx.GetContext(ctx, data, query, whereDataList...)`)
				g.P(`if err != nil {`)
				g.P(`if err != `, sqlPackage.Ident("ErrNoRows"), ` {`)
				g.P(`return`)
				g.P(`}`)
				g.P(`err = nil`)
				g.P(`return`)
				g.P(`}`)
				g.P(`exist = true`)
			case "list":

			}
		}
		g.P(`return`)
		g.P(`}`)
	}
	generateFuncHeader(g, selectType, flag, m.GoIdent.GoName, key)
	generateFuncBody(g, m.GoIdent.GoName, key)
	generateFuncFooter(g, selectType, flag)
}

func (gen Generator) generateDAOOperate(
	g *protogen.GeneratedFile,
	m *protogen.Message,
	op string,
	key string,
	generateFuncBody func(g *protogen.GeneratedFile, batch bool, structName string, key string),
) {
	var generateFuncHeader = func(g *protogen.GeneratedFile, batch bool, structName string, key string) {
		funcName := "func "
		in := "tmp "
		out := "data "
		if batch {
			funcName += "Batch"
			in += "[]"
			out = "list []"
		}
		funcName += op + structName
		if key != "" {
			funcName += "By" + key
		}
		in += "*" + structName
		out += "*" + structName
		switch op {
		case "Insert", "Update", "Replace":
			out += ", rowsAffected int64, err error"
		case "Delete", "SoftDelete":
			out = "rowsAffected int64, err error"
		}
		g.P(funcName, `(ctx `, contextPackage.Ident("Context"), `, tx *`, sqlxPackage.Ident("Tx"), `, `, in, `) (`, out, `) {`)
	}
	var generateFuncFooter = func(g *protogen.GeneratedFile, batch bool) {
		g.P(`var result sql.Result`)
		switch op {
		case "Insert", "Update", "Replace":
			if batch {
				g.P(`result, err = tx.NamedExecContext(ctx, query, list)`)
			} else {
				g.P(`result, err = tx.NamedExecContext(ctx, query, data)`)
			}
		case "Delete", "SoftDelete":
			g.P(`result, err = tx.NamedExecContext(ctx, query, tmp)`)
		}
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`rowsAffected, err = result.RowsAffected()`)
		g.P(`if err != nil {`)
		g.P(`return`)
		g.P(`}`)
		g.P(`return`)
		g.P(`}`)
	}
	structName := m.GoIdent.GoName
	generateFuncHeader(g, false, structName, key)
	generateFuncBody(g, false, structName, key)
	generateFuncFooter(g, false)
	g.P()
	generateFuncHeader(g, true, structName, key)
	generateFuncBody(g, true, structName, key)
	generateFuncFooter(g, true)
}
