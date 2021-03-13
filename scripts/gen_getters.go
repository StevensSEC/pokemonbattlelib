// Auto generate getters on one type to grab data from another type.
// example usage: go run ./scripts/gen_getters.go -for Move -data AllMoves -output move_getters_GEN.go
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	forTypeName    = flag.String("for", "", "The type to create getter functions for.")
	dataTypeName   = flag.String("data", "", "The variable that contains all the constant data for the type")
	outputFileName = flag.String("output", "", "output file name; default <for>_getters_GEN.go")
)

func createCodeOutput(s string) {
	file, err := os.Create(*outputFileName)
	if err != nil {
		log.Panicln(err)
	}
	_, err = file.WriteString("// Code generated - DO NOT EDIT.\n" +
		"// Regenerate with `go generate`.\n\n" +
		"package pokemonbattlelib\n\n" + s)
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	flag.Parse()
	if *outputFileName == "" {
		*outputFileName = fmt.Sprintf("%s_getters_GEN.go", strings.ToLower(*forTypeName))
	}
	log.Printf("Generating getters on type %s for data %s", *forTypeName, *dataTypeName)

	fset := token.NewFileSet()

	log.Println("Parsing current directory...")
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("packages: %d - %v", len(pkgs), pkgs)

	pkg := pkgs["pokemonbattlelib"]
	libFile := ast.MergePackageFiles(pkg, ast.FilterUnassociatedComments)
	// log.Printf("obj: %+v", libFile)

	// find the forType and the data
	var forType *ast.TypeSpec
	var dataType *ast.TypeSpec
	for _, d := range libFile.Decls {
		switch n := d.(type) {
		case *ast.GenDecl:
			for _, s := range n.Specs {
				switch n := s.(type) {
				case *ast.TypeSpec:
					if forType == nil && n.Name.Name == *forTypeName {
						log.Printf("found for type declaration: %v", n)
						forType = n
					} else if dataType == nil && n.Name.Name == *dataTypeName {
						log.Printf("found data type declaration: %v", n)
						dataType = n
					}
				}
			}
		}
		if forType != nil && dataType != nil {
			break
		}
	}

	// extract the fields
	log.Printf("data var type: %#v", dataType.Type)

	output := ""
	for _, field := range dataType.Type.(*ast.StructType).Fields.List {
		log.Printf("field: %v %v", field.Names[0], field.Type)
		output += fmt.Sprintf("func (n %[1]s) %[2]s() %[3]s { return n.Data().%[2]s }\n", *forTypeName, field.Names[0], field.Type)
	}

	createCodeOutput(output)

	// run gofmt on generated code
	log.Println("Formatting generated code...")
	cmd := exec.Command("gofmt", "-w", *outputFileName)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to format generated code: %v", err)
	}
}
