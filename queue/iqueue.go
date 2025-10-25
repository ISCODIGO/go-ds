package queue

type Queue interface {
	Clear()    // O(1)
	Size() int // O(1)
	Enqueue(e int) error   // O(1)
	Dequeue() (int, error) // O(1)
	Front() (int, error)   // O(1)
}
