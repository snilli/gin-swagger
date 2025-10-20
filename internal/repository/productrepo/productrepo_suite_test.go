package productrepo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductRepo Suite")
}
