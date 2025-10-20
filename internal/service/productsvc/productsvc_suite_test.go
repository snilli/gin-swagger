package productsvc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductSvc Suite")
}
