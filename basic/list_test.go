package basic

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()
	l.Insert(0, 1)
	l.PrintAll()
	l.Insert(1, 2)
	l.PrintAll()
	l.Insert(0, 0)
	l.PrintAll()
	l.Insert(2, 4)
	l.PrintAll()
	fmt.Println(l.Index(2))
	l.RemoveIndex(2)
	l.PrintAll()
	l.RemoveIndex(2)
	l.PrintAll()
	l.RemoveIndex(0)
	l.PrintAll()
}
