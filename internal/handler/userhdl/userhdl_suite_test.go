package userhdl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserhdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Userhdl Suite")
}
