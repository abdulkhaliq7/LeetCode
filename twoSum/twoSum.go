package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	for i, _ := range nums {
		for j, _ := range nums {
			if nums[i] != nums[j] {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			} else {
				continue
			}
		}
	}

	return []int{}
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	target := 6

	result := twoSum(nums, target)

	fmt.Println(result)

}
