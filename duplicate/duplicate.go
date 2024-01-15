package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {

	seen := make(map[int]bool)

	for _, k := range nums {
		if seen[k] {
			return true
		}

		seen[k] = true
	}

	return false

}

func main() {

	var nums = []int{2, 3, 4, 7}

	result := containsDuplicate(nums)

	if result {
		fmt.Println("There is a duplicate")
	} else {
		fmt.Println("There is no duplicate")
	}

}
