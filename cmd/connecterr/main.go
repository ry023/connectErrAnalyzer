package main

import (
	"github.com/ry023/connecterr"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(connecterr.Analyzer) }
