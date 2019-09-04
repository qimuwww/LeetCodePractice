package main

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

思路：
	两数相加 a+b=c
	c%10 获得该位的值
	c/10 获得进位的值

	结果保存在li中 免去新的链表的生成

	需要判断两个链表的长度是否相同
	l1=[0]
	l2=[3,7]
	||
	l1=[3,7]
	l2=[0]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var i int
	var tl1, tl2 = l1, l2
	for {
		temp := (tl1.Val + tl2.Val + i) % 10
		i = (tl1.Val + tl2.Val + i) / 10
		tl1.Val = temp
		if tl1.Next == nil && tl2.Next != nil {
			n := &ListNode{}
			tl1.Next = n
		} else if tl1.Next != nil && tl2.Next == nil {
			n := &ListNode{}
			tl2.Next = n
		} else if tl1.Next == nil && tl2.Next == nil {
			if i > 0 {
				n := &ListNode{
					Val: i,
				}
				tl1.Next = n
			}
			break
		}
		tl1 = tl1.Next
		tl2 = tl2.Next
	}
	return l1
}

/*
func main() {
	var l1, l2, l3, l4, l5, l6 ListNode
	l1.Val = 2
	l1.Next = &l2
	l2.Val = 4
	l2.Next = &l3
	l3.Val = 3
	l3.Next = nil

	l4.Val = 5
	l4.Next = &l5
	l5.Val = 6
	l5.Next = &l6
	l6.Val = 4
	l6.Next = nil

	ll := addTwoNumbers(&l1, &l4)
	for {
		fmt.Println(ll.Val)
		if ll.Next == nil {
			break
		}
		ll = ll.Next
	}

}
*/
