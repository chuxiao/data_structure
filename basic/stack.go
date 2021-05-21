package basic

import (
	"fmt"
	"github.com/chuxiao/data_structure/comm"
)

type Stack struct {
	base []ElementType
	top int
	cap int
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func (s *Stack) IsFull() bool {
	return s.top > s.cap - 1
}

func (s *Stack) Capacity() int {
	return s.cap
}

func (s *Stack) Clear() {
	s.top = -1
}

func (s *Stack) Push(e ElementType) {
	comm.Assert(!s.IsFull(), "stack is full")
	s.top++
	s.base[s.top] = e
}

func (s *Stack) Pop() {
	comm.Assert(!s.IsEmpty(), "stack is empty")
	s.top--
}

func (s *Stack) Top() ElementType {
	comm.Assert(!s.IsEmpty(), "stack is empty")
	return s.base[s.top]
}

func NewStack(cap int) *Stack {
	comm.Assert(cap > 0, fmt.Sprintf("invalid stack cap: %d", cap))
	return &Stack{
		base: make([]ElementType, cap, cap),
		top: -1,
		cap: cap,
	}
}

