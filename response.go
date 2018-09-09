package sophos

import (
	"encoding/json"
	"net/http"
)

// Response contains the http.Response from the API
type Response struct {
	*http.Response
}

// Errors is a slice of type Error which is returned from the api
type Errors []Error

// An Error is returned from the Endpoint when errors occur
type Error struct {
	Oattrs    []string      `json:"Oattrs"`
	Attrs     []interface{} `json:"attrs"`
	Class     string        `json:"class"`
	DelObject string        `json:"del_object"`
	Fatal     int64         `json:"fatal"`
	Format    string        `json:"format"`
	Msgtype   string        `json:"msgtype"`
	Name      string        `json:"name"`
	NeverHide int64         `json:"never_hide"`
	Objname   string        `json:"objname"`
	Perms     string        `json:"perms"`
	Ref       string        `json:"ref"`
	Rights    string        `json:"rights"`
	Type      string        `json:"type"`
}

// Errors returns any errors returned
func (r *Response) Errors() []Error {
	var errs Errors
	_ = r.MarshalTo(&errs)
	return errs
}

// MarshalTo marshals the response's body to the provided interface
func (r *Response) MarshalTo(x interface{}) error {
	return json.NewDecoder(r.Body).Decode(&x)
}
