package basic

import (
	"fmt"
	"github.com/chuxiao/data_structure/comm"
)

type Queue struct {
	base []ElementType
	front, tail int
	cap int
	size int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.cap
}

func (q *Queue) Capacity() int {
	return q.cap
}

func (q *Queue) Clear() {
	q.front, q.tail = 0, 0
	q.size = 0
}

func (q *Queue) Enqueue(e ElementType) {
	comm.Assert(!q.IsFull(), "queue is full")
	q.base[q.tail] = e
	q.tail = (q.tail + 1) % q.cap
	q.size++
}

func (q *Queue) Dequeue() {
	comm.Assert(!q.IsEmpty(), "queue is empty")
	q.front = (q.front + 1) % q.cap
	q.size--
}

func (q *Queue) Front() ElementType {
	comm.Assert(!q.IsEmpty(), "queue is empty")
	return q.base[q.front]
}

func NewQueue(cap int) *Queue {
	comm.Assert(cap > 0, fmt.Sprintf("invalid queue cap: %d", cap))
	return &Queue{
		base: make([]ElementType, cap, cap),
		cap: cap,
		front: 0,
		tail: 0,
		size: 0,
	}
}

