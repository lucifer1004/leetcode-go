package minwindow

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 0076: Minimum Window Substring")
}

var _ = Describe("Minimum Window Substring for", func() {
	Describe("'ab' & ''", func() {
		s := "ab"
		t := ""
		It("should be ''", func() {
			Expect(minWindow(s, t)).To(Equal(""))
		})
	})

	Describe("'ab' & 'b'", func() {
		s := "ab"
		t := "b"
		It("should be 'b'", func() {
			Expect(minWindow(s, t)).To(Equal("b"))
		})
	})

	Describe("'abc' & 'bdfe'", func() {
		s := "abc"
		t := "bdfe"
		It("should be ''", func() {
			Expect(minWindow(s, t)).To(Equal(""))
		})
	})

	Describe("'ADOBECODEBANC' & 'ABC'", func() {
		s := "ADOBECODEBANC"
		t := "ABC"
		It("should be 'BANC'", func() {
			Expect(minWindow(s, t)).To(Equal("BANC"))
		})
	})
})
