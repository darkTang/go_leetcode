package main

import (
	"container/list"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var recursion func(node *TreeNode, s string)
	var res []string
	recursion = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			recursion(node.Left, s)
		}
		if node.Right != nil {
			recursion(node.Right, s)
		}
	}
	recursion(root, "")
	return res
}

func binaryTreePaths2(root *TreeNode) []string {
	st := list.New()
	paths := list.New()
	st.PushBack(root)
	paths.PushBack("")
	var curNode *TreeNode
	var res []string
	for st.Len() > 0 {
		path := paths.Remove(paths.Front()).(string)
		curNode = st.Remove(st.Front()).(*TreeNode)
		if curNode.Left == nil && curNode.Right == nil {
			v := path + strconv.Itoa(curNode.Val)
			res = append(res, v)
		}
		if curNode.Left != nil {
			st.PushBack(curNode.Left)
			paths.PushBack(path + strconv.Itoa(curNode.Val) + "->")
		}
		if curNode.Right != nil {
			st.PushBack(curNode.Right)
			paths.PushBack(path + strconv.Itoa(curNode.Val) + "->")
		}
	}
	return res
}
