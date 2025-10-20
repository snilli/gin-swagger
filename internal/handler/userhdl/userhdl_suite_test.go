package userhdl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserHdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserHdl Suite")
}
