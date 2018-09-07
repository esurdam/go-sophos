package sophos_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/esurdam/go-sophos/types"

	"github.com/esurdam/go-sophos"
)

var client *sophos.Client

var errOption = func(r *http.Request) error {
	return fmt.Errorf("this is a fake error")
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasPrefix(r.URL.Path, "/api/definitions/") {
			json.NewEncoder(w).Encode(map[string]sophos.MethodMap{})
			return
		}
		json.NewEncoder(w).Encode(types.Dns{})
	}))
	sophos.DefaultHTTPClient = ts.Client()
	clientF, err := sophos.New(ts.URL, sophos.WithAPIToken("abc"))
	if err == nil {
		client = clientF
	}
	if client == nil {
		t.Error("errror setting up client, client is nil")
	}
	return func(t *testing.T) {
		ts.Close()
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

func TestDefinition_Get(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	d := types.Nodes{}.Definition()
	_, err := d.GetSwag(client)
	if err != nil {
		t.Error(err)
	}

	_, err = d.GetSwag(client, errOption)
	if err == nil {
		t.Error(err)
	}
}

func TestClient_Ping(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	if _, err := client.Ping(errOption); err == nil {
		t.Error("error should not be nil with errOption")
	}

	if _, err := client.Ping(); err != nil {
		t.Error(err)
	}
}

func TestClient_Delete(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Delete("/api", errOption)
	if err == nil {
		t.Error("error should not be nil with errOption")
	}

	if r.Request.Method != http.MethodDelete {
		t.Error("method should be DELETE")
	}
}

func TestClient_GetObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var dns types.Dns
	err := client.GetObject(&dns)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	if err := client.GetObject(&dns, errOption); err == nil {
		t.Error("error should not be nil with errOption")
	}

	var d types.DhcpServer
	err = client.GetObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PostObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = types.DhcpServer{Reference: "abc"}
	err := client.PostObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}
}

func TestClient_DeleteObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = types.DhcpServer{Reference: "abc"}
	err := client.DeleteObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = types.DhcpServer{Reference: ""}
	err = client.DeleteObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PatchObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = types.DhcpServer{Reference: "abc"}
	err := client.PatchObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = types.DhcpServer{}
	err = client.PatchObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PutObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = types.DhcpServer{Reference: "abc"}
	err := client.PutObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = types.DhcpServer{}
	err = client.PutObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_Get(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Get("/api", errOption)
	if err == nil {
		t.Error("error should not be nil with errOption")
	}

	if r.Request.Method != http.MethodGet {
		t.Error("method should be GET")
	}
}

func TestClient_Put(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Put("/api", nil, errOption)
	if err == nil {
		t.Error("error should not be nil with errOption")
	}

	if r.Request.Method != http.MethodPut {
		t.Error("method should be PUT")
	}
}

func TestClient_Post(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Post("/api", nil, errOption)
	if err == nil {
		t.Error("error should not be nil with errOption")
	}

	if r.Request.Method != http.MethodPost {
		t.Error("method should be DELETE")
	}
}

func TestClient_Patch(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	r, err := client.Patch("/api", nil, errOption)
	if err == nil {
		t.Error("error should not be nil with errOption")
	}

	if r.Request.Method != http.MethodPatch {
		t.Error("method should be DELETE")
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
		want    string
		wantErr bool
	}{
		{"testNewNoEndpoint", args{endpoint: ""}, "", true},
		{"testNewEndpointNoHttp", args{endpoint: "boom.org"}, "https://boom.org", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sophos.New(tt.args.endpoint, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == "" {
				return
			}
			if !reflect.DeepEqual(got.Endpoint(), tt.want) {
				t.Errorf("New() = %v, want %v", got.Endpoint(), tt.want)
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
