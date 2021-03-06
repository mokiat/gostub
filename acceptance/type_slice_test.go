package acceptance_test

import (
	. "github.com/mokiat/gostub/acceptance"
	"github.com/mokiat/gostub/acceptance/acceptance_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeSlice", func() {
	var stub *acceptance_stubs.SliceSupportStub

	BeforeEach(func() {
		stub = new(acceptance_stubs.SliceSupportStub)
	})

	It("stub is assignable to interface", func() {
		_, assignable := interface{}(stub).(SliceSupport)
		Ω(assignable).Should(BeTrue())
	})
})
