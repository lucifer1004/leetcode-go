package dividearray

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 1121: Divide Array Into Increasing Sequences")
}

var _ = Describe("Array", func() {
	Describe("[1,2,2,3,3,4,4]", func() {
		It("can be divided into sequences with length >= 3", func() {
			Expect(canDivideIntoSubsequences([]int{1, 2, 2, 3, 3, 4, 4}, 3)).To(Equal(true))
		})
	})

	Describe("[5,6,6,7,8]", func() {
		It("cannot be divided into sequences with length >= 3", func() {
			Expect(canDivideIntoSubsequences([]int{5, 6, 6, 7, 8}, 3)).To(Equal(false))
		})
	})
})
