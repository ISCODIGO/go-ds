package list

import (
	"fmt"
)

type Node struct {
	Element int   // dato almacenado
	Next    *Node // enlace al siguiente nodo
}

// Función que permite inicializar un objeto de tipo Node
func NewNode(element int) *Node {
	return &Node{
		Element: element,
		Next:    nil,
	}
}

// Representa la lista enlazada
type LinkedList struct {
	head    *Node // referencia al primer nodo de la lista
	tail    *Node // referencia al último nodo de la lista
	current *Node // referencia al nodo actual
	size    int   // tamaño de la lista
	currPos int   // posición actual
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:    nil,
		tail:    nil,
		current: nil,
		size:    0,
		currPos: 0,
	}
}

func (list LinkedList) Length() int {
	return list.size
}

func (list LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.current = nil
	list.size = 0
	list.currPos = 0
}

// Crear el primer nodo en la lista (privada)
func (list *LinkedList) createFirstNode(element int) *Node {
	newNode := NewNode(element)
	list.head = newNode
	list.tail = newNode
	list.current = newNode
	list.size = 1
	list.currPos = 0
	return newNode
}

// Obtener nodo en una posición específica (privada)
func (list *LinkedList) getNodeAtPos(pos int) *Node {
	if pos < 0 || pos >= list.size {
		return nil
	}

	current := list.head
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	return current
}

// Establecer la posición actual al inicio de la lista
func (list *LinkedList) MoveToStart() {
	if !list.IsEmpty() {
		list.current = list.head
		list.currPos = 0
	}
}

// Establecer la posición actual al final de la lista
func (list *LinkedList) MoveToEnd() {
	if !list.IsEmpty() {
		list.current = list.tail
		list.currPos = list.size - 1
	}
}

// Mover la posición actual un paso a la izquierda, sin cambio si ya está al principio
func (list *LinkedList) Prev() {
	if list.currPos > 0 {
		list.currPos--
		list.current = list.getNodeAtPos(list.currPos)
	}
}

// Mover la posición actual un paso a la derecha, sin cambio si ya está al final
func (list *LinkedList) Next() {
	if list.currPos < list.size-1 {
		list.currPos++
		list.current = list.current.Next
	}
}

// Devolver la posición del elemento actual
func (list *LinkedList) CurrPos() int {
	return list.currPos
}

// Establecer la posición actual a "pos"
func (list *LinkedList) MoveToPos(pos int) bool {
	if pos < 0 || pos >= list.size {
		return false
	}
	list.currPos = pos
	list.current = list.getNodeAtPos(pos)
	return true
}

// Devolver true si la posición actual está al final de la lista
func (list *LinkedList) IsAtEnd() bool {
	return list.currPos == list.size-1
}

// Devolver el elemento actual
func (list *LinkedList) GetValue() (int, error) {
	if list.IsEmpty() || list.current == nil {
		return 0, ErrNoSuchElement
	}
	return list.current.Element, nil
}

// Insertar "it" en la ubicación actual
func (list *LinkedList) Insert(it int) bool {
	if list.IsEmpty() {
		list.createFirstNode(it)
		return true
	}

	newNode := NewNode(it)
	if list.currPos == 0 {
		// Insertar al inicio
		newNode.Next = list.head
		list.head = newNode
		list.current = newNode
	} else {
		// Insertar en posición actual
		prev := list.getNodeAtPos(list.currPos - 1)
		newNode.Next = prev.Next
		prev.Next = newNode
		list.current = newNode
	}

	list.size++
	// Actualizar posiciones de los elementos posteriores
	return true
}

// Agregar "it" al final de la lista
func (list *LinkedList) Append(it int) bool {
	if list.IsEmpty() {
		list.createFirstNode(it)
		return true
	}

	newNode := NewNode(it)
	list.tail.Next = newNode
	list.tail = newNode
	list.size++
	return true
}

