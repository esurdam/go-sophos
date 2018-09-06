package sophos_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/esurdam/go-sophos"

	"github.com/esurdam/go-sophos/types"
)

func TestResponse_MarshalTo(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r := httptest.NewRecorder()
	byt, _ := json.Marshal(types.Dns{
		Email: "test@test.com",
	})
	r.Body.Write(byt)

	var res sophos.Response
	var dns types.Dns

	res.Response = r.Result()
	err := res.MarshalTo(&dns)
	if err != nil {
		t.Error(err)
	}

	if dns.Email != "test@test.com" {
		t.Errorf("wanted test@test.com, got %s", dns.Email)
	}
}
