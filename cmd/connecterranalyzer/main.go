package main

import (
	"connecterranalyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(connecterranalyzer.Analyzer) }
