package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
    
    position := 0

    for i:=0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[position], nums[i] = nums[i], nums[position]
            position++
        }
    }

	fmt.Println(nums)
    
}

	

func main() {

	nums := []int{4, 1, 2, 0, 12}

	moveZeroes(nums)

}
