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
	ArrayStub        func(arg1 []alias1.Job) (result1 []alias1.Job)
	arrayMutex       sync.RWMutex
	arrayArgsForCall []struct {
		arg1 []alias1.Job
	}
	arrayReturns struct {
		result1 []alias1.Job
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
func (stub *MismatchedReferenceStub) Array(arg1 []alias1.Job) []alias1.Job {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayArgsForCall = append(stub.arrayArgsForCall, struct {
		arg1 []alias1.Job
	}{arg1})
	if stub.ArrayStub != nil {
		return stub.ArrayStub(arg1)
	} else {
		return stub.arrayReturns.result1
	}
}
func (stub *MismatchedReferenceStub) ArrayCallCount() int {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return len(stub.arrayArgsForCall)
}
func (stub *MismatchedReferenceStub) ArrayArgsForCall(index int) []alias1.Job {
	stub.arrayMutex.RLock()
	defer stub.arrayMutex.RUnlock()
	return stub.arrayArgsForCall[index].arg1
}
func (stub *MismatchedReferenceStub) ArrayReturns(result1 []alias1.Job) {
	stub.arrayMutex.Lock()
	defer stub.arrayMutex.Unlock()
	stub.arrayReturns = struct {
		result1 []alias1.Job
	}{result1}
}
