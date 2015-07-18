package acceptance_test

import (
	. "github.com/momchil-atanasov/gostub/acceptance"
	"github.com/momchil-atanasov/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeChannel", func() {
	var stub *acceptance_stubs.ChannelSupportStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.ChannelSupportStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(ChannelSupport)
		Ω(assignable).Should(BeTrue())
	})
})
