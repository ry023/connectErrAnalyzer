package main

import (
	"github.com/ry023/connecterr"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{connecterr.Analyzer}, nil
}
