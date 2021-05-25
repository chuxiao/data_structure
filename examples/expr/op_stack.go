package main

import (
	"fmt"
	"github.com/chuxiao/data_structure/comm"
)

type opStack struct {
	base []rune
	top int
	cap int
}

func (s *opStack) IsEmpty() bool {
	return s.top < 0
}

func (s *opStack) IsFull() bool {
	return s.top > s.cap - 1
}

func (s *opStack) Capacity() int {
	return s.cap
}

func (s *opStack) Clear() {
	s.top = -1
}

func (s *opStack) Push(e rune) {
	comm.Assert(!s.IsFull(), "opStack is full")
	s.top++
	s.base[s.top] = e
}

func (s *opStack) Pop() {
	comm.Assert(!s.IsEmpty(), "opStack is empty")
	s.top--
}

func (s *opStack) Top() rune {
	comm.Assert(!s.IsEmpty(), "opStack is empty")
	return s.base[s.top]
}

func NewOpStack(cap int) *opStack {
	comm.Assert(cap > 0, fmt.Sprintf("invalid opStack cap: %d", cap))
	return &opStack{
		base: make([]rune, cap, cap),
		top: -1,
		cap: cap,
	}
}