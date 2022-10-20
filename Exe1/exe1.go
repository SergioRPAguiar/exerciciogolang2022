package main

import "fmt"

func main() {
	arvoreDeNatal()
}

//Criar um laço que crie com asterisco uma arvore de natal
func arvoreDeNatal() {
	fmt.Println("Digite com quantas linhas dejesa que o corpo da arvore tenha:")
	var ast string = "*"
	var esp string = " "
	var a int
	fmt.Scan(&a)
	var h int = a - 1

	for x := 0; x < a; x++ {
		for y := 0; y < h; y++ {
			fmt.Print(esp)

		}
		fmt.Println(ast)
		ast = ast + "**"
		h -= 1

	}

	ast = "*"

	for i := 0; i < a; i += 4 {
		for y := 1; y < a; y++ {
			fmt.Print(esp)
		}
		fmt.Println(ast)
	}

}
