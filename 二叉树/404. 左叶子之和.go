package main

import (
	"container/list"
)

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil{
			sum += node.Left.Val
		}
		helper(node.Left)
		helper(node.Right)
	}
  helper(root)
	return sum
}

func sumOfLeftLeaves2(root *TreeNode) int {
	sum := 0
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	for st.Len() > 0 {
		curNode = st.Remove(st.Front()).(*TreeNode)
		if curNode.Left != nil {
			st.PushBack(curNode.Left)
		}
		if curNode.Right != nil {
			st.PushBack(curNode.Right)
		}
		if curNode.Left != nil && curNode.Left.Left == nil && curNode.Left.Right == nil {
			sum += curNode.Left.Val
		}
	}
	return sum
}