package idcard

import "testing"

type IDTest struct {
	ID       string
	Expected bool
}

func TestValidation(t *testing.T) {
	IDs := []IDTest{
		IDTest{"1235123", false},
		IDTest{"11111111111", false},
		IDTest{"0544179943", true},
		IDTest{"853245304", true},
		IDTest{"۸۵۳۲۴۵۳۰۴", true},
		IDTest{"52032108", true},
	}

	for _, test := range IDs {
		if ok, _ := IsValid(test.ID); ok != test.Expected {
			t.Errorf("Output of %q is %t while expected result is %t", test.ID, ok, test.Expected)
		}
	}
}

func BenchmarkValidation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsValid("1235123")
    }
}
