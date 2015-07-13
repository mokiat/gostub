package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/mismatch"
)

type MismatchedReferenceStub struct {
	MismatchedStub        func(arg1 alias1.Job)
	mismatchedMutex       sync.RWMutex
	mismatchedArgsForCall []struct {
		arg1 alias1.Job
	}
}

func (stub *MismatchedReferenceStub) Mismatched(arg1 alias1.Job) {
	stub.mismatchedMutex.Lock()
	defer stub.mismatchedMutex.Unlock()
	stub.mismatchedArgsForCall = append(stub.mismatchedArgsForCall, struct {
		arg1 alias1.Job
	}{arg1})
	if stub.MismatchedStub != nil {
		stub.MismatchedStub(arg1)
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