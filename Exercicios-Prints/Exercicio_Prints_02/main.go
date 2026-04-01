package main

import (
	"fmt" // Importação do pacote "fmt"
)

var x int    //O Var é sempre declarado em package-level scope
var y string // E é daclarado usando o sinal de atribuição "="
var z bool   // Nem sempre precisa receber algum valor declarado

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z) //Sempre que for usar formatação no print, precisa de ser "Printf"
}

//Quando a variável não recebe valores, o compilador atribui um valor padrão, chamado valor "ZERO".
