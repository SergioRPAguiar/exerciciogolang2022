// Criar uma função que converta anos para minutos e segundos.
package main

import "fmt"

func main() {
	anosMinSeg()
}

func anosMinSeg() {
	fmt.Println("Digite quantos anos dejesa converter para minutos e segundos")
	var anos int
	fmt.Scan(&anos)
	var dias int
	var horas int
	var minutos int
	var segundos int

	dias = anos * 365
	horas = dias * 24
	minutos = horas * 60
	segundos = minutos * 60

	fmt.Println(anos, "anos tem", minutos, "minutos")
	fmt.Println(anos, "anos tem", segundos, "segundos")
}
