package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmptyInterface", func() {
	var stub *acceptance_stubs.EmptyInterfaceStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.EmptyInterfaceStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(EmptyInterface)
		Ω(assignable).Should(BeTrue())
	})
})
