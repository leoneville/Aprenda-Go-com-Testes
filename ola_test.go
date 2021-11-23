package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Neville")
		esperado := "Olá, Neville"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrao para 'string' vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
