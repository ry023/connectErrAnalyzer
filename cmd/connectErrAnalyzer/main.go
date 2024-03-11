package main

import (
	"connectErrAnalyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(connectErrAnalyzer.Analyzer) }
