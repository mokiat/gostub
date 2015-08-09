// Generated by 'github.com/momchil-atanasov/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
	alias2 "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"
)

type PointerSupportStub struct {
	StubGUID          int
	MethodStub        func(arg1 *alias2.Address) (result1 *alias2.Address)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 *alias2.Address
	}
	methodReturns struct {
		result1 *alias2.Address
	}
}

var _ alias1.PointerSupport = new(PointerSupportStub)

func (stub *PointerSupportStub) Method(arg1 *alias2.Address) *alias2.Address {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 *alias2.Address
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *PointerSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *PointerSupportStub) MethodArgsForCall(index int) *alias2.Address {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *PointerSupportStub) MethodReturns(result1 *alias2.Address) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 *alias2.Address
	}{result1}
}
