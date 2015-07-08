package generator

import (
	"go/ast"
	"go/token"

	"github.com/momchil-atanasov/gostub/util"
)

func NewStubMethodBuilder() *StubMethodBuilder {
	return &StubMethodBuilder{
		selfName: "stub",
		params:   make([]*ast.Field, 0),
		results:  make([]*ast.Field, 0),
	}
}

// StubMethodBuilder is reponsible for creating a method that implements
// the original method from the interface and does all the tracking
// logic used by this framework.
//
// Example:
//     func (stub *StubStruct) Sum(a int, b int) int {
//         // ...
//     }
type StubMethodBuilder struct {
	selfName       string
	selfType       string
	methodName     string
	argsFieldName  string
	mutexFieldName string
	stubFieldName  string
	params         []*ast.Field
	results        []*ast.Field
}

func (b *StubMethodBuilder) SetReceiverName(name string) {
	b.selfName = name
}

func (b *StubMethodBuilder) SetReceiverType(name string) {
	b.selfType = name
}

func (b *StubMethodBuilder) SetMethodName(name string) {
	b.methodName = name
}

func (b *StubMethodBuilder) SetArgsFieldName(name string) {
	b.argsFieldName = name
}

func (b *StubMethodBuilder) SetMutexFieldName(name string) {
	b.mutexFieldName = name
}

func (b *StubMethodBuilder) SetStubFieldName(name string) {
	b.stubFieldName = name
}

// SetParams specifies the parameters that the original method
// uses. These parameters need to have been normalized and resolved
// in advance.
func (b *StubMethodBuilder) SetParams(params []*ast.Field) {
	b.params = params
}

// SetResults specifies the results that the original method
// returns. These results need to have been normalized and resolved
// in advance.
func (b *StubMethodBuilder) SetResults(results []*ast.Field) {
	b.results = results
}

func (b *StubMethodBuilder) Build() *ast.FuncDecl {
	mutexLockBuilder := NewMutexLockBuilder()
	mutexLockBuilder.SetReceiverName(b.selfName)
	mutexLockBuilder.SetMutexField(b.mutexFieldName)
	mutexLockBuilder.SetAction("Lock")

	mutexUnlockBuilder := NewMutexUnlockBuilder()
	mutexUnlockBuilder.SetReceiverName(b.selfName)
	mutexUnlockBuilder.SetMutexField(b.mutexFieldName)
	mutexUnlockBuilder.SetAction("Unlock")

	method := NewMethodModel()
	method.SetName(b.methodName)
	method.SetReceiver(b.selfName, b.selfType)
	method.SetType(&ast.FuncType{
		Params: &ast.FieldList{
			List: b.params,
		},
		Results: &ast.FieldList{
			List: util.GetFieldsAsAnonymous(b.results),
		},
	})
	method.AddStatement(mutexLockBuilder.Build())
	method.AddStatement(mutexUnlockBuilder.Build())

	paramSelectors := []ast.Expr{}
	for _, param := range b.params {
		paramSelectors = append(paramSelectors, ast.NewIdent(param.Names[0].String()))
	}

	method.AddStatement(&ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.SelectorExpr{
				X:   ast.NewIdent(b.selfName),
				Sel: ast.NewIdent(b.argsFieldName),
			},
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: ast.NewIdent("append"),
				Args: []ast.Expr{
					&ast.SelectorExpr{
						X:   ast.NewIdent(b.selfName),
						Sel: ast.NewIdent(b.argsFieldName),
					},
					&ast.CompositeLit{
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: b.params,
							},
						},
						Elts: paramSelectors,
					},
				},
			},
		},
	})

	method.AddStatement(&ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X: &ast.SelectorExpr{
				X:   ast.NewIdent(b.selfName),
				Sel: ast.NewIdent(b.stubFieldName),
			},
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent(b.selfName),
							Sel: ast.NewIdent(b.stubFieldName),
						},
						Args: paramSelectors,
					},
				},
			},
		},
	})

	return method.BuildASTFuncDecl()
}
