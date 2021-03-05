package main

import (
	"fmt"
)

type Xi []int
type Xs []string
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (self Xi) Len() int               { return len(self) }
func (self Xi) Less(i int, j int) bool { return self[i] < self[j] }
func (self Xi) Swap(i int, j int)      { self[i], self[j] = self[j], self[i] }

func Sort(s Sorter) {
	for i := 0; i < s.Len()-1; i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(i, j) {
				s.Swap(i, j)
			}
		}
	}
}

func main() {
	s := Xi{1, 4, 7, 2, 9, 3, 8, 5}
	Sort(s)
	fmt.Printf("排序之后的s：%v\n", s)
}
