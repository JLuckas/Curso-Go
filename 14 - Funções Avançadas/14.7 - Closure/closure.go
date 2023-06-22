package main

import "fmt"

func closure() func() {
	texto := "Dentro da Função closure :)"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Funções Closure")

	texto := "Dentro da função main :)"

	fmt.Println(texto)

	func2 := closure()
	func2()
}
