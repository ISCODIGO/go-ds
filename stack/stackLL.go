package stack

import "go-ds/list"

type StackLL struct {
	data *list.LinkedList
}

func New(capacidad int) StackLL {
	return StackLL{
		data: list.NewLinkedList(),
	}
}

func (pila StackLL) IsEmpty() bool {
	return pila.data.IsEmpty()
}

func (pila StackLL) IsFull() bool {
	return false // una lista enlazada no tiene limite de tamaño
}

func (pila *StackLL) Clear() {
	pila.data.Clear()
}

func (pila *StackLL) Push(e int) error {
	pila.data.AddHead(e)
	return nil
}

func (pila StackLL) Top() (e int, err error) {
	if pila.IsEmpty() {
		err = ErrStackUnderflow
		return // devuelve e y err
	}

	headValue := pila.data.GetHead()
	return headValue, nil
}

func (pila *StackLL) Pop() (e int, err error) {
	e, err = pila.Top()

	if err != nil {
		return
	}

	pila.data.RemoveHead()
	return
}

func (pila StackLL) Size() int {
	return pila.data.Length()
}
