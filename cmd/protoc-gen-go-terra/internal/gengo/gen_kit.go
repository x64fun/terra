package gengo

import (
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

// GenerateFile generates the contents of a _kit.pb.go file.
func (gen Generator) GenerateFile(file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join("terra", file.GeneratedFilenamePrefix+"_kit.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(g, gen.Plugin, file)
	gen.generateFileContent(file, g)
	return g
}

// generateFileContent generates the go-kit definitions, excluding the package statement.
func (gen Generator) generateFileContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the go-kit package it is being compiled against.")
	g.P("// Requires go-kit v0.12.0.")
	g.P()
	for _, service := range file.Services {
		gen.genKitClient(file, g, service)
		gen.genKitServer(file, g, service)
		gen.genKitEndpointSet(file, g, service)
		gen.genKitHTTPTransport(file, g, service)
		gen.genKitGRPCTransport(file, g, service)
		gen.genKitHTTPHandlerSet(file, g, service)
		gen.genKitGRPCHandlerSet(file, g, service)
	}
}

func (gen Generator) genKitClient(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	clientName := service.GoName + "Client"
	kitEndpointSet := service.GoName + "EndpointSet"

	g.P("// ", clientName, " is the go-kit client API for ", service.GoName, " service.")

	// Client interface.
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.Annotate(clientName, service.Location)
	g.P("type ", clientName, " interface {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.Annotate(clientName+"."+method.GoName, method.Location)
			if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
				g.P(deprecationComment)
			}
			g.P(method.Comments.Leading,
				clientSignature(g, method))
		}
	}
	g.P("}")
	g.P()

	// Client structure.
	g.P("type ", unexport(clientName), " struct {")
	g.P("s ", kitEndpointSet)
	g.P("}")
	g.P()

	// NewClient factory.
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("func New", clientName, " (s ", kitEndpointSet, ") ", clientName, " {")
	g.P("return &", unexport(clientName), "{s}")
	g.P("}")
	g.P()

	// Client method implementations.
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			gen.genKitClientMethod(file, g, method)
		}
	}
}
func (gen Generator) genKitServer(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	mustOrShould := "must"
	if !gen.requireUnimplemented {
		mustOrShould = "should"
	}

	// Server interface.
	serverName := service.GoName + "Server"
	g.P("// ", serverName, " is the server API for ", service.GoName, " service.")
	g.P("// All implementations ", mustOrShould, " embed Unimplemented", serverName)
	g.P("// for forward compatibility")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.Annotate(serverName, service.Location)
	g.P("type ", serverName, " interface {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.Annotate(serverName+"."+method.GoName, method.Location)
			if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
				g.P(deprecationComment)
			}
			g.P(method.Comments.Leading,
				serverSignature(g, method))
		}
	}
	if gen.requireUnimplemented {
		g.P("mustEmbedUnimplemented", serverName, "()")
	}
	g.P("}")
	g.P()

	// Server Unimplemented struct for forward compatibility.
	g.P("// Unimplemented", serverName, " ", mustOrShould, " be embedded to have forward compatible implementations.")
	g.P("type Unimplemented", serverName, " struct {")
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("func (Unimplemented", serverName, ") ", serverSignature(g, method), "{")
			g.P("return nil, ", errorsPackage.Ident("New"), `("method `, method.GoName, ` not implemented")`)
			g.P("}")
		}
	}
	if gen.requireUnimplemented {
		g.P("func (Unimplemented", serverName, ") mustEmbedUnimplemented", serverName, "() {}")
	}
	g.P()

	// Unsafe Server interface to opt-out of forward compatibility.
	g.P("// Unsafe", serverName, " may be embedded to opt out of forward compatibility for this service.")
	g.P("// Use of this interface is not recommended, as added methods to ", serverName, " will")
	g.P("// result in compilation errors.")
	g.P("type Unsafe", serverName, " interface {")
	g.P("mustEmbedUnimplemented", serverName, "()")
	g.P("}")
	g.P()
}
func (gen Generator) genKitEndpointSet(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	// define Endpoint.
	kitEndpointSet := service.GoName + "EndpointSet"
	serverName := service.GoName + "Server"
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("// ", kitEndpointSet, " is realize kit/endpoint.Endpoint")
	g.P("type ", kitEndpointSet, " struct {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.Annotate(serverName+"."+method.GoName, method.Location)
			if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
				g.P(deprecationComment)
			}
			g.P(method.Comments.Leading,
				" "+method.GoName+"Endpoint ", kitEndpointPackage.Ident("Endpoint"))
		}
	}
	g.P("}")
	g.P("func New", kitEndpointSet, " (srv ", serverName, ", mws ...", kitEndpointPackage.Ident("Middleware"), ") ", kitEndpointSet, " {")
	g.P("set := ", kitEndpointSet, "{}")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("set.", method.GoName, "Endpoint = func(ctx ", contextPackage.Ident("Context"), ", request interface{}) (response interface{}, err error) {")
			g.P("return srv.", method.GoName, "(ctx, request.(*", method.Input.GoIdent, "))")
			g.P("}")
		}
	}
	g.P("for _, mw := range mws {")
	methodCount := 0
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			methodCount++
			g.P("set.", method.GoName, "Endpoint = ", "mw(set.", method.GoName, "Endpoint)")
		}
	}
	if methodCount == 0 {
		g.P("_ = mw")
	}
	g.P("}")
	g.P("return set")
	g.P("}")
	g.P()
}
func (gen Generator) genKitHTTPTransport(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	// define http transport.
	kitHTTPTransport := service.GoName + "HTTPTransport"
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("// ", kitHTTPTransport, " defined from kit/transport/http.")
	g.P("type ", kitHTTPTransport, " interface {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("Decode", method.GoName, "Request(", contextPackage.Ident("Context"), ", *", netHTTPPackage.Ident("Request"), ") (interface{}, error)")
			g.P("Encode", method.GoName, "Response(", contextPackage.Ident("Context"), ", ", netHTTPPackage.Ident("ResponseWriter"), ", interface{}) error")
			g.P("Encode", method.GoName, "Request(", contextPackage.Ident("Context"), ", *", netHTTPPackage.Ident("Request"), ", interface{}) error")
			g.P("Decode", method.GoName, "Response(", contextPackage.Ident("Context"), ", *", netHTTPPackage.Ident("Response"), ") (interface{}, error)")
		}
	}
	g.P("}")
	g.P()
}
func (gen Generator) genKitGRPCTransport(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	// define grpc transport.
	kitGRPCTransport := service.GoName + "GRPCTransport"
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("// ", kitGRPCTransport, " defined from kit/transport/grpc.")
	g.P("type ", kitGRPCTransport, " interface {")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("Decode", method.GoName, "Request(", contextPackage.Ident("Context"), ", interface{}) (interface{}, error)")
			g.P("Encode", method.GoName, "Response(", contextPackage.Ident("Context"), ", interface{}) (interface{}, error)")
			g.P("Encode", method.GoName, "Request(", contextPackage.Ident("Context"), ", interface{}) (interface{}, error)")
			g.P("Decode", method.GoName, "Response(", contextPackage.Ident("Context"), ", interface{}) (interface{}, error)")
		}
	}
	g.P("}")
	g.P()
}
func (gen Generator) genKitHTTPHandlerSet(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	kitEndpointSet := service.GoName + "EndpointSet"
	kitHTTPTransport := service.GoName + "HTTPTransport"
	// define http handler set.
	serviceHTTPHandlerSet := service.GoName + "HTTPHandlerSet"
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("// ", serviceHTTPHandlerSet, " is realize kit/transport/http.Server set.")
	g.P("type ", serviceHTTPHandlerSet, " struct {")
	for _, method := range service.Methods {
		g.P(method.GoName, "Handler *", kitTransportHTTPPackage.Ident("Server"))
	}
	g.P("}")
	g.P("func New", serviceHTTPHandlerSet, "(")
	g.P("set ", kitEndpointSet, ",")
	g.P("transport ", kitHTTPTransport, ",")
	g.P("opts ...", kitTransportHTTPPackage.Ident("ServerOption"), ",")
	g.P(") ", serviceHTTPHandlerSet, " {")
	g.P("s := ", serviceHTTPHandlerSet, "{}")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("s.", method.GoName, "Handler = ", kitTransportHTTPPackage.Ident("NewServer"), "(")
			g.P("set.", method.GoName, "Endpoint,")
			g.P("transport.Decode", method.GoName, "Request,")
			g.P("transport.Encode", method.GoName, "Response,")
			g.P("opts...")
			g.P(")")
		}
	}
	g.P("return s")
	g.P("}")
	g.P()
}
func (gen Generator) genKitGRPCHandlerSet(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	kitEndpointSet := service.GoName + "EndpointSet"
	kitGRPCTransport := service.GoName + "GRPCTransport"
	// define grpc handler set.
	serviceGRPCHandlerSet := service.GoName + "GRPCHandlerSet"
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("// ", serviceGRPCHandlerSet, " is realize kit/transport/grpc.Server set.")
	g.P("type ", serviceGRPCHandlerSet, " struct {")
	for _, method := range service.Methods {
		g.P(method.GoName, "Handler *", kitTransportGRPCPackage.Ident("Server"))
	}
	g.P("}")
	g.P("func New", serviceGRPCHandlerSet, "(")
	g.P("set ", kitEndpointSet, ",")
	g.P("transport ", kitGRPCTransport, ",")
	g.P("opts ...", kitTransportGRPCPackage.Ident("ServerOption"), ",")
	g.P(") ", serviceGRPCHandlerSet, " {")
	g.P("s := ", serviceGRPCHandlerSet, "{}")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			g.P("s.", method.GoName, "Handler = ", kitTransportGRPCPackage.Ident("NewServer"), "(")
			g.P("set.", method.GoName, "Endpoint,")
			g.P("transport.Decode", method.GoName, "Request,")
			g.P("transport.Encode", method.GoName, "Response,")
			g.P("opts...")
			g.P(")")
		}
	}
	g.P("return s")
	g.P("}")
	g.P()
}

func (gen Generator) genKitClientMethod(file *protogen.File, g *protogen.GeneratedFile, method *protogen.Method) {
	service := method.Parent
	if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("func (c *", unexport(service.GoName), "Client) ", clientSignature(g, method), "{")
	g.P("var response interface{}")
	g.P("response, err = c.s." + method.GoName + "Endpoint(ctx, in)")
	g.P("if err != nil { return }")
	g.P("out = response.(*", method.Output.GoIdent, ")")
	g.P("return")
	g.P("}")
	g.P()
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	return method.GoName + "(ctx " + g.QualifiedGoIdent(contextPackage.Ident("Context")) + ", in *" + g.QualifiedGoIdent(method.Input.GoIdent) + ") (out *" + g.QualifiedGoIdent(method.Output.GoIdent) + ", err error)"
}
func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	return method.GoName + "(" + g.QualifiedGoIdent(contextPackage.Ident("Context")) + ", *" + g.QualifiedGoIdent(method.Input.GoIdent) + ") (*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
}
func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }
