package sophos_test

import (
	"testing"

	"github.com/esurdam/go-sophos"
)

func TestResponse_MarshalTo(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var nodes sophos.Dns
	res, err := client.Get(nodes.GetPath())
	if err != nil {
		t.Error(err)
	}
	err = res.MarshalTo(&nodes)
	if err != nil {
		t.Error(err)
	}
}
