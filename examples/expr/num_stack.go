package main

import (
	"fmt"
	"github.com/chuxiao/data_structure/comm"
)

type numStack struct {
	base []int
	top int
	cap int
}

func (s *numStack) IsEmpty() bool {
	return s.top < 0
}

func (s *numStack) IsFull() bool {
	return s.top > s.cap - 1
}

func (s *numStack) Capacity() int {
	return s.cap
}

func (s *numStack) Clear() {
	s.top = -1
}

func (s *numStack) Push(e int) {
	comm.Assert(!s.IsFull(), "numStack is full")
	s.top++
	s.base[s.top] = e
}

func (s *numStack) Pop() {
	comm.Assert(!s.IsEmpty(), "numStack is empty")
	s.top--
}

func (s *numStack) Top() int {
	comm.Assert(!s.IsEmpty(), "numStack is empty")
	return s.base[s.top]
}

func NewNumStack(cap int) *numStack {
	comm.Assert(cap > 0, fmt.Sprintf("invalid numStack cap: %d", cap))
	return &numStack{
		base: make([]int, cap, cap),
		top: -1,
		cap: cap,
	}
}