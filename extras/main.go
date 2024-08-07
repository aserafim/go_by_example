package main

import (
	"fmt"
	"errors"
)

var DivideByZero = errors.New("division by zero")
var PesoNaoAceito = errors.New("peso deve ser maior que zero")
var AlturaNaoAceita = errors.New("altura deve ser maior que zero")

func Divide(a, b float32) (ret float32, err error) {

	if b == 0 {
		return 0, DivideByZero
	}

	return a / b, nil
}

func calculaIMC(peso float64, altura float64) (res float64, err error) {

	if peso <= 0 {
		return 0, PesoNaoAceito
	}

	if altura <= 0 {
		return 0, AlturaNaoAceita
	}

	var imc = peso / (altura * altura)

	return imc, nil

}

func main() {

	fmt.Println(Divide(2,3))

	altura := 1.83
	peso := 60.0
	var resultado,_ = calculaIMC(peso, altura)
	fmt.Printf("O IMC é: %.2f.",resultado)

}