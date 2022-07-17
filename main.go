package main

import (
	"fmt"
	"07_listas/array"
	"07_listas/linked"
)

func main() {
	lista := array.NewListaArray()
	lista.Append(10)
	lista.Append(20)
	lista.Append(30)
	lista.Insert(5)
	lista.Next()
	lista.Insert(4)
	fmt.Println(lista)

	node := linked.NewNode(10)
	fmt.Println(node)
}