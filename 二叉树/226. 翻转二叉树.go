package main

import "container/list"

// 递归遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 迭代遍历
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	st := list.New()
	var curNode *TreeNode = root
	for curNode != nil || st.Len() > 0 {
		for curNode != nil {
			curNode.Left, curNode.Right = curNode.Right, curNode.Left
			st.PushBack(curNode)
			curNode = curNode.Left
		}
		curNode = st.Remove(st.Back()).(*TreeNode)
		curNode = curNode.Right
	}
	return root
}
