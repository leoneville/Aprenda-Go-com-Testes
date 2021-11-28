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
	testesArea := []struct { // struct anônima
		forma    Forma
		esperado float64
	}{
		{Retangulo{12, 6}, 72.0},
		{Circulo{10}, 314.1592653589793},
	}

	for _, tableTests := range testesArea {
		resultado := tableTests.forma.Area()
		if resultado != tableTests.esperado {
			t.Errorf("resultado: '%.2f', esperado: '%.2f'", resultado, tableTests.esperado)
		}
	}
}
