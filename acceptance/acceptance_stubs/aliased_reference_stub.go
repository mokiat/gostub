package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/aliased"
)

type AliasedReferenceStub struct {
	AliasedStub        func(arg1 alias1.User) (result1 alias1.User)
	aliasedMutex       sync.RWMutex
	aliasedArgsForCall []struct {
		arg1 alias1.User
	}
	aliasedReturns struct {
		result1 alias1.User
	}
	ArrayStub        func(arg1 []alias1.User) (result1 []alias1.User)
	arrayMutex       sync.RWMutex
	arrayArgsForCall []struct {
		arg1 []alias1.User
	}
	arrayReturns struct {
		result1 []alias1.User
	}
}

func (stub *AliasedReferenceStub) Aliased(arg1 alias1.User) alias1.User {
	stub.aliasedMutex.Lock()
	defer stub.aliasedMutex.Unlock()
	stub.aliasedArgsForCall = append(stub.aliasedArgsForCall, struct {
		arg1 alias1.User
	}{arg1})
	if stub.AliasedStub != nil {
		return stub.AliasedStub(arg1)
	} else {
		return stub.aliasedReturns.result1
	}
}
func (stub *AliasedReferenceStub) AliasedCallCount() int {
	stub.aliasedMutex.RLock()
	defer stub.aliasedMutex.RUnlock()
	return len(stub.aliasedArgsForCall)
}
func (stub *AliasedReferenceStub) AliasedArgsForCall(index int) alias1.User {
	stub.aliasedMutex.RLock()
	defer stub.aliasedMutex.RUnlock()
	return stub.aliasedArgsForCall[index].arg1
}
func (stub *AliasedReferenceStub) AliasedReturns(result1 alias1.User) {
	stub.aliasedMutex.Lock()
	defer stub.aliasedMutex.Unlock()
	stub.aliasedReturns = struct {
		result1 alias1.User
	}{result1}
}
func (stub *AliasedReferenceStub) Array(arg1 []alias1.User) []alias1.User {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayArgsForCall = append(stub.arrayArgsForCall, struct {
		arg1 []alias1.User
	}{arg1})
	if stub.ArrayStub != nil {
		return stub.ArrayStub(arg1)
	} else {
		return stub.arrayReturns.result1
	}
}
func (stub *AliasedReferenceStub) ArrayCallCount() int {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return len(stub.arrayArgsForCall)
}
func (stub *AliasedReferenceStub) ArrayArgsForCall(index int) []alias1.User {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return stub.arrayArgsForCall[index].arg1
}
func (stub *AliasedReferenceStub) ArrayReturns(result1 []alias1.User) {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayReturns = struct {
		result1 []alias1.User
	}{result1}
}
