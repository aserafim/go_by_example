package main

import (
	"fmt"
	"errors"
)

var DivideByZero = errors.New("division by zero")

func Divide(a, b float32) (ret float32, err error) {

	if b == 0 {
		return 0, DivideByZero
	}

	return a / b, nil
}

func DefineIMC(peso float32, altura float32) (result float32, err error) {
	
}

func main() {

	fmt.Println(Divide(2,3))

}