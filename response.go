package sophos

import (
	"encoding/json"
	"errors"
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
		return errors.New(fmt.Sprintf("response: could not read response body: %s", err.Error()))
	}
	if err = json.Unmarshal(bodyText, x); err != nil {
		return errors.New(fmt.Sprintf("response: could not unmsarshal response to interface: %s", err.Error()))
	}
	return nil
}
