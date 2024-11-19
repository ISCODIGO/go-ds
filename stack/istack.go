package stack

type Stack interface {
	Clear()            // Limpiar la estructura -> O(1)
	Push(e int) error  // Agregar un nuevo elemento (cima) -> O(1)
	Top() (int, error) // Mostrar el elemento de la cima -> O(1)
	Pop() (int, error) // Mostrar el elemento de la cima. Remueve la cima actual -> O(1)
	Size() int         // Cantidad de elementos en la pila -> O(1)
}
