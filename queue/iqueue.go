package queue

type List interface {
	Clear()  // O(1)
	Length() // O(1)

	Enqueue(e int)  //
	Dequeue(e int)  //
	Front()  // 
}