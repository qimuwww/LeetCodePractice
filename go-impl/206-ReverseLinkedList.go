package ReverseLinkedList

/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// iteratively
// 循环读取链表上的每个节点
// 每次将取到的节点的next指向rh,并将rh更新指到该节点
// rh为指向已经逆序后的链表头
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	var rh *ListNode
	for tmp != nil {
		node := tmp
		tmp = tmp.Next
		node.Next = rh
		rh = node
	}
	return rh
}
