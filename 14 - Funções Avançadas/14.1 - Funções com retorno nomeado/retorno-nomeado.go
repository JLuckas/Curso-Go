package main

import "fmt"

func calculos(n1, n2 int) (soma int, subs int) {
	soma = n1 + n2
	subs = n1 - n2
	return
}

func main() {
	soma, subtracao := calculos(10, 20)
	fmt.Println(soma, subtracao)
}
