package main

import (
	"fmt"
	"log"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT2(N int) [][]TreeNode {
	if N%2 == 0 {
		return [][]TreeNode{}
	}

	if N == 1 {
		return [][]TreeNode{{TreeNode{0, nil, nil}}}
	}

	trees := make([][]TreeNode, 0)

	sub := allPossibleFBT2(N - 2)

	for _, tree := range sub {

		for i, node := range tree {
			if node.Left == nil && node.Right == nil {
				newTree := make([]TreeNode, len(tree))
				copy(newTree, tree)
				newTree = append(newTree, []TreeNode{TreeNode{0, nil, nil}, TreeNode{0, nil, nil}}...)
				newTree[i].Left = &newTree[len(newTree)-2]
				newTree[i].Right = &newTree[len(newTree)-1]
				trees = append(trees, newTree)
			}
		}
	}

	return uniq(&trees)

}

func uniq(T *[][]TreeNode) [][]TreeNode {
	mark := make(map[int]int)
	trees := make([][]TreeNode, 0)

	for _, tree := range *T {
		m := 1
		for _, p := range tree {
			if p.Left == nil && p.Right == nil {
				m <<= 1
			} else {
				m <<= 1
				m++
			}
		}
		_, ok := mark[m]
		if !ok {
			mark[m] = 1
			trees = append(trees, tree)
		}
	}

	return trees

}

func main() {
	trees := allPossibleFBT2(7)
	fmt.Printf("%d\n", len(trees))
	for _, t := range trees {
		log.Println(t)
		fmt.Println()
	}
}
