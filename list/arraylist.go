package list

// ArrayList representa una implementación de arreglo dinámico de IList
type ArrayList struct {
	elements []int
	size     int
	capacity int
	current  int // posición actual del cursor
}

// NewArrayList crea un nuevo ArrayList
func NewArrayList() *ArrayList {
	initialCapacity := 10
	return &ArrayList{
		elements: make([]int, initialCapacity),
		size:     0,
		capacity: initialCapacity,
		current:  0,
	}
}

// NewArrayListWithCapacity crea un nuevo ArrayList con capacidad inicial especificada
func NewArrayListWithCapacity(capacity int) *ArrayList {
	if capacity < 0 {
		capacity = 10
	}
	return &ArrayList{
		elements: make([]int, capacity),
		size:     0,
		capacity: capacity,
		current:  0,
	}
}

// Capacity devuelve la capacidad actual de la lista
func (al *ArrayList) Capacity() int {
	return al.capacity
}

// ensureCapacity duplica la capacidad cuando es necesario
func (al *ArrayList) ensureCapacity() {
	if al.size >= al.capacity {
		al.capacity *= 2
		newElements := make([]int, al.capacity)
		copy(newElements, al.elements)
		al.elements = newElements
	}
}

// Clear elimina todo el contenido de la lista, para que esté vacía de nuevo
func (al *ArrayList) Clear() {
	al.size = 0
	al.current = 0
}

// Insert inserta "it" en la ubicación actual
func (al *ArrayList) Insert(it int) bool {
	if al.current < 0 || al.current > al.size {
		return false
	}

	al.ensureCapacity()

	// Mover elementos hacia la derecha para hacer espacio
	for i := al.size; i > al.current; i-- {
		al.elements[i] = al.elements[i-1]
	}

	al.elements[al.current] = it
	al.size++
	return true
}

// Append agrega "it" al final de la lista
func (al *ArrayList) Append(it int) bool {
	al.ensureCapacity()
	al.elements[al.size] = it
	al.size++
	return true
}

// Remove elimina y retorna el elemento actual
func (al *ArrayList) Remove() (int, error) {
	if al.size == 0 {
		return 0, ErrEmptyList
	}
	if al.current < 0 || al.current >= al.size {
		return 0, ErrOutOfRangeList
	}

	removedValue := al.elements[al.current]

	// Mover elementos hacia la izquierda
	for i := al.current; i < al.size-1; i++ {
		al.elements[i] = al.elements[i+1]
	}

	al.size--

	// Ajustar la posición actual si es necesario
	if al.current >= al.size && al.size > 0 {
		al.current = al.size - 1
	}
	if al.size == 0 {
		al.current = 0
	}

	return removedValue, nil
}

// MoveToStart establece la posición actual al inicio de la lista
func (al *ArrayList) MoveToStart() {
	al.current = 0
}

// MoveToEnd establece la posición actual al final de la lista
func (al *ArrayList) MoveToEnd() {
	if al.size == 0 {
		al.current = 0
	} else {
		al.current = al.size - 1
	}
}

// Prev mueve la posición actual un paso a la izquierda
func (al *ArrayList) Prev() {
	if al.current > 0 {
		al.current--
	}
}

// Next mueve la posición actual un paso a la derecha
func (al *ArrayList) Next() {
	if al.current < al.size-1 {
		al.current++
	}
}

// Length retorna el número de elementos en la lista
func (al *ArrayList) Length() int {
	return al.size
}

// CurrPos retorna la posición del elemento actual
func (al *ArrayList) CurrPos() int {
	return al.current
}

// MoveToPos establece la posición actual a "pos"
func (al *ArrayList) MoveToPos(pos int) bool {
	if pos < 0 || pos >= al.size {
		return false
	}
	al.current = pos
	return true
}

// IsAtEnd retorna true si la posición actual está al final de la lista
func (al *ArrayList) IsAtEnd() bool {
	return al.current == al.size-1 || al.size == 0
}

// GetValue retorna el elemento actual
func (al *ArrayList) GetValue() (int, error) {
	if al.size == 0 {
		return 0, ErrEmptyList
	}
	if al.current < 0 || al.current >= al.size {
		return 0, ErrOutOfRangeList
	}
	return al.elements[al.current], nil
}

// IsEmpty retorna true si la lista está vacía
func (al *ArrayList) IsEmpty() bool {
	return al.size == 0
}
