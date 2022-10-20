// Crie uma função que some duas notas
package main

import "fmt"

func main() {
	somaNotas()
}

func somaNotas() {
	fmt.Println("Digite as duas notas que deseja somar:")
	var nota1 float32
	var nota2 float32
	var soma float32
	fmt.Print("Nota 1: ")
	fmt.Scan(&nota1)
	fmt.Print("Nota 2: ")
	fmt.Scan(&nota2)

	soma = nota1 + nota2

	fmt.Println("A soma das notas é", soma)

}
