package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/twoSum"
	"testing"
)

func TestFunction(t *testing.T) {
	fmt.Println("TestFunction")
	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := twoSum.Function(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}

func TestFunction1(t *testing.T) {
	fmt.Println("TestFunction1")
	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := twoSum.Function1(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}
