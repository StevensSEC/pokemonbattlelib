// Auto generate MarshallJSON on all transactions.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

const outputFileName = "transactions_GEN.go"

func createCodeOutput(s string) {
	file, err := os.Create(outputFileName)
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
	if _, err := os.Stat(outputFileName); err == nil {
		log.Println("Found existing generated code, deleting to regenerate cleanly...")
		err := os.Remove(outputFileName)
		if err != nil {
			log.Panicf("Failed to delete old generated code: %s", err)
		}
	}
	log.Printf("Generating MarshallJSON for transactions")

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
	transactionTypes := []*ast.TypeSpec{}
	for _, d := range libFile.Decls {
		switch n := d.(type) {
		case *ast.GenDecl:
			for _, s := range n.Specs {
				switch n := s.(type) {
				case *ast.TypeSpec:
					if strings.Contains(n.Name.Name, "Transaction") && unicode.IsUpper(rune(n.Name.Name[0])) {
						switch n.Type.(type) {
						case *ast.StructType:
							transactionTypes = append(transactionTypes, n)
						}
					}
				}
			}
		}
	}

	output := ""
	for id, t := range transactionTypes {
		output += fmt.Sprintf("func (t %[1]s) MarshalJSON() ([]byte, error) {\ntype alias %[1]s\n", t.Name.Name) +
			"return json.Marshal(&struct {\n" +
			"Type int `json:\"type\"`\n" +
			"Name string `json:\"name\"`\n" +
			"Args *alias `json:\"args\"`\n" +
			"}{\n" +
			fmt.Sprintf("Type: %d,\n", id) +
			fmt.Sprintf("Name: \"%s\",\n", t.Name.Name) +
			"Args: (*alias)(&t),\n" +
			"})\n" +
			"}\n\n"
	}

	createCodeOutput(output)

	// run gofmt on generated code
	log.Println("Formatting generated code...")
	cmd := exec.Command("gofmt", "-w", outputFileName)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to format generated code: %v", err)
	}

	log.Println("Auto import packages for generated code...")
	cmd = exec.Command("goimports", "-w", outputFileName)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to auto import packages generated code: %v", err)
	}
}
