package acceptance_stubs

import (
	sync "sync"
)

type PrimitiveResultsStub struct {
	UserStub        func() (name string, age int, height float32)
	userMutex       sync.RWMutex
	userArgsForCall []struct {
	}
	userReturns struct {
		name   string
		age    int
		height float32
	}
}

func (stub *PrimitiveResultsStub) User() (string, int, float32) {
	stub.userMutex.Lock()
	defer stub.userMutex.Unlock()
	stub.userArgsForCall = append(stub.userArgsForCall, struct {
	}{})
	if stub.UserStub != nil {
		return stub.UserStub()
	} else {
		return stub.userReturns.name, stub.userReturns.age, stub.userReturns.height
	}
}
func (stub *PrimitiveResultsStub) UserCallCount() int {
	stub.userMutex.RLock()
	defer stub.userMutex.RUnlock()
	return len(stub.userArgsForCall)
}
func (stub *PrimitiveResultsStub) UserReturns(name string, age int, height float32) {
	stub.userMutex.Lock()
	defer stub.userMutex.Unlock()
	stub.userReturns = struct {
		name   string
		age    int
		height float32
	}{name, age, height}
}
