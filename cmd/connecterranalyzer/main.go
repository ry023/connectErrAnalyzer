package main

import (
	"github.com/ry023/connecterranalyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(connecterranalyzer.Analyzer) }
