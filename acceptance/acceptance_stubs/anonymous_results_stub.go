// Generated by 'github.com/momchil-atanasov/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
)

type AnonymousResultsStub struct {
	StubGUID              int
	ActiveUserStub        func() (result1 int, result2 string)
	activeUserMutex       sync.RWMutex
	activeUserArgsForCall []struct {
	}
	activeUserReturns struct {
		result1 int
		result2 string
	}
}

var _ alias1.AnonymousResults = new(AnonymousResultsStub)

func (stub *AnonymousResultsStub) ActiveUser() (int, string) {
	stub.activeUserMutex.Lock()
	defer stub.activeUserMutex.Unlock()
	stub.activeUserArgsForCall = append(stub.activeUserArgsForCall, struct {
	}{})
	if stub.ActiveUserStub != nil {
		return stub.ActiveUserStub()
	} else {
		return stub.activeUserReturns.result1, stub.activeUserReturns.result2
	}
}
func (stub *AnonymousResultsStub) ActiveUserCallCount() int {
	stub.activeUserMutex.RLock()
	defer stub.activeUserMutex.RUnlock()
	return len(stub.activeUserArgsForCall)
}
func (stub *AnonymousResultsStub) ActiveUserReturns(result1 int, result2 string) {
	stub.activeUserMutex.Lock()
	defer stub.activeUserMutex.Unlock()
	stub.activeUserReturns = struct {
		result1 int
		result2 string
	}{result1, result2}
}
