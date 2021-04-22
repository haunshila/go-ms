package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "haunshila",
		Price: 2.0,
		SKU:   "haaaa-adfa-asdfa",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
