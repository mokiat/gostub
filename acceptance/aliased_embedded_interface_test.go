package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AliasedEmbeddedInterface", func() {
	var stub *acceptance_stubs.AliasedEmbeddedInterfaceSupportStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.AliasedEmbeddedInterfaceSupportStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(AliasedEmbeddedInterfaceSupport)
		Ω(assignable).Should(BeTrue())
	})
})
