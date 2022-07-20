package list

import (
	"errors"
	"fmt"
)

var ErrEmptyList = errors.New("lista vacia")
var ErrFullList = errors.New("lista llena")
var ErrOutOfRangeList = errors.New("fuera de rango")

type ArrayList struct {
	data     []int
	size     int
	capacity int
	curr     int
}

// constructor
func NewArrayList(capacidad int) *ArrayList {
	return &ArrayList{
		size:     0,
		capacity: capacidad,
		curr:     0,
		data:     make([]int, capacidad),
	}
}

// metodos
func (lista ArrayList) Length() int {
	return lista.size
}

func (lista *ArrayList) Clear() {
	lista.size = 0
	lista.curr = 0
}

func (lista *ArrayList) MoveToPos(pos int) (err error) {
	if pos < 0 || pos >= lista.size {
		err = ErrOutOfRangeList
	} else {		
		lista.curr = pos
	}
	
	return err
}

func (lista *ArrayList) MoveToStart() {
	lista.curr = 0
}

func (lista *ArrayList) MoveToEnd() {
	lista.curr = lista.size - 1
}

func (lista *ArrayList) Prev() {
	if lista.curr != 0 {
		lista.curr--
	}
}

func (lista *ArrayList) Next() {
	if lista.curr < lista.size-1 {
		lista.curr++
	}
}

func (lista ArrayList) CurrentElement() (valor int, err error) {
	// O(1)
	if lista.size == 0 {
		err = ErrEmptyList
	} else {
		valor = lista.data[lista.curr]
	}
	return
}

func (lista *ArrayList) Append(valor int) (err error) {
	// O(1)
	if lista.size > lista.capacity {
		err = ErrFullList
		return
	}

	lista.data[lista.size] = valor
	lista.size++
	return
}

func (lista *ArrayList) Remove() (valor int, err error) {
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

func (lista *ArrayList) Insert(valor int) (err error) {
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

func (lista *ArrayList) Find(valor int) int {
	// implementa linear search -> O(n)
	lista.MoveToStart()
	for i := lista.curr; i < lista.size; i++ {
		valor_actual, _ := lista.CurrentElement()

		if valor_actual == valor {
			return lista.curr
		}
		lista.Next()
	}

	return -1
}

func (lista ArrayList) String() string {
	// metodo similar al toString() de Java
	return fmt.Sprintf("Data: %v\nCurrent: %v", lista.data, lista.curr)
}
