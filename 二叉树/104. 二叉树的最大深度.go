package main

import (
	"container/list"
	"math"
)

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
}

// 迭代层序遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	for st.Len() > 0 {
		for l := st.Len(); l > 0; l-- {
			curNode = st.Remove(st.Front()).(*TreeNode)
			if curNode.Left != nil {
				st.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				st.PushBack(curNode.Right)
			}
		}
		depth++
	}
	return depth
}
