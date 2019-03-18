package main

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
)

func main() {
	fmt.Println("Hello World")

	myMap:=make(map[int]int, 8)
	myMap[20]=10
	i,ok:=myMap[20]
	fmt.Println(i,ok)
	i,ok=myMap[100]
	fmt.Println(i,ok)
	if _, ok := myMap[20]; ok {
		//yes
	}else {
		//no
	}


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
