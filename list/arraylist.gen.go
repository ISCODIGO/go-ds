package list

import (
	"fmt"
	"github.com/ISCODIGO/go-ds/generic"
)

type ArrayListGeneric[T comparable] struct {
	data     []T
	size     int
	capacity int
	curr     int
}

// constructor
func NewArrayListGeneric[T comparable](capacidad int) *ArrayListGeneric[T] {
	return &ArrayListGeneric[T] {
		size:     0,
		capacity: capacidad,
		curr:     0,
		data:     make([]T, capacidad),
	}
}

// metodos
func (lista ArrayListGeneric[T]) Length() int {
	return lista.size
}

func (lista *ArrayListGeneric[T]) Clear() {
	lista.size = 0
	lista.curr = 0
}

func (lista *ArrayListGeneric[T]) MoveToPos(pos int) (err error) {
	if pos < 0 || pos >= lista.size {
		err = ErrOutOfRangeList
	} else {		
		lista.curr = pos
	}
	
	return err
}

func (lista *ArrayListGeneric[T]) MoveToStart() {
	lista.curr = 0
}

func (lista *ArrayListGeneric[T]) MoveToEnd() {
	lista.curr = lista.size - 1
}

func (lista *ArrayListGeneric[T]) Prev() {
	if lista.curr != 0 {
		lista.curr--
	}
}

func (lista *ArrayListGeneric[T]) Next() {
	if lista.curr < lista.size-1 {
		lista.curr++
	}
}

func (lista ArrayListGeneric[T]) CurrentElement() (valor T, err error) {
	// O(1)
	if lista.size == 0 {
		err = ErrEmptyList
	} else {
		valor = lista.data[lista.curr]
	}
	return
}

func (lista *ArrayListGeneric[T]) Append(valor T) (err error) {
	// O(1)
	if lista.size > lista.capacity {
		err = ErrFullList
		return
	}

	lista.data[lista.size] = valor
	lista.size++
	return
}

func (lista *ArrayListGeneric[T]) Remove() (valor T, err error) {
	// O(n)
	if lista.size == 0 {
		err = ErrEmptyList
		return
	}

	valor = lista.data[lista.curr]
	for i := lista.curr; i <= lista.size; i++ {
		lista.data[i] = lista.data[i+1]
	}
	lista.size--
	return
}

func (lista *ArrayListGeneric[T]) Insert(valor T) (err error) {
	// O(n)
	if lista.size >= lista.capacity {
		err = ErrFullList
		return
	}

	for i := lista.size; i > lista.curr; i-- {
		lista.data[i] = lista.data[i-1]
	}

	lista.data[lista.curr] = valor
	lista.size++
	return
}

func (lista *ArrayListGeneric[T]) Find(valor T) int {
	// implementa linear search -> O(n)
	lista.MoveToStart()
	for i := lista.curr; i < lista.size; i++ {
		valor_actual, _ := lista.CurrentElement()

		if generic.Equals(valor_actual, valor) {
			return lista.curr
		}
		lista.Next()
	}

	return -1
}

func (lista ArrayListGeneric[T]) String() string {
	// metodo similar al toString() de Java
	return fmt.Sprintf("Data: %v\nCurrent: %v", lista.data, lista.curr)
}
