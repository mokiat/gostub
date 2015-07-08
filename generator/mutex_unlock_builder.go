package generator

import "go/ast"

func NewMutexUnlockBuilder() *MutexUnlockBuilder {
	return &MutexUnlockBuilder{}
}

// MutexUnlockBuilder is reponsible for creating a deferred unlock
// statement for a mutex of a given stub method.
//
// Example:
//     func (stub *StubStruct) SumCallCount() int {
//         // ...
//         stub.sumMutex.RUnlock()
//         // ...
//     }
type MutexUnlockBuilder struct {
	receiverName   string
	mutexFieldName string
	action         string
}

func (b *MutexUnlockBuilder) SetReceiverName(name string) {
	b.receiverName = name
}

func (b *MutexUnlockBuilder) SetMutexField(name string) {
	b.mutexFieldName = name
}

func (b *MutexUnlockBuilder) SetAction(action string) {
	b.action = action
}

func (b *MutexUnlockBuilder) Build() ast.Stmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(b.receiverName),
					Sel: ast.NewIdent(b.mutexFieldName),
				},
				Sel: ast.NewIdent(b.action),
			},
		},
	}
}
