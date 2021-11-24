package soma

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	mensagemdeErro := func(t *testing.T, resultado, esperado int, numeros []int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	}

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		mensagemdeErro(t, resultado, esperado, numeros[:])
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		mensagemdeErro(t, resultado, esperado, numeros)
	})
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado: '%v', esperado: '%v'", resultado, esperado)
	}
}
