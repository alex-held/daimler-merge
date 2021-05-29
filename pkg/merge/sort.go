package merge

// ByRangeStart implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Ranges value.
type ByRangeStart Ranges

// Len is part of sort.Interface.
func (s ByRangeStart) Len() int { return len(s) }

// Less is part of sort.Interface.
// It compares the start value of both Range
func (s ByRangeStart) Less(i, j int) bool { return s[i][0] < s[j][0] }

// Swap is part of sort.Interface.
func (s ByRangeStart) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
