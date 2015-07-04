package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGostub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gostub Suite")
}
