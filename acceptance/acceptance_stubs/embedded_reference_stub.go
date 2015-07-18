package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/embedded"
)

type EmbeddedReferenceStub struct {
	EmbeddedStub        func(arg1 alias1.Resource) (result1 alias1.Resource)
	embeddedMutex       sync.RWMutex
	embeddedArgsForCall []struct {
		arg1 alias1.Resource
	}
	embeddedReturns struct {
		result1 alias1.Resource
	}
}

func (stub *EmbeddedReferenceStub) Embedded(arg1 alias1.Resource) alias1.Resource {
	stub.embeddedMutex.Lock()
	defer stub.embeddedMutex.Unlock()
	stub.embeddedArgsForCall = append(stub.embeddedArgsForCall, struct {
		arg1 alias1.Resource
	}{arg1})
	if stub.EmbeddedStub != nil {
		return stub.EmbeddedStub(arg1)
	} else {
		return stub.embeddedReturns.result1
	}
}
func (stub *EmbeddedReferenceStub) EmbeddedCallCount() int {
	stub.embeddedMutex.RLock()
	defer stub.embeddedMutex.RUnlock()
	return len(stub.embeddedArgsForCall)
}
func (stub *EmbeddedReferenceStub) EmbeddedArgsForCall(index int) alias1.Resource {
	stub.embeddedMutex.RLock()
	defer stub.embeddedMutex.RUnlock()
	return stub.embeddedArgsForCall[index].arg1
}
func (stub *EmbeddedReferenceStub) EmbeddedReturns(result1 alias1.Resource) {
	stub.embeddedMutex.Lock()
	defer stub.embeddedMutex.Unlock()
	stub.embeddedReturns = struct {
		result1 alias1.Resource
	}{result1}
}
