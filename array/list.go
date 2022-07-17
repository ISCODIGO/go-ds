package array

import (	
	"errors"
	"fmt"
)

type ListaArray struct {
	data []int
	size int
	capacity int
	curr int
}

// constructor
func NewListaArray() (*ListaArray) {
	return &ListaArray{
		size: 0,
		capacity: 10,
		curr: 0,
		data: make([]int, 10),
	}
}

// metodos
func (lista ListaArray) Length() (int) {
	return lista.size
}

func (lista *ListaArray) Clear() {
	lista.size = 0
	lista.curr = 0
	// 10 | 20 | 30 | 40 curr=2 size=4
	// 10 | 20 | 30 | 40 curr=0 size=0
}

func (lista *ListaArray) MoveToPos(pos int) (err error) {
	if pos < 0 || pos > lista.size {
		err = errors.New("posicion fuera de rango")
	}
	
	lista.curr = pos
	return
}

func (lista *ListaArray) MoveToStart() {
	lista.curr = 0
}

func (lista *ListaArray) MoveToEnd() {
	lista.curr = lista.size - 1
}

func (lista *ListaArray) Prev() {
	if lista.curr != 0 {
		lista.curr--
	}
}

func (lista *ListaArray) Next() {
	if lista.curr < lista.size - 1 {
		lista.curr++
	}
}

func (lista ListaArray) GetValue() (valor int, err error) {
	// O(1)
	if lista.size == 0 {
		err = errors.New("lista vacia")
	} else {
		valor = lista.data[lista.curr]
	}
	return
}

func (lista *ListaArray) Append(valor int) (err error) {
	// O(1)
	if lista.size > lista.capacity {
		err = errors.New("capacidad excedida")
	}

	lista.data[lista.size] = valor
	lista.size++
	return
}

func (lista *ListaArray) Remove() (valor int, err error) {
	// O(n)
	if lista.size == 0 {
		err = errors.New("lista vacia")
	} 

	valor = lista.data[lista.curr]
	for i := lista.curr; i <= lista.size; i++ {
		lista.data[i] = lista.data[i + 1]
	} 
	lista.size--  
	return
}

func (lista *ListaArray) Insert(valor int) (err error) {
	// O(n)
	if lista.size >= lista.capacity {
		err = errors.New("capacidad excedida")
	}

	for i := lista.size; i > lista.curr; i-- {
		lista.data[i] = lista.data[i - 1]
	}

	lista.data[lista.curr] = valor
	lista.size++
	return
}

func (lista *ListaArray) Find(valor int) (int) {
	// implementar linear search -> O(n)
	lista.MoveToStart()
	for i := lista.curr; i < lista.size; i++ {
		valor_actual, _ := lista.GetValue()

		if valor_actual == valor {
			return lista.curr
		}
		lista.Next()
	}

	return -1
}

func (lista ListaArray) String() (string) {
	return fmt.Sprintf("Data: %v\nCurrent: %v", lista.data, lista.curr)
}