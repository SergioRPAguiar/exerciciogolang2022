// Faça um programa que mostra na tela os número de 1 a 100 e print em uma matriz quais deles sao numero de Ouro ?
package main

import "fmt"

func main() {
	fibonacci()
}

func fibonacci() {
	var ant int
	var armazena [11]int
	var numeros [100]int
	fibo := 11
	prox := 1
	result := 1

	for i := 0; i < 100; i++ {
		numeros[i] = i + 1
	}
	fmt.Println("Dentre os numeros", numeros)

	for i := 0; i < fibo; i++ {
		armazena[i] = result
		result = ant + prox
		ant = prox
		prox = result

	}
	fmt.Println("esses sao numeros de Ouro", armazena)
}
