package merge

import (
	"math"
	"sort"
)

// Range is an interval with a start integer and a end integer.
// It `always` consists of exactly 2 integers.
type Range [2]int

// Merge returns a new Range with the minimal start and end values
func (a Range) Merge(b Range) Range {
	return Range{int(math.Min(float64(a[0]), float64(b[0]))), int(math.Max(float64(a[1]), float64(b[1])))}
}

// Overlaps returns a boolean indicating whether the ranges overlap
func (a Range) Overlaps(b Range) bool {
	if a[0] > b[0] {
		a, b = b, a
	}

	return a[1] >= b[0]
}

// Ranges is an interval of Range
type Ranges []Range

// Merge returns Ranges where overlapping Range values have been merged
func Merge(ranges Ranges) (result Ranges) {
	// sort the ranges by the starting value
	sort.Sort(ByRangeStart(ranges))

	// early return if the input is empty
	if len(ranges) == 0 {
		return ranges
	}

	// seed the first compare with the smallest range
	result = append(result, ranges[0])

	for _, current := range ranges {
		// get the last range
		last := result[len(result)-1]

		// check if the current range overlaps with the last
		if !last.Overlaps(current) {
			// no overlap means current starts a new (result) range
			result = append(result, current)

			continue
		}

		// overlap means we can merge current with last
		result[len(result)-1] = last.Merge(current)
	}

	return result
}
