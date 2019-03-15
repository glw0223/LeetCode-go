package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question31_40"
	"testing"
)

func TestA33_search(t *testing.T) {
	arr:=[]int{4,5,6,7,0,1,2}
	result:=question31_40.A33_search(arr,0)
	fmt.Println(result)
	result=question31_40.A33_search(arr,3)
	fmt.Println(result)
	result=question31_40.A33_search([]int{1},1)
	fmt.Println(result)
	result=question31_40.A33_search([]int{1,3},0)
	fmt.Println(result)
	result=question31_40.A33_search([]int{8,1,2,3,4,5,6,7},6)
	fmt.Println(result)
}
