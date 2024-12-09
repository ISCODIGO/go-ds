package stack

import (
	"testing"
)

func TestStackLL_EmptyStack(t *testing.T) {
	stack := New(0)

	// La pila debería estar vacía al inicio
	if !stack.IsEmpty() {
		t.Errorf("La pila debería estar vacía al ser recién creada")
	}

	// Top en una pila vacía debería retornar error
	if _, err := stack.Top(); err == nil {
		t.Errorf("Top debería retornar error en una pila vacía")
	}

	// Pop en una pila vacía debería retornar error
	if _, err := stack.Pop(); err == nil {
		t.Errorf("Pop debería retornar error en una pila vacía")
	}
}

func TestStackLL_PushAndTop(t *testing.T) {
	stack := New(0)

	// Insertar un elemento
	err := stack.Push(10)
	if err != nil {
		t.Errorf("No se esperaba error al hacer Push: %v", err)
	}

	// Verificar Top
	val, err := stack.Top()
	if err != nil {
		t.Errorf("No se esperaba error al hacer Top: %v", err)
	}

	if val != 10 {
		t.Errorf("Se esperaba Top() = 10, se obtuvo %d", val)
	}

	if stack.IsEmpty() {
		t.Errorf("La pila no debería estar vacía después de hacer Push")
	}

	// Insertar otro elemento
	err = stack.Push(20)
	if err != nil {
		t.Errorf("No se esperaba error al hacer Push(20): %v", err)
	}

	val, err = stack.Top()
	if err != nil {
		t.Errorf("No se esperaba error al hacer Top luego de Push(20): %v", err)
	}

	if val != 20 {
		t.Errorf("Se esperaba Top() = 20, se obtuvo %d", val)
	}
}

func TestStackLL_Pop(t *testing.T) {
	stack := New(0)

	// Agregamos elementos
	stack.Push(5)
	stack.Push(15)
	stack.Push(25)

	// Probamos el tamaño
	if stack.Size() != 3 {
		t.Errorf("Se esperaba tamaño 3, se obtuvo %d", stack.Size())
	}

	// Probamos Pop
	val, err := stack.Pop()
	if err != nil {
		t.Errorf("No se esperaba error al hacer Pop: %v", err)
	}
	if val != 25 {
		t.Errorf("Se esperaba valor 25 al hacer Pop, se obtuvo %d", val)
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("No se esperaba error al hacer Pop: %v", err)
	}
	if val != 15 {
		t.Errorf("Se esperaba valor 15 al hacer Pop, se obtuvo %d", val)
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("No se esperaba error al hacer Pop: %v", err)
	}
	if val != 5 {
		t.Errorf("Se esperaba valor 5 al hacer Pop, se obtuvo %d", val)
	}

	// Ahora la pila está vacía, Pop debería retornar error
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Se esperaba error al hacer Pop en una pila vacía")
	}
}

func TestStackLL_Clear(t *testing.T) {
	stack := New(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Limpiamos la pila
	stack.Clear()

	if !stack.IsEmpty() {
		t.Errorf("La pila debería estar vacía después de Clear")
	}

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Se esperaba error al hacer Pop en una pila vacía después de Clear")
	}
}

func TestStackLL_IsFull(t *testing.T) {
	stack := New(0)

	// Dado que es una lista enlazada, IsFull siempre debería retornar falso
	if stack.IsFull() {
		t.Errorf("IsFull no debería ser true para una lista enlazada")
	}

	stack.Push(10)
	if stack.IsFull() {
		t.Errorf("IsFull no debería ser true incluso después de hacer Push")
	}
}

func TestStackLL_Size(t *testing.T) {
	stack := New(0)
	if stack.Size() != 0 {
		t.Errorf("Tamaño inicial esperado 0, obtenido %d", stack.Size())
	}

	stack.Push(100)
	stack.Push(200)
	if stack.Size() != 2 {
		t.Errorf("Se esperaba tamaño 2, obtenido %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Se esperaba tamaño 1 después de Pop, obtenido %d", stack.Size())
	}
}
