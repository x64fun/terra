package main

import (
	"flag"
	"fmt"

	"github.com/x64fun/terra/cmd/protoc-gen-go-kit/internal/gengo"
	"github.com/x64fun/terra/cmd/protoc-gen-go-kit/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

var requireUnimplemented *bool

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-kit %v\n", version.String())
		return
	}

	var flags flag.FlagSet
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		g := gengo.NewGenerator(gen, *requireUnimplemented)
		for _, f := range gen.Files {
			if f.Generate {
				g.GenerateFile(f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
