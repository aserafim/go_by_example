package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = -5

	if x > 0 {
		fmt.Printf("%d é maior que 0.\n", x)
	} else {
		fmt.Printf("%d não é maior que 0. \n", x)
	}

	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Printf("Hoje é dia %d.\n", today.Day())
	case 6:
		fmt.Printf("Hoje é o dia %d.\n", today.Day())
		fallthrough //continua a execução
	default:
		fmt.Println("Não conseguimos identificar o dia.")
	}

	switch {
	case today.Month() <= 6:
		fmt.Println("Estamos no primeiro semestre.")
	default:
		fmt.Println("Estamos no segundo semestre.")

	}

}
