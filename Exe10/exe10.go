package main

//uma função para dizer se o numero inserido é primo ou n
import "fmt"

func main() {
	numPrimo()
}

func numPrimo() {
	fmt.Println("Digite qual numero deseja verificar se é primo: ")
	var num1 int
	fmt.Scan(&num1)
	var result int

	for i := 2; i < num1; i++ {
		if num1%i == 0 {
			result++
		}
	}
	if result == 0 {
		fmt.Println("O numero", num1, "é primo")
	} else {
		fmt.Println("O numero", num1, "não é primo")
	}
}
