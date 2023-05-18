package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"go-func-param-linter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
