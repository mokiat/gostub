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
})
