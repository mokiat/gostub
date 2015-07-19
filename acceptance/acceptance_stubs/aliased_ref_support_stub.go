package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/aliased"
)

type AliasedRefSupportStub struct {
	MethodStub        func(arg1 alias1.User) (result1 alias1.User)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 alias1.User
	}
	methodReturns struct {
		result1 alias1.User
	}
}

func (stub *AliasedRefSupportStub) Method(arg1 alias1.User) alias1.User {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 alias1.User
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *AliasedRefSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *AliasedRefSupportStub) MethodArgsForCall(index int) alias1.User {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *AliasedRefSupportStub) MethodReturns(result1 alias1.User) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 alias1.User
	}{result1}
}