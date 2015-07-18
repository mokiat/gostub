package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/mismatch"
)

type MismatchedReferenceStub struct {
	MismatchedStub        func(arg1 alias1.Job) (result1 alias1.Job)
	mismatchedMutex       sync.RWMutex
	mismatchedArgsForCall []struct {
		arg1 alias1.Job
	}
	mismatchedReturns struct {
		result1 alias1.Job
	}
	SliceStub        func(arg1 []alias1.Job) (result1 []alias1.Job)
	sliceMutex       sync.RWMutex
	sliceArgsForCall []struct {
		arg1 []alias1.Job
	}
	sliceReturns struct {
		result1 []alias1.Job
	}
	PointerStub        func(arg1 *alias1.Job) (result1 *alias1.Job)
	pointerMutex       sync.RWMutex
	pointerArgsForCall []struct {
		arg1 *alias1.Job
	}
	pointerReturns struct {
		result1 *alias1.Job
	}
}

func (stub *MismatchedReferenceStub) Mismatched(arg1 alias1.Job) alias1.Job {
	stub.mismatchedMutex.Lock()
	defer stub.mismatchedMutex.Unlock()
	stub.mismatchedArgsForCall = append(stub.mismatchedArgsForCall, struct {
		arg1 alias1.Job
	}{arg1})
	if stub.MismatchedStub != nil {
		return stub.MismatchedStub(arg1)
	} else {
		return stub.mismatchedReturns.result1
	}
}
func (stub *MismatchedReferenceStub) MismatchedCallCount() int {
	stub.mismatchedMutex.RLock()
	defer stub.mismatchedMutex.RUnlock()
	return len(stub.mismatchedArgsForCall)
}
func (stub *MismatchedReferenceStub) MismatchedArgsForCall(index int) alias1.Job {
	stub.mismatchedMutex.RLock()
	defer stub.mismatchedMutex.RUnlock()
	return stub.mismatchedArgsForCall[index].arg1
}
func (stub *MismatchedReferenceStub) MismatchedReturns(result1 alias1.Job) {
	stub.mismatchedMutex.Lock()
	defer stub.mismatchedMutex.Unlock()
	stub.mismatchedReturns = struct {
		result1 alias1.Job
	}{result1}
}
func (stub *MismatchedReferenceStub) Slice(arg1 []alias1.Job) []alias1.Job {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceArgsForCall = append(stub.sliceArgsForCall, struct {
		arg1 []alias1.Job
	}{arg1})
	if stub.SliceStub != nil {
		return stub.SliceStub(arg1)
	} else {
		return stub.sliceReturns.result1
	}
}
func (stub *MismatchedReferenceStub) SliceCallCount() int {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return len(stub.sliceArgsForCall)
}
func (stub *MismatchedReferenceStub) SliceArgsForCall(index int) []alias1.Job {
	stub.sliceMutex.RLock()
	defer stub.sliceMutex.RUnlock()
	return stub.sliceArgsForCall[index].arg1
}
func (stub *MismatchedReferenceStub) SliceReturns(result1 []alias1.Job) {
	stub.sliceMutex.Lock()
	defer stub.sliceMutex.Unlock()
	stub.sliceReturns = struct {
		result1 []alias1.Job
	}{result1}
}
func (stub *MismatchedReferenceStub) Pointer(arg1 *alias1.Job) *alias1.Job {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerArgsForCall = append(stub.pointerArgsForCall, struct {
		arg1 *alias1.Job
	}{arg1})
	if stub.PointerStub != nil {
		return stub.PointerStub(arg1)
	} else {
		return stub.pointerReturns.result1
	}
}
func (stub *MismatchedReferenceStub) PointerCallCount() int {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return len(stub.pointerArgsForCall)
}
func (stub *MismatchedReferenceStub) PointerArgsForCall(index int) *alias1.Job {
	stub.pointerMutex.RLock()
	defer stub.pointerMutex.RUnlock()
	return stub.pointerArgsForCall[index].arg1
}
func (stub *MismatchedReferenceStub) PointerReturns(result1 *alias1.Job) {
	stub.pointerMutex.Lock()
	defer stub.pointerMutex.Unlock()
	stub.pointerReturns = struct {
		result1 *alias1.Job
	}{result1}
}
