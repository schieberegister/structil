package structil_test

import (
	"testing"

	. "github.com/goldeneggg/structil"
)

func TestVersion(t *testing.T) {
	exp := "0.7.3"

	if VERSION != exp {
		t.Errorf("expected: %#v, but actual: %#v", exp, VERSION)
	}
}
