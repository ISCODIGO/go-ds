package queue

import "testing"

func TestPriorityQueue_empty(t *testing.T) {
	pq := NewPriorityQueue(3) // Create with 3 priority levels
	_, err := pq.Front()

	if pq.Size() != 0 {
		t.Errorf("Esperado: %v / Obtenido: %v", 0, pq.Size())
	}

	if err == nil {
		t.Error("Debe existir un error")
	}
}

func TestPriorityQueue_add(t *testing.T) {
	pq := NewPriorityQueue(3) // Create with 3 priority levels
	
	// Add elements with different priorities (0=low, 1=medium, 2=high)
	pq.EnqueueWithPriority(10, 0) // Low priority
	pq.EnqueueWithPriority(30, 2) // High priority
	pq.EnqueueWithPriority(20, 1) // Medium priority

	// Size should be 3
	if pq.Size() != 3 {
		t.Errorf("Esperado: %v / Obtenido: %v", 3, pq.Size())
	}

	// Highest priority should be returned first (30 with priority 2)
	frente, _ := pq.Front()
	if frente != 30 {
		t.Errorf("Esperado: %v / Obtenido: %v", 30, frente)
	}
}

func TestPriorityQueue_remove(t *testing.T) {
	pq := NewPriorityQueue(3) // Create with 3 priority levels
	
	// Add elements with different priorities (0=low, 1=medium, 2=high)
	pq.EnqueueWithPriority(10, 0) // Low priority
	pq.EnqueueWithPriority(30, 2) // High priority
	pq.EnqueueWithPriority(20, 1) // Medium priority

	// Dequeue should return 30 first (highest priority)
	removido1, _ := pq.Dequeue()
	if removido1 != 30 {
		t.Errorf("Esperado: %v / Obtenido: %v", 30, removido1)
	}

	if pq.Size() != 2 {
		t.Errorf("Esperado: %v / Obtenido: %v", 2, pq.Size())
	}

	// Now front should be 20 (medium priority)
	frente, _ := pq.Front()
	if frente != 20 {
		t.Errorf("Esperado: %v / Obtenido: %v", 20, frente)
	}

	// Dequeue 20 (medium priority)
	removido2, _ := pq.Dequeue()
	if removido2 != 20 {
		t.Errorf("Esperado: %v / Obtenido: %v", 20, removido2)
	}

	// Now 10 (low priority)
	removido3, _ := pq.Dequeue()
	if removido3 != 10 {
		t.Errorf("Esperado: %v / Obtenido: %v", 10, removido3)
	}

	// Empty
	_, err := pq.Dequeue()
	if err == nil {
		t.Error("cola debe estar vacia")
	}
}

func TestPriorityQueue_defaultPriority(t *testing.T) {
	pq := NewPriorityQueue(5) // Create with 5 priority levels
	
	// Use the default Enqueue method which should use middle priority (2)
	pq.Enqueue(10)
	pq.Enqueue(20)
	
	// Add something with explicit higher priority
	pq.EnqueueWithPriority(30, 4)
	
	// Highest priority should come out first
	val, _ := pq.Dequeue()
	if val != 30 {
		t.Errorf("Esperado: %v / Obtenido: %v", 30, val)
	}
	
	// Then the default priority items in FIFO order
	val, _ = pq.Dequeue()
	if val != 10 {
		t.Errorf("Esperado: %v / Obtenido: %v", 10, val)
	}
	
	val, _ = pq.Dequeue()
	if val != 20 {
		t.Errorf("Esperado: %v / Obtenido: %v", 20, val)
	}
}

func TestPriorityQueue_clear(t *testing.T) {
	pq := NewPriorityQueue(3) // Create with 3 priority levels
	pq.EnqueueWithPriority(10, 0)
	pq.EnqueueWithPriority(20, 2)
	pq.Clear()

	if pq.Size() != 0 {
		t.Errorf("Esperado: %v / Obtenido: %v", 0, pq.Size())
	}
}
