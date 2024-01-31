package main

import "fmt"



func missingNumber(nums []int) int {

    length := len(nums)

    expectedSum := (length*(length + 1))/2
    calculatedSum := 0

    for i:=0; i < length; i++ {
        calculatedSum +=nums[i]
    }

    return expectedSum - calculatedSum
    
}


func main() {

	nums := []int{3, 0, 1}

	fmt.Println(missingNumber(nums))
}