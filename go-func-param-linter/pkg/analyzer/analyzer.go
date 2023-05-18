package analyzer

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gofuncparamlint",
	Doc:  "Checks if func params or ordered alphabetically",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDec, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		var funcParams Alphabetic
		var ints Alphabetic

		params := funcDec.Type.Params.List
		for _, param := range params {
			strParam := param.Names[0].String()

			if strings.HasSuffix(strParam, "Int") {
				ints = append(ints, strParam)
			} else if strings.HasSuffix(strParam, "Client") || strings.Contains(strParam, "log") {
				continue
			} else {
				funcParams = append(funcParams, strParam)
			}
		}

		var errors []error

		if paramErr := isAlphabeticalOrder(funcParams); paramErr != nil {
			errors = append(errors, paramErr)
		}

		if intsErr := isAlphabeticalOrder(ints); intsErr != nil {
			errors = append(errors, intsErr)
		}

		if len(errors) > 0 {
			for _, err := range errors {
				pass.Reportf(node.Pos(), "%s", err.Error())
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}

type Alphabetic []string

func (list Alphabetic) Len() int      { return len(list) }
func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Alphabetic) Less(i, j int) bool {
	si := list[i]
	sj := list[j]

	siLower := strings.ToLower(si)
	sjLower := strings.ToLower(sj)

	if siLower == sjLower {
		return si < sj
	}

	return siLower < sjLower
}

func isAlphabeticalOrder(params Alphabetic) error {
	if !sort.SliceIsSorted(params, func(i, j int) bool {
		return params[i] < params[j]
	}) {
		return fmt.Errorf("parameters: %s are not sorted alphabetically", params)
	}

	return nil
}
