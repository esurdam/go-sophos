package sophos

import (
	"encoding/json"
	"net/http"
)

// Response contains the http.Response from the API
type Response struct {
	*http.Response
}

// MarshalTo marshals the response's body to the provided interface
func (r *Response) MarshalTo(x interface{}) error {
	return json.NewDecoder(r.Body).Decode(&x)
}
