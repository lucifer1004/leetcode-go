package ismatch

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 0010: Regular Expression Matching")
}

var _ = Describe("", func() {
	Describe("'' & 'a******'", func() {
		s := ""
		t := "a******"
		It("match", func() {
			Expect(isMatch(s, t)).To(Equal(true))
		})
	})

	Describe("'abc' & 'c*abc'", func() {
		s := "abc"
		t := "c*abc"
		It("match", func() {
			Expect(isMatch(s, t)).To(Equal(true))
		})
	})

	Describe("'abcdef' & 'abcde.f'", func() {
		s := "abcdef"
		t := "abcde.f"
		It("do not match", func() {
			Expect(isMatch(s, t)).To(Equal(false))
		})
	})
})
