//Criar uma função que faça média entre notas

package main

import "fmt"

func main() {
	mediaNotas()
}

func mediaNotas() {
	fmt.Println("Digite as notas deseja fazer media:")
	var nota1 float32
	var nota2 float32
	var media float32
	fmt.Print("Nota 1: ")
	fmt.Scan(&nota1)
	fmt.Print("Nota 2: ")
	fmt.Scan(&nota2)
	media = (nota1 + nota2) / 2

	fmt.Println("A media das notas é:", media)

}
