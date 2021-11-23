package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Neville")
	esperado := "Olá, Neville"

	if resultado != esperado {
		t.Errorf("Resultado '%s', esperado '%s'", resultado, esperado)
	}
}
