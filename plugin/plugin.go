package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/ry023/connecterr"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("example", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return Plugin{}, nil
}

type Plugin struct {
}

func (p Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func (p Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{connecterr.Analyzer}, nil
}
