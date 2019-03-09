package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/binarySearch"
	"testing"
)

func TestBinarySearch(t *testing.T)  {
	arr:=[]int{1,2,3,4,7,8,9}
	fmt.Println("------------------非递归测试------------------")
	key:=8
	result:=binarySearch.BinarySearch1(arr, key)
	fmt.Println(1,result)
	key=6
	result=binarySearch.BinarySearch1(arr, key)
	fmt.Println(2,result)
	key=10
	result=binarySearch.BinarySearch1(arr, key)
	fmt.Println(3,result)
	key=0
	result=binarySearch.BinarySearch1(arr, key)
	fmt.Println(4,result)

	fmt.Println("------------------递归测试------------------")
	key=8
	result=binarySearch.BinarySearch2(arr, 0, len(arr),key)
	fmt.Println(1,result)
	key=6
	result=binarySearch.BinarySearch2(arr, 0, len(arr),key)
	fmt.Println(2,result)
	key=10
	result=binarySearch.BinarySearch2(arr, 0, len(arr),key)
	fmt.Println(3,result)
	key=0
	result=binarySearch.BinarySearch2(arr, 0, len(arr),key)
	fmt.Println(4,result)
}
