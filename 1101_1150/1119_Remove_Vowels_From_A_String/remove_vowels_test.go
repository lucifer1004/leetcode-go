package removevowels

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 1119: Remove Vowels From A String")
}

var _ = Describe("Processed result of", func() {
	Describe("book", func() {
		It("should be bk", func() {
			Expect(removeVowels("book")).To(Equal("bk"))
		})
	})

	Describe("zukdoor", func() {
		It("should be zkdr", func() {
			Expect(removeVowels("zukdoor")).To(Equal("zkdr"))
		})
	})
})
