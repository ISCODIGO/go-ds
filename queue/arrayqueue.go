package queue

import (
	"errors"
)

type ArrayQueue struct {
	capacity int
	size int
	front int // frente: posicion para eliminar
	rear int // fin: posicion a insertar
	data []int
}

var ErrQueueEmpty = errors.New("cola vacia")
var ErrQueueFull = errors.New("cola llena")

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		capacity: capacity,
		size: 0,
		data: make([]int, capacity),
		front: 0,
		rear: -1,
	}
}

func (q ArrayQueue) next(pos int) int {
	return (pos + 1) % q.capacity
}

func (q *ArrayQueue) Clear() {
	q.front = 0
	q.rear = -1
	q.size = 0
}

func (q *ArrayQueue) Enqueue(e int) (err error) {
	if q.Length() == q.capacity {
		err = ErrQueueFull
		return
	}

	q.rear = q.next(q.rear)
	q.data[q.rear] = e
	q.size++
	return err
}

func (q *ArrayQueue) Dequeue() (e int, err error) {
	if q.Length() == 0 {
		err = ErrQueueEmpty
		return e, err
	}

	e = q.data[q.front]
	q.front = q.next(q.front)
	q.size--
	return e, err
}

func (q *ArrayQueue) Length() int {
	return q.size
}

func (q *ArrayQueue) Front() (e int, err error) {
	if q.Length() == 0 {
		err = ErrQueueEmpty
		return
	}

	e = q.data[q.front]
	return e, err
}
