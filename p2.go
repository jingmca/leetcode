package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	more := 0
	res := new(ListNode)
	head := res

	s1, s2, s := 0, 0, 0

	for l1 != nil || l2 != nil {

		if l1 == nil {
			s1 = 0
			s2 = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			s2 = 0
			s1 = l1.Val
			l1 = l1.Next
		} else {
			s1 = l1.Val
			s2 = l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		s = s1 + s2 + more
		res.Val = s % 10
		more = s / 10

		if l1 == nil && l2 == nil {
			if more == 0 {
				res.Next = nil
			} else {
				res.Next = new(ListNode)
				res = res.Next
				res.Val = more
			}
			break
		}

		res.Next = new(ListNode)
		res = res.Next

	}

	return head
}
