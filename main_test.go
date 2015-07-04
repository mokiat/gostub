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
})
