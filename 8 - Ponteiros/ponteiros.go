package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 500
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
