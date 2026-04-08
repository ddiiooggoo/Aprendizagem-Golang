package main

import (
	"fmt"
)

const (
	_ = 1986 + iota //O iota é usado para declarar um bloco de constantes e usar uma sequencia esperada.
	b               //O _ é usado quando não queremos usar o valor atribuido a variável.
	c
	d
	e
)

func main() {
	fmt.Print(b, c, d, e)
}
