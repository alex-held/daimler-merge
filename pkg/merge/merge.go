package merge

import (
	"sort"
)

// Range is an interval with a start integer and a end integer.
// It `always` consists of exactly 2 integers.
type Range [2]int

// Ranges is an interval of Range
type Ranges []Range

func (s Ranges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Ranges) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s Ranges) Len() int {
	return len(s)
}

// Merge returns Ranges where overlapping Range values have been merged
func Merge(ranges Ranges) (result Ranges) {
	// sort the ranges by the starting value
	sort.Sort(ranges)

	// early return if the input is empty
	if len(ranges) == 0 {
		return ranges
	}

	return result
}
