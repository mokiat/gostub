package acceptance_test

import (
	"github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Note: These tests should be executed through `scripts/test`
// which assures that the necessary stubs are created.

var _ = Describe("Main", func() {
	const threshold = 0.0001

	Describe("Empty interface", func() {
		var stub *acceptance_stubs.EmptyInterfaceStub

		BeforeEach(func() {
			stub = new(acceptance_stubs.EmptyInterfaceStub)
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(acceptance.EmptyInterface)
			Ω(assignable).Should(BeTrue())
		})
	})

	Describe("NoParamsNoResults", func() {
		var stub *acceptance_stubs.NoParamsNoResultsStub
		var runWasCalled bool

		BeforeEach(func() {
			stub = new(acceptance_stubs.NoParamsNoResultsStub)
			runWasCalled = false
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(acceptance.NoParamsNoResults)
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
		var stub *acceptance_stubs.PrimitiveParamsStub
		var runWasCalled bool
		var runCountArg int
		var runLocationArg string
		var runTimeoutArg float32

		BeforeEach(func() {
			stub = new(acceptance_stubs.PrimitiveParamsStub)
			runWasCalled = false
			runCountArg = 0
			runLocationArg = ""
			runTimeoutArg = 0.0
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(acceptance.PrimitiveParams)
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

	Describe("PrimitiveResults", func() {
		var stub *acceptance_stubs.PrimitiveResultsStub

		BeforeEach(func() {
			stub = new(acceptance_stubs.PrimitiveResultsStub)
		})

		It("stub is assignable to interface", func() {
			_, assignable := interface{}(stub).(acceptance.PrimitiveResults)
			Ω(assignable).Should(BeTrue())
		})

		It("is possible to stub the behavior", func() {
			stub.UserStub = func() (name string, age int, height float32) {
				return "John", 31, 1.83
			}
			name, age, height := stub.User()
			Ω(name).Should(Equal("John"))
			Ω(age).Should(Equal(31))
			Ω(height).Should(BeNumerically("~", 1.83, threshold))
		})

		It("is possible to get call count", func() {
			stub.User()
			stub.User()
			Ω(stub.UserCallCount()).Should(Equal(2))
		})

		It("is possible to stub results", func() {
			stub.UserReturns("Jack", 53, 1.69)

			name, age, height := stub.User()
			Ω(name).Should(Equal("Jack"))
			Ω(age).Should(Equal(53))
			Ω(height).Should(BeNumerically("~", 1.69, threshold))
		})
	})
})
