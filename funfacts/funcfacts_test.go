package funfacts

import (
  "testing"
  "reflect"
)

/**
  Mal for FunFacts funksjonen.
  Definer korrekte typer for input og want,
  og sette korrekte testverdier i slices tests.
*/
func TestFunFacts(t *testing.T) {
    type test struct {
        input // her må du skrive riktig type for input
        want // her må du skrive riktig type for returverdien
    }

    // Her må du legge inn korrekte testverdier
    tests := []test{
      {input: , want: },
    }

    for _, tc := range tests {
        got := FunFacts(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}
