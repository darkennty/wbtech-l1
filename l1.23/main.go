package main

import "fmt"

func main() {
	i := 2

	// Удаление с помощью append
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before deleting with append: %#v\n", nums)
	fmt.Printf("Length: %d Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums[:i], nums[i+1:]...) // слайс от 0 до i невключительно + слайс от i+1 до конца
	fmt.Printf("After deleting with append:  %#v\n", nums)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(nums), cap(nums))

	// Удаление с помощью copy
	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("Before deleting with copy: %#v\n", nums)
	fmt.Printf("Length: %d Capacity: %d\n", len(nums), cap(nums))

	nums = nums[:i+copy(nums[i:], nums[i+1:])] // слайс от 0 до i + количество скопированных элементов
	fmt.Printf("After deleting with copy:  %#v\n", nums)
	fmt.Printf("Length: %d Capacity: %d\n", len(nums), cap(nums))
}

/* Results:
Before deleting with append: []int{1, 2, 3, 4, 5}
Length: 5 Capacity: 5
After deleting with append:  []int{1, 2, 4, 5}
Length: 4 Capacity: 5

Before deleting with copy: []int{1, 2, 3, 4, 5}
Length: 5 Capacity: 5
After deleting with copy:  []int{1, 2, 4, 5}
Length: 4 Capacity: 5
*/
