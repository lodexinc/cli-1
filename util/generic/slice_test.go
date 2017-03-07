package generic_test

import (
	"code.cloudfoundry.org/cli/util/generic"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slice", func() {
	Describe("IsSliceable", func() {
		It("returns false if the type is nil", func() {
			Expect(generic.IsSliceable(nil)).To(BeFalse())
		})

		It("should return false when the type cannot be sliced", func() {
			Expect(generic.IsSliceable("bad slicing")).To(BeFalse())
		})

		It("should return true if the type can be sliced", func() {
			Expect(generic.IsSliceable([]string{"a string"})).To(BeTrue())
		})

		It("should return true if the type can be sliced", func() {
			Expect(generic.IsSliceable([]interface{}{1, 2, 3})).To(BeTrue())
		})
	})

	Describe("SortAndUniquifyStringSlice", func() {
		It("returns an empty slice if the provided slice is empty", func() {
			Expect(generic.SortAndUniquifyStringSlice([]string{})).To(BeEmpty())
		})

		It("returns a slice of one element when provided a single-element slice", func() {
			inputSlice := []string{"y", "y", "y", "y", "y", "y", "y", "y", "y", "y", "y"}
			Expect(generic.SortAndUniquifyStringSlice(inputSlice)).To(Equal([]string{"y"}))
		})

		It("returns a slice of one element when provided an all-duplicate slice", func() {
			Expect(generic.SortAndUniquifyStringSlice([]string{"y"})).To(Equal([]string{"y"}))
		})

		It("returns a sorted slice", func() {
			inputSlice := []string{"y", "a", "c", "z", "d"}
			expectedSlice := []string{"a", "c", "d", "y", "z"}

			Expect(generic.SortAndUniquifyStringSlice(inputSlice)).To(Equal(expectedSlice))
		})

		It("returns a sorted slice of unique values", func() {
			inputSlice := []string{"y", "a", "c", "z", "d", "c", "c", "c", "c", "c", "a", "z", "d", "c"}
			expectedSlice := []string{"a", "c", "d", "y", "z"}

			Expect(generic.SortAndUniquifyStringSlice(inputSlice)).To(Equal(expectedSlice))
		})
	})
})
