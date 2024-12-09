package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()

	// Prueba AddHead
	list.AddHead(10)
	if list.head == nil || list.head.Element != 10 {
		t.Errorf("AddHead falló: se esperaba que el elemento de la cabeza fuera 10, se obtuvo %v", list.head.Element)
	}

	// Prueba AddTail
	list.AddTail(20)
	if list.tail == nil || list.tail.Element != 20 {
		t.Errorf("AddTail falló: se esperaba que el elemento de la cola fuera 20, se obtuvo %v", list.tail.Element)
	}

	// Prueba Add
	node := list.head // Nodo con valor 10
	list.Add(15, node)
	if list.head.Next == nil || list.head.Next.Element != 15 {
		t.Errorf("Add falló: se esperaba que el elemento después de la cabeza fuera 15, se obtuvo %v", list.head.Next.Element)
	}

	// Prueba RemoveHead
	list.RemoveHead()
	if list.head == nil || list.head.Element != 15 {
		t.Errorf("RemoveHead falló: se esperaba que el elemento de la cabeza fuera 15, se obtuvo %v", list.head.Element)
	}

	// Prueba RemoveTail
	list.RemoveTail()
	if list.tail == nil || list.tail.Element != 15 {
		t.Errorf("RemoveTail falló: se esperaba que el elemento de la cola fuera 15, se obtuvo %v", list.tail.Element)
	}

	// Prueba Remove elemento específico
	list.AddTail(25)
	list.Remove(15)
	if list.head == nil || list.head.Element != 25 {
		t.Errorf("Remove falló: se esperaba que el elemento de la cabeza fuera 25, se obtuvo %v", list.head.Element)
	}
}
