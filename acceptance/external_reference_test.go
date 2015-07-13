package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExternalReference", func() {
	var stub *acceptance_stubs.ExternalReferenceStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.ExternalReferenceStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(ExternalReference)
		Ω(assignable).Should(BeTrue())
	})
})
