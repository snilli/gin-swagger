package usersvc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsersvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usersvc Suite")
}
