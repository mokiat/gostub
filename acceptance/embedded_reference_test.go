package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmbeddedReference", func() {
	var stub *acceptance_stubs.EmbeddedReferenceStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.EmbeddedReferenceStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(EmbeddedReference)
		Ω(assignable).Should(BeTrue())
	})
})
