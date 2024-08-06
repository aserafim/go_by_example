package main

import "fmt"

// variáveis de escopo global
// ou package level
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)

	//variáveis locais
	f := "apple"
	fmt.Println(f)
}
