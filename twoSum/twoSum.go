package twoSum

import (
	"errors"
	"fmt"
	"strconv"
)

func Function(nums []int, target int) (one, two int, err error) {
	fmt.Println(nums)
	err = errors.New("not found")
	mymap := make(map[int]string)

	for i, v := range nums {
		mymap[v] = fmt.Sprintf("%d", i)
	}

	for i, v := range nums {
		temp := target - v
		if mymap[temp] != "" {
			one = i
			two, _ = strconv.Atoi(mymap[temp])
			err = nil
			break
		}
	}
	return
}
func Function1(nums []int, target int) (one, two int, err error) {
	fmt.Println(nums)
	err = errors.New("not found")
	mymap := make(map[int]string)

	for i, v := range nums {
		temp := target - v
		if mymap[temp] != "" {
			one = i
			two, _ = strconv.Atoi(mymap[temp])
			err = nil
			break
		}
		mymap[v] = fmt.Sprintf("%d", i)
	}
	return
}
