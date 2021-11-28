package perimetro

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado: '%.2f', esperado: '%.2f'", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	mensagemdeErro := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: '%.2f', esperado: '%.2f'", resultado, esperado)
		}
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0

		mensagemdeErro(t, resultado, esperado)
	})

	t.Run("circulos", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := 314.1592653589793

		mensagemdeErro(t, resultado, esperado)
	})
}
