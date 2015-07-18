package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/external"
)

type ExternalReferenceStub struct {
	ExternalStub        func(arg1 alias1.Address) (result1 alias1.Address)
	externalMutex       sync.RWMutex
	externalArgsForCall []struct {
		arg1 alias1.Address
	}
	externalReturns struct {
		result1 alias1.Address
	}
	ArrayStub        func(arg1 [3]alias1.Address) (result1 [3]alias1.Address)
	arrayMutex       sync.RWMutex
	arrayArgsForCall []struct {
		arg1 [3]alias1.Address
	}
	arrayReturns struct {
		result1 [3]alias1.Address
	}
	SliceStub        func(arg1 []alias1.Address) (result1 []alias1.Address)
	sliceMutex       sync.RWMutex
	sliceArgsForCall []struct {
		arg1 []alias1.Address
	}
	sliceReturns struct {
		result1 []alias1.Address
	}
	PointerStub        func(arg1 *alias1.Address) (result1 *alias1.Address)
	pointerMutex       sync.RWMutex
	pointerArgsForCall []struct {
		arg1 *alias1.Address
	}
	pointerReturns struct {
		result1 *alias1.Address
	}
}

func (stub *ExternalReferenceStub) External(arg1 alias1.Address) alias1.Address {
	stub.externalMutex.Lock()
	defer stub.externalMutex.Unlock()
	stub.externalArgsForCall = append(stub.externalArgsForCall, struct {
		arg1 alias1.Address
	}{arg1})
	if stub.ExternalStub != nil {
		return stub.ExternalStub(arg1)
	} else {
		return stub.externalReturns.result1
	}
}
func (stub *ExternalReferenceStub) ExternalCallCount() int {
	stub.externalMutex.RLock()
	defer stub.externalMutex.RUnlock()
	return len(stub.externalArgsForCall)
}
func (stub *ExternalReferenceStub) ExternalArgsForCall(index int) alias1.Address {
	stub.externalMutex.RLock()
	defer stub.externalMutex.RUnlock()
	return stub.externalArgsForCall[index].arg1
}
func (stub *ExternalReferenceStub) ExternalReturns(result1 alias1.Address) {
	stub.externalMutex.Lock()
	defer stub.externalMutex.Unlock()
	stub.externalReturns = struct {
		result1 alias1.Address
	}{result1}
}
func (stub *ExternalReferenceStub) Array(arg1 [3]alias1.Address) [3]alias1.Address {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayArgsForCall = append(stub.arrayArgsForCall, struct {
		arg1 [3]alias1.Address
	}{arg1})
	if stub.ArrayStub != nil {
		return stub.ArrayStub(arg1)
	} else {
		return stub.arrayReturns.result1
	}
}
func (stub *ExternalReferenceStub) ArrayCallCount() int {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return len(stub.arrayArgsForCall)
}
func (stub *ExternalReferenceStub) ArrayArgsForCall(index int) [3]alias1.Address {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return stub.arrayArgsForCall[index].arg1
}
func (stub *ExternalReferenceStub) ArrayReturns(result1 [3]alias1.Address) {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayReturns = struct {
		result1 [3]alias1.Address
	}{result1}
}
func (stub *ExternalReferenceStub) Slice(arg1 []alias1.Address) []alias1.Address {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceArgsForCall = append(stub.sliceArgsForCall, struct {
		arg1 []alias1.Address
	}{arg1})
	if stub.SliceStub != nil {
		return stub.SliceStub(arg1)
	} else {
		return stub.sliceReturns.result1
	}
}
func (stub *ExternalReferenceStub) SliceCallCount() int {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return len(stub.sliceArgsForCall)
}
func (stub *ExternalReferenceStub) SliceArgsForCall(index int) []alias1.Address {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return stub.sliceArgsForCall[index].arg1
}
func (stub *ExternalReferenceStub) SliceReturns(result1 []alias1.Address) {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceReturns = struct {
		result1 []alias1.Address
	}{result1}
}
func (stub *ExternalReferenceStub) Pointer(arg1 *alias1.Address) *alias1.Address {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerArgsForCall = append(stub.pointerArgsForCall, struct {
		arg1 *alias1.Address
	}{arg1})
	if stub.PointerStub != nil {
		return stub.PointerStub(arg1)
	} else {
		return stub.pointerReturns.result1
	}
}
func (stub *ExternalReferenceStub) PointerCallCount() int {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return len(stub.pointerArgsForCall)
}
func (stub *ExternalReferenceStub) PointerArgsForCall(index int) *alias1.Address {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return stub.pointerArgsForCall[index].arg1
}
func (stub *ExternalReferenceStub) PointerReturns(result1 *alias1.Address) {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerReturns = struct {
		result1 *alias1.Address
	}{result1}
}
