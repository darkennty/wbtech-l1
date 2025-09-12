package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	union := make([]int, 0)

	if len(a) <= len(b) {
		unite(a, b, &union)
	} else {
		unite(b, a, &union)
	}

	fmt.Println("Union: ", union)
}

func unite(first, second []int, union *[]int) {
	firstSliceMap := make(map[int]struct{})

	for _, v := range first {
		firstSliceMap[v] = struct{}{}
	}

	for _, v := range second {
		if _, ok := firstSliceMap[v]; ok {
			*union = append(*union, v)
		}
	}
}
