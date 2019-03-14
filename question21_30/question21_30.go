package question21_30

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func A21_mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1==nil{
        return l2
    }
    if l2==nil {
        return l1
    }
    head:=ListNode{0,nil}
    p:=&ListNode{0,nil}
    head.Next=p
    l1value:=l1.Val;l2value:=l2.Val
    for {
        if l1value<l2value{
            p.Val=l1value
            p.Next=&ListNode{0,nil}
            p=p.Next
            if l1!=nil{
                l1=l1.Next
            }
        }else {
            p.Val=l2value
            p.Next=&ListNode{0,nil}
            p=p.Next
            if l2!=nil{
                l2=l2.Next
            }
        }
        if l1!=nil{
            l1value=l1.Val
        }else {
            for l2!=nil{
                p.Val=l2value
                l2=l2.Next
                if l2!=nil{
                    p.Next=&ListNode{0,nil}
                    p=p.Next
                }
            }
        }
        if l2!=nil{
            l2value=l2.Val
        }else {
            for l1!=nil{
                p.Val=l1value
                l1=l1.Next
                if l1!=nil{
                    p.Next=&ListNode{0,nil}
                    p=p.Next
                }
            }
        }
        if l1==nil&&l2==nil{
            break
        }
    }
    return head.Next
}
/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */
func A22_generateParenthesis(n int) (result []string) {
    return
}

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func A23_mergeKLists(lists []*ListNode) (result *ListNode) {
    return
}