package stack

import "errors"

// Sugerencia: Go Effective
var ErrStackUnderflow = errors.New("pila vacia")
var ErrStackOverflow = errors.New("pila llena")

type ArrayStack struct {
	top int
	size int
	capacity int
	data []int  // slice
}

func NewArrayStack(capacidad int) (*ArrayStack) {
	return &ArrayStack{
		top: -1,
		size: 0,
		capacity: capacidad,
		data: make([]int, capacidad),
	}
}

func (pila *ArrayStack) isEmpty() (bool) {
	return pila.top == -1
}

func (pila *ArrayStack) isFull() (bool){
	return pila.top == pila.capacity - 1
}

func (pila *ArrayStack) Clear() {
	pila.top = -1
}

func (pila *ArrayStack) Push(e int) (err error) {
	if pila.isFull() {
		err = ErrStackOverflow
		return err
	}

	pila.top++
	pila.size++
	pila.data[pila.top] = e

	return err  // devuelve nil al no inicializarse
}

func (pila ArrayStack) Top() (e int, err error) {
	if pila.isEmpty() {
		err = ErrStackUnderflow
		return  // devuelve e y err
	}
	
	e = pila.data[pila.top]
	return e, err
}

func (pila *ArrayStack) Pop() (e int, err error) {
	e, err = pila.Top()
	pila.top--
	pila.size--
	return e, err
}