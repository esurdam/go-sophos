package sophos_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/esurdam/go-sophos"
)

var client *sophos.Client

func init() {
	ep := os.Getenv("ENDPOINT")
	token := os.Getenv("TOKEN")
	if ep == "" || token == "" {
		panic("need endpoint and token")
	}

	clientF, err := sophos.New(ep, sophos.WithAPIToken(token))
	if err == nil {
		client = clientF
	}
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	if client == nil {
		t.Error("errror setting up client, client is nil")
	}
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestClient_Request(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Request("GET", "/api/status/version", nil)
	if err != nil {
		t.Error(err)
	}

	if r.Header.Get(sophos.Authorization) == "" {
		t.Error(errors.New("authorization is empty"))
	}

	wanted := fmt.Sprintf("%s/api/status/version", client.Endpoint())
	if r.URL.String() != wanted {
		t.Error(fmt.Errorf("incorrect URL: wanted %s, got %s", wanted, r.URL.String()))
	}
}

func TestClient_Ping(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	if _, err := client.Ping(); err != nil {
		t.Error(err)
	}
}
