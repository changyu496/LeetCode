package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2020-4-1
// 思路：
// 两个指针分别指向链表的头节点，只要有一个不为nil，就持续相加，将结果存在结果链表里，要注意进位的问题
// 如何利用头节点构建一个链表，这是需要多训练
// 熟练理解和使用链表，是基础，要多联系
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	cur := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10, Next: nil}
		cur = cur.Next
	}
	return dummyHead.Next
}
