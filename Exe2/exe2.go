package main

import "fmt"

func main() {
	minSeg()
}

//Criar uma função que converta segundos para minutos e segundos.
func minSeg() {
	fmt.Println("Digite quantos sec deseja converter para minutos")
	var seg int
	var min int
	var resto int
	fmt.Scan(&seg)
	min = seg / 60
	resto = seg % 60

	fmt.Println(seg, "segundos é igual a", min, "Minutos e", resto, "segundos")
}
