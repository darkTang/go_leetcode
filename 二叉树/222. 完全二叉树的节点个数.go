package main

import "container/list"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	num := 0
	for st.Len() > 0 {
			curNode = st.Remove(st.Front()).(*TreeNode)
			num++
			if curNode.Left != nil {
				st.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				st.PushBack(curNode.Right)
			}
	}
	return num
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}