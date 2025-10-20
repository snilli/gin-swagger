package userrepo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserRepo Suite")
}
