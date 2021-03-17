package main

import "testing"

type listNode struct {
	Val  int
	Next *listNode
}

func reverseLinkedList(head *listNode) *listNode {
	var prevNode *listNode
	currentNode := head
	for currentNode != nil {
		tmp := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = tmp
	}
	return prevNode
}

func equal(l1 *listNode, l2 *listNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return true
	}
	return false
}

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		input   *listNode
		output  *listNode
		isEqual bool
	}{
		{
			input:   &listNode{Val: 1, Next: &listNode{Val: 2, Next: &listNode{Val: 3, Next: &listNode{Val: 4, Next: nil}}}},
			output:  &listNode{Val: 4, Next: &listNode{Val: 3, Next: &listNode{Val: 2, Next: &listNode{Val: 1, Next: nil}}}},
			isEqual: true,
		},
		{
			input:   &listNode{Val: 4, Next: &listNode{Val: 3, Next: &listNode{Val: 2, Next: &listNode{Val: 1, Next: nil}}}},
			output:  &listNode{Val: 1, Next: &listNode{Val: 2, Next: &listNode{Val: 3, Next: &listNode{Val: 4, Next: nil}}}},
			isEqual: true,
		},
		{
			input:   &listNode{Val: 5, Next: &listNode{Val: 2, Next: nil}},
			output:  &listNode{Val: 2, Next: &listNode{Val: 5, Next: nil}},
			isEqual: true,
		},
	}

	for _, tc := range tests {
		output := reverseLinkedList(tc.input)
		if !equal(output, tc.output) {
			t.Errorf("want %v, got %v", tc.input, output)
		}
	}
}
