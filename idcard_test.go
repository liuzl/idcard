package idcard

import (
	"testing"
)

func TestVerify(t *testing.T) {
	cases := []struct {
		id    string
		valid bool
	}{
		{"110225196403026127", true},
		{"110225196403026227", false},
		{"110225196403026122", false},
	}
	for _, c := range cases {
		ret, _ := Verify(c.id)
		if c.valid != ret {
			t.Errorf("Verify(%s)=%v != expected %v", c.id, ret, c.valid)
		}
	}
}

func TestParse(t *testing.T) {
	ret, err := Parse("110225196403026127")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

	_, err = Parse("")
	t.Log(err)
}
