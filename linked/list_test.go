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

func crearError(t *testing.T, esperado string, obtenido interface{}) {
	t.Error("Esperado", esperado, "Obtenido", fmt.Sprint(obtenido))
}

func TestLLVacio(t *testing.T) {
	// el comportamiento de LL cuando esta vacia
	lista := crearLL(0)

	if lista.Length() != 0 {
		t.Error("La LL no debe tener nodos. Obtenido", lista.Length())
	}

	if lista.head != nil {
		t.Error("Head no es nil")
	}

	if lista.tail != nil {
		t.Error("Tail no es nil")
	}

	if lista.curr != nil {
		t.Error("Current no es nil")
	}
}

func TestLLAppend(t *testing.T) {
	lista := crearLL(3)  // 10 -> 20 -> 30x

	if lista.Length() != 3 {
		crearError(t, "Size: 3", lista.Length())
	}

	lista.MoveToStart()
	primero, _ := lista.GetValue()
	if primero != 10 {
		crearError(t, "GetValue start: 10", primero)
	}

	lista.MoveToEnd()
	ultimo, _ := lista.GetValue()
	if ultimo != 30 {
		crearError(t, "GetValue end 30", ultimo)
	}	
}

func TestLLInsert(t *testing.T) {


}

func TestClear(t *testing.T) {
	lista := crearLL(3)
	lista.Clear()

	if lista.Length() != 0 {
		crearError(t, "Length: 0", lista.Length())
	}

	if lista.head != nil {
		crearError(t, "Head: nil", lista.head)
	}

	if lista.tail != nil {
		crearError(t, "Tail: nil", lista.tail)
	}
}

func TestLLRemove(t *testing.T) {

}

func TestMove(t *testing.T) {

}