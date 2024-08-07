package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

// Função que divide dois inteiros
// retorna um inteiro (resultado)
// e um erro
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	num1, num2 := 10, 10
	result, err := Divide(num1, num2)

	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", num1, num2, result)
}
