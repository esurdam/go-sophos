package sophos

import (
	"encoding/json"
	"net/http"
)

// Response contains the http.Response from the API
type Response struct {
	*http.Response
	// Errors is a slice of type Error that != nil when http.StatusCode >= 400 <= 421
	Errors *Errors
}

// MarshalTo marshals the response's body to the provided interface
func (r *Response) MarshalTo(x interface{}) error {
	return json.NewDecoder(r.Body).Decode(x)
}
