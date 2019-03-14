package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question21_30"
	"testing"
)

func TestA21_mergeTwoLists(t *testing.T) {
	l1:=question21_30.ListNode{1,nil}
	l1.Next=&question21_30.ListNode{2,nil}
	l1.Next.Next=&question21_30.ListNode{4,nil}

	l2:=question21_30.ListNode{1,nil}
	l2.Next=&question21_30.ListNode{3,nil}
	l2.Next.Next=&question21_30.ListNode{4,nil}
	result:=question21_30.A21_mergeTwoLists(&l1,&l2)
	var str string
	for result!=nil{
		str+=fmt.Sprintf("%d",result.Val)
		result=result.Next
	}
	fmt.Println(str)
}

func TestA24_swapPairs(t *testing.T) {
	head:=question21_30.ListNode{1,nil}
	head.Next=&question21_30.ListNode{2,nil}
	head.Next.Next=&question21_30.ListNode{3,nil}
	head.Next.Next.Next=&question21_30.ListNode{4,nil}
	result:=question21_30.A24_swapPairs(&head)
	str:="开始"
	for result!=nil{
		str+=fmt.Sprintf("->%d",result.Val)
		result=result.Next
	}
	fmt.Println(str)
}