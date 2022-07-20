package linked

import (
	"errors"
)

var ErrEmptyLinkedList = errors.New("lista enlazada vacia")
var ErrFullLinkedList = errors.New("lista enlazada llena")
var ErrOutOfRangeLinkedList = errors.New("fuera de rango")

type Node struct {
	element int  // dato almacenado
	next *Node   // enlace al siguiente nodo
}

// Funcion que permite inicializar un objeto de tipo Node
func NewNode(element int) (*Node) {
	return &Node{
		element: element,
		next: nil,
	}
}

type LinkedList struct {
	size int  // la cantidad de nodos de la lista
	curr *Node  // el nodo actual (el que se recorre actualmente)
	head *Node  // un enlace al primer nodo
	tail *Node	// un enlace al ultimo nodo
}

// constructor (inicializador)
func NewLinkedList() (*LinkedList) {
	return &LinkedList{
		size: 0,
		curr: nil,
		head: nil,
		tail: nil,
	}
}


func (lista LinkedList) isEmpty() (bool) {
	return lista.head == nil
}

// Remueve todos los nodos de la lista
func (lista *LinkedList) Clear() {
	// O(1)
	lista.curr = nil
	lista.head = nil
	lista.tail = nil
	lista.size = 0
}

// Agrega un nodo al final de la lista (tail)
func (lista *LinkedList) Append(e int) (node *Node) {
	// O(1)
	node = NewNode(e)

	if lista.isEmpty() {
		lista.head = node
		lista.curr = node  // opcional: para no tener un current nil
	} else {
		lista.tail.next = node
	}
	
	lista.tail = node
	lista.size++
	return node
}

// Agrega un nodo desplazando al nodo actual (curr)
func (lista *LinkedList) Insert(e int) (node *Node){
	// O(n)
	if lista.curr == lista.tail {		
		return lista.Append(e)  // la funcionalidad ya existe
	}

	node = NewNode(e)
	node.next = lista.curr

	if lista.curr == lista.head {
		// esto significa que hay que un nuevo head
		lista.head = node
	} else {
		// nodos interiores
		lista.Prev()  // O(n)
		lista.curr.next = node
	}

	lista.size++
	return node
}

// Remueve el nodo actual (curr)
func (lista *LinkedList) Remove() (node *Node, err error) {
	if lista.isEmpty() {
		err = ErrEmptyLinkedList
		return  // como las salidas llevan nombre esto es posible
	} 
	
	node = lista.curr

	if lista.curr == lista.head {
		// Si es head el eliminado
		lista.curr = lista.curr.next
		lista.head = lista.curr
	} else {
		// Cualquier otro nodo: encontrar el nodo previo a curr
		lista.Prev()
		lista.curr.next = lista.curr.next.next // eliminar el nodo: lista.curr.next
		lista.curr = lista.curr.next  // opcional: para que no haga referencia a un nodo que "no existe"
	}

	lista.size--
	return node, err
}

// Coloca al [curr] en el nodo head
func (lista *LinkedList) MoveToStart() {
	if !lista.isEmpty() {
		lista.curr = lista.head
	}
}

// Coloca [curr] en el tail
func (lista *LinkedList) MoveToEnd() {
	if !lista.isEmpty() {
		lista.curr = lista.tail
	}
}

// Mueve [curr] una cantidad de nodos (pos) a partir de head
func (lista *LinkedList) MoveToPos(pos int) (err error) {
	if pos < 0 || pos >= lista.size {
		err = errors.New("posicion fuera de rango")
	} else {
		lista.curr = lista.head
		for i := 1; i <= pos; i++ {
			lista.curr = lista.curr.next
		}
	}
	return err
}

// Retrocede un nodo
func (lista *LinkedList) Prev() {
	// O(n)
	if lista.isEmpty() || lista.curr == lista.head {
		return  // aqui return es una secuencia de escape
	}

	temp := lista.head
	for temp.next != lista.curr {
		temp = temp.next
	}

	lista.curr = temp
}

// Se adelanta un nodo
func (lista *LinkedList) Next() {
	if lista.isEmpty() || lista.curr == lista.tail {
		return
	}

	lista.curr = lista.curr.next
}

func (lista LinkedList) Length() (int) {
	return lista.size
}

func (lista LinkedList) CurrentPosition() (int) {
	if lista.isEmpty() {
		return -1
	}

	if lista.curr == lista.tail {
		return lista.size - 1
	}

	temp := lista.head
	pos := 0
	if temp.next != lista.curr {
		temp = temp.next
		pos++
	}

	return pos
}

func (lista LinkedList) CurrentElement() (e int, err error) {
	if lista.isEmpty() {
		err = errors.New("lista vacia")
	}
	e = lista.curr.element
	return e, err
}


