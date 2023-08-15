package main

import (
	"container/list"
)

func findBottomLeftValue(root *TreeNode) int {
	st := list.New()
	st.PushBack(root)
	var curNode *TreeNode
	for st.Len() > 0 {
			curNode = st.Remove(st.Front()).(*TreeNode)
			if curNode.Right != nil {
				st.PushBack(curNode.Right)
			}
			if curNode.Left != nil {
				st.PushBack(curNode.Left)
			}
	}
	if curNode != nil {
		return curNode.Val
	}
	return root.Val
}

func findBottomLeftValue2(root *TreeNode) int {
	res,maxDepth := root.Val,0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				res = node.Val
			}
		}
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right !=nil {
			dfs(node.Right, depth+1)
		}
	}
	dfs(root, 1)
	return res
}