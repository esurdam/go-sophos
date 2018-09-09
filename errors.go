package sophos

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
