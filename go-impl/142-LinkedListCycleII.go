package LinkedListCycleII

/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.


Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.


Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.

*/

/*
				O<--O<--O<--O -->cross
				|			^ |
				|	R		| |
O-->O-->O-->O-->O-->O-->O-->O |
|------L--------|-------------M

假设快慢指针在cross处相遇,则慢指针走过的长度为:L+M,快指针走过的长度为:L+M+nR
由于快指针的速度是慢指针速度的2倍,则:
	2(L+M) = L+M+nR 即:
	L+M = nR
	说明当两指针相遇时快指针已经在环内循环了nR+M步,回退M步就是环的入口,也就是说快指针从cross处再走
	R-M步就到达环的入口,相当于慢指针从起点走L步.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			//break
			slow = head
			for {
				slow = slow.Next
				fast = fast.Next
				if slow == fast {
					return slow
				}
			}
		}
	}
	return nil
}
