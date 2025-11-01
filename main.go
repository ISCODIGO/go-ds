package main

import (
	"fmt"
	"go-ds/stack"
)

func validarParentesis() {
	frase := "()()()"
	pila := stack.NewArrayStack(len(frase))
	revisar := true

	for _, v := range frase {
		if v == '(' {
			pila.Push(1)
			// push a la pila
		}

		if v == ')' {
			// pop a la pila
			if pila.IsEmpty() {
				revisar = false
				break
			}
			pila.Pop()
		}
	}

	if !pila.IsEmpty() {
		revisar = false
	}

	fmt.Print("La frase ", frase, "... ")
	if revisar {
		fmt.Println("Esta bien")
	} else {
		fmt.Println("Incorrecta")
	}
}

func main() {
	validarParentesis()
}
