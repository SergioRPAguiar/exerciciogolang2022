// Crie uma função que retira as vogais de uma frase
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	retiraVogais()
}

func retiraVogais() {

	fmt.Println("Digite uma frase que queira remover as vogais: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, h := range vogais {
		frase = strings.ReplaceAll(frase, h, "")
	}
	fmt.Println("A sua frase sem vogais é: ")
	fmt.Println(frase)

}
