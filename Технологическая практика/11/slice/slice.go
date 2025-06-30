package main

import (
	"fmt"
	"sort"
)

func find(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func Sort(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func filter(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 8, 5, 2, 7, 4}

	fmt.Println(find(nums, 8))

	sorted := Sort(nums)
	fmt.Println(sorted)

	evens := filter(nums)
	fmt.Println(evens)
}
