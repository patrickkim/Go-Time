package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGosample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gosample Suite")
}

var _ = Describe("Animal", func() {

	BeforeEach(func() {
		spot := Animal{
			species: "dog",
			name:    "Spot",
			age:     3,
		}
	})

	Describe("Animal", func() {
		Context("Canine", func() {
			It("should be a dog", func() {
				Expect(spot.isType()).To(Equal("dog"))
			})
		})
	})
})
