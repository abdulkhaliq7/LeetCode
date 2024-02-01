package main

import "fmt"

func majorityElement(nums []int) int {

    seen := make(map[int]int)

    var majority int

    for _, value := range nums {
        seen[value]++
    }

    for key, count := range seen {
        if count > len(nums)/2 {
            majority = key
        }
    }

    return majority
    
}


func main () {

	nums := []int {3,2,3}

	fmt.Println(majorityElement(nums))
}