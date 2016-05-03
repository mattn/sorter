package sorter_test

import (
	"fmt"
	"sort"

	"github.com/mattn/sorter"
)

func ExampleWrapper() {
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

func ExampleNewWrapperWith() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(sorter.NewWrapperWith(
		s,
		func(i, j int) bool {
			return s[i] < s[j]
		},
	))

	for _, v := range s {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleNewWrapper() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(sorter.NewWrapper(
		func() int {
			return len(s)
		},
		func(i, j int) bool {
			return s[i] < s[j]
		},
		func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	))

	for _, v := range s {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}
