package factoryMethod_test

import (
	"testing"

	factoryMethod "github.com/kiki-ki/design_pattern_go/factory_method"
)

func TestFactoryMethod(t *testing.T) {
	wants := []string{"A", "B", "C"}
	factory := &factoryMethod.IDCardFactory{}
	products := []factoryMethod.Product{
		factory.Create(factory, "A"),
		factory.Create(factory, "B"),
		factory.Create(factory, "C"),
	}
	for i, p := range products {
		if got := p.Use(); wants[i] != got {
			t.Errorf("want %s but got %s", wants[i], got)
		}
	}
}
