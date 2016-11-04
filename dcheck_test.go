package dcheck_test

import (
	"errors"
	. "github.com/redforks/dcheck"

	"github.com/redforks/life"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redforks/hal"
	"github.com/redforks/testing/reset"
)

var _ = Describe("Dcheck", func() {
	var exitCode int

	BeforeEach(func() {
		reset.Enable()
		exitCode = 0
		hal.Exit = func(n int) {
			exitCode += n
		}
	})

	AfterEach(func() {
		reset.Disable()
	})

	It("Check succeed", func() {
		var called1, called2, called3 int
		Add(func() error {
			called1++
			return nil
		})
		Add(func() error {
			called2++
			return nil
		})
		Add(func() error {
			called3++
			return nil
		})

		life.Start()
		Ω(called1).Should(Equal(1))
		Ω(called2).Should(Equal(1))
		Ω(called3).Should(Equal(1))
	})

	It("Check failed", func() {
		Add(func() error {
			return nil
		})
		Add(func() error {
			return errors.New("foo")
		})
		life.Start()
		Ω(exitCode).Should(Equal(13))
	})
})
