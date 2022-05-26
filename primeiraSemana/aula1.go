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

func exercicio3() {
	var nome string      //retirado o 1 da frente do nome
	var sobrenome string //estava certa
	var idade int        //coloquei o tipo na frente do nome
	sobrenome = "6"      //tirei o 1 na frente do sobrenome
	//tirei os dois pontos
	//coloquei aspas no 6
	var licenca_para_dirigir = true //estava correta
	var estaturaDaPessoa int        //tirei os espaços
	quantidadeDeFilhos := 2         //estava correta

	nome = "Sheila"
	idade = 20
	estaturaDaPessoa = 200

	fmt.Printf("%s %s %v %v %d %d", nome, sobrenome, idade, licenca_para_dirigir, estaturaDaPessoa, quantidadeDeFilhos)
}

func main() {

	exercicio1()
	fmt.Println()
	exercicio2()
	fmt.Println()
	exercicio3()
	fmt.Println()
}
