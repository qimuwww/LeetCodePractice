package SList

import "errors"

var (
	errIndexOutOfRange = errors.New("index out of range")
)

type (
	SLLNode struct {
		Data interface{}
		//Prev *Node
		Next *SLLNode
	}

	SList struct {
		len  int
		head *SLLNode
		tail *SLLNode
	}
)

func New() *SList {
	list := SList{}
	return &list
}

func (sl *SList) Len() int {
	return sl.len
}

func (sl *SList) InsertAt(index int, data interface{}) error {
	if index < 0 || index > sl.len {
		return errIndexOutOfRange
	}
	newN := &SLLNode{Data: data}
	if sl.len == 0 {
		sl.head = newN
		sl.tail = newN
		sl.len++
		newN = nil
		return nil
	}
	if index == 0 {
		newN.Next = sl.head
		sl.head = newN
		sl.len++
		newN = nil
		return nil
	}
	if index == sl.len {
		sl.tail.Next = newN
		sl.tail = newN
		sl.len++
		newN = nil
		return nil
	}
	current := sl.getNode(index - 1)
	if current == nil {
		return errIndexOutOfRange
	}
	newN.Next = current.Next
	current.Next = newN
	sl.len++
	current = nil
	newN = nil
	return nil
}

func (sl *SList) getNode(index int) *SLLNode {
	if index < 0 || index > sl.len {
		return nil
	}
	current := sl.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}

func (sl *SList) GetAllNodes() []interface{} {
	data := make([]interface{}, 0, sl.len)
	current := sl.head
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	return data
}

func (sl *SList) Reserve() {
	current := sl.head
	next := sl.head.Next
	// 先处理头节点
	sl.head = current
	sl.tail = current
	current.Next = nil
	// 将头节点之后的所有节点全部前项插入
	for next != nil {
		current = next
		next = next.Next
		current.Next = sl.head
		sl.head = current
	}
}

//Input: 1->2->3->4->5->NULL, k = 2
//Output: 4->5->1->2->3->NULL
func (sl *SList) Rotate(index int) {
	index = index % sl.len
	if index < 1 {
		return
	}
	prev := sl.head
	current := sl.head
	for i := 0; i < sl.len-index; i++ {
		prev = current
		current = current.Next
	}
	prev.Next = nil
	sl.tail.Next = sl.head
	sl.head = current
	sl.tail = prev

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var index int
	tail := head
	for {
		index++
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	k = k % index
	if k < 1 {
		return head
	}
	prev := head
	current := head
	for i := 0; i < index-k; i++ {
		prev = current
		current = current.Next
	}
	prev.Next = nil
	tail.Next = head
	return current
}
