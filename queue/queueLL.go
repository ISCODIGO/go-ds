package queue

import "go-ds/list"

type QueueLL struct {
	data *list.LinkedList
}

func NewQueueLL() *QueueLL {
	return &QueueLL{
		data: list.NewLinkedList(),
	}
}

func (q *QueueLL) Clear() {
	q.data.Clear()
}

func (q *QueueLL) Size() int {
	return q.data.Length()
}

func (q *QueueLL) Enqueue(e int) error {
	q.data.Append(e)
	return nil
}

func (q *QueueLL) Dequeue() (int, error) {
	if q.Size() == 0 {
		return 0, ErrQueueEmpty
	}
	return q.data.RemoveHead()
}

func (q *QueueLL) Front() (int, error) {
	if q.Size() == 0 {
		return 0, ErrQueueEmpty
	}
	return q.data.GetHead(), nil
}