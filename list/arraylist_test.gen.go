package list

import "testing"

func TestALG_Vacia(t *testing.T) {
	lista := NewArrayListGeneric[int](5)

	// Tendra un size == 0
	if lista.Length() != 0 {
		t.Error("Elementos esperados 0, obtenidos", lista.Length())
	}

}

func TestALG_Append(t *testing.T) {
	lista := NewArrayListGeneric[int](5)

	lista.Append(10)
	lista.Append(20)

	// Tendra 2 elementos
	if lista.Length() != 2 {
		t.Error("Elementos esperados 2, obtenidos", lista.Length())
	}

	lista.MoveToPos(0) // lista.MoveToStart()
	valor_1, _ := lista.CurrentElement()
	if valor_1 != 10 {
		t.Error("Esperado 10, obtenido", valor_1)
	}

	lista.MoveToPos(1)
	valor_2, _ := lista.CurrentElement()
	if valor_2 != 20 {
		t.Error("Esperado 20, obtenido", valor_2)
	}
}

func TestALG_Remove(t *testing.T) {
	lista := NewArrayListGeneric[int](5)

	lista.Append(10)
	lista.Append(20)

	lista.MoveToStart()
	removido_1, error := lista.Remove() // Remueve el primero

	if error != nil {
		t.Error("Error inesperado", error)
	}

	if removido_1 != 10 {
		t.Error("Esperado 10, obtenido", removido_1)
	}

	if lista.Length() != 1 {
		t.Error("Esperado 1 elemento, obtenido", lista.Length())
	}

	removido_2, error := lista.Remove()

	if error != nil {
		t.Error("Error inesperado:", error)
	}

	if removido_2 != 20 {
		t.Error("Esperadoo 20, obtenido", removido_2)
	}

	if lista.Length() != 0 {
		t.Error("Esperado 0, obtenido", lista.Length())
	}

	_, error = lista.Remove()

	if error == nil {
		t.Error("Debio generarse un error")
	}
}

func TestALG_Clear(t *testing.T) {
	lista := NewArrayListGeneric[int](5)

	lista.Append(10)
	lista.Append(20)
	lista.Clear()

	if lista.Length() != 0 {
		t.Error("Debia estar vacio pero hay ", lista.Length(), "elementos")
	}
}

func TestALG_Next(t *testing.T) {
	lista := NewArrayListGeneric[int](5)

	lista.Append(10)
	lista.Append(20)
	lista.Append(30)

	lista.MoveToStart()
	lista.Next()
	primero, _ := lista.CurrentElement()
	if primero != 20 {
		t.Error("Esperado 20, obtenido", primero)
	}

	lista.Next()
	segundo, _ := lista.CurrentElement()
	if segundo != 30 {
		t.Error("Esperado 30, obtenido", segundo)
	}

	lista.Next()
	tercero, _ := lista.CurrentElement()
	if tercero != 30 {
		t.Error("Esperado 30, obtenido ", tercero)
	}
}

func TestALG_Prev(t *testing.T) {

	lista := NewArrayListGeneric[int](5)

	lista.Append(10)
	lista.Append(20)
	lista.Append(30)
	lista.MoveToEnd()

	primero, _ := lista.CurrentElement()
	if primero != 30 {
		t.Error("Esperado 30, obtenido", primero)
	}

	lista.Prev()
	segundo, _ := lista.CurrentElement()
	if segundo != 20 {
		t.Error("Esperado 20, obtenido", segundo)
	}

	lista.Prev()
	tercero, _ := lista.CurrentElement()
	if tercero != 10 {
		t.Error("Esperado 10 obtenido", tercero)
	}

	// Ese es el comportamiento del libro: al llegar al principio no retrocede mas
	lista.Prev()
	tercero_2, _ := lista.CurrentElement()
	if tercero_2 != 10 {
		t.Error("Esperado 10 obtenido", tercero_2)
	}
}

func TestALG_Insert(t *testing.T) {
	lista := NewArrayListGeneric[int](5)
	lista.Append(10)
	lista.Append(30)

	lista.MoveToPos(1)
	lista.Insert(20)

	valor, _ := lista.CurrentElement()
	if lista.Length() != 3 {
		t.Error("Esperados 3 elementos, obtenido", lista.Length())
	}

	if valor != 20 {
		t.Error("Esperado 20, obtenido ", valor)
	}

	lista.MoveToEnd()
	valor, _ = lista.CurrentElement()
	if valor != 30 {
		t.Error("Esperado 30, obtenido ", valor)
	}
}

func TestALG_Find(t *testing.T) {
	lista := NewArrayListGeneric[int](5)
	lista.Append(10)  // 0
	lista.Append(20)  // 1
	lista.Append(30)  // 2

	hallado_1 := lista.Find(20)

	if hallado_1 != 1 {
		t.Error("Posicion requerida 1, obtenida", hallado_1)
	}

	hallado_2 := lista.Find(50)

	if hallado_2 != -1 {
		t.Error("Posicion requerida -1, obtenida", hallado_2)
	}
}