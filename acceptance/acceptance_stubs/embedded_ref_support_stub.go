// Generated by 'github.com/momchil-atanasov/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
	alias2 "github.com/momchil-atanasov/gostub/acceptance/embedded"
)

type EmbeddedRefSupportStub struct {
	StubGUID          int
	MethodStub        func(arg1 alias2.Resource) (result1 alias2.Resource)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 alias2.Resource
	}
	methodReturns struct {
		result1 alias2.Resource
	}
}

var _ alias1.EmbeddedRefSupport = new(EmbeddedRefSupportStub)

func (stub *EmbeddedRefSupportStub) Method(arg1 alias2.Resource) alias2.Resource {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 alias2.Resource
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
func (stub *EmbeddedRefSupportStub) MethodArgsForCall(index int) alias2.Resource {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *EmbeddedRefSupportStub) MethodReturns(result1 alias2.Resource) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 alias2.Resource
	}{result1}
}
