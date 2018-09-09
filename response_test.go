package sophos_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/esurdam/go-sophos"
)

func TestResponse_MarshalTo(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r := httptest.NewRecorder()
	byt, _ := json.Marshal(dnsMock{
		Email: "test@test.com",
	})
	r.Body.Write(byt)

	var res sophos.Response
	var dns dnsMock

	res.Response = r.Result()
	err := res.MarshalTo(&dns)
	if err != nil {
		t.Error(err)
	}

	if dns.Email != "test@test.com" {
		t.Errorf("wanted test@test.com, got %s", dns.Email)
	}
}

func TestResponse_Errors(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	res, _ := client.Get("/api/errorjson")
	errs := res.Errors
	if len(errs) != 1 {
		t.Error("TestResponse_Errors should have returned Error")
		return
	}
	if !errs.IsFatal() {
		t.Error("TestResponse_Errors should be fatal")
	}
	if !errs[0].IsFatal() {
		t.Error("TestResponse_Errors should be fatal")
	}
	if errs[0].Msgtype != "DATATYPE_OBJECT_ATTRIBUTE" {
		t.Error("TestResponse_Errors should have MsgType DATATYPE_OBJECT_ATTRIBUTE")
	}
	errs[0].Fatal = 0
	if errs.IsFatal() {
		t.Error("TestResponse_Errors should not be fatal")
	}
}
