package sophos_test

import (
	"fmt"
	"testing"

	"github.com/esurdam/go-sophos/types"
)

func TestResponse_MarshalTo(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var dns types.Dns
	res, err := client.Get(dns.GetPath())
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode > 300 {
		t.Error(res.Body)
		return
	}
	err = res.MarshalTo(&dns)
	if err != nil {
		t.Error(err)
	}

	err = res.MarshalTo(&dns)
	if err == nil {
		t.Error(fmt.Errorf("res.MarshalTo(&dns) should have errored with body already read"))
	}
}
