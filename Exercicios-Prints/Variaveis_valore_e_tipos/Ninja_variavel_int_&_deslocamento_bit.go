package main

import (
	"fmt"
)

var x int = 200 //Variável do tipo int

func main() {
	fmt.Printf("%d\n%b\n%#x\n\n", x, x, x) //Demonstra os valores em decimal, binário e hexadecimal
	y := x >> 1                            //Desloca >> um bit para a esquerda
	fmt.Printf("%d\n%b\n%#x", y, y, y)
}
