package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Rhiad",
		Price: 22.0,
		SKU:   "aaa-bbb-ccc",
	}
	if err := p.Validate(); err != nil {
		t.Fatal(err)
	}
}
