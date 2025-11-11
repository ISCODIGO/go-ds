package queue

type Queue interface {
	Clear()    // O(1)
	Size() int // O(1)
	Enqueue(e int) error   // Agregar al final de la cola O(1)
	Dequeue() (int, error) // Remover un elemento al frente O(1)
	Front() (int, error)   // Obtiene el valor del frente O(1)
}
