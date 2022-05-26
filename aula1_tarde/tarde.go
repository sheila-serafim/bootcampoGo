package main

import "fmt"

func letrasDeUmaPalavra() {

	palavra := "soletrando"

	for i := 0; i < len(palavra); i++ {

		fmt.Printf("%c \n", palavra[i])
	}
	fmt.Println(len(palavra))
}

func main() {

	letrasDeUmaPalavra()
}
