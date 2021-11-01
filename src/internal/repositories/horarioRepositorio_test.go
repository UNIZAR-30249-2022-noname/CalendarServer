package horarioRepositorio_test

import (
	"strings"
	"testing"
)

func IsSuperAnimal(animal string) bool {
	return strings.ToLower(animal) == "gopher"
}

func TestIsSuperAnimal(t *testing.T) {
	expected := true
	got := IsSuperAnimal("gopher")
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}