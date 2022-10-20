// Crie uma função que receba 10 números, faça a média e mostre o maior número
package main

import "fmt"

func main() {
	mediaMaior()
}

func mediaMaior() {
	var numeros [10]float64
	var maior float64
	var media float64

	fmt.Println("Digites 10 numeros: ")

	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
		if numeros[i] > maior {
			maior = numeros[i]
		}

		media = media + numeros[i]

	}
	media = media / 10

	fmt.Println("A media dos 10 numeros é: ", media)
	fmt.Println("O maior numero dentre os 10 é o numero: ", maior)

}
