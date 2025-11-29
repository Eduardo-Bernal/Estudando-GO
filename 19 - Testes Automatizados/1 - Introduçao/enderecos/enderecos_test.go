package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel() //Rodar as funções em paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel() //Rodar as funções em paralelo

	if 1 > 2 {
		t.Errorf("teste quebrou!")
	}
}

// 	enderecoParaTeste := "Avenida Paulista"
// TipoDeEnderecoEsperado := "Avenida"
// TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if TipoDeEnderecoEsperado != TipoDeEnderecoRecebido {
// 		t.Error("O tipo Recebido e diferente do esperado!")
// 	}

// }
