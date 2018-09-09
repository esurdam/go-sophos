package sophos_test

import (
	"github.com/esurdam/go-sophos"
	"testing"
)

func TestIsReference(t *testing.T) {
	ref := sophos.Reference("abc")
	if ref.IsReference() {
		t.Error("abc should not be a Reference")
	}
}
