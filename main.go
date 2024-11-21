package main

import (
	"fmt"
	s "skiplist/skiplist"
	"slices"
)

func main() {
	list, _ := s.NewSkipList[float32](8)

	list.Insert(3, 12.7)
	list.Insert(6, 44.4)
	list.Insert(9, 723.9)

	list.Insert(1, 2.2)
	list.Insert(4, 999.99)
	list.Delete(3)
	list.Insert(3, 32.5)

	list.Insert(20, 0.78)
	fmt.Println(slices.Collect(list.Values()))
}
