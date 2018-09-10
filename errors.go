package sophos

import "fmt"

// An Error is returned from the Endpoint when errors occur
// Confd validates all nodes and objects on change operations (e.g., create, update,
// delete). When you execute a RESTful API call, you may trigger some validation errors.
// Validation errors are:
//  - fatal
//  - non-fatal
//
// Fatal errors indicate a programming error like wrong input or wrong type. Non-fatal
// errors indicate that the current operation has impact on other objects or nodes, which
// might be harmful.
//
// For example, if you create a rule for “packetfilter/packetfilter”, you would need to ref
// erence the node “packetfiler/rules” and enable the status. If you deleted the rule after
// wards, confd would detect that a referenced rule of “packetfilter/rules” was still
// enabled and report an error.
// For more information on validation errors, see chapter X-Restd-Err-Ack header.
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

// IsFatal returns true if Error is fatal
func (e Error) IsFatal() bool {
	return e.Fatal == 1
}

// IsFatalErr returns true if the error is an Error and is fatal
func IsFatalErr(err error) bool {
	switch err.(type) {
	case *Errors:
		return err.(*Errors).IsFatal()
	case Errors:
		return err.(Errors).IsFatal()
	case *Error:
		return err.(*Error).Fatal == 1
	case Error:
		return err.(Error).Fatal == 1
	default:
		return false
	}
}

// Error implements error interface
func (e Error) Error() string {
	if e.IsFatal() {
		return fmt.Sprintf("client do: FATAL error from UTM server: %s", e.Name)
	}
	return fmt.Sprintf("client do: error from UTM server: %s", e.Name)
}

// Errors is a slice of type Error which is returned from the api
type Errors []Error

// IsFatal returns true if any Error in the slice is fatal
func (ee Errors) IsFatal() bool {
	for _, e := range ee {
		if e.IsFatal() {
			return true
		}
	}
	return false
}

// Error implements error interface
func (ee Errors) Error() string {
	if len(ee) == 0 {
		// we should not get here since we will only return Errors when len(Errors) > 0
		// but in the event we do
		return fmt.Sprintf("error accessing UTM interface: check status code. No Errors were retuned in response body.")
	}
	for _, e := range ee {
		// return the first Fatal error message since it is most relevant
		if e.IsFatal() {
			return e.Error()
		}
	}

	// Otherwise just the first Error
	return ee[0].Error()
}
