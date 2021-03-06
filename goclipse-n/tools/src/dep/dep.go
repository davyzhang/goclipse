///////////////////////////////////////////////////////////////////
// IMPORTANT!
// If you make changes to this file, please increment the value of:
// DEP_TOOL_VERSION in Environment.java
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"flag"
	"go/parser"
	"go/token"
	"go/ast"
	"os"
)

func main() {
	flag.Parse()
	narg := flag.NArg()
	if narg != 1 {
		fmt.Printf("- must have a Go source as parameter\n")
		return
	}
	filename := flag.Arg(0)
	astf, error := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ImportsOnly)
	if error == nil  {
		fmt.Printf("p:%s\n", astf.Name);
		for _, d:= range astf.Decls {
			if gd, ok:=d.(*ast.GenDecl); ok {
				if gd.Tok == token.IMPORT {
					for _,s:= range gd.Specs {
						if v, ok := s.(*ast.ImportSpec); ok {
							fmt.Printf("%s\n",string(v.Path.Value))
						}
					}
				}
			}
		}
	} else {
		fmt.Printf("- %v\n", error)
		os.Exit(-1)
	}
}