// Remover y devolver el elemento actual
func (list *LinkedList) Remove() (int, error) {
	if list.IsEmpty() || list.current == nil {
		return 0, ErrNoSuchElement
	}

	element := list.current.Element

	if list.size == 1 {
		// Único elemento
		list.Clear()
		return element, nil
	}

	if list.currPos == 0 {
		// Remover el primer elemento
		list.head = list.head.Next
		list.current = list.head
	} else {
		// Remover elemento en posición actual
		prev := list.getNodeAtPos(list.currPos - 1)
		prev.Next = list.current.Next

		if list.current == list.tail {
			list.tail = prev
		}

		// Mover current al siguiente elemento o al anterior si estamos al final
		if list.current.Next != nil {
			list.current = list.current.Next
		} else {
			list.currPos--
			list.current = prev
		}
	}

	list.size--
	return element, nil
}

// Métodos de compatibilidad - mantener para que las pruebas sigan funcionando
func (list *LinkedList) AddHead(element int) *Node {
	if list.IsEmpty() {
		return list.createFirstNode(element)
	}

	newNode := NewNode(element)
	newNode.Next = list.head
	list.head = newNode
	list.size++

	// Actualizar current si estaba apuntando al head anterior
	if list.current == list.head.Next {
		list.currPos++
	}

	return newNode
}

func (list *LinkedList) AddTail(element int) *Node {
	return list.appendNode(element)
}

func (list *LinkedList) appendNode(element int) *Node {
	if list.IsEmpty() {
		return list.createFirstNode(element)
	}

	newNode := NewNode(element)
	list.tail.Next = newNode
	list.tail = newNode
	list.size++
	return newNode
}

// Métodos de compatibilidad (para mantener pruebas funcionando)
func (list *LinkedList) RemoveHead() (int, error) {
	if list.head == nil {
		return 0, ErrEmptyList
	}

	element := list.head.Element
	list.head = list.head.Next

	// actualizacion de los punteros
	if list.head == nil {
		list.tail = nil
		list.current = nil
		list.currPos = 0
	} else {
		// Actualizar current si estaba apuntando al head que se eliminó
		if list.currPos == 0 {
			list.current = list.head
		} else {
			list.currPos--
		}
	}
	list.size--
	return element, nil
}

func (list *LinkedList) RemoveTail() (int, error) {
	if list.IsEmpty() {
		return 0, ErrEmptyList
	}

	if list.head == list.tail {
		element := list.head.Element
		list.Clear()
		return element, nil
	}

	// Encontrar el nodo anterior al tail
	previous := list.head
	for previous.Next != list.tail {
		previous = previous.Next
	}

	element := list.tail.Element
	previous.Next = nil
	list.tail = previous
	list.size--

	// Actualizar current si estaba apuntando al tail
	if list.current == nil || list.currPos >= list.size {
		list.currPos = list.size - 1
		list.current = list.tail
	}

	return element, nil
}

func (list *LinkedList) RemoveByValue(element int) error {
	if list.IsEmpty() {
		return ErrEmptyList
	}

	// Si es el primer elemento
	if list.head.Element == element {
		_, err := list.RemoveHead()
		return err
	}

	// Si es el último elemento
	if list.tail.Element == element {
		_, err := list.RemoveTail()
		return err
	}

	// Buscar en el medio
	current := list.head
	pos := 0
	for current.Next != nil && current.Next.Element != element {
		current = current.Next
		pos++
	}

	if current.Next != nil {
		// Actualizar current position si es necesario
		if list.currPos > pos {
			list.currPos--
		} else if list.currPos == pos+1 {
			list.current = current
			list.currPos = pos
		}

		current.Next = current.Next.Next
		if current.Next == nil {
			list.tail = current
		}
		list.size--
	}
	return nil
}

// Mostrar los elementos de la lista enlazada
func (list *LinkedList) ToString() string {
	if list.IsEmpty() {
		return ""
	}

	result := ""
	current := list.head
	for current != nil {
		result += fmt.Sprintf("%v -> ", current.Element)
		current = current.Next
	}

	return result
}

func (list *LinkedList) GetHead() int {
	if list.head == nil {
		return 0
	}
	return list.head.Element
}

func (list *LinkedList) GetTail() int {
	if list.tail == nil {
		return 0
	}
	return list.tail.Element
}
