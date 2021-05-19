package basic

import "fmt"

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
		panic(fmt.Sprintf("out of range, idx: %d", idx))
	}
	if idx == 0 {
		n := &node{
			e: e,
			prev: l.head.prev,
			next: l.head,
		}
		l.head = n
	} else {
		p := l.head
		for i := 0; i < idx - 1; i++ {
			p = p.next
		}
		n := &node{
			e: e,
			prev: p,
			next: p.next,
		}
		p.next = n
	}
	l.size++
}

func (l *List) Append(e ElementType) {
	n := &node{
		e: e,
	}
	if l.tail == nil {
		n.prev, n.next = n, n
		l.head, l.tail = n, n
	} else {
		n.prev, n.next = l.tail, l.tail.next
		l.tail = n
	}
	l.size++
}

func (l *List) RemoveIndex(idx int) {
	size := l.Size()
	if idx >= size || idx < 0 {
		panic(fmt.Sprintf("out of range, idx: %d", idx))
	}
	if size == 1 {
		l.head, l.tail = nil, nil
		l.size = 0
		return
	}
	if idx == 0 {
		l.head = l.head.next
		l.head.prev = l.tail
	} else if idx == size - 1 {
		l.tail = l.tail.prev
		l.tail.next = l.head
	} else {
		p := l.head
		for i := 0; i < idx - 1; i++ {
			p = p.next
		}
		p.next = p.next.next
		p.next.prev = p
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
		prev := p.prev
		prev.next = p.next
		prev.next.prev = prev
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

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}
