package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeMap", func() {
	var stub *acceptance_stubs.MapSupportStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.MapSupportStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(MapSupport)
		Ω(assignable).Should(BeTrue())
	})
})
