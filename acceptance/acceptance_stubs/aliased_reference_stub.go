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
	SliceStub        func(arg1 []alias1.User) (result1 []alias1.User)
	sliceMutex       sync.RWMutex
	sliceArgsForCall []struct {
		arg1 []alias1.User
	}
	sliceReturns struct {
		result1 []alias1.User
	}
	PointerStub        func(arg1 *alias1.User) (result1 *alias1.User)
	pointerMutex       sync.RWMutex
	pointerArgsForCall []struct {
		arg1 *alias1.User
	}
	pointerReturns struct {
		result1 *alias1.User
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
func (stub *AliasedReferenceStub) Slice(arg1 []alias1.User) []alias1.User {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceArgsForCall = append(stub.sliceArgsForCall, struct {
		arg1 []alias1.User
	}{arg1})
	if stub.SliceStub != nil {
		return stub.SliceStub(arg1)
	} else {
		return stub.sliceReturns.result1
	}
}
func (stub *AliasedReferenceStub) SliceCallCount() int {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return len(stub.sliceArgsForCall)
}
func (stub *AliasedReferenceStub) SliceArgsForCall(index int) []alias1.User {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return stub.sliceArgsForCall[index].arg1
}
func (stub *AliasedReferenceStub) SliceReturns(result1 []alias1.User) {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceReturns = struct {
		result1 []alias1.User
	}{result1}
}
func (stub *AliasedReferenceStub) Pointer(arg1 *alias1.User) *alias1.User {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerArgsForCall = append(stub.pointerArgsForCall, struct {
		arg1 *alias1.User
	}{arg1})
	if stub.PointerStub != nil {
		return stub.PointerStub(arg1)
	} else {
		return stub.pointerReturns.result1
	}
}
func (stub *AliasedReferenceStub) PointerCallCount() int {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return len(stub.pointerArgsForCall)
}
func (stub *AliasedReferenceStub) PointerArgsForCall(index int) *alias1.User {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return stub.pointerArgsForCall[index].arg1
}
func (stub *AliasedReferenceStub) PointerReturns(result1 *alias1.User) {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerReturns = struct {
		result1 *alias1.User
	}{result1}
}
