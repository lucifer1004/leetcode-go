package parseboolexpr

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 1106: Parsing A Boolean Expression")
}

var _ = Describe("Parsed result of", func() {
	Describe("!(f)", func() {
		It("should be true", func() {
			Expect(parseBoolExpr("!(f)")).To(Equal(true))
		})
	})

	Describe("|(&(t,f,t),!(t))", func() {
		It("should be false", func() {
			Expect(parseBoolExpr("|(&(t,f,t),!(t))")).To(Equal(false))
		})
	})
})
