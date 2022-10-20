// Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada
package main

import "fmt"

func main() {
	calculadora()
}

func calculadora() {
	var num1 float64
	var num2 float64
	var resul float64
	var menu int
	var cont int = 10
	var n int

	fmt.Println("Bem vindo a calculadora: ")
	for i := 0; i < cont; {
		fmt.Println("qual operação deseja fazer?")
		fmt.Println("1: soma ")
		fmt.Println("2: subtração ")
		fmt.Println("3: multiplicação")
		fmt.Println("4: divisão")

		fmt.Scan(&menu)

		if menu == 1 {
			fmt.Print("Digite o primeiro numero: ")
			fmt.Scan(&num1)
			fmt.Print("Digite o segundo numero: ")
			fmt.Scan(&num2)
			resul = num1 + num2
			fmt.Println("O resultado da soma é:", resul)

			fmt.Println("Deseja fazer outra operação?")
			fmt.Println("1: sim")
			fmt.Println("2: nao")
			fmt.Scan(&n)
			if n == 2 {
				i = cont
			}

		} else if menu == 2 {
			fmt.Print("Digite o primeiro numero: ")
			fmt.Scan(&num1)
			fmt.Print("Digite o segundo numero: ")
			fmt.Scan(&num2)
			resul = num1 - num2
			fmt.Println("O resultado da subtração é:", resul)
			fmt.Println("Deseja fazer outra operação?")
			fmt.Println("1: sim")
			fmt.Println("2: nao")
			fmt.Scan(&n)
			if n == 2 {
				i = cont
			}
		} else if menu == 3 {
			fmt.Print("Digite o primeiro numero: ")
			fmt.Scan(&num1)
			fmt.Print("Digite o segundo numero: ")
			fmt.Scan(&num2)
			resul = num1 * num2
			fmt.Println("O resultado da multiplicação é:", resul)
			fmt.Println("Deseja fazer outra operação?")
			fmt.Println("1: sim")
			fmt.Println("2: nao")
			fmt.Scan(&n)
			if n == 2 {
				i = cont
			}
		} else if menu == 4 {
			fmt.Print("Digite o primeiro numero: ")
			fmt.Scan(&num1)
			fmt.Print("Digite o segundo numero: ")
			fmt.Scan(&num2)
			resul = num1 / num2
			fmt.Println("O resultado da divisao é:", resul)
			fmt.Println("Deseja fazer outra operação?")
			fmt.Println("1: sim")
			fmt.Println("2: nao")
			fmt.Scan(&n)
			if n == 2 {
				i = cont
			}
		} else {
			fmt.Println("Digito invalido")
		}
	}
	fmt.Println("Calculadora encerrada!")
}
