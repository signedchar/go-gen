package mult

import "testing"

func Testmult(t *testing.T) {
	if mult(3, 3) != true {
		t.Fail()
	}

	if mult(3,5) == true {
		t.Fail()
	}

	if mult(5,5) != true {
		t.Fail()
	}

	if mult(5,3) == true {
		t.Fail()
	}
}
