package acceptance_stubs

import (
	sync "sync"
)

type ReusedParamsStub struct {
	ConcatStub        func(arg1 string, arg2 string)
	concatMutex       sync.RWMutex
	concatArgsForCall []struct {
		arg1 string
		arg2 string
	}
}

func (stub *ReusedParamsStub) Concat(arg1 string, arg2 string) {
	stub.concatMutex.Lock()
	defer stub.concatMutex.Unlock()
	stub.concatArgsForCall = append(stub.concatArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if stub.ConcatStub != nil {
		stub.ConcatStub(arg1, arg2)
	}
}
func (stub *ReusedParamsStub) ConcatCallCount() int {
	stub.concatMutex.RLock()
	defer stub.concatMutex.RUnlock()
	return len(stub.concatArgsForCall)
}
func (stub *ReusedParamsStub) ConcatArgsForCall(index int) (string, string) {
	stub.concatMutex.RLock()
	defer stub.concatMutex.RUnlock()
	return stub.concatArgsForCall[index].arg1, stub.concatArgsForCall[index].arg2
}
