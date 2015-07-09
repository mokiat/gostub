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
	mutexFieldSelector *ast.SelectorExpr
	action             string
}

func (b *MutexUnlockBuilder) SetMutexFieldSelector(selector *ast.SelectorExpr) {
	b.mutexFieldSelector = selector
}

func (b *MutexUnlockBuilder) SetAction(action string) {
	b.action = action
}

func (b *MutexUnlockBuilder) Build() ast.Stmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   b.mutexFieldSelector,
				Sel: ast.NewIdent(b.action),
			},
		},
	}
}
