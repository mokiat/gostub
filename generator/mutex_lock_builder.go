package generator

import "go/ast"

func NewMutexLockBuilder() *MutexLockBuilder {
	return &MutexLockBuilder{}
}

// MutexLockBuilder is reponsible for creating a lock statement for
// a mutex of a given stub method.
//
// Example:
//     func (stub *StubStruct) SumCallCount() int {
//         // ...
//         stub.sumMutex.RLock()
//         // ...
//     }
type MutexLockBuilder struct {
	receiverName   string
	mutexFieldName string
	action         string
}

func (b *MutexLockBuilder) SetReceiverName(name string) {
	b.receiverName = name
}

func (b *MutexLockBuilder) SetMutexField(name string) {
	b.mutexFieldName = name
}

func (b *MutexLockBuilder) SetAction(action string) {
	b.action = action
}

func (b *MutexLockBuilder) Build() ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
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
