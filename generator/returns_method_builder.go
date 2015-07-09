package generator

import (
	"go/ast"
	"go/token"
)

func NewReturnsMethodBuilder() *ReturnsMethodBuilder {
	return &ReturnsMethodBuilder{
		selfName: "stub",
		results:  make([]*ast.Field, 0),
	}
}

// ReturnsMethodBuilder is reponsible for creating a method on the stub
// structure that allows you to specify the results to be returned by
// default when the stub method is called.
//
// Example:
//     func (stub *StubStruct) AddressReturns(name string, number int) {
//         // ...
//     }
type ReturnsMethodBuilder struct {
	selfName         string
	selfType         string
	methodName       string
	returnsFieldName string
	mutexFieldName   string
	results          []*ast.Field
}

func (b *ReturnsMethodBuilder) SetReceiverName(name string) {
	b.selfName = name
}

func (b *ReturnsMethodBuilder) SetReceiverType(name string) {
	b.selfType = name
}

func (b *ReturnsMethodBuilder) SetMethodName(name string) {
	b.methodName = name
}

func (b *ReturnsMethodBuilder) SetReturnsFieldName(name string) {
	b.returnsFieldName = name
}

func (b *ReturnsMethodBuilder) SetMutexFieldName(name string) {
	b.mutexFieldName = name
}

// SetResults specifies the results that the original method
// uses. These results need to have been normalized and resolved
// in advance.
func (b *ReturnsMethodBuilder) SetResults(results []*ast.Field) {
	b.results = results
}

func (b *ReturnsMethodBuilder) Build() *ast.FuncDecl {
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
			List: b.results,
		},
	})
	method.AddStatement(mutexLockBuilder.Build())
	method.AddStatement(mutexUnlockBuilder.Build())

	resultSelectors := []ast.Expr{}
	for _, result := range b.results {
		resultSelectors = append(resultSelectors, ast.NewIdent(result.Names[0].String()))
	}
	method.AddStatement(&ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.SelectorExpr{
				X:   ast.NewIdent(b.selfName),
				Sel: ast.NewIdent(b.returnsFieldName),
			},
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CompositeLit{
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: b.results,
					},
				},
				Elts: resultSelectors,
			},
		},
	})
	return method.BuildASTFuncDecl()
}
