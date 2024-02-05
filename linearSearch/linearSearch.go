package main

import (
	"fmt"
)


func linearSearch(data []int, target int) *int {

	for idx, item := range data {
		if item == target {
			return &idx
		}
	}

	return nil
}

func main() {

	target := 7

	data := []int{1,2,7,5,7}

	if linearSearch(data, target) != nil {
		result := linearSearch(data, target)
		fmt.Println("Target found at index :",*result)
	}else {
		fmt.Println("The target not found in the list")
	}

}
