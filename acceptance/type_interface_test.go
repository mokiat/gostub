package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeInterface", func() {
	var stub *acceptance_stubs.InterfaceSupportStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.InterfaceSupportStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(InterfaceSupport)
		Ω(assignable).Should(BeTrue())
	})
})
