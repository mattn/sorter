package sorter_test

import (
	"fmt"
	"sort"

	"github.com/mattn/sorter"
)

func ExampleSorter() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(&sorter.Wrapper{
		LenFunc: func() int {
			return len(s)
		},
		LessFunc: func(i, j int) bool {
			return s[i] < s[j]
		},
		SwapFunc: func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	})

	for _, v := range s {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}
