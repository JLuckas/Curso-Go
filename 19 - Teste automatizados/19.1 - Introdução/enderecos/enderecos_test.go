package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado string
}
// ---------- TESTE DE UNIDADE ----------
// Funções de teste SEMPRE devem ser declaradas com a palavra "Test" no início do nome
/*func TestTipoEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida João Dutra Filho"

	tipoEnderecoEsperado := "Rua"

	tipoEnderecoRecebido := TipoEndereco(enderecoParaTeste)
	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido não é o esperado. Era esperado o tipo %s mas foi recebido o tipo %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}

}*/

//TESTE DE UNIDADE COM VÁRIOS CENÁRIOS
func TestTipoEndereco(t *testing.T) {
	t.Parallel()
	
	cenariosTeste := []cenarioTeste {
		{"Rua 1", "Rua"},
		{"Avenida 1", "Avenida"},
		{"Rodovia 1", "Rodovia"},
		{"Praça 1", "Tipo Inválido"},
		{"Estrada 1", "Estrada"},
		{"RUA 1", "Rua"},
		{"avenida 1", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1>2 {
		t.Errorf("Teste quebrou!")
	}
}