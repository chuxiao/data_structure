package basic

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(6)
	fmt.Printf("front: %d, tail: %d, size: %d\n", q.front, q.tail, q.size)
}