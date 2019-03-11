package main

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
)

func main() {
	fmt.Println("Hello World")

	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := question1_10.TwoNumFunction(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}

	one, twe, err = question1_10.TwoNumFunction1(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}
