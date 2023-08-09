package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return res
}

// 迭代遍历(栈的思想)
func preorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	var nodeList []*TreeNode
	var curNode *TreeNode
	nodeList = append(nodeList, root)
	for len(nodeList) > 0 {
		curNode = nodeList[len(nodeList)-1]
		nodeList = nodeList[:len(nodeList)-1]
		res = append(res, curNode.Val)
		if curNode.Right != nil {
			nodeList = append(nodeList, curNode.Right)
		}
		if curNode.Left != nil {
			nodeList = append(nodeList, curNode.Left)
		}
	}
	return
}

// 迭代遍历(双向链表)
func preorderTraversal3(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	st := list.New()
	var curNode *TreeNode
	st.PushBack(root)
	for st.Len() > 0 {
		curNode = st.Remove(st.Back()).(*TreeNode)
		res = append(res, curNode.Val)
		if curNode.Right != nil {
			st.PushBack(curNode.Right)
		}
		if curNode.Left != nil {
			st.PushBack(curNode.Left)
		}
	}
	return
}
