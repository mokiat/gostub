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
	ArrayStub        func(arg1 [3]alias1.Resource) (result1 [3]alias1.Resource)
	arrayMutex       sync.RWMutex
	arrayArgsForCall []struct {
		arg1 [3]alias1.Resource
	}
	arrayReturns struct {
		result1 [3]alias1.Resource
	}
	SliceStub        func(arg1 []alias1.Resource) (result1 []alias1.Resource)
	sliceMutex       sync.RWMutex
	sliceArgsForCall []struct {
		arg1 []alias1.Resource
	}
	sliceReturns struct {
		result1 []alias1.Resource
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
func (stub *EmbeddedReferenceStub) Array(arg1 [3]alias1.Resource) [3]alias1.Resource {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayArgsForCall = append(stub.arrayArgsForCall, struct {
		arg1 [3]alias1.Resource
	}{arg1})
	if stub.ArrayStub != nil {
		return stub.ArrayStub(arg1)
	} else {
		return stub.arrayReturns.result1
	}
}
func (stub *EmbeddedReferenceStub) ArrayCallCount() int {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return len(stub.arrayArgsForCall)
}
func (stub *EmbeddedReferenceStub) ArrayArgsForCall(index int) [3]alias1.Resource {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return stub.arrayArgsForCall[index].arg1
}
func (stub *EmbeddedReferenceStub) ArrayReturns(result1 [3]alias1.Resource) {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayReturns = struct {
		result1 [3]alias1.Resource
	}{result1}
}
func (stub *EmbeddedReferenceStub) Slice(arg1 []alias1.Resource) []alias1.Resource {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceArgsForCall = append(stub.sliceArgsForCall, struct {
		arg1 []alias1.Resource
	}{arg1})
	if stub.SliceStub != nil {
		return stub.SliceStub(arg1)
	} else {
		return stub.sliceReturns.result1
	}
}
func (stub *EmbeddedReferenceStub) SliceCallCount() int {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return len(stub.sliceArgsForCall)
}
func (stub *EmbeddedReferenceStub) SliceArgsForCall(index int) []alias1.Resource {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return stub.sliceArgsForCall[index].arg1
}
func (stub *EmbeddedReferenceStub) SliceReturns(result1 []alias1.Resource) {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceReturns = struct {
		result1 []alias1.Resource
	}{result1}
}
