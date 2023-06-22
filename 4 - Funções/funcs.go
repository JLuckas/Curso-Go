package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int, text string) (string, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	var txt string = text

	return txt, soma, subtracao
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto referente a Função F")
	fmt.Println(resultado)

	//Todos os retornos:
	resultadosSoma, resultadosSubtracao, textoCalculos := calculosMatematicos(10, 15, "Os resultados de soma e subtração são, respectivamente:")
	fmt.Println(resultadosSoma, resultadosSubtracao, textoCalculos)

	//Ocultando um dos retornos:
	resultsSoma, textCalculos, _ := calculosMatematicos(10, 15, "Os resultados de soma e subtração são, respectivamente:")
	fmt.Println(resultsSoma, textCalculos)
}
