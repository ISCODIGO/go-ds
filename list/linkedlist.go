package list

import (
	"fmt"
)

type Node struct {
	Element int   // dato almacenado
	Next    *Node // enlace al siguiente nodo
}

// Funcion que permite inicializar un objeto de tipo Node
func NewNode(element int) *Node {
	return &Node{
		Element: element,
		Next:    nil,
	}
}

// Representa la lista enlazada
type LinkedList struct {
	head *Node // referencia al primer nodo de la lista
	tail *Node // referencia al último nodo de la lista
	size int   // tamaño de la lista
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list LinkedList) Size() int {
	return list.size
}

func (list LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Crear el primer nodo en la lista (privada)
func (list *LinkedList) createFirstNode(element int) *Node {
	newNode := NewNode(element)
	list.head = newNode
	list.tail = newNode
	list.size = 1
	return newNode
}

// Obtener el nodo previo a un nodo dado (privada)
func (list *LinkedList) getPreviousNode(node *Node) (*Node, error) {
	// si la lista esta vacia
	if list.IsEmpty() {
		return nil, ErrEmptyList
	}

	// si el nodo es el primero no tiene previo
	if list.head == node {
		return nil, ErrOutOfRangeList
	}

	aux := list.head
	for aux.Next != nil && aux.Next != node {
		aux = aux.Next
	}

	if aux.Next == node {
		return aux, nil
	}
	return nil, nil
}

// Agregar un nuevo nodo al inicio de la lista
func (list *LinkedList) AddHead(element int) *Node {
	if list.IsEmpty() {
		return list.createFirstNode(element)
	}

	newNode := NewNode(element)
	newNode.Next = list.head
	list.head = newNode
	list.size++
	return newNode
}

// Agregar un nuevo nodo al final de la lista
func (list *LinkedList) AddTail(element int) *Node {
	if list.IsEmpty() {
		return list.createFirstNode(element)
	}
	newNode := NewNode(element)
	list.tail.Next = newNode
	list.tail = newNode
	list.size++
	return newNode
}

// Agregar un nodo después de un nodo específico
func (list *LinkedList) Add(element int, node *Node) *Node {
	if list.IsEmpty() || node == nil {
		return list.createFirstNode(element)
	}

	if node == list.tail {
		return list.AddTail(element)
	}

	newNode := NewNode(element)
	newNode.Next = node.Next
	node.Next = newNode
	list.size++
	return newNode
}

// Eliminar el primer nodo de la lista
func (list *LinkedList) RemoveHead() (int, error) {
	if list.head == nil {
		return 0, ErrEmptyList
	}

	element := list.head.Element
	list.head = list.head.Next
	if list.head == nil {
		list.tail = nil
	}
	list.size--
	return element, nil
}

// Eliminar el último nodo de la lista
func (list *LinkedList) RemoveTail() (int, error) {
	if list.IsEmpty() {
		return 0, ErrEmptyList
	}

	if list.head == list.tail {
		list.Clear()
		return 0, nil
	}

	previous, err := list.getPreviousNode(list.tail)
	if err != nil {
		return 0, err
	}

	element := list.tail.Element
	previous.Next = nil
	list.tail = previous
	list.size--
	return element, nil
}

// Eliminar un nodo específico por valor
func (list *LinkedList) Remove(element int) (err error) {
	if list.IsEmpty() {
		return ErrEmptyList
	}

	if list.head.Element == element {
		_, err = list.RemoveHead()
		return
	}

	if list.tail.Element == element {
		_, err = list.RemoveTail()
		return
	}

	current := list.head
	for current.Next != nil && current.Next.Element != element {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
		if current.Next == nil {
			list.tail = current
		}
		list.size--
	}
	return
}

// Mostrar los elementos de la lista enlazada
func (list *LinkedList) ToString() string {
	if list.IsEmpty() {
		return ""
	}

	result := ""
	current := list.head
	for current != nil {
		result += fmt.Sprintf("%d -> ", current.Element)
		current = current.Next
	}

	return result
}

func (list *LinkedList) GetHead() int {
	return list.head.Element
}

func (list *LinkedList) GetTail() int {
	return list.tail.Element
}
