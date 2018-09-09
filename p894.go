package main

import "fmt"

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var forest = make(map[int][]*TreeNode)

func allPossibleFBT(N int) []*TreeNode {

	if N%2 == 0 {
		return []*TreeNode{}
	}

	if N == 1 {
		return []*TreeNode{&TreeNode{0, nil, nil}}
	}

	if forest[N] == nil {
		trees := make([]*TreeNode, 0)

		for i := 1; i < N; i++ {
			j := N - 1 - i
			for _, left := range allPossibleFBT(i) {
				for _, right := range allPossibleFBT(j) {
					root := &TreeNode{0, left, right}
					trees = append(trees, root)
				}
			}
		}
		forest[N] = trees
	}

	return forest[N]

}

func main() {
	fmt.Println(len(allPossibleFBT(17)))
}
