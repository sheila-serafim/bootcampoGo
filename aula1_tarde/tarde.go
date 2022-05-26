package main

import "fmt"

func letrasDeUmaPalavra() {

	palavra := "soletrando"

	for i := 0; i < len(palavra); i++ {

		fmt.Printf("%c \n", palavra[i])
	}
	fmt.Println(len(palavra))
}

func emprestimo(idade int, empregado bool, tempoAtividade int, salario float32) {

	if idade > 22 && empregado && tempoAtividade > 1 {
		if salario > 100000 {
			fmt.Println("Empréstimo concedido sem juros")
		} else {
			fmt.Println("Empréstimo concedido com juros.")
		}
	} else {
		fmt.Println("Empréstimo não concedido.")
	}
}

func aQueMesCorresponde(mes int) {

	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês não encontrado")
	}
}

func main() {

	letrasDeUmaPalavra()
	fmt.Println()
	emprestimo(23, true, 2, 200000)
	fmt.Println()
	aQueMesCorresponde(2)
}
