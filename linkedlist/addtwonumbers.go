package linkedlist

/*
Definition for singly-linked list.

	type ListNode struct {
	    Val int
	    Next *ListNode

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// need to pass pointer of result btw, because probably best to use recursion
	// why? because the output order is in reverse.
	// but it should also handle the remainder. Think about how that should work.

	return recursiveSolution(l1, l2, 0)
}

func recursiveSolution(l1, l2 *ListNode, carry int) *ListNode {
	// paths
	// 1. case when only l1 has value
	// 2. case when only l2 has value
	// 3. case when both l1 l2 has value
	// 4. case when none of them has value
	// in all cases nodes should point to next when its not nil
	// recursion should last until both nodes are currently nil

	// exit case
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	return &ListNode{
		Val:  sum % 10,
		Next: recursiveSolution(l1, l2, sum/10),
	}
}
