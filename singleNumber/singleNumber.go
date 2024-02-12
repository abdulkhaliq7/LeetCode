package main

import (
	"fmt"
)

func singleNumber(nums []int) int {

	seen := make(map[int]int)

	number := 0

	for i := 0; i < len(nums); i++ {

		seen[nums[i]]++
	}

	for key, value := range seen {
		if value == 1 {
			number = key
		}
	}

	return number

}

func main() {

	nums := []int{1, 2, 1, 2, 4}

	fmt.Println(singleNumber(nums))

}
