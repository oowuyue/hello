package ds

// Stack ...
type Stack struct {
	stack []int
}

// Push ...
func (s *Stack) Push(data int) {
	s.stack = append([]int{data}, s.stack...)
}

// Pop ...
func (s *Stack) Pop() int {
	popValue := s.stack[0]
	s.stack = s.stack[1:]
	return popValue
}
