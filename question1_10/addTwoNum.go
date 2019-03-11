package question1_10

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (list *ListNode) {
	l1 = l1.Next
	l2 = l2.Next

	list = &ListNode{
		-1,
		nil,
	}
	p := list

	carry := 0
	for l1 != nil || l2 != nil {

		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		v3 := v1 + v2
		v4 := v3 % 10

		t := &ListNode{
			carry + v4,
			nil,
		}
		p.Next = t
		p = p.Next

		carry = v3 / 10

		l1 = l1.Next
		l2 = l2.Next
	}

	list = list.Next
	return
}
