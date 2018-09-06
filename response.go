package sophos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response contains the http.Response from the API
type Response struct {
	*http.Response
}

// MarshalTo marshals the response's body to the provided interface
func (r *Response) MarshalTo(x interface{}) error {
	bodyText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("response: could not read response body: %s", err.Error())
	}
	return json.Unmarshal(bodyText, x)
}
