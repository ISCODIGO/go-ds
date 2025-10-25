package list

import (
	"testing"
)

func TestLinkedListCompatibility(t *testing.T) {
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

	// Prueba RemoveHead
	removed, _ := list.RemoveHead()
	if removed != 10 {
		t.Errorf("RemoveHead falló: se esperaba 10, se obtuvo %v", removed)
	}
	if list.head == nil || list.head.Element != 20 {
		t.Errorf("RemoveHead falló: se esperaba que el elemento de la cabeza fuera 20, se obtuvo %v", list.head.Element)
	}

	// Agregar más elementos para probar RemoveTail
	list.AddHead(15)
	// Lista: 15 -> 20

	// Prueba RemoveTail
	removed, _ = list.RemoveTail()
	if removed != 20 {
		t.Errorf("RemoveTail falló: se esperaba 20, se obtuvo %v", removed)
	}
	if list.tail == nil || list.tail.Element != 15 {
		t.Errorf("RemoveTail falló: se esperaba que el elemento de la cola fuera 15, se obtuvo %v", list.tail.Element)
	}

	// Prueba Remove elemento específico
	list.AddTail(25)
	// Lista: 15 -> 25
	list.RemoveByValue(15)
	if list.head == nil || list.head.Element != 25 {
		t.Errorf("RemoveByValue falló: se esperaba que el elemento de la cabeza fuera 25, se obtuvo %v", list.head.Element)
	}
}

func TestNewListInterface(t *testing.T) {
	list := NewLinkedList()

	// Test Append
	if !list.Append(100) {
		t.Error("Append should return true")
	}
	if !list.Append(200) {
		t.Error("Append should return true")
	}
	if !list.Append(300) {
		t.Error("Append should return true")
	}

	// Test Length
	if list.Length() != 3 {
		t.Errorf("Expected length 3, got %d", list.Length())
	}

	// Test navigation
	list.MoveToStart()
	val, err := list.GetValue()
	if err != nil || val != 100 {
		t.Errorf("Expected 100 at start, got %d", val)
	}

	list.MoveToEnd()
	val, err = list.GetValue()
	if err != nil || val != 300 {
		t.Errorf("Expected 300 at end, got %d", val)
	}

	// Test position methods
	if list.CurrPos() != 2 {
		t.Errorf("Expected position 2, got %d", list.CurrPos())
	}

	if !list.IsAtEnd() {
		t.Error("Should be at end")
	}

	// Test MoveToPos
	if !list.MoveToPos(1) {
		t.Error("MoveToPos should return true")
	}
	val, _ = list.GetValue()
	if val != 200 {
		t.Errorf("Expected 200 at position 1, got %d", val)
	}

	// Test Insert
	if !list.Insert(150) {
		t.Error("Insert should return true")
	}
	if list.Length() != 4 {
		t.Errorf("Expected length 4 after insert, got %d", list.Length())
	}

	// Test Remove
	removed, err := list.Remove()
	if err != nil || removed != 150 {
		t.Errorf("Expected to remove 150, got %d", removed)
	}
	if list.Length() != 3 {
		t.Errorf("Expected length 3 after remove, got %d", list.Length())
	}

	// Test Prev and Next
	list.MoveToStart()
	list.Next()
	val, _ = list.GetValue()
	if val != 200 {
		t.Errorf("Expected 200 after Next from start, got %d", val)
	}

	list.Prev()
	val, _ = list.GetValue()
	if val != 100 {
		t.Errorf("Expected 100 after Prev, got %d", val)
	}

	// Test Clear
	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after Clear")
	}
	if list.Length() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", list.Length())
	}
}
