package main

import (
	"fmt"
)

//Node have rank for tree height and key for its name
type Node struct {
	key    int
	parent pNode
	rank   int
}

type pNode *Node

func find(x pNode) pNode {

	for x != x.parent {
		x = find(x.parent)
	}

	return x
}

func union(a pNode, b pNode) {
	rootA := find(a)
	rootB := find(b)

	if rootA == rootB {
		fmt.Printf("%d and %d are connected!\n", a.key, b.key)
		return
	}

	if rootA.rank > rootB.rank {
		rootB.parent = rootA
		fmt.Printf("%d is linking to %d\n", b.key, a.key)
	} else {
		rootA.parent = rootB
		if rootA.rank == rootB.rank {
			rootB.rank++
		}
		fmt.Printf("%d is linking to %d\n", a.key, b.key)
	}

}

func main() {
	tree := make([]Node, 10)
	for i := range tree {
		tree[i] = Node{key: i, rank: 0}
		tree[i].parent = &tree[i]
	}

	fmt.Println(tree)

	union(&tree[1], &tree[2])
	union(&tree[0], &tree[3])
	union(&tree[2], &tree[9])
	union(&tree[7], &tree[9])
	union(&tree[9], &tree[3])

	fmt.Println(tree)

	fmt.Println(15/10, 15%10)
}
