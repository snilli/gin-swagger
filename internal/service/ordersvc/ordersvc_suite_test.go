package ordersvc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOrderSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderSvc Suite")
}
