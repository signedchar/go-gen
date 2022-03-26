package mult

import "testing"

func Testmult3(t *testing.T) {
	if mult3(3) != true {
		t.Fail()
	}

	if mult3(5) == true {
		t.Fail()
	}
}
