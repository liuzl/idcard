package idcard

import (
	"testing"
)

func TestVerify(t *testing.T) {
	cases := []struct {
		id    string
		valid bool
	}{
		{"130524198308184533", true},
		{"130524198308184532", false},
		{"130524098318184532", false},
	}
	for _, c := range cases {
		ret := Verify(c.id)
		if c.valid != ret {
			t.Errorf("Verify(%s)=%v != expected %v", c.id, ret, c.valid)
		}
	}
}

func TestParse(t *testing.T) {
	ret, err := Parse("130524198308184533")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
