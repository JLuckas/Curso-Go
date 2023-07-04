package enderecos

import ("strings")


// TipoEndereco = verifica se o enderaço é válido.
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecosMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecosMinusculo, " ")[0]

	enderecoTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}