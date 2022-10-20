// Escrever algoritmo em Go que receba a data de nascimento do usuário como entrada,  e retorne o dia da semana que o usuário nasceu.
package main

import (
	"fmt"
	"time"
)

func main() {
	diaSemana()
}

func diaSemana() {
	dia := 0
	mes := 0
	ano := 0

	fmt.Println("Digite sua data de nascimento em forma numerica ex 01/01/1999")

	fmt.Print("Dia: ")
	fmt.Scan(&dia)

	fmt.Print("Mes: ")
	fmt.Scan(&mes)

	fmt.Print("Ano: ")
	fmt.Scan(&ano)

	data := time.Date(ano, time.Month(mes), dia, 00, 00, 00, 0, time.Local).Weekday()

	fmt.Print("O dia da semana que voce nasceu foi: ")

	if data.String() == "Monday" {
		fmt.Println("Segunda-feira")
	}
	if data.String() == "Tuesday" {
		fmt.Println("Terça-feira")
	}
	if data.String() == "Wednesday" {
		fmt.Println("Quarta-feira")
	}
	if data.String() == "Thursday" {
		fmt.Println("Quinta-feira")
	}
	if data.String() == "Friday" {
		fmt.Println("Sexta-feira")
	}
	if data.String() == "Saturday" {
		fmt.Println("Sábado")
	}
	if data.String() == "Sunday" {
		fmt.Println("Domingo")
	}
}
