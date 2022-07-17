package linked

import (
	"fmt"
	"testing"
)

func crearLL(elementos int) (*LinkedList) {
	lista := NewLinkedList()

	for i := 1; i <= elementos; i++ {
		lista.Append( i * 10 )
	}

	return lista
}

func crearError(t *testing.T, esperado interface{}, obtenido interface{}) {
	t.Errorf("Esperado: %v Obtenido: %v", fmt.Sprint(esperado), fmt.Sprint(obtenido))
}

func TestLLVacio(t *testing.T) {
	// el comportamiento de LL cuando esta vacia
	lista := crearLL(0)

	if lista.Length() != 0 {
		crearError(t, 0, lista.Length())
	}

	if lista.head != nil {
		t.Error("Head debe ser nil")
	}

	if lista.tail != nil {
		t.Error("Tail debe ser nil")
	}

	if lista.curr != nil {
		t.Error("Current debe ser nil")
	}
}

func TestLLAppend(t *testing.T) {
	lista := crearLL(3)  // 10 -> 20 -> 30x

	if lista.Length() != 3 {
		crearError(t, 3, lista.Length())
	}

	lista.MoveToStart()
	primero, _ := lista.GetValue()
	if primero != 10 {
		crearError(t, 10, primero)
	}

	lista.MoveToEnd()
	ultimo, _ := lista.GetValue()
	if ultimo != 30 {
		crearError(t, 30, ultimo)
	}	
}

func TestLLInsert(t *testing.T) {


}

func TestClear(t *testing.T) {
	lista := crearLL(3)
	lista.Clear()

	if lista.Length() != 0 {
		crearError(t, 0, lista.Length())
	}

	if lista.head != nil {
		crearError(t, "nil", lista.head)
	}

	if lista.tail != nil {
		crearError(t, "nil", lista.tail)
	}
}

func TestLLRemove(t *testing.T) {

}

func TestMove(t *testing.T) {

}