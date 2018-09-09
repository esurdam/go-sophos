package sophos_test

import (
	"testing"

	"github.com/esurdam/go-sophos"
)

func TestIsReference(t *testing.T) {
	ref := sophos.Reference("abc")
	if ref.IsReference() {
		t.Error("abc should not be a Reference")
	}
}
