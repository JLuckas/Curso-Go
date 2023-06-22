package main

import "fmt"

func inversor(n1 int) int {
	return n1 * -1
}

func inverteSinalPonteiro(n2 *int) {
	*n2 = *n2 * -1
}

func main() {
	fmt.Println("Ponteiros")

	n1 := 50
	inverteSinal := inversor(n1)

	fmt.Println(inverteSinal)
	fmt.Println(n1)

	novoNumero := 30
	fmt.Println(novoNumero)
	inverteSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
