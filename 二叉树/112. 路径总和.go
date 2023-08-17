package main

import "container/list"

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum-=root.Val
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			return true
		} else {
			return false
		}
	}
	return hasPathSum(root.Left, targetSum) || 	hasPathSum(root.Right, targetSum)
}

// 迭代法 每次push进去为一个节点+节点值
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	st := list.New()
	st.PushBack(root)
	st.PushBack(root.Val)
	for st.Len() > 0 {
		curNode := st.Remove(st.Front()).(*TreeNode)
		sum := st.Remove(st.Front()).(int)
		if curNode.Left == nil && curNode.Right == nil && sum == targetSum {
			return true
		}
		if curNode.Left !=nil {
			st.PushBack(curNode.Left)
			st.PushBack(sum+curNode.Left.Val)
		}
		if curNode.Right !=nil {
			st.PushBack(curNode.Right)
			st.PushBack(sum+curNode.Right.Val)
		}
	}
	return false
}