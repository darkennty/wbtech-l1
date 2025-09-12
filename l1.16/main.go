package main

import "fmt"

func main() {
	arr := []int{3, 5, 2, -9, 11, 44, 1, -3, 5, 4, 19, -110}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := make([]int, 0)
	middle := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result := append(left, middle...)
	result = append(result, right...)
	return result
}
