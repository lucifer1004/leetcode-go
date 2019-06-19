package restoreipaddresses

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for LeetCode 0093: Restore IP Addresses")
}

var _ = Describe("Possible IP addresses for", func() {
	Describe("'25525511135'", func() {
		s := "25525511135"
		It("should be 255.255.11.135 and 255.255.111.35", func() {
			Expect(restoreIPAddresses(s)).To(Equal([]string{"255.255.11.135", "255.255.111.35"}))
		})
	})

	Describe("'0000'", func() {
		s := "0000"
		It("should be 0.0.0.0", func() {
			Expect(restoreIPAddresses(s)).To(Equal([]string{"0.0.0.0"}))
		})
	})
})
