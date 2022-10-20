package main

//Crie uma função recebe uma temperatura em graus celsius e converta para Fahrenheit.
import "fmt"

func main() {
	celFarh()
}

func celFarh() {
	fmt.Println("Digite quantos graus Celsius deseja converter para Fahrenheit: ")
	var num1 float64
	var fahr float64
	fmt.Scan(&num1)

	fahr = (num1 * 9 / 5) + 32

	fmt.Println(num1, "Graus Celsius é igual a", fahr, "Fahreinheit")

}
