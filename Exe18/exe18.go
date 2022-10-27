/*A Cifra de César é a técnica mais simples usada pela criptografia. Usada por Júlio César para
a transmissão de mensagens secretas,
a cifra é baseada na troca de uma letra por outra sempre obedecendo um código (ou melhor, uma chave).
Faça um programa com base no que leu a cima que é criptografia utilizando da Cifra de César.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cifraDeCesar()
}

func cifraDeCesar() {
	fmt.Println("Digite a frase que deseja codificar usando a cifra de César:")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	frase := scan.Text()

	fmt.Println("Digite quantas letras deseja pular:")
	var pula rune
	fmt.Scan(&pula)

	for i := 0; i < 2; {
		if pula > 26 {
			pula = pula - 26
		}
		if pula < 27 && pula >= 0 {
			i = 2
		}
		if pula < 0 {
			fmt.Println("Digito Invalido")
			i = 2
		}
	}
	fmt.Println("A frase:", frase, ", depois de ser criptogafada com a cifra de César fica:")

	frase = strings.ToUpper(frase)

	for _, i := range strings.ToUpper(frase) {
		if i == 32 {
			print(" ")
		}
		if i > 64 && i < 91 {
			if i+pula > 90 {
				i = 64
			}
			i = i + pula
			fmt.Printf("%c", i)
		}
	}
}
