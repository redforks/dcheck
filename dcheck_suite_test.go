package dcheck_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDcheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dcheck Suite")
}
