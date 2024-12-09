package list

import "errors"

type List interface {
	Clear()  // O(1)
	Size() // O(1)
	IsEmpty() bool  // O(1)

	Append(e int)  // O(1)
	Insert(e int)  // O(n)
	Remove()  // O(n)
}

// Errores comunes en las listas
var ErrEmptyList = errors.New("lista vacia")
var ErrFullList = errors.New("lista llena")
var ErrOutOfRangeList = errors.New("fuera de rango")