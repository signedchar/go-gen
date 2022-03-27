package mult

import "testing"

func Testmult(t *Testing.T) {
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

func Testmult3(t *testing.T) {
	if mult3(3) != true {
		t.Fail()
	}

	if mult3(5) == true {
		t.Fail()
	}
}

func Testmult5(t *testing.T) {
	if mult5(5) != true {
		t.Fail()
	}

	if mult5(3) == true {
		t.Fail()
	}
}
