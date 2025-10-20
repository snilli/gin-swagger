package usersvc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserSvc Suite")
}
