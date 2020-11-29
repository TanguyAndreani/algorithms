package decorator_pattern

import (
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	coffee := BeverageWithSugar{BeverageWithChocolate{Coffee{}}}

	if coffee.cost() != 6 {
		t.Error("Expected cost to be 3! Got", coffee.cost())
	}
}
