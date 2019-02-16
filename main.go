package main

import (
	"fmt"
	"github.com/LeetCode-go/twoSum"
)

func main() {
	fmt.Println("Hello World")

	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := twoSum.Function(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}

	one, twe, err = twoSum.Function1(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}
