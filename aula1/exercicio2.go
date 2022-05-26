package main

import "fmt"

var umidade int = 40
var temperatura float32 = 21.5
var pressao int = 10

func main() {
	fmt.Printf("Temperatura: %.2f\n", temperatura)
	fmt.Printf("Umidade: %d\n", umidade)
	fmt.Printf("Pressão: %d\n", pressao)
}
