package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//O primeiro parametro de map é do tipo da chave(Ex.: se tiver como string, a chave tem que ser do tipo string, se tiver como int, a chave tem que ser um numero inteiro)
	//e o segundo é o tipo de dado do map.
	usuario := map[string]string{
		"nome":      "Kléber",
		"sobrenome": "Chicletinho",
	}

	fmt.Println(usuario["nome"], usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Pablo",
			"sobrenome": "Guiñazu",
		},
		"profissao": {
			"tipo":   "Jogador de futebol",
			"função": "Volante",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"], usuario2["profissao"]["tipo"])
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["nome"] = map[string]string{
		"primeiro":  "Pablo",
		"sobrenome": "Guiñazu",
	}

	fmt.Println(usuario2)
}
