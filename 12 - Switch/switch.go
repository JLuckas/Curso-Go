package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número da semana é inválido"
	}
}

func diaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough //caso entre neste case o fallthrough manda para o próximo case, no caso para o retorno do próximo case
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número da semana é inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaSemana(60)
	fmt.Println(dia)

	fmt.Println("...............................")

	dia1 := diaSemana2(1)
	fmt.Println(dia1)
}
