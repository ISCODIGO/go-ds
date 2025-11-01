package list

import (
	"testing"
)

func TestArrayListBasicOperations(t *testing.T) {
	list := NewArrayList()

	// Test IsEmpty
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	// Test Append
	if !list.Append(10) {
		t.Error("Append should succeed")
	}
	if !list.Append(20) {
		t.Error("Append should succeed")
	}
	if !list.Append(30) {
		t.Error("Append should succeed")
	}

	// Test Length
	if list.Length() != 3 {
		t.Errorf("Expected length 3, got %d", list.Length())
	}

	// Test IsEmpty after adding elements
	if list.IsEmpty() {
		t.Error("List should not be empty after adding elements")
	}

	// Test GetValue at start
	list.MoveToStart()
	val, err := list.GetValue()
	if err != nil {
		t.Errorf("GetValue failed: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// Test Next and GetValue
	list.Next()
	val, err = list.GetValue()
	if err != nil {
		t.Errorf("GetValue failed: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	// Test MoveToEnd
	list.MoveToEnd()
	val, err = list.GetValue()
	if err != nil {
		t.Errorf("GetValue failed: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	// Test IsAtEnd
	if !list.IsAtEnd() {
		t.Error("Should be at end")
	}
}

func TestArrayListInsert(t *testing.T) {
	list := NewArrayList()
	list.Append(10)
	list.Append(30)

	// Insert at position 1
	list.MoveToPos(1)
	if !list.Insert(20) {
		t.Error("Insert should succeed")
	}

	// Check sequence: 10, 20, 30
	list.MoveToStart()
	val, _ := list.GetValue()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	list.Next()
	val, _ = list.GetValue()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	list.Next()
	val, _ = list.GetValue()
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
}

func TestArrayListRemove(t *testing.T) {
	list := NewArrayList()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	// Remove middle element
	list.MoveToPos(1)
	val, err := list.Remove()
	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected removed value 20, got %d", val)
	}

	// Check length
	if list.Length() != 2 {
		t.Errorf("Expected length 2, got %d", list.Length())
	}

	// Check remaining elements
	list.MoveToStart()
	val, _ = list.GetValue()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	list.Next()
	val, _ = list.GetValue()
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
}

func TestArrayListClear(t *testing.T) {
	list := NewArrayList()
	list.Append(10)
	list.Append(20)

	list.Clear()

	if !list.IsEmpty() {
		t.Error("List should be empty after Clear")
	}

	if list.Length() != 0 {
		t.Errorf("Expected length 0, got %d", list.Length())
	}

	if list.CurrPos() != 0 {
		t.Errorf("Expected current position 0, got %d", list.CurrPos())
	}
}

func TestArrayListInterfaceCompliance(t *testing.T) {
	// Verify that ArrayList implements List interface
	var _ List = (*ArrayList)(nil)
}
