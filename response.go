package sophos

import (
	"encoding/json"
	"net/http"
)

// Response contains the http.Response from the API
type Response struct {
	*http.Response
}

// Errors returns any errors returned
func (r *Response) Errors() Errors {
	var errs Errors
	_ = r.MarshalTo(&errs)
	return errs
}

// MarshalTo marshals the response's body to the provided interface
func (r *Response) MarshalTo(x interface{}) error {
	return json.NewDecoder(r.Body).Decode(&x)
}
