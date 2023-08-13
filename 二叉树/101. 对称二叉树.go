package main

import "container/list"

func isSymmetric(root *TreeNode) bool {
	var recursion func(left *TreeNode, right *TreeNode) bool
	recursion = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
	}
	return recursion(root.Left, root.Right)
}

func isSymmetric2(root *TreeNode) bool {
	st := list.New()
	st.PushBack(root.Left)
	st.PushBack(root.Right)
	for st.Len() > 0 {
		left := st.Remove(st.Front()).(*TreeNode)
		right := st.Remove(st.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		st.PushBack(left.Left)
		st.PushBack(right.Right)
		st.PushBack(left.Right)
		st.PushBack(right.Left)
	}
	return true
}
