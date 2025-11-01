package stack

import "errors"

type Stack interface {
	Clear()            // Limpiar la estructura -> O(1)
	Push(e int) error  // Agregar un nuevo elemento (cima) -> O(1)
	Top() (int, error) // Mostrar el elemento de la cima -> O(1)
	Pop() (int, error) // Remueve la cima actual y la muestra -> O(1)
	Size() int         // Cantidad de elementos en la pila -> O(1)
}


// Sugerencia: Go Effective
var ErrStackUnderflow = errors.New("pila vacia")
var ErrStackOverflow = errors.New("pila llena")