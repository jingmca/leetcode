package main

import (
	"fmt"
)

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {

	fmt.Println(root)

	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	left := root.Left
	right := root.Right

	head := increasingBST(left)
	next := increasingBST(right)

	fmt.Println(left, right)

	if left != nil {
		tail := left
		for tail.Right != nil {
			tail = tail.Right
		}
		tail.Right = root
	} else {
		head = root
	}

	if right != nil {
		root.Right = next
	}

	root.Left = nil

	fmt.Println(head, next)

	return head
}
