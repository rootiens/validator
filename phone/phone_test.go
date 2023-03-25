package phone

import "testing"

type PhoneTest struct {
	Phone       string
	Expected bool
}

func TestValidation(t *testing.T) {
	phoneNumbers := []PhoneTest{
		PhoneTest{"9123456789", true},
		PhoneTest{"123512hasd", false},
		PhoneTest{"09123456789", true},
		PhoneTest{"00989123456789", true},
		PhoneTest{"۹۱۲۳۴۵۶۷۸۹", true},
		PhoneTest{"+۹۸۹۱۲۳۴۵۶۷۸۹", true},
		PhoneTest{"912345678", false},
	}

	for _, test := range phoneNumbers {
		if ok  := IsValid(test.Phone); ok != test.Expected {
			t.Errorf("Output of %q is %t while expected result is %t", test.Phone, ok, test.Expected)
		}
	}
}

func BenchmarkValidation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsValid("1235123")
    }
}
