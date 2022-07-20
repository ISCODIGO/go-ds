package list

type List interface {
	Clear()
	Length()
	Append(e int)
	Insert(e int)
	Remove()
	MoveToStart()
	MoveToEnd()
	MoveToPos(pos int)
	Prev()
	Next()
	CurrentPosition()
	CurrentElement()
}