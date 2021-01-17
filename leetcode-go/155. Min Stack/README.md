# [155. Min Stack](https://leetcode-cn.com/problems/min-stack/)

## 题目

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

- push(x) —— 将元素 x 推入栈中。
- pop() —— 删除栈顶的元素。
- top() —— 获取栈顶元素。
- getMin() —— 检索栈中的最小元素。

示例:

```c
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

提示：

- pop、top 和 getMin 操作总是在 非空栈 上调用。

## 题目大意

栈的基本操作，这道题的关键是在常数时间内获取到栈内元素的最小值

## 解题思路

### 解法一，使用辅助栈

可以使用辅助栈，即用两个相同的数组，一个正常存储入栈的元素，一个每次入栈的时候与栈顶元素比较，将较小元素入栈，出栈的时候与正常的栈一起出栈即可，这样最小栈的栈顶都是最小的元素，常数时间就可以获得最小值

### 解法二，不使用辅助栈，不用额外的空间

可以使用链表存储栈的元素，定义一个链表节点，一个元素，一个最小值，每次在入栈的时候都跟栈顶元素的最小值比较，当栈顶元素的最小值小于当前入栈的元素时，将当前入栈的节点的最小值改为栈顶元素的最小值。
