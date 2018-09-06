package sophos_test

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/esurdam/go-sophos"
)

var client *sophos.Client

var errOption = func(r *http.Request) error {
	return fmt.Errorf("this is a fake error")
}

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

	r, err = client.Request("é", "/api/status/version", nil)
	if err == nil {
		t.Error(fmt.Errorf("request should have errored"))
	}
}

func TestClient_Ping(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	if _, err := client.Ping(); err != nil {
		t.Error(err)
	}

	if _, err := client.Ping(errOption); err == nil {
		t.Error("error should not be nil with errOption")
	}
}

func TestNew(t *testing.T) {
	type args struct {
		endpoint string
		opts     []sophos.Option
	}
	tests := []struct {
		name    string
		args    args
		want    *sophos.Client
		wantErr bool
	}{
		{"testNewNoEndpoint", args{endpoint: ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sophos.New(tt.args.endpoint, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	_, err := client.Do("é", "/api", nil)
	if err == nil {
		t.Error("should have error since bad method")
	}

	_, err = client.Do("GET", "/api", nil, errOption)
	if err == nil {
		t.Error("should have error since bad errOption")
	}
}
