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

	if lista.head != lista.tail || lista.head != lista.curr {
		t.Error("head == tail == curr")
	}
}

func TestLLAppend(t *testing.T) {
	lista := crearLL(3)  // 10 -> 20 -> 30x

	if lista.Length() != 3 {
		crearError(t, 3, lista.Length())
	}

	lista.MoveToStart()
	primero, _ := lista.CurrentElement()
	if primero != 10 {
		crearError(t, 10, primero)
	}

	lista.MoveToEnd()
	ultimo, _ := lista.CurrentElement()
	if ultimo != 30 {
		crearError(t, 30, ultimo)
	}	
}

func TestMove(t *testing.T) {
	lista := crearLL(4)  // 10->20->30->40

	lista.MoveToStart()  // 10
	lista.Next()  // 20
	
	valor1, _ := lista.CurrentElement()
	if valor1 != 20 {
		crearError(t, 20, valor1)
	}

	lista.Next()  // 30
	lista.Next()  // 40
	lista.Next()  // 40: debe seguir siendo 40

	valor2, _ := lista.CurrentElement()
	if valor2 != 40 {
		crearError(t, 40, valor2)
	}

	lista.Prev()  // 30

	valor3, _ := lista.CurrentElement()
	if valor3 != 30 {
		crearError(t, 30, valor3)
	}

	lista.Prev()  // 20
	lista.Prev()  // 10
	lista.Prev()  // 10: debe seguir siendo 10

	valor4, _ := lista.CurrentElement()
	if valor4 != 10 {
		crearError(t, 10, valor4)
	}
}

func TestMovePos(t *testing.T) {
	lista := crearLL(5)  // 10->20->30->40->50
	lista.MoveToStart()

	lista.MoveToPos(1)
	valor1, _ := lista.CurrentElement()
	if valor1 != 20 {
		crearError(t, 20, valor1)
	}

	err := lista.MoveToPos(10)
	if err == nil {
		t.Error("MoveToPos(10): Debio generarse un error")
	}

	lista.MoveToPos(4)
	valor2, _ := lista.CurrentElement()
	if valor2 != 50 {
		crearError(t, 50, valor2)
	}
}

func TestLLInsert(t *testing.T) {
	lista := crearLL(3)  // 10->20->30

	// 11->10->20->30
	lista.MoveToStart()
	lista.Insert(11)

	if lista.Length() != 4 {
		crearError(t, 4, lista.Length())
	}

	lista.MoveToStart()  // 11
	valor1, _ := lista.CurrentElement()
	if valor1 != 11 {
		crearError(t, 11, valor1)
	}

	lista.Next()  // 10
	valor2, _ := lista.CurrentElement()
	if valor2 != 10 {
		crearError(t, 10, valor2)
	}

	// 11->10->20->30->22
	lista.MoveToEnd()  // 30
	lista.Insert(22)

	if lista.Length() != 5 {
		crearError(t, 5, lista.Length())
	}

	lista.MoveToPos(4)
	valor3, _ := lista.CurrentElement()
	if valor3 != 22 {
		fmt.Println(lista)
		crearError(t, 22, valor3)
	}

	// 1->10->20->33->30->2
	lista.MoveToPos(3)
	lista.Insert(33)

	if lista.Length() != 6 {
		crearError(t, 6, lista.Length())
	}

	lista.MoveToPos(3)
	valor4, _ := lista.CurrentElement()
	if valor4 != 33 {
		crearError(t, 33, lista.Length())
	}
}

func TestClear(t *testing.T) {
	lista := crearLL(3)
	lista.Clear()

	if lista.Length() != 0 {
		crearError(t, 0, lista.Length())
	}

	if lista.head != lista.tail {
		t.Error("head == tail : ", lista.head, ",", lista.tail)
	}
}

func TestLLRemove(t *testing.T) {
	lista := crearLL(5)  // 10->20->30->40->50

	// 20->30->40->50
	lista.MoveToStart()
	remove1, err1 := lista.Remove()

	if err1 != nil {
		crearError(t, "nil", err1)
	}

	if remove1.element != 10 {
		crearError(t, 10, remove1.element)
	}

	// 20->30->40
	lista.MoveToEnd()
	remove2, err2 := lista.Remove()

	if err2 != nil {
		crearError(t, "nil", err2)
	}

	if remove2.element != 50 {
		crearError(t, 50, remove2.element)
	}

	// 20->40
	lista.MoveToStart()  // 20
	lista.Next()  // 30

	remove3, err3 := lista.Remove()

	if err3 != nil {
		crearError(t, "nil", err3)
	}

	if remove3.element != 30 {
		crearError(t, 30, remove3.element)
	}

	// dejar sin elementos (borrar dos veces)
	lista.MoveToStart()
	remove4, _ := lista.Remove()
	remove5, _ := lista.Remove()

	if lista.Length() != 0 {
		crearError(t, 0, lista.Length())
	}

	if remove4.element != 20 {
		crearError(t, 20, remove4.element)
	}

	if remove5.element != 40 {
		crearError(t, 40, remove5.element)
	}

	// por ultimo intentar borrar con lista vacia
	_, err4 := lista.Remove()  // no importa el nodo borrado
	
	if err4 == nil {
		t.Error("Se esperaba un error pero se obtuvo nil")
	}
}
