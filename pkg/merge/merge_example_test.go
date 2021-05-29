package merge

import (
	"fmt"
)

// ExampleMerge 
func ExampleMerge() {
	fmt.Println(Merge(Ranges{{25, 30}, {2, 19}, {14, 23}, {4, 8}}))
	// Output: [[2 23] [25 30]]
}
