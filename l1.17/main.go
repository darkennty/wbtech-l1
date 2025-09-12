package main

import "fmt"

func main() {
	slice := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	elem := 9
	fmt.Println(binarySearch(slice, elem))
}

func binarySearch(slice []int, elem int) int {
	if len(slice) == 0 {
		return -1
	}
	middle := len(slice) / 2

	if slice[middle] == elem {
		return middle
	} else if elem < slice[middle] {
		index := binarySearch(slice[:middle], elem)
		if index == -1 {
			return -1
		}
		return index
	} else {
		index := binarySearch(slice[middle+1:], elem)
		if index == -1 {
			return -1
		}
		return middle + index + 1
	}
}
