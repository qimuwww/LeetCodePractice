package leetcode

import "container/list"

/*
	Given a binary tree, return the preorder traversal of its nodes' values.

	Example:

	Input: [1,null,2,3]
	   1
	    \
	     2
	    /
	   3

	Output: [1,2,3]
	Follow up: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代前序遍历
// 根据前序遍历的顺序，优先访问根结点，然后在访问左子树和右子树。
// 所以，对于任意结点node，第一部分即直接访问之，之后在判断左子树是否为空，
// 不为空时即重复上面的步骤，直到其为空。若为空，则需要访问右子树。
// 在访问过左孩子之后，需要反过来访问其右孩子，所以，需要栈这种数据结构的支持。
func preorderTraversalIteratively(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := list.New()
	stack.PushBack(root) // 也可以使用slice代替
	for stack.Len() > 0 {
		node := stack.Back()
		stack.Remove(node)
		ret = append(ret, (node.Value).(*TreeNode).Val)
		if (node.Value).(*TreeNode).Right != nil {
			stack.PushBack((node.Value).(*TreeNode).Right)
		}
		if (node.Value).(*TreeNode).Left != nil {
			stack.PushBack((node.Value).(*TreeNode).Left)
		}
	}
	return ret
}

// 递归前序遍历
func preorderTraversalRecursive(root *TreeNode) []int {
	ret := make([]int, 0)
	helper(root, &ret)
	return ret
}

// 前序遍历顺序: 根->左->右
func helper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	helper(root.Left, ret)
	helper(root.Right, ret)
}
