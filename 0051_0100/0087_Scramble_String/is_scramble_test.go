package isscramble

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 0087: Scramble String")
}

var _ = Describe("", func() {
	// abab => a-bab => a-(b-ab) => a-(ab-b)
	Describe("'abab' & 'aabb'", func() {
		s := "abab"
		t := "aabb"
		It("are scramble strings", func() {
			Expect(isScramble(s, t)).To(Equal(true))
		})
	})

	// great => gr-eat => (g-r)-eat => (r-g)eat
	Describe("'great' & 'rgeat'", func() {
		s := "great"
		t := "rgeat"
		It("are scramble strings", func() {
			Expect(isScramble(s, t)).To(Equal(true))
		})
	})

	Describe("'great' & 'rtgea'", func() {
		s := "great"
		t := "rtgea"
		It("are not scramble strings", func() {
			Expect(isScramble(s, t)).To(Equal(false))
		})
	})
})
