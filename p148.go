package main

import (
	"log"
	"sort"
)

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := make([]*ListNode, 0)

	for head != nil {
		h = append(h, head)
		head = head.Next
	}

	log.Println(h)

	sort.Slice(h, func(i, j int) bool {
		return h[i].Val < h[j].Val
	})
	log.Println(h)

	for i := 0; i < len(h); i++ {
		if i == len(h)-1 {
			h[i].Next = nil
		} else {
			h[i].Next = h[i+1]
		}
	}

	return h[0]
}
