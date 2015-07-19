package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/embedded"
)

type EmbeddedRefSupportStub struct {
	MethodStub        func(arg1 alias1.Resource) (result1 alias1.Resource)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 alias1.Resource
	}
	methodReturns struct {
		result1 alias1.Resource
	}
}

func (stub *EmbeddedRefSupportStub) Method(arg1 alias1.Resource) alias1.Resource {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 alias1.Resource
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *EmbeddedRefSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *EmbeddedRefSupportStub) MethodArgsForCall(index int) alias1.Resource {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *EmbeddedRefSupportStub) MethodReturns(result1 alias1.Resource) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 alias1.Resource
	}{result1}
}