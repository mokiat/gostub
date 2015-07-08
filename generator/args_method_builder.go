package generator

import (
	"go/ast"

	"github.com/momchil-atanasov/gostub/util"
)

func NewArgsMethodBuilder() *ArgsMethodBuilder {
	return &ArgsMethodBuilder{
		selfName: "stub",
		params:   make([]*ast.Field, 0),
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
	selfName       string
	selfType       string
	methodName     string
	argsFieldName  string
	mutexFieldName string
	params         []*ast.Field
}

func (b *ArgsMethodBuilder) SetReceiverName(name string) {
	b.selfName = name
}

func (b *ArgsMethodBuilder) SetReceiverType(name string) {
	b.selfType = name
}

func (b *ArgsMethodBuilder) SetMethodName(name string) {
	b.methodName = name
}

func (b *ArgsMethodBuilder) SetArgsFieldName(name string) {
	b.argsFieldName = name
}

func (b *ArgsMethodBuilder) SetMutexFieldName(name string) {
	b.mutexFieldName = name
}

// SetParams specifies the parameters that the original method
// uses. These parameters need to have been normalized and resolved
// in advance.
func (b *ArgsMethodBuilder) SetParams(params []*ast.Field) {
	b.params = params
}

func (b *ArgsMethodBuilder) Build() *ast.FuncDecl {
	mutexLockBuilder := NewMutexLockBuilder()
	mutexLockBuilder.SetReceiverName(b.selfName)
	mutexLockBuilder.SetMutexField(b.mutexFieldName)
	mutexLockBuilder.SetAction("RLock")

	mutexUnlockBuilder := NewMutexUnlockBuilder()
	mutexUnlockBuilder.SetReceiverName(b.selfName)
	mutexUnlockBuilder.SetMutexField(b.mutexFieldName)
	mutexUnlockBuilder.SetAction("RUnlock")

	method := NewMethodModel()
	method.SetName(b.methodName)
	method.SetReceiver(b.selfName, b.selfType)
	method.SetType(&ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				util.CreateField("index", ast.NewIdent("int")),
			},
		},
		Results: &ast.FieldList{
			List: util.GetFieldsAsAnonymous(b.params),
		},
	})
	method.AddStatement(mutexLockBuilder.Build())
	method.AddStatement(mutexUnlockBuilder.Build())

	results := []ast.Expr{}
	for _, param := range b.params {
		results = append(results, &ast.SelectorExpr{
			X: &ast.IndexExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(b.selfName),
					Sel: ast.NewIdent(b.argsFieldName),
				},
				Index: ast.NewIdent("index"),
			},
			Sel: ast.NewIdent(param.Names[0].String()),
		})
	}
	method.AddStatement(&ast.ReturnStmt{
		Results: results,
	})
	return method.BuildASTFuncDecl()
}
