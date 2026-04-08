package main

import "fmt"

var a = 2 == 200
var b = 2 != 200
var c = 2 <= 200
var d = 2 < 200
var e = 2 >= 200
var f = 200 > 200

func main() {
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
