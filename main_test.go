package main_test

import (
	"github.com/momchil-atanasov/gostub/testables"
	"github.com/momchil-atanasov/gostub/testables/testables_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Note: These tests should be executed through `scripts/test`
// which assures that the necessary stubs are created.

var _ = Describe("Main", func() {
	const threshold = 0.0001

	Describe("Empty interface", func() {
		var stub *testables_stubs.EmptyInterfaceStub

		BeforeEach(func() {
			stub = new(testables_stubs.EmptyInterfaceStub)
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(testables.EmptyInterface)
			Ω(assignable).Should(BeTrue())
		})
	})

	Describe("NoParamsNoResults", func() {
		var stub *testables_stubs.NoParamsNoResultsStub
		var runWasCalled bool

		BeforeEach(func() {
			stub = new(testables_stubs.NoParamsNoResultsStub)
			runWasCalled = false
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(testables.NoParamsNoResults)
			Ω(assignable).Should(BeTrue())
		})

		It("is possible to stub the behavior", func() {
			stub.RunStub = func() {
				runWasCalled = true
			}
			stub.Run()
			Ω(runWasCalled).Should(BeTrue())
		})

		It("is possible to get call count", func() {
			stub.Run()
			stub.Run()
			Ω(stub.RunCallCount()).Should(Equal(2))
		})
	})

	Describe("PrimitiveParams", func() {
		var stub *testables_stubs.PrimitiveParamsStub
		var runWasCalled bool
		var runCountArg int
		var runLocationArg string
		var runTimeoutArg float32

		BeforeEach(func() {
			stub = new(testables_stubs.PrimitiveParamsStub)
			runWasCalled = false
			runCountArg = 0
			runLocationArg = ""
			runTimeoutArg = 0.0
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(testables.PrimitiveParams)
			Ω(assignable).Should(BeTrue())
		})

		It("is possible to stub the behavior", func() {
			stub.SaveStub = func(count int, location string, timeout float32) {
				runWasCalled = true
				runCountArg = count
				runLocationArg = location
				runTimeoutArg = timeout
			}
			stub.Save(10, "/tmp", 3.14)
			Ω(runWasCalled).Should(BeTrue())
			Ω(runCountArg).Should(Equal(10))
			Ω(runLocationArg).Should(Equal("/tmp"))
			Ω(runTimeoutArg).Should(BeNumerically("~", 3.14, threshold))
		})

		It("is possible to get call count", func() {
			stub.Save(1, "/first", 3.3)
			stub.Save(2, "/second", 5.5)
			Ω(stub.SaveCallCount()).Should(Equal(2))
		})

		It("is possible to get arguments for call", func() {
			stub.Save(1, "/first", 3.3)
			stub.Save(2, "/second", 5.5)
			argCount, argLocation, argTimeout := stub.SaveArgsForCall(0)
			Ω(argCount).Should(Equal(1))
			Ω(argLocation).Should(Equal("/first"))
			Ω(argTimeout).Should(BeNumerically("~", 3.3, threshold))

			argCount, argLocation, argTimeout = stub.SaveArgsForCall(1)
			Ω(argCount).Should(Equal(2))
			Ω(argLocation).Should(Equal("/second"))
			Ω(argTimeout).Should(BeNumerically("~", 5.5, threshold))
		})
	})
})
