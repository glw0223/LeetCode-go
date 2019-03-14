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

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func A24_swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next==nil {
        return head
    }
    p := head
    head = head.Next
    p.Next = head.Next
    head.Next = p
    p.Next = A24_swapPairs(p.Next)
    return head
}

/*
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func A25_reverseKGroup(head *ListNode, k int) *ListNode {
    return head
}
/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
 */
func A26_removeDuplicates(nums []int) int {
    if len(nums)==0 {
        return 0
    }
    i:=0
    for j:=1; j<len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    return i + 1
}
/*
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。
示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。
 */
func A27_removeElement(nums []int, val int) int {
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}