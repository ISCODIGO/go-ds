package linked

type Node struct {
	element int  // dato almacenado
	next *Node   // enlace al siguiente nodo
}

// Funcion que permite inicializar un objeto de tipo Node
func NewNode(element int) (*Node) {
	return &Node{
		element: element,
		next: nil,
	}
}

func (node *Node) SetElement(element int) {
	node.element = element
}

func (node *Node) SetNext(n *Node) {
	node.next = n
}

func (node Node) GetElement() (int) {
	return node.element
}

func (node Node) Next() (*Node) {
	return node.next
}