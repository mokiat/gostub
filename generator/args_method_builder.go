package generator

import (
	"go/ast"

	"github.com/momchil-atanasov/gostub/util"
)

func NewArgsMethodBuilder(methodBuilder *MethodBuilder) *ArgsMethodBuilder {
	return &ArgsMethodBuilder{
		methodBuilder: methodBuilder,
		params:        make([]*ast.Field, 0),
	}
}

// ArgsMethodBuilder is reponsible for creating a method on the stub
// structure that allows you to check what arguments were used during
// a specific call on the stub method.
//
// Example:
//     func (stub *StubStruct) SumArgsForCall(index int) (int, int) {
//         // ...
//     }
type ArgsMethodBuilder struct {
	methodBuilder      *MethodBuilder
	mutexFieldSelector *ast.SelectorExpr
	argsFieldSelector  *ast.SelectorExpr
	params             []*ast.Field
}

func (b *ArgsMethodBuilder) SetMutexFieldSelector(selector *ast.SelectorExpr) {
	b.mutexFieldSelector = selector
}

func (b *ArgsMethodBuilder) SetArgsFieldSelector(selector *ast.SelectorExpr) {
	b.argsFieldSelector = selector
}

// SetParams specifies the parameters that the original method
// uses. These parameters need to have been normalized and resolved
// in advance.
func (b *ArgsMethodBuilder) SetParams(params []*ast.Field) {
	b.params = params
}

func (b *ArgsMethodBuilder) Build() *ast.FuncDecl {
	mutexLockBuilder := NewMutexLockBuilder()
	mutexLockBuilder.SetMutexFieldSelector(b.mutexFieldSelector)
	mutexLockBuilder.SetAction("RLock")

	mutexUnlockBuilder := NewMutexUnlockBuilder()
	mutexUnlockBuilder.SetMutexFieldSelector(b.mutexFieldSelector)
	mutexUnlockBuilder.SetAction("RUnlock")

	b.methodBuilder.SetType(&ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				util.CreateField("index", ast.NewIdent("int")),
			},
		},
		Results: &ast.FieldList{
			List: util.GetFieldsAsAnonymous(b.params),
		},
	})
	b.methodBuilder.AddStatement(mutexLockBuilder.Build())
	b.methodBuilder.AddStatement(mutexUnlockBuilder.Build())

	results := []ast.Expr{}
	for _, param := range b.params {
		results = append(results, &ast.SelectorExpr{
			X: &ast.IndexExpr{
				X:     b.argsFieldSelector,
				Index: ast.NewIdent("index"),
			},
			Sel: ast.NewIdent(param.Names[0].String()),
		})
	}
	b.methodBuilder.AddStatement(&ast.ReturnStmt{
		Results: results,
	})
	return b.methodBuilder.Build()
}
