package util_test

import (
	"os"

	. "github.com/momchil-atanasov/gostub/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gopath", func() {
	var testDir string

	// Note: Ginkgo sets PWD to the test directory, which allows us
	// to do the following tests that way. If GOPATH has symbolic links,
	// they might fail

	BeforeEach(func() {
		var err error
		testDir, err = os.Getwd()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("is possible to get directory from import", func() {
		Ω(ImportToDir("github.com/momchil-atanasov/gostub/util")).Should(Equal(testDir))
	})

	It("is possible to get import from directory", func() {
		Ω(DirToImport(testDir)).Should(Equal("github.com/momchil-atanasov/gostub/util"))
	})
})
