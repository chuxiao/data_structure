package basic

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Pop()
}
