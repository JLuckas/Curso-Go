package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada!")
	fmt.Println("Verificando se o aluno está aprovado...")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {

	fmt.Println(alunoEstaAprovado(7,4))
}