package main

import (
	"fmt"
)

// Tudo o que estiver dentro de “ ficará literalmente com a formatação imposta.
//Tipo de string raw literal, fica exatamente com a formatação do texto.

func main() {
	x := `Diogo Bastos 				
			42
		110 quilos`
	fmt.Print(x)
}
