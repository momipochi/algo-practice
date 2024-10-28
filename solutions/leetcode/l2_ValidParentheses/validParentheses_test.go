package l2validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	if !isValid("()") {
		t.Error("Wrong answer")
	}
}
