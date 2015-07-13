package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/aliased"
)

type AliasedReferenceStub struct {
	AliasedStub        func(arg1 alias1.User)
	aliasedMutex       sync.RWMutex
	aliasedArgsForCall []struct {
		arg1 alias1.User
	}
}

func (stub *AliasedReferenceStub) Aliased(arg1 alias1.User) {
	stub.aliasedMutex.Lock()
	defer stub.aliasedMutex.Unlock()
	stub.aliasedArgsForCall = append(stub.aliasedArgsForCall, struct {
		arg1 alias1.User
	}{arg1})
	if stub.AliasedStub != nil {
		stub.AliasedStub(arg1)
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