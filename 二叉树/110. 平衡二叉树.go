package main

import (
	"container/list"
	"math"
)

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

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	for st.Len() > 0 {
			curNode = st.Remove(st.Front()).(*TreeNode)
			if math.Abs(float64(maxDepth3(curNode.Left))-float64(maxDepth3(curNode.Right))) > 1 {
				return false
			}
			if curNode.Left != nil {
				st.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				st.PushBack(curNode.Right)
			}
	}
	return true
}

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
}