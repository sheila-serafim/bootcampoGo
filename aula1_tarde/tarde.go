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

func main() {

	letrasDeUmaPalavra()
	fmt.Println()
	emprestimo(23, true, 2, 200000)
}
