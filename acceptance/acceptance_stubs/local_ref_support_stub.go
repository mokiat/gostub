package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
)

type LocalRefSupportStub struct {
	MethodStub        func(arg1 alias1.Customer) (result1 alias1.Customer)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 alias1.Customer
	}
	methodReturns struct {
		result1 alias1.Customer
	}
}

func (stub *LocalRefSupportStub) Method(arg1 alias1.Customer) alias1.Customer {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 alias1.Customer
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *LocalRefSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *LocalRefSupportStub) MethodArgsForCall(index int) alias1.Customer {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *LocalRefSupportStub) MethodReturns(result1 alias1.Customer) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 alias1.Customer
	}{result1}
}
