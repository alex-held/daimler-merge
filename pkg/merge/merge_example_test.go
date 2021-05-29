package merge_test

import (
	"fmt"

	. "github.com/alex-held/daimler-merge/pkg/merge"
)

// ExampleMerge
func ExampleMerge() {
	fmt.Println(Merge(Ranges{{25, 30}, {2, 19}, {14, 23}, {4, 8}}))
	// Output: [[2 23] [25 30]]
}
