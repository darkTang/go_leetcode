package main

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var temp []int
	if root == nil {
		return res
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		for l := st.Len(); l > 0; l-- {
			curNode := st.Remove(st.Front()).(*TreeNode)
			temp = append(temp, curNode.Val)
			if curNode.Left != nil {
				st.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				st.PushBack(curNode.Right)
			}
		}
		res = append(res, temp)
		temp = []int{}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
