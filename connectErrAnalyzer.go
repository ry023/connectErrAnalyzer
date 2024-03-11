package connectErrAnalyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "connectErrAnalyzer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "connectErrAnalyzer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Find connect method
	var methods []*ast.FuncDecl
	filter := []ast.Node{(*ast.FuncDecl)(nil)}
	inspect.Preorder(filter, func(n ast.Node) {
		if isConnectMethod(pass, n.(*ast.FuncDecl)) {
			methods = append(methods, n.(*ast.FuncDecl))
		}
	})

	for _, m := range methods {
		ast.Inspect(m, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.ReturnStmt); ok {
				errResult := stmt.Results[1]
				if !isConnectErrorWrap(pass, errResult) {
					pass.Reportf(errResult.Pos(), "error must be wrap by connect.NewError method")
				}
			}
			return true
		})
	}

	return nil, nil
}

func isConnectMethod(pass *analysis.Pass, n *ast.FuncDecl) bool {
	resultList := n.Type.Results.List

	// num of returning var must be 2
	if len(resultList) != 2 {
		return false
	}

	// first var must be connect.Result[T]
	resultField1 := resultList[0]
	if !isConnectResultField(pass, resultField1) {
		return false
	}

	// next var must be error
	resultField2 := resultList[1]
	typ := pass.TypesInfo.TypeOf(resultField2.Type)
	return typ != nil && typ.String() == "error"
}

func isConnectResultField(pass *analysis.Pass, f *ast.Field) bool {
	if starExpr, ok := f.Type.(*ast.StarExpr); ok { // `*Something` format
		if indexExpr, ok := starExpr.X.(*ast.IndexExpr); ok { // `*X[Y]` format
			if selectorExpr, ok := indexExpr.X.(*ast.SelectorExpr); ok { // *pkg.Obj[Y] format
				id := selectorExpr.Sel
				obj := pass.TypesInfo.ObjectOf(id)

				return obj != nil && obj.Pkg().Path() == "connectrpc.com/connect" && obj.Name() == "Response"
			}
		}
	}
	return false
}

func isConnectErrorWrap(pass *analysis.Pass, e ast.Expr) bool {
	// is nil ?
	if id, ok := e.(*ast.Ident); ok {
		obj := pass.TypesInfo.ObjectOf(id)
		if _, ok := obj.(*types.Nil); ok {
			return true
		}
	}

	// is connect.Error ?
	typ := pass.TypesInfo.TypeOf(e)
	return typ.String() == "*connectrpc.com/connect.Error"
}
