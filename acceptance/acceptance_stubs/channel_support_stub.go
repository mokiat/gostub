// Generated by 'github.com/mokiat/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/mokiat/gostub/acceptance"
	alias2 "github.com/mokiat/gostub/acceptance/external/external_dup"
)

type ChannelSupportStub struct {
	StubGUID          int
	MethodStub        func(arg1 chan alias2.Address) (result1 chan alias2.Address)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 chan alias2.Address
	}
	methodReturns struct {
		result1 chan alias2.Address
	}
}

var _ alias1.ChannelSupport = new(ChannelSupportStub)

func (stub *ChannelSupportStub) Method(arg1 chan alias2.Address) chan alias2.Address {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 chan alias2.Address
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *ChannelSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *ChannelSupportStub) MethodArgsForCall(index int) chan alias2.Address {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *ChannelSupportStub) MethodReturns(result1 chan alias2.Address) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 chan alias2.Address
	}{result1}
}
