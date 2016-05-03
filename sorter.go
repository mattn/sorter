package sorter

import (
	"sort"
)

type Wrapper struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (w *Wrapper) Len() int {
	return w.LenFunc()
}

func (w *Wrapper) Less(i, j int) bool {
	return w.LessFunc(i, j)
}

func (w *Wrapper) Swap(i, j int) {
	w.SwapFunc(i, j)
}
