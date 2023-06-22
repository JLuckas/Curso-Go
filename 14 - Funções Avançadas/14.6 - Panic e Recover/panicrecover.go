package main

import "fmt"

func recuperaExec() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoSitu(n1, n2 float32) bool {
	defer recuperaExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É 6!!")
}

func main() {
	fmt.Println(alunoSitu(6, 6))
	fmt.Println("Pós execução!")
}