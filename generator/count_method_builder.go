package generator

import "go/ast"

func NewCountMethodBuilder() *CountMethodBuilder {
	return &CountMethodBuilder{
		selfName: "stub",
	}
}

// CountMethodBuilder is reponsible for creating a method on the stub
// structure that allows you to check how many times the stubbed method
// was called.
//
// Example:
//     func (stub *StubStruct) SumCallCount() int {
//         // ...
//     }
type CountMethodBuilder struct {
	selfName       string
	selfType       string
	methodName     string
	argsFieldName  string
	mutexFieldName string
}

func (b *CountMethodBuilder) SetReceiverName(name string) {
	b.selfName = name
}

func (b *CountMethodBuilder) SetReceiverType(name string) {
	b.selfType = name
}

func (b *CountMethodBuilder) SetMethodName(name string) {
	b.methodName = name
}

func (b *CountMethodBuilder) SetArgsFieldName(name string) {
	b.argsFieldName = name
}

func (b *CountMethodBuilder) SetMutexFieldName(name string) {
	b.mutexFieldName = name
}

func (b *CountMethodBuilder) Build() *ast.FuncDecl {
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
		Params: &ast.FieldList{},
		Results: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Type: ast.NewIdent("int"),
				},
			},
		},
	})
	method.AddStatement(mutexLockBuilder.Build())
	method.AddStatement(mutexUnlockBuilder.Build())
	method.AddStatement(&ast.ReturnStmt{
		Results: []ast.Expr{
			&ast.CallExpr{
				Fun: ast.NewIdent("len"),
				Args: []ast.Expr{
					&ast.SelectorExpr{
						X:   ast.NewIdent(b.selfName),
						Sel: ast.NewIdent(b.argsFieldName),
					},
				},
			},
		},
	})
	return method.BuildASTFuncDecl()
}
