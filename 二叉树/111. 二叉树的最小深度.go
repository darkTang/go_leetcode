package main

import (
	"container/list"
	"math"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return  1
	}
	if root.Left == nil || root.Right == nil {
		return int(math.Max(float64(minDepth(root.Left)), float64(minDepth(root.Right))) +1)
	}
	return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))) +1)
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	depth := 0
	for st.Len() > 0 {
		for l := st.Len(); l > 0; l-- {
			curNode = st.Remove(st.Front()).(*TreeNode)
			if curNode.Left == nil && curNode.Right == nil {
				return depth + 1
			}
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