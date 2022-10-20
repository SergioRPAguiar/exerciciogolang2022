// Receber três números e faça a multiplicação entre eles
package main

import "fmt"

func main() {
	multiNum()
}

func multiNum() {
	fmt.Println("Digite três numero que deseja multiplicar")
	var num1 int
	var num2 int
	var num3 int

	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)

	result := num1 * num2 * num3

	fmt.Println("O resultado da multiplicação é:", result)
}
