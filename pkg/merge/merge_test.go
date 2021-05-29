package merge_test

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/alex-held/daimler-merge/pkg/merge"
)

var _ = Describe("merge", func() {
	Describe("Ranges", func() {
		Context("merge range with another range", func() {
			It("returns merged range with min start and max end values", func() {
				Expect(Range{1, 4}.Merge(Range{3, 6})).Should(Equal(Range{1, 6}))
			})
		})

		Describe("check whether two ranges overlap", func() {
			DescribeTable("merge range intervals",
				func(a, b Range, expected bool) {
					Expect(a.Overlaps(b)).Should(Equal(expected))
					Expect(b.Overlaps(a)).Should(Equal(expected))
				},
				Entry("{1,5} {4,10} overlap", Range{1, 5}, Range{4, 10}, true),
				Entry("{1,1} {1,4} overlap", Range{1, 1}, Range{1, 4}, true),
				Entry("{1,1} {1,1} overlap", Range{1, 1}, Range{1, 1}, true),
				Entry("{4,10} {1,5} overlap", Range{4, 10}, Range{1, 5}, true),
				Entry("{1,5} {6,10} do not overlap", Range{1, 5}, Range{6, 10}, false))
		})

		DescribeTable(
			"sorting merge.Ranges by ascending range start values",
			func(in, expected Ranges) {
				sort.Sort(in)
				Expect(in).Should(Equal(expected))
			},
			Entry("most negative left", Ranges{{-1, 5}, {-20, -10}}, Ranges{{-20, -10}, {-1, 5}}),
			Entry("lesser left", Ranges{{1, 1}, {0, 1}}, Ranges{{0, 1}, {1, 1}}),
			Entry("already ordered", Ranges{{1, 5}, {8, 10}}, Ranges{{1, 5}, {8, 10}}),
			Entry("order overlap", Ranges{{3, 6}, {1, 5}}, Ranges{{1, 5}, {3, 6}}),
		)
	})

	Describe("merge.Merge(range)", func() {
		When("input is empty", func() {
			It("returns empty interval", func() {
				Expect(Merge(Ranges{})).Should(BeEmpty())
			})
		})
		When("input contains `negative` mergable intervals", func() {
			It("returns intervals with merged intervals", func() {
				Expect(Merge(Ranges{{-1, 5}, {4, 10}, {-20, -10}})).Should(Equal(Ranges{{-20, -10}, {-1, 10}}))
			})
		})
		When("input contains mergable intervals", func() {
			It("returns intervals with merged intervals", func() {
				Expect(Merge(Ranges{{25, 30}, {2, 19}, {14, 23}, {4, 8}})).Should(Equal(Ranges{{2, 23}, {25, 30}}))
			})
		})
		When("input contains no mergable intervals", func() {
			It("returns input without merged intervals", func() {
				Expect(Merge(Ranges{{1, 3}, {4, 6}, {7, 9}})).Should(Equal(Ranges{{1, 3}, {4, 6}, {7, 9}}))
			})
		})
	})
})
