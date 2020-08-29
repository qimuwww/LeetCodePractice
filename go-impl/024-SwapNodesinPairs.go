package leetcode

/*
	Given a linked list, swap every two adjacent nodes and return its head.

	You may not modify the values in the list's nodes, only nodes itself may be changed.



	Example:

	Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

type ListNode struct {
	val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	next := head.Next.Next
	tmp := head
	prev := tmp
	tmp.Next.Next = tmp
	tmp.Next = nil // 防止这两个节点成环
	tmp = next
	for tmp != nil {
		if tmp.Next == nil {
			// 最后一次循环为一个节点
			prev.Next = tmp
			break
		}
		next = tmp.Next.Next
		prev.Next = tmp.Next
		tmp.Next.Next = tmp
		tmp.Next = nil // 防止这两个节点成环
		prev = tmp
		tmp = next
	}
	return newHead
}

/*
func main() {
	n6 := &ListNode{
		Val:  6,
		Next: nil,
	}
	n5 := &ListNode{
		Val:  5,
		Next: n6,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	head := n1
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	head = swapPairs(n1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
*/
