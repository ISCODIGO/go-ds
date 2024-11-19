package queue

import "testing"

const mensaje_test string = "Esperado: %v / Obtenido: %v"

func TestArrayQueue_empty(t *testing.T) {
	cola := New(10)
	_, err := cola.Front()

	if cola.Size() != 0 {
		t.Errorf(mensaje_test, 0, cola.Size())
	}

	if err == nil {
		t.Error("Debe existir un error")
	}

}

func TestArrayQueue_add(t *testing.T) {
	cola := New(3)
	cola.Enqueue(10)
	cola.Enqueue(20)
	cola.Enqueue(30)

	// 10 (frente) | 20 | 30

	if cola.Size() != 3 {
		t.Errorf(mensaje_test, 3, cola.Size())
	}

	frente, _ := cola.Front()
	if frente != 10 {
		t.Errorf(mensaje_test, 10, frente)
	}

	err := cola.Enqueue(40)
	if err == nil {
		t.Error("cola debe estar llena")
	}
}

func TestArrayQueue_remove(t *testing.T) {
	cola := New(2)
	cola.Enqueue(10)
	cola.Enqueue(20)

	// 10 (frente) | 20

	removido1, _ := cola.Dequeue()
	if removido1 != 10 {
		t.Errorf(mensaje_test, 10, removido1)
	}

	if cola.Size() != 1 {
		t.Errorf(mensaje_test, 1, cola.Size())
	}

	// 20 |
	frente, _ := cola.Front()
	if frente != 20 {
		t.Errorf(mensaje_test, 20, frente)
	}

	// 20 |
	removido2, _ := cola.Dequeue()
	if removido2 != 20 {
		t.Errorf(mensaje_test, 20, removido2)
	}

	// cola vacia
	_, err := cola.Dequeue()
	if err == nil {
		t.Error("cola debe estar vacia")
	}
}
