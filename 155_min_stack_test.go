package main

import (
	"fmt"
	"testing"
)

type MinStack struct {
	items [][2]int
}

/** initialize your data structure here. */
func constructor() MinStack {
	return MinStack{items: [][2]int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.items) == 0 {
		this.items = append(this.items, [2]int{x, x})
	} else {
		min := this.items[len(this.items)-1][1]
		if x < min {
			min = x
		}
		this.items = append(this.items, [2]int{x, min})
	}
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	if len(this.items) == 0 {
		return 0
	}
	return this.items[len(this.items)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1][1]
}

func TestMinStack(t *testing.T) {
	ms := constructor()
	ms.Push(0)
	ms.Push(5)
	ms.Push(-3)
	ms.Push(6)
	fmt.Println(ms.GetMin())
	ms.Pop()
	ms.Pop()
	fmt.Println(ms.GetMin())
}
