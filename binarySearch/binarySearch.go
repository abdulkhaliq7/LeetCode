package main

import (
	"fmt"
)

func binarySearch(data []int, target int) *int {
	first := 0
	last := len(data) - 1

	for first <= last {
		midpoint := (first + last) / 2

		if data[midpoint] == target {
			return &midpoint
		} else if data[midpoint] < target {
			first = midpoint + 1
		} else if data[midpoint] > target {
			last = midpoint - 1
		}
	}
	return nil
}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 10

	if binarySearch(data, target) != nil {
		result := binarySearch(data, target)
		fmt.Println("Targt found at position :", *result)
	} else {
		fmt.Println("Target not Found in the List")
	}
}
