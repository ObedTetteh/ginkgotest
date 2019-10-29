package adder

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ObedTetteh/ginkgotest/adder"
)

var _ = Describe("Adder", func() {
	Describe("Add", func() {

		It("adds two numbers", func() {
			sum := Add(2, 3)
			Expect(sum).To(Equal(5))
		})
	})
})
