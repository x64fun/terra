package gengo

import (
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (gen Generator) GenerateMux(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+"_mux.go")
	if len(file.Services) == 0 {
		return nil
	}
	importPath := getCurrentImportPath(file.GoImportPath.String(), file.GeneratedFilenamePrefix, file.GoImportPath.String())
	file.GoImportPath = protogen.GoImportPath(importPath)
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(file.GoImportPath))
	generateFileHeader(g, gen.Plugin, file)
	g.P()
	for _, service := range file.Services {
		gen.generateMuxService(g, file, service)
	}
	return g
}
func (gen Generator) generateMuxService(g *protogen.GeneratedFile, file *protogen.File, service *protogen.Service) {
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			_router := getMethodOptionMuxRouter(method)
			_swagger := getMethodOptionSwagger(method)
			if _router != nil {
				if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
					g.P(deprecationComment)
				}
				g.P("// ", method.GoName, " godoc")
				if _swagger != nil {
					g.P("// @Summary ", _swagger.GetSummary())
					g.P("// @Description ", _swagger.GetDescription())
					g.P("// @Tags ", _swagger.GetTags())
					g.P("// @Accept ", _swagger.GetAccept())
					g.P("// @Produce ", _swagger.GetProduce())
					for _, param := range _swagger.GetParam() {
						g.P("// @Param ", param)
					}
					g.P("// @Success ", _swagger.GetSuccess())
					g.P("// @Router ", _swagger.GetRouter())
				}
				g.P("func RegisterHTTP", service.GoName, method.GoName, "Handler(r *", muxPackage.Ident("Router"), ", set ", service.GoName, "HTTPHandlerSet) {")
				g.P(`var _r *mux.Route`)
				flag := false
				if _router.GetHost() != "" {
					if flag {
						g.P(`_r = _r.Host("`, _router.GetHost(), `")`)
					} else {
						g.P(`_r = r.Host("`, _router.GetHost(), `")`)
						flag = true
					}
				}
				if _router.GetPath() != "" {
					if flag {
						g.P(`_r = _r.Path("`, _router.GetPath(), `")`)
					} else {
						g.P(`_r = r.Path("`, _router.GetPath(), `")`)
						flag = true
					}
				}
				if _router.GetPathPrefix() != "" {
					if flag {
						g.P(`_r = _r.PathPrefix("`, _router.GetPathPrefix(), `")`)
					} else {
						g.P(`_r = r.PathPrefix("`, _router.GetPathPrefix(), `")`)
						flag = true
					}
				}
				if len(_router.GetMethods()) > 0 {
					methods := []string{}
					for _, _method := range _router.GetMethods() {
						methods = append(methods, `"`+_method.String()+`"`)
					}
					if flag {
						g.P("_r = _r.Methods(", strings.Join(methods, ", "), ")")
					} else {
						g.P("_r = r.Methods(", strings.Join(methods, ", "), ")")
						flag = true
					}
				}
				if len(_router.GetHeaders()) > 0 {
					if flag {
						g.P(`_r = _r.Headers("`, strings.Join(_router.GetHeaders(), `", "`), `")`)
					} else {
						g.P(`_r = r.Headers("`, strings.Join(_router.GetHeaders(), `", "`), `")`)
						flag = true
					}
				}
				if len(_router.GetQueries()) > 0 {
					if flag {
						g.P(`_r = _r.Queries("`, strings.Join(_router.GetQueries(), `", "`), `")`)
					} else {
						g.P(`_r = r.Queries("`, strings.Join(_router.GetQueries(), `", "`), `")`)
						flag = true
					}
				}
				if len(_router.GetSchemes()) > 0 {
					schemas := []string{}
					for _, _schema := range _router.GetSchemes() {
						schemas = append(schemas, `"`+_schema.String()+`"`)
					}
					if flag {
						g.P(`_r = _r.Schemes(`, strings.Join(schemas, ", "), `)`)
					} else {
						g.P(`_r = r.Schemes(`, strings.Join(schemas, ", "), `)`)
						flag = true
					}
				}
				g.P(`_r.Handler(set.`, method.GoName, `Handler)`)
				g.P("}")
			}
		}
	}
	g.P("// RegisterHTTP", service.GoName, "HandlerSet - regist all handler")
	g.P("func RegisterHTTP", service.GoName, "HandlerSet(r *", muxPackage.Ident("Router"), ", set ", service.GoName, "HTTPHandlerSet) {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			_router := getMethodOptionMuxRouter(method)
			if _router != nil {
				g.P("RegisterHTTP", service.GoName, method.GoName, "Handler(r, set)")
			}
		}
	}
	g.P("}")
}
