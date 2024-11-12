package stack

type Stack interface {
	Clear()     // Limpiar la estructura -> O(1)
	Push(e int) // Agregar un nuevo elemento (cima) -> O(1)
	Top()       // Mostrar el elemento de la cima -> O(1)
	Pop()       // Mostrar el elemento de la cima. Remueve la cima actual -> O(1)
}
