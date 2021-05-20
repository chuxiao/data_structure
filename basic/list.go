package basic

import (
	"fmt"
	"strconv"
	"strings"
)

// type ElementType int

type node struct {
	e ElementType
	prev, next *node
}

type List struct {
	head, tail *node
	size int
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Insert(idx int, e ElementType) {
	size := l.Size()
	if idx == size {
		l.Append(e)
		return
	}
	if idx >= size || idx < 0 {
		panic(fmt.Sprintf("idx out of range: %d", idx))
	}
	n := &node{
		e: e,
	}
	if idx == 0 {
		n.prev, n.next = l.head.prev, l.head
		l.head.prev, l.tail.next = n, n
		l.head = n
	} else {
		p := l.head
		for i := 0; i < idx - 1; i++ {
			p = p.next
		}
		n.prev, n.next = p, p.next
		p.next.prev, p.next = n, n
	}
	l.size++
}

func (l *List) Append(e ElementType) {
	n := &node{
		e: e,
	}
	if l.tail != nil {
		n.prev, n.next = l.tail, l.tail.next
		l.head.prev, l.tail.next = n, n
		l.tail = n
	} else {
		n.prev, n.next = n, n
		l.head, l.tail = n, n
	}
	l.size++
}

func (l *List) RemoveIndex(idx int) {
	size := l.Size()
	if idx >= size || idx < 0 {
		panic(fmt.Sprintf("idx out of range: %d", idx))
	}
	if size == 1 {
		l.head, l.tail = nil, nil
	} else {
		p := l.head
		for i := 0; i < idx; i++ {
			p = p.next
		}
		p.prev.next, p.next.prev = p.next, p.prev
		if p == l.head {
			l.head = p.next
		}
		if p == l.tail {
			l.tail = p.prev
		}
	}
	l.size--
}

func (l *List) RemoveFirst(e ElementType) {
	size := l.Size()
	if size == 0 {
		return
	}
	p := l.head
	for p.e != e {
		if p == l.tail {
			return
		}
		p = p.next
	}
	if size == 1 {
		l.head, l.tail = nil, nil
	} else {
		p.prev.next, p.next.prev = p.next, p.prev
		if p == l.head {
			l.head = p.next
		}
		if p == l.tail {
			l.tail = p.prev
		}
	}
	l.size--
}

func (l *List) Index(idx int) ElementType {
	size := l.Size()
	if idx >= size || idx < 0 {
		panic(fmt.Sprintf("out of range, idx: %d", idx))
	}
	p := l.head
	for i := 0; i < idx; i++ {
		p = p.next
	}
	return p.e
}

func (l *List) PrintAll() {
	if l.Size() == 0 {
		fmt.Println("no element")
	}
	var ss []string
	p := l.head
	for {
		ss = append(ss, strconv.Itoa(int(p.e)))
		if p == l.tail {
			break
		}
		p = p.next
	}
	fmt.Printf("%d elements: %s, first: %d, last: %d, first_prev: %d, last_next: %d\n", l.size, strings.Join(ss, "    "), l.head.e, l.tail.e, l.head.prev.e, l.tail.next.e)
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}
