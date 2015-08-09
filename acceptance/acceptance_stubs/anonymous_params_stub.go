// Generated by 'github.com/momchil-atanasov/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
)

type AnonymousParamsStub struct {
	StubGUID            int
	RegisterStub        func(arg1 string, arg2 int)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 string
		arg2 int
	}
}

var _ alias1.AnonymousParams = new(AnonymousParamsStub)

func (stub *AnonymousParamsStub) Register(arg1 string, arg2 int) {
	stub.registerMutex.Lock()
	defer stub.registerMutex.Unlock()
	stub.registerArgsForCall = append(stub.registerArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	if stub.RegisterStub != nil {
		stub.RegisterStub(arg1, arg2)
	}
}
func (stub *AnonymousParamsStub) RegisterCallCount() int {
	stub.registerMutex.RLock()
	defer stub.registerMutex.RUnlock()
	return len(stub.registerArgsForCall)
}
func (stub *AnonymousParamsStub) RegisterArgsForCall(index int) (string, int) {
	stub.registerMutex.RLock()
	defer stub.registerMutex.RUnlock()
	return stub.registerArgsForCall[index].arg1, stub.registerArgsForCall[index].arg2
}
