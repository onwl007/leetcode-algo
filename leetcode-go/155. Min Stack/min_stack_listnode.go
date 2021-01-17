package main

type ListNode struct {
	val int
	min int
}

// 最小栈结构体
type MStack struct {
	stack []ListNode
}

/** initialize your data structure here. */
func NewMStack() MStack {
	return MStack{}
}

func (this *MStack) Push(x int) {
	node := ListNode{val: x, min: x}
	if len(this.stack) > 0 && this.stack[len(this.stack)-1].min < x {
		node.min = this.stack[len(this.stack)-1].min
	}
	this.stack = append(this.stack, node)
}

func (this *MStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
