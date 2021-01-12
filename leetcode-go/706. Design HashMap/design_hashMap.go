package main

const Len int = 100000

type ListNode struct {
	key  int
	val  int
	next *ListNode
}

func (node *ListNode) Put(key int, value int) {
	if node.key == key {
		node.val = value
		return
	}
	if node.next == nil {
		node.next = &ListNode{key: key, val: value, next: nil}
	}
	node.next.Put(key, value)
}

func (node *ListNode) Get(key int) int {
	if node.key == key {
		return node.val
	}
	if node.next == nil {
		return -1
	}
	return node.next.Get(key)
}

func (node *ListNode) Remove(key int) *ListNode {
	if node.key == key {
		p := node.next
		node.next = nil
		return p
	}
	if node.next != nil {
		return node.Remove(key)
	}
	return nil
}

type MyHashMap struct {
	content [Len]*ListNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		this.content[this.Hash(key)] = &ListNode{key: key, val: value, next: nil}
		return
	}
	node.Put(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	listNode := this.content[this.Hash(key)]
	if listNode == nil {
		return -1
	}
	return listNode.Get(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	listNode := this.content[this.Hash(key)]
	if listNode == nil {
		return
	}
	this.content[this.Hash(key)] = listNode.Remove(key)
}

func (this *MyHashMap) Hash(value int) int {
	return value % Len
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
