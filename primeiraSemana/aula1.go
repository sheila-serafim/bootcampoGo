package main

import "fmt"

func exercicio1() {

	var nome string = "Sheila"
	var idade int = 36

	fmt.Println("Nome:	", nome)
	fmt.Println("Idade:	", idade)
}

func exercicio2() {

	var umidade int = 40
	var temperatura float32 = 21.5
	var pressao int = 10

	fmt.Printf("Temperatura: %.2f\n", temperatura)
	fmt.Printf("Umidade: %d\n", umidade)
	fmt.Printf("Pressão: %d\n", pressao)
}

func main() {

	exercicio1()
	fmt.Println()
	exercicio2()
}
