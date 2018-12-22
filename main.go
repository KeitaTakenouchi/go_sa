package main

import (
	"go/ast"
	"go/parser"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
)

func main() {
	analyses := []*analysis.Analyzer{
		unusedresult.Analyzer,
		nilness.Analyzer,
		printf.Analyzer,
	}

	_ = analyses

	exp, _ := parser.ParseExpr("1 + 3")
	ast.Print(nil, exp)

}
