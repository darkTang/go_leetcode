package main

import "math"

func isBalanced(root *TreeNode) bool {
	var getDepth func(node *TreeNode) float64
	getDepth = func(node *TreeNode) float64 {
		if node == nil {
			return 0
		}
		l,r := getDepth(node.Left),getDepth(node.Right)
		if l == -1 || r == -1 {
			return -1
		}
		if math.Abs(l-r) > 1 {
			return -1
		}
		return math.Max(l,r)+1
	}
	return getDepth(root) != -1
}