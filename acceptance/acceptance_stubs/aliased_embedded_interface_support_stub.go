package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"
)

type AliasedEmbeddedInterfaceSupportStub struct {
	MethodStub        func(arg1 int) (result1 int)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 int
	}
	methodReturns struct {
		result1 int
	}
	RunStub        func(arg1 alias1.Address) (result1 error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 alias1.Address
	}
	runReturns struct {
		result1 error
	}
}

func (stub *AliasedEmbeddedInterfaceSupportStub) Method(arg1 int) int {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *AliasedEmbeddedInterfaceSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *AliasedEmbeddedInterfaceSupportStub) MethodArgsForCall(index int) int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *AliasedEmbeddedInterfaceSupportStub) MethodReturns(result1 int) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 int
	}{result1}
}
func (stub *AliasedEmbeddedInterfaceSupportStub) Run(arg1 alias1.Address) error {
	stub.runMutex.Lock()
	defer stub.runMutex.Unlock()
	stub.runArgsForCall = append(stub.runArgsForCall, struct {
		arg1 alias1.Address
	}{arg1})
	if stub.RunStub != nil {
		return stub.RunStub(arg1)
	} else {
		return stub.runReturns.result1
	}
}
func (stub *AliasedEmbeddedInterfaceSupportStub) RunCallCount() int {
	stub.runMutex.RLock()
	defer stub.runMutex.RUnlock()
	return len(stub.runArgsForCall)
}
func (stub *AliasedEmbeddedInterfaceSupportStub) RunArgsForCall(index int) alias1.Address {
	stub.runMutex.RLock()
	defer stub.runMutex.RUnlock()
	return stub.runArgsForCall[index].arg1
}
func (stub *AliasedEmbeddedInterfaceSupportStub) RunReturns(result1 error) {
	stub.runMutex.Lock()
	defer stub.runMutex.Unlock()
	stub.runReturns = struct {
		result1 error
	}{result1}
}
