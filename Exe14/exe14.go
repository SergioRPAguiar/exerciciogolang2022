// crie uma função que calcula o IMC
package main

import "fmt"

func main() {
	imc()
}

func imc() {
	var peso float64
	var altura float64

	fmt.Println("Digite o seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Digite a sua altura em metros: ")
	fmt.Scan(&altura)

	result := peso / (altura * altura)

	fmt.Println("O seu imc é:", result)
}
