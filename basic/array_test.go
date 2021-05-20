package basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := NewArray()
	fmt.Println(a.IsEmpty())
	for i := 0; i < 10; i++ {
		a.Append(ElementType(i))
	}
	a.PrintAll()
	a.Append(ElementType(10))
	a.PrintAll()
	a.Insert(5, ElementType(5))
	a.PrintAll()
	fmt.Println(a.Find(5))
	fmt.Println(a.Find(11))
	fmt.Println(a.Index(3))
	a.RemoveIndex(7)
	a.PrintAll()
	a.RemoveFirst(5)
	a.PrintAll()
}
