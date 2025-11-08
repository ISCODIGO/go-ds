package queue

import "testing"

func TestQueueLL_empty(t *testing.T) {
	cola := NewQueueLL()
	_, err := cola.Front()

	if cola.Size() != 0 {
		t.Errorf("Esperado: %v / Obtenido: %v", 0, cola.Size())
	}

	if err == nil {
		t.Error("Debe existir un error")
	}
}

func TestQueueLL_add(t *testing.T) {
	cola := NewQueueLL()
	cola.Enqueue(10)
	cola.Enqueue(20)
	cola.Enqueue(30)

	if cola.Size() != 3 {
		t.Errorf("Esperado: %v / Obtenido: %v", 3, cola.Size())
	}

	frente, _ := cola.Front()
	if frente != 10 {
		t.Errorf("Esperado: %v / Obtenido: %v", 10, frente)
	}
}

func TestQueueLL_remove(t *testing.T) {
	cola := NewQueueLL()
	cola.Enqueue(10)
	cola.Enqueue(20)

	removido1, _ := cola.Dequeue()
	if removido1 != 10 {
		t.Errorf("Esperado: %v / Obtenido: %v", 10, removido1)
	}

	if cola.Size() != 1 {
		t.Errorf("Esperado: %v / Obtenido: %v", 1, cola.Size())
	}

	frente, _ := cola.Front()
	if frente != 20 {
		t.Errorf("Esperado: %v / Obtenido: %v", 20, frente)
	}

	removido2, _ := cola.Dequeue()
	if removido2 != 20 {
		t.Errorf("Esperado: %v / Obtenido: %v", 20, removido2)
	}

	_, err := cola.Dequeue()
	if err == nil {
		t.Error("cola debe estar vacia")
	}
}