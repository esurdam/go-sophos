package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Right is a generated struct representing the Sophos Right Endpoint
// GET /api/nodes/right
type Right struct {
	RightGroup RightGroup `json:"right_group"`
	RightRight RightRight `json:"right_right"`
}

var defsRight = map[string]sophos.RestObject{
	"RightGroup": &RightGroup{},
	"RightRight": &RightRight{},
}

// RestObjects implements the sophos.Node interface and returns a map of Right's Objects
func (Right) RestObjects() map[string]sophos.RestObject { return defsRight }

// GetPath implements sophos.RestGetter
func (*Right) GetPath() string { return "/api/nodes/right" }

// RefRequired implements sophos.RestGetter
func (*Right) RefRequired() (string, bool) { return "", false }

var defRight = &sophos.Definition{Description: "right", Name: "right", Link: "/api/definitions/right"}

// Definition returns the /api/definitions struct of Right
func (Right) Definition() sophos.Definition { return *defRight }

// ApiRoutes returns all known Right Paths
func (Right) ApiRoutes() []string {
	return []string{
		"/api/objects/right/group/",
		"/api/objects/right/group/{ref}",
		"/api/objects/right/group/{ref}/usedby",
		"/api/objects/right/right/",
		"/api/objects/right/right/{ref}",
		"/api/objects/right/right/{ref}/usedby",
	}
}

// References returns the Right's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Right) References() []string {
	return []string{
		"REF_RightGroup",
		"REF_RightRight",
	}
}

// RightGroup is an Sophos Endpoint subType and implements sophos.RestObject
type RightGroup []interface{}

// GetPath implements sophos.RestObject and returns the RightGroup GET path
// Returns all available right/group objects
func (*RightGroup) GetPath() string { return "/api/objects/right/group/" }

// RefRequired implements sophos.RestObject
func (*RightGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RightGroup DELETE path
// Creates or updates the complete object group
func (*RightGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/right/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RightGroup PATCH path
// Changes to parts of the object group types
func (*RightGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RightGroup POST path
// Create a new right/group object
func (*RightGroup) PostPath() string {
	return "/api/objects/right/group/"
}

// PutPath implements sophos.RestObject and returns the RightGroup PUT path
// Creates or updates the complete object group
func (*RightGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*RightGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/group/%s/usedby", ref)
}

// RightRights is an Sophos Endpoint subType and implements sophos.RestObject
type RightRights []RightRight

// RightRight is a generated Sophos object
type RightRight struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the RightRights GET path
// Returns all available right/right objects
func (*RightRights) GetPath() string { return "/api/objects/right/right/" }

// RefRequired implements sophos.RestObject
func (*RightRights) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the RightRights GET path
// Returns all available right types
func (r *RightRight) GetPath() string { return fmt.Sprintf("/api/objects/right/right/%s", r.Reference) }

// RefRequired implements sophos.RestObject
func (r *RightRight) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the RightRight DELETE path
// Creates or updates the complete object right
func (*RightRight) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/right/right/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RightRight PATCH path
// Changes to parts of the object right types
func (*RightRight) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/right/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RightRight POST path
// Create a new right/right object
func (*RightRight) PostPath() string {
	return "/api/objects/right/right/"
}

// PutPath implements sophos.RestObject and returns the RightRight PUT path
// Creates or updates the complete object right
func (*RightRight) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/right/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*RightRight) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/right/right/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *RightRight) GetType() string { return r._type }
