package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MismatchedReference", func() {
	var stub *acceptance_stubs.MismatchedReferenceStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.MismatchedReferenceStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(MismatchedReference)
		Ω(assignable).Should(BeTrue())
	})
})
