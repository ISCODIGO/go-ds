package array

import "testing"

const mensaje_test string = "Esperado: %v / Obtenido: %v"

func TestArrayStack_empty(t *testing.T) {
	pila := NewArrayStack(10)

	_, err := pila.Top()
	if err == nil {
		t.Error("Debe existir un error")
	}
}

func TestArrayStack_add(t *testing.T) {
	pila := NewArrayStack(3)
	pila.Push(10)
	pila.Push(20)
	pila.Push(30)

	cima, _ := pila.Top()
	if cima != 30 {
		t.Errorf(mensaje_test, 30, cima)
	}

	err := pila.Push(40)
	if err == nil {
		t.Error("Pila debe estar llena")
	}
}

func TestArrayStack_remove(t *testing.T) {
	pila := NewArrayStack(2)
	pila.Push(10)
	pila.Push(20)

	removido1, _ := pila.Pop()
	if removido1 != 20 {
		t.Errorf(mensaje_test, 20, removido1)
	}

	cima, _ := pila.Top()
	if cima != 10 {
		t.Errorf(mensaje_test, 10, cima)
	}

	removido2, _ := pila.Pop()
	if removido2 != 10 {
		t.Errorf(mensaje_test, 10, removido2)
	}

	_, err := pila.Pop()
	if err == nil {
		t.Error("Pila debe estar vacia")
	}
}

