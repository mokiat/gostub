package acceptance_stubs

import (
	sync "sync"
)

type PrimitiveParamsStub struct {
	SaveStub        func(count int, location string, timeout float32)
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		count    int
		location string
		timeout  float32
	}
}

func (stub *PrimitiveParamsStub) Save(count int, location string, timeout float32) {
	stub.saveMutex.Lock()
	defer stub.saveMutex.Unlock()
	stub.saveArgsForCall = append(stub.saveArgsForCall, struct {
		count    int
		location string
		timeout  float32
	}{count, location, timeout})
	if stub.SaveStub != nil {
		stub.SaveStub(count, location, timeout)
	}
}
func (stub *PrimitiveParamsStub) SaveCallCount() int {
	stub.saveMutex.RLock()
	defer stub.saveMutex.RUnlock()
	return len(stub.saveArgsForCall)
}
func (stub *PrimitiveParamsStub) SaveArgsForCall(index int) (int, string, float32) {
	stub.saveMutex.RLock()
	defer stub.saveMutex.RUnlock()
	return stub.saveArgsForCall[index].count, stub.saveArgsForCall[index].location, stub.saveArgsForCall[index].timeout
}
