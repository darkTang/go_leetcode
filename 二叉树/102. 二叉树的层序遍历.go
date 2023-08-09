package main

import (
	"container/list"
)

func main() {
}

func levelOrder(root *TreeNode) [][]int {
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
	return res
}
