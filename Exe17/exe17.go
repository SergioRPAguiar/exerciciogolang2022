package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	ordemAlfabetica()
}

func ordemAlfabetica() {
	fmt.Println("Digite quantos nomes deseja digitar para colocar em ordem alfabética:")
	num1 := 0
	fmt.Scan(&num1)
	nomes := make([]string, num1)

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Println("Digite os", num1, "nomes:")
	for i := 0; i < num1; i++ {

		fmt.Print("Nome", i+1, ": ")
		scan.Scan()
		armazena := scan.Text()

		nomes[i] = armazena

	}

	fmt.Println("Os nomes na ordem que você digitou: ", nomes)
	sort.Strings(nomes)

	fmt.Println("Os nomes em ordem alfabetica: ", nomes)
}
