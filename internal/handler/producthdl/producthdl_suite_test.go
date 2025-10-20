package producthdl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductHdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductHdl Suite")
}
