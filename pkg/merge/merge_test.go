package merge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/alex-held/daimler-merge/pkg/merge"
)

var _ = Describe("merge intervals", func() {
	Context("merge.Merge()", func() {
		When("input is empty", func() {
			It("returns empty interval", func() {
				Expect(Merge([][]int{})).Should(BeEmpty())
			})
		})
		When("input contains `negative` mergable intervals", func() {
			It("returns intervals with merged intervals", func() {
				Expect(Merge([][]int{{-1, 5}, {4, 10}, {-20, -10}})).Should(Equal([][]int{{-20, -10}, {-1, 10}}))
			})
		})
		When("input contains mergable intervals", func() {
			It("returns intervals with merged intervals", func() {
				Expect(Merge([][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}})).Should(Equal([][]int{{2, 23}, {25, 30}}))
			})
		})
		When("input contains no mergable intervals", func() {
			It("returns input without merged intervals", func() {
				Expect(Merge([][]int{{1, 3}, {4, 6}, {7, 9}})).Should(Equal([][]int{{1, 3}, {4, 6}, {7, 9}}))
			})
		})
	})
})
