package main

import (
	"fmt"
	"testing"
)

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stack1: []int{}, stack2: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			item := this.stack1[len(this.stack1)-1]
			this.stack2 = append(this.stack2, item)
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
		first := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return first
	} else {
		first := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return first
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stack2) != 0 {
		return this.stack2[len(this.stack2)-1]
	} else {
		for len(this.stack1) > 0 {
			item := this.stack1[len(this.stack1)-1]
			this.stack2 = append(this.stack2, item)
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
		return this.stack2[len(this.stack2)-1]
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Peek()
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Push(6)
	q.Push(7)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
