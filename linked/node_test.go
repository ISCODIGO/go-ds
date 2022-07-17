package linked

import "testing"

func TestNodeCreate(t *testing.T) {
	n := NewNode(10)

	if n.GetElement() != 10 {
		t.Error("Esperado 20, obtenido", n.GetElement())
	}

	if n.Next() != nil {
		t.Error("Se esperaba nil, obtenido", n.Next())
	}
}