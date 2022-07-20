package stack

type Stack interface {
	Clear()
	Push(e int)
	Top()
	Pop()
}