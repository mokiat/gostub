package testables_stubs

import (
	sync "sync"
)

type NoParamsNoResultsStub struct {
	RunStub        func()
	runMutex       sync.RWMutex
	runArgsForCall []struct {
	}
}

func (stub *NoParamsNoResultsStub) Run() {
	stub.runMutex.Lock()
	defer stub.runMutex.Unlock()
	stub.runArgsForCall = append(stub.runArgsForCall, struct {
	}{})
	if stub.RunStub != nil {
		stub.RunStub()
	}
}
func (stub *NoParamsNoResultsStub) RunCallCount() int {
	stub.runMutex.RLock()
	defer stub.runMutex.RUnlock()
	return len(stub.runArgsForCall)
}
