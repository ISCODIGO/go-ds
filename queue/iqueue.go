package queue

type List interface {
	Clear()  // O(1)
	Length() // O(1)

	Enqueue(e int) // O(1)
	Dequeue(e int) // O(1)
	Front()        // O(1)
}
