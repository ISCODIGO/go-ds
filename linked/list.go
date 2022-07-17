package linked

import (
	"errors"
	_ "errors"
)

type LinkedList struct {
	size int  // la cantidad de nodos de la lista
	curr *Node  // el nodo actual (el que se recorre actualmente)
	head *Node  // un enlace al primer nodo
	tail *Node	// un enlace al ultimo nodo
}

// constructor (inicializador)
func NewLinkedList() (*LinkedList) {
	return &LinkedList{}
}

// metodos

func (lista LinkedList) isEmpty() (bool) {
	return lista.head == nil
}

// para insertar el primero de la lista enlazada
func (lista *LinkedList) insertFirstNode(node *Node) {
	lista.head = node
	lista.tail = node
	lista.curr = node
	return
}

// Remueve todos los nodos de la lista
func (lista *LinkedList) Clear() {
	lista.curr = nil
	lista.head = nil
	lista.tail = nil
	lista.size = 0
}

// Agrega un nodo desplazando al nodo actual (curr)
func (lista *LinkedList) Insert(e int) (node *Node){
	return nil
}

// Agrega un nodo al final de la lista (tail)
func (lista *LinkedList) Append(e int) (node *Node) {
	node = NewNode(e)

	if lista.isEmpty() {
		lista.insertFirstNode(node)
	} else {
		lista.tail.next = node
		lista.tail = node
	}

	lista.size++
	return
}

// Remueve el nodo actual (curr)
func (lista *LinkedList) Remove() (node *Node, err error) {
	return nil, nil
}

// Coloca al [curr] en el nodo head
func (lista *LinkedList) MoveToStart() {
	if !lista.isEmpty() {
		lista.curr = lista.head
	}
}

// Coloca [curr] en el nodo tail
func (lista *LinkedList) MoveToEnd() {
	if !lista.isEmpty() {
		lista.curr = lista.tail
	}
}

func (lista *LinkedList) Prev() {
}

func (lista *LinkedList) Next() {
}

func (lista LinkedList) Length() (int) {
	return lista.size
}

func (lista LinkedList) GetCurrentPosition() (int) {
	return 0
}

func (lista LinkedList) GetValue() (e int, err error) {
	if lista.isEmpty() {
		err = errors.New("lista vacia")
	}
	e = lista.curr.element
	return
}


// 10 -> 20 -> 30