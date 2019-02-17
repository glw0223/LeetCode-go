package twoSum

import (
	"errors"
	"fmt"
	"strconv"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
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
