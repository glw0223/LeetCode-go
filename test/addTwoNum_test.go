package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/addTwoNum"
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	fmt.Println("TestaddTwoNumbers")
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}

	fmt.Println("输入1：", arr1, "输入2：", arr2)

	head1 := &addTwoNum.ListNode{
		Val:  -1,
		Next: nil,
	}
	head2 := &addTwoNum.ListNode{
		Val:  -1,
		Next: nil,
	}
	temp := head1
	for _, v := range arr1 {
		temp.Next = &addTwoNum.ListNode{v, nil}
		temp = temp.Next
	}
	temp = head2
	for _, v := range arr2 {
		temp.Next = &addTwoNum.ListNode{v, nil}
		temp = temp.Next
	}

	v := addTwoNum.AddTwoNumbers(head1, head2)
	var str string
	for v != nil {
		str = str + strconv.Itoa(v.Val)
		v = v.Next
	}
	fmt.Println("结果是:", str)
}
