package stack

import "errors"

// Sugerencia: Go Effective
var ErrStackUnderflow = errors.New("pila vacia")
var ErrStackOverflow = errors.New("pila llena")

type ArrayStack struct {
	top      int  // posicion de la cima
	size     int  // tamaño de la pila
	capacity int  // capacidad del slice
	data     []int // slice
}

func NewArrayStack(capacidad int) ArrayStack {
	return ArrayStack{
		top:      -1,
		size:     0,
		capacity: capacidad,
		data:     make([]int, capacidad),
	}
}

func (pila ArrayStack) isEmpty() bool {
	return pila.size == 0
}

func (pila ArrayStack) isFull() bool {
	return pila.top == pila.capacity-1
}

func (pila *ArrayStack) Clear() {
	pila.top = -1
	pila.size = 0
}

func (pila *ArrayStack) Push(e int) (err error) {
	if pila.isFull() {
		err = ErrStackOverflow
		return  // devuelve la variable err
	}

	pila.top++
	pila.size++
	pila.data[pila.top] = e

	return err // devuelve nil al no inicializarse
}

func (pila ArrayStack) Top() (e int, err error) {
	if pila.isEmpty() {
		err = ErrStackUnderflow
		return // devuelve e y err
	}

	e = pila.data[pila.top]
	return  // devuelve e y err
}

func (pila *ArrayStack) Pop() (e int, err error) {
	e, err = pila.Top()

	if err != nil {
		return
	}

	pila.top--
	pila.size--
	return
}


func (pila ArrayStack) Size() int {
	return pila.size
}