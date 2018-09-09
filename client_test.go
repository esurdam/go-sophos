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

		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/error" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(dnsMock{})
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

	var dns dnsMock
	err := client.GetObject(&dns)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	if err := client.GetObject(&dns, errOption); err == nil {
		t.Error("error should not be nil with errOption")
	}

	var d dhcpServerMock
	err = client.GetObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PostObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = dhcpServerMock{Reference: "abc"}
	err := client.PostObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}
}

func TestClient_DeleteObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = dhcpServerMock{Reference: "abc"}
	err := client.DeleteObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = dhcpServerMock{Reference: ""}
	err = client.DeleteObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PatchObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = dhcpServerMock{Reference: "abc"}
	err := client.PatchObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = dhcpServerMock{}
	err = client.PatchObject(&d)
	if err == nil {
		t.Error("type should require ref")
	}
}

func TestClient_PutObject(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var d = dhcpServerMock{Reference: "abc"}
	err := client.PutObject(&d)
	if err != nil {
		t.Error("error should be nil", err.Error())
	}

	d = dhcpServerMock{}
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

	_, err = client.Do("GET", "/api/error", nil)
	if err == nil {
		t.Error("should have error since bad errOption")
	}
}

// dnsMock is a generated struct representing the Sophos dnsMock Endpoint
// GET /api/nodes/dns
type dnsMock struct {
	AllowedNetworks []string      `json:"allowed_networks"`
	Axfr            []interface{} `json:"axfr"`
	Dnssec          int64         `json:"dnssec"`
	Email           string        `json:"email"`
	EmptyZones      string        `json:"empty_zones"`
	FwdDynamic      int64         `json:"fwd_dynamic"`
	FwdStatic       []string      `json:"fwd_static"`
	RecheckInterval int64         `json:"recheck_interval"`
	Routes          []string      `json:"routes"`
}

var defsDns = map[string]sophos.RestObject{}

// RestObjects implements the sophos.Node interface and returns a map of dnsMock's Objects
func (dnsMock) RestObjects() map[string]sophos.RestObject { return defsDns }

// GetPath implements sophos.RestGetter
func (*dnsMock) GetPath() string { return "/api/nodes/dns" }

// RefRequired implements sophos.RestGetter
func (*dnsMock) RefRequired() (string, bool) { return "", false }

var defDns = &sophos.Definition{Description: "dns", Name: "dns", Link: "/api/definitions/dns"}

// Definition returns the /api/definitions struct of dnsMock
func (dnsMock) Definition() sophos.Definition { return *defDns }

// ApiRoutes returns all known dnsMock Paths
func (dnsMock) ApiRoutes() []string {
	return []string{
		"/api/objects/dns/axfr/",
		"/api/objects/dns/axfr/{ref}",
		"/api/objects/dns/axfr/{ref}/usedby",
		"/api/objects/dns/group/",
		"/api/objects/dns/group/{ref}",
		"/api/objects/dns/group/{ref}/usedby",
		"/api/objects/dns/route/",
		"/api/objects/dns/route/{ref}",
		"/api/objects/dns/route/{ref}/usedby",
	}
}

// References returns the dnsMock's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (dnsMock) References() []string {
	return []string{
		"REF_DnsAxfr",
		"REF_DnsGroup",
		"REF_DnsRoute",
	}
}

// dhcpServerMock is a generated Sophos object
type dhcpServerMock struct {
	Locked          string   `json:"_locked"`
	Reference       string   `json:"_ref"`
	_type           string   `json:"_type"`
	Address         string   `json:"address"`
	Comment         string   `json:"comment"`
	Custom          string   `json:"custom"`
	DefaultGateway  string   `json:"default_gateway"`
	DenyUnknown     bool     `json:"deny_unknown"`
	DNS1            string   `json:"dns1"`
	DNS2            string   `json:"dns2"`
	Domain          string   `json:"domain"`
	Interface       string   `json:"interface"`
	LeaseTime       int64    `json:"lease_time"`
	Mappings        []string `json:"mappings"`
	Name            string   `json:"name"`
	Netmask         int64    `json:"netmask"`
	ProxyAutoconfig bool     `json:"proxy_autoconfig"`
	RangeEnd        string   `json:"range_end"`
	RangeStart      string   `json:"range_start"`
	RelayMode       bool     `json:"relay_mode"`
	Status          bool     `json:"status"`
	Wins            string   `json:"wins"`
	WinsNodeType    string   `json:"wins_node_type"`
}

// GetPath implements sophos.RestObject and returns the DhcpServers GET path
// Returns all available server types
func (d *dhcpServerMock) GetPath() string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *dhcpServerMock) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the dhcpServerMock DELETE path
// Creates or updates the complete object server
func (*dhcpServerMock) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the dhcpServerMock PATCH path
// Changes to parts of the object server types
func (*dhcpServerMock) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the dhcpServerMock POST path
// Create a new dhcp/server object
func (*dhcpServerMock) PostPath() string {
	return "/api/objects/dhcp/server/"
}

// PutPath implements sophos.RestObject and returns the dhcpServerMock PUT path
// Creates or updates the complete object server
func (*dhcpServerMock) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*dhcpServerMock) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s/usedby", ref)
}

// GetType implements sophos.Object
func (d *dhcpServerMock) GetType() string { return d._type }
