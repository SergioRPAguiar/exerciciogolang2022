// Crie uma função  que classifique o  nível de satisfação do cliente usando METRICA NPS
package main

import "fmt"

func main() {
	metricaNps()
}

func metricaNps() {
	fmt.Println("Digite quantas notas da empresa deseja inserir:")
	var num1 int
	var num2 int
	fmt.Scan(&num1)
	var promot int
	var neutro int
	var np int
	var detrat int
	var dp int
	var result int
	var notas = make([]int, num1)

	fmt.Println("Agora digite as notas: ")

	for i := 0; i < num1; i++ {
		fmt.Print("Nota ", i+1, ": ")
		fmt.Scan(&num2)
		if num2 >= 0 && num2 <= 10 {
			result = result + num2
			notas[i] = num2
			if num2 <= 6 {
				detrat = detrat + num2
				dp++
			}
			if num2 >= 9 {
				promot = promot + num2
				np++
			} else {
				neutro++
			}

		} else {
			fmt.Println("Nota invalida")
			i -= 1
		}

	}

	num1 = num1 - neutro

	np = (np * 10) / num1
	fmt.Println(np)
	dp = (dp * 10) / num1
	fmt.Println(dp)
	result = (np - dp) * 10
	fmt.Println(result)

	if result > 74 {
		fmt.Println("O NPS da sua empresa é", result, "considerada excelente")
	} else if result >= 50 {
		fmt.Println("O NPS da sua empresa é", result, "considerada muito bom")
	} else if result >= 0 {
		fmt.Println("O NPS da sua empresa é", result, "considerada Razoável")
	} else if result >= -100 {
		fmt.Println("O NPS da sua empresa é", result, "considerada Ruim")
	}

}
