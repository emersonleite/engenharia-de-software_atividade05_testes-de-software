package identifier_test

import (
	identifier "atividade05/identifier"
	"fmt"
	"testing"
)

type TestCase struct {
	Input    string
	Expected bool
}

func formatNumber(index int) string {
	if index < 10 {
		return "0" + fmt.Sprint(index)
	}
	return fmt.Sprint(index)
}

func TestIdentifierValidation(t *testing.T) {
	identifier := identifier.Identifier{}

	casosDeTeste := []TestCase{
		{"Emerso", true},
		{"", false},
		{"Emerson", false},
		{"Stella", true},
		{"1Mari", false},
		{"Emer41", true},
		{"Em&r%", false},
		{"`aemer", false},
		{"a`emer", false},
		{"a12345", true},
		{"aaemer", true},
		{"b12345", true},
		{"bbemer", true},
		{"y12345", true},
		{"yyemer", true},
		{"z12345", true},
		{"zzemer", true},
		{"{12345", false},
		{"z{emer", false},
		{"@12345", false},
		{"@{emer", false},
		{"A12345", true},
		{"aAemer", true},
		{"B12345", true},
		{"aBemer", true},
		{"Y12345", true},
		{"aYemer", true},
		{"Z12345", true},
		{"aZemer", true},
		{"[12345", false},
		{"a[emer", false},
		{"a/emer", false},
		{"a0emer", true},
		{"a1emer", true},
		{"a8emer", true},
		{"a9emer", true},
		{"a:emer", false},
		{"a", true},
		{"Emers", true},
	}

	for index, teste := range casosDeTeste {
		t.Run(fmt.Sprintf("Caso de teste %s para %q", formatNumber(index+1), teste.Input), func(t *testing.T) {
			recebido := identifier.ValidateIdentifier(teste.Input)

			if teste.Expected != recebido {
				t.Errorf("O valor esperado %t é diferente do valor recebido %t", teste.Expected, recebido)
			}
		})
	}
}
