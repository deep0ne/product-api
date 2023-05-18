package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "some garbage",
		Price: 1.00,
		SKU:   "abs-cde-fdg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
