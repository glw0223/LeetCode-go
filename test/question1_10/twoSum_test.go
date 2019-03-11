package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestFunction(t *testing.T) {
	fmt.Println("TestFunction")
	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := question1_10.TwoNumFunction(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}

func TestFunction1(t *testing.T) {
	fmt.Println("TestFunction1")
	nums := []int{1, 2, 4, 5, 9}
	one, twe, err := question1_10.TwoNumFunction1(nums, 9)
	if err == nil {
		fmt.Println(one, twe)
	} else {
		fmt.Println(err.Error())
	}
}
