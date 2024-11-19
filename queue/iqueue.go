package queue

type List interface {
	Clear()    // O(1)
	Size() int // O(1)

	Enqueue(e int) error        // O(1)
	Dequeue(e int) (int, error) // O(1)
	Front() (int, error)        // O(1)
}
