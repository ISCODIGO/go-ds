package array

import "errors"

var ErrStackUnderflow = errors.New("lista vacia")
var ErrStackOverflow = errors.New("lista llena")

type ArrayStack struct {
	top int
	size int
	capacity int
	data []int
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
	pila.data[pila.top] = e

	return err  // devuelve nil al no inicializarse
}

func (pila ArrayStack) Top() (e int, err error) {
	if pila.isEmpty() {
		err = ErrStackUnderflow
		return  // devuelve e y err
	}
	
	e = pila.data[pila.top]
	pila.size++
	return e, err
}

func (pila *ArrayStack) Pop() (e int, err error) {
	if pila.isEmpty() {
		err = ErrStackUnderflow
		return  // devuelve e y err
	}

	e = pila.data[pila.top]
	pila.top--
	pila.size--
	return e, err
}