package structures

// 声明栈的结构体
// 用于存放 int 的栈
type Stack struct {
	nums []int
}

// 返回栈的实例
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// 入栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// 栈顶元素出栈
func (s *Stack) Pop() int {
	n := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return n
}

// 返回栈的长度
func (s *Stack) Length() int {
	return len(s.nums)
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}
