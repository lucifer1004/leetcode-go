package numenclaves

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 1020: Number of Enclaves")
}

var _ = Describe("Example", func() {
	Describe("One", func() {
		square := [][]int{[]int{0, 0, 0, 0}, []int{1, 0, 1, 0}, []int{0, 1, 1, 0}, []int{0, 0, 0, 0}}
		It("has 3 enclaves", func() {
			Expect(numEnclaves(square)).To(Equal(3))
		})
	})

	Describe("Two", func() {
		square := [][]int{[]int{0, 1, 1, 0}, []int{0, 0, 1, 0}, []int{0, 0, 1, 0}, []int{0, 0, 0, 0}}
		It("has no enclaves", func() {
			Expect(numEnclaves(square)).To(Equal(0))
		})
	})
})
