package main

import "fmt"

func main() {
	/*
		 OPERADORES ARITIMÉTICOS{
			+ = SOMA
			- = SUBTRAÇÃO
			/ = DIVISÃO
			* = MULTIPLICAÇÃO
			% = MOD(RESTO DA DIVISÃO)
		 }
	*/

	soma := 2 + 2
	subtracao := 10 - 5
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2
	fmt.Println("\n== OPERADORES ARITIMÉTICOS ==\n")
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)
	fmt.Println("\n== OPERAÇÕES COM VARIÁVEIS DE TIPOS DIFERENTES NÃO PODEM SER REALIZADAS ==\n")

	/*
		OPERAÇÕES COM TIPOS DE VARIÁVEIS DIFERENTES INT16, INT32 ETC... NÃO É POSSÍVEL POIS GOLANG É FORTEMENTE TIPADA
	*/

	var numero1 int16 = 10
	var numero2 int16 = 50
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	fmt.Println("\n== OPERADORES DE ATRIBUIÇÃO ==\n")

	/*
		OPERADORES DE ATRIBUIÇÃO
	*/

	var var1 string = "String 1"
	var2 := "String 2"

	fmt.Println(var1, var2)
	fmt.Println("\n== OPERADORES RELACIONAIS ==\n")

	/*
		OPERADORES RELACIONAIS{
			> = MAIOR QUE
			< = MENOR QUE
			== = IGUAL
			<= = MENOR OU IGUAL QUE
			>= = MAIOR OU IGUAL QUE
			!= = DIFERENTE DE
		}
	*/

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println("\n== OPERADORES LÓGICOS ==\n")
	/*
		OPERADORES LÓGICOS{
			&& = AND
			|| = OR
			! = NEGAÇÃO (INVERTE O VALOR BOOL DA VARIÁVEL)
		}
	*/

	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println("\n== OPERADORES UNÁRIOS ==\n")

	/*
		OPERADORES UNÁRIOS
	*/

	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero -= 20
	fmt.Println(numero)
	numero *= 10
	fmt.Println(numero)
	numero /= 4
	fmt.Println(numero)
	numero %= 5
	fmt.Println(numero)
	fmt.Println("\n== OPERADORES TERNÁRIOS NÃO EXISTEM EM GO, APENAS O IF()==\n")

	var texto string
	if numero > 5 {
		texto = "O número é maior que 5."
	} else {
		texto = "O número é menor que 5."
	}

	fmt.Println(texto)
	fmt.Println("\n=========================\n")
}
