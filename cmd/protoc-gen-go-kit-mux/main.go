package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/x64fun/terra/cmd/protoc-gen-go-kit-mux/internal/gengo"
	"github.com/x64fun/terra/cmd/protoc-gen-go-kit-mux/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-kit-mux %v\n", version.String())
		return
	}

	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gengo.GeneratePower(gen)
		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
