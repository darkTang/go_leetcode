package main

import "container/list"

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return res
}

func inorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	var curNode *TreeNode = root
	for curNode != nil || len(stack) > 0 {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			curNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, curNode.Val)
			curNode = curNode.Right
		}
	}
	return
}

func inorderTraversal3(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	st := list.New()
	var curNode *TreeNode = root
	for curNode != nil || st.Len() > 0 {
		if curNode != nil {
			st.PushBack(curNode)
			curNode = curNode.Left
		} else {
			curNode = st.Remove(st.Back()).(*TreeNode)
			res = append(res, curNode.Val)
			curNode = curNode.Right
		}
	}
	return
}
