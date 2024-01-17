package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	var x int
	var y int

	for i, v := range nums {
		_, exists := seen[target - v]

		if !exists {
			seen[v]=i
		} else {
			x = i
			y = seen[target -v]
			break
		}
	}
	return []int{y, x}
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	target := 5

	result := twoSum(nums, target)

	fmt.Println(result)

}
