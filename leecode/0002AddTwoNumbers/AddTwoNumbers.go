package leetcode

import (
	"fmt"
)

//```
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//```
//Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Printf("twoSum l1 = %+v,l2= %+v\n", l1, l2)
	head := &ListNode{Val: 0}
	carry, n1, n2, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {

		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{(n2 + n1 + carry) % 10, nil}
		current = current.Next
		carry = (n2 + n1 + carry) / 10
	}
	fmt.Printf("twoSum result list = %+v", head.Next)
	fmt.Printf("twoSum result [] = %+v", List2Ints(head.Next))
	return head.Next
}
