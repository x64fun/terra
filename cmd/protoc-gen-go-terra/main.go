package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/x64fun/terra/cmd/protoc-gen-go-terra/internal/gengo"
	"github.com/x64fun/terra/cmd/protoc-gen-go-terra/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

// var requireUnimplemented *bool

var (
	showVersion             = flag.Bool("version", false, "print the version and exit")
	requireUnimplemented    = flag.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")
	generateConvertProtobuf = flag.Bool("generate_convert_protobuf", false, "generate terra struct convert to protobuf struct")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	if *showVersion {
		fmt.Printf("protoc-gen-go-terra %v\n", version.String())
		return
	}

	// var flags flag.FlagSet
	// requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		g := gengo.NewGenerator(gen, *requireUnimplemented, *generateConvertProtobuf)

		for _, f := range gen.Files {
			if f.Generate {
				glog.V(1).Infof("NewGeneratedFile %q in %s", f.GoPackageName, *f.Proto.Name)
				g.GenerateStruct(f)
				g.GenerateFile(f)
				// DAO
				g.GenerateDAO(f)
				g.GenerateDAOMySQL(f)
				g.GenerateDAOPostgres(f)

				// Mux
				g.GenerateMux(f)

				glog.V(1).Info("Processed code generator request")
			}
		}
		daoDefaultMap := make(map[string]bool)
		for _, f := range gen.Files {
			if f.Generate {
				g.GenerateDAODefault(f, daoDefaultMap)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
