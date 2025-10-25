package list

import "errors"

type List interface {
	// Eliminar todo el contenido de la lista, para que esté vacía de nuevo
	Clear()

	// Insertar "it" en la ubicación actual
	// El cliente debe asegurar que la capacidad de la lista no sea excedida
	Insert(it int) bool

	// Agregar "it" al final de la lista
	// El cliente debe asegurar que la capacidad de la lista no sea excedida
	Append(it int) bool

	// Eliminar y retornar el elemento actual
	Remove() (int, error)

	// Establecer la posición actual al inicio de la lista
	MoveToStart()

	// Establecer la posición actual al final de la lista
	MoveToEnd()

	// Mover la posición actual un paso a la izquierda, sin cambio si ya está al principio
	Prev()

	// Mover la posición actual un paso a la derecha, sin cambio si ya está al final
	Next()

	// Retornar el número de elementos en la lista
	Length() int

	// Retornar la posición del elemento actual
	CurrPos() int

	// Establecer la posición actual a "pos"
	MoveToPos(pos int) bool

	// Retornar true si la posición actual está al final de la lista
	IsAtEnd() bool

	// Retornar el elemento actual
	GetValue() (int, error)

	// Retornar true si la lista está vacía
	IsEmpty() bool
}

// Errores comunes en listas
var ErrEmptyList = errors.New("empty list")
var ErrFullList = errors.New("full list")
var ErrOutOfRangeList = errors.New("out of range")
var ErrNoSuchElement = errors.New("no such element")
