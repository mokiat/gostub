package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AliasedReference", func() {
	var stub *acceptance_stubs.AliasedReferenceStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.AliasedReferenceStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(AliasedReference)
		Ω(assignable).Should(BeTrue())
	})
})
