package main

import "strconv"

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

