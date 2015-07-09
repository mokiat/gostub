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
	mutexFieldSelector *ast.SelectorExpr
	action             string
}

func (b *MutexLockBuilder) SetMutexFieldSelector(selector *ast.SelectorExpr) {
	b.mutexFieldSelector = selector
}

func (b *MutexLockBuilder) SetAction(action string) {
	b.action = action
}

func (b *MutexLockBuilder) Build() ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   b.mutexFieldSelector,
				Sel: ast.NewIdent(b.action),
			},
		},
	}
}
