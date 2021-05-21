package basic

import (
	"fmt"
	"github.com/chuxiao/data_structure/comm"
	"strconv"
	"strings"
)

type ElementType int

type Array struct {
	base []ElementType
	cap int
}

func (a *Array) IsEmpty() bool {
	return a.Size() == 0
}

func (a *Array) Size() int {
	return len(a.base)
}

func (a *Array) Index(idx int) ElementType {
	comm.Assert(idx >= 0 && idx < a.Size(), fmt.Sprintf("out of range, idx: %d", idx))
	return a.base[idx]
}

func (a *Array) Find(e ElementType) int {
	for idx, v := range a.base {
		if v == e {
			return idx
		}
	}
	return -1
}

func (a *Array) Insert(idx int, e ElementType) {
	if idx == a.Size() {
		a.Append(e)
		return
	}
	size := a.Size()
	comm.Assert(idx >= 0 && idx < a.Size(), fmt.Sprintf("out of range, idx: %d", idx))
	if size == a.cap {
		newCap := a.cap * 2
		newBase := make([]ElementType, 0, newCap)
		newBase = append(newBase, a.base...)
		a.base = newBase
		a.cap = newCap
	}
	a.base = append(a.base, a.base[size - 1])
	for i := size - 1; i > idx; i-- {
		a.base[i] = a.base[i - 1]
	}
	a.base[idx] = e
}

func (a *Array) Append(e ElementType) {
	if a.Size() == a.cap {
		newCap := a.cap * 2
		newBase := make([]ElementType, 0, newCap)
		newBase = append(newBase, a.base...)
		a.base = newBase
		a.cap = newCap
	}
	a.base = append(a.base, e)
}

func (a *Array) RemoveIndex(idx int) {
	size := a.Size()
	comm.Assert(idx >= 0 && idx < a.Size(), fmt.Sprintf("out of range, idx: %d", idx))
	for i := idx; i < size - 1; i++ {
		a.base[i] = a.base[i + 1]
	}
	a.base = a.base[0:size-1]
}

func (a *Array) RemoveFirst(e ElementType) {
	idx := 0
	for a.base[idx] != e {
		idx++
	}
	size := a.Size()
	if idx == size {
		return
	}
	for i := idx; i < size - 1; i++ {
		a.base[i] = a.base[i + 1]
	}
	a.base = a.base[0:size-1]
}

func (a *Array) PrintAll() {
	ss := make([]string, 0)
	for _, v := range a.base {
		ss = append(ss, strconv.Itoa(int(v)))
	}
	fmt.Println(strings.Join(ss, "    "))
}

func NewArray() *Array {
	cap := 32
	a := &Array{
		cap: cap,
	}
	a.base = make([]ElementType, 0, cap)
	return a
}
