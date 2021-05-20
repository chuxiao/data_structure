package basic

import "fmt"

type Stack struct {
	base []ElementType
	top int
	cap int
}

func (s *Stack) IsEmpty() bool {
	// TODO
	return false
}

func (s *Stack) IsFull() bool {
	// TODO
	return false
}

func (s *Stack) Capacity() int {
	// TODO
	return 0
}

func (s *Stack) Clear() {
	// TODO
}

func (s *Stack) Push(e ElementType) {
	// TODO
}

func (s *Stack) Pop() {
	// TODO
}

func (s *Stack) Top() ElementType {
	// TODO
	return 0
}

func NewStack(cap int) *Stack {
	if cap <= 0 {
		panic(fmt.Sprintf("invalid stack cap: %d", cap))
	}
	return &Stack{
		base: make([]ElementType, 0, cap),
		top: -1,
		cap: cap,
	}
}

