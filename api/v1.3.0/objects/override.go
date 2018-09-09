package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Override is a generated struct representing the Sophos Override Endpoint
// GET /api/nodes/override
type Override struct {
	OverrideGroup  OverrideGroup  `json:"override_group"`
	OverrideObjref OverrideObjref `json:"override_objref"`
}

var defsOverride = map[string]sophos.RestObject{
	"OverrideGroup":  &OverrideGroup{},
	"OverrideObjref": &OverrideObjref{},
}

// RestObjects implements the sophos.Node interface and returns a map of Override's Objects
func (Override) RestObjects() map[string]sophos.RestObject { return defsOverride }

// GetPath implements sophos.RestGetter
func (*Override) GetPath() string { return "/api/nodes/override" }

// RefRequired implements sophos.RestGetter
func (*Override) RefRequired() (string, bool) { return "", false }

var defOverride = &sophos.Definition{Description: "override", Name: "override", Link: "/api/definitions/override"}

// Definition returns the /api/definitions struct of Override
func (Override) Definition() sophos.Definition { return *defOverride }

// ApiRoutes returns all known Override Paths
func (Override) ApiRoutes() []string {
	return []string{
		"/api/objects/override/group/",
		"/api/objects/override/group/{ref}",
		"/api/objects/override/group/{ref}/usedby",
		"/api/objects/override/objref/",
		"/api/objects/override/objref/{ref}",
		"/api/objects/override/objref/{ref}/usedby",
	}
}

// References returns the Override's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Override) References() []string {
	return []string{
		"REF_OverrideGroup",
		"REF_OverrideObjref",
	}
}

// OverrideGroup is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideGroup []interface{}

// GetPath implements sophos.RestObject and returns the OverrideGroup GET path
// Returns all available override/group objects
func (*OverrideGroup) GetPath() string { return "/api/objects/override/group/" }

// RefRequired implements sophos.RestObject
func (*OverrideGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OverrideGroup DELETE path
// Creates or updates the complete object group
func (*OverrideGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OverrideGroup PATCH path
// Changes to parts of the object group types
func (*OverrideGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OverrideGroup POST path
// Create a new override/group object
func (*OverrideGroup) PostPath() string {
	return "/api/objects/override/group/"
}

// PutPath implements sophos.RestObject and returns the OverrideGroup PUT path
// Creates or updates the complete object group
func (*OverrideGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*OverrideGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s/usedby", ref)
}

// OverrideObjref is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideObjref []interface{}

// GetPath implements sophos.RestObject and returns the OverrideObjref GET path
// Returns all available override/objref objects
func (*OverrideObjref) GetPath() string { return "/api/objects/override/objref/" }

// RefRequired implements sophos.RestObject
func (*OverrideObjref) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OverrideObjref DELETE path
// Creates or updates the complete object objref
func (*OverrideObjref) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OverrideObjref PATCH path
// Changes to parts of the object objref types
func (*OverrideObjref) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OverrideObjref POST path
// Create a new override/objref object
func (*OverrideObjref) PostPath() string {
	return "/api/objects/override/objref/"
}

// PutPath implements sophos.RestObject and returns the OverrideObjref PUT path
// Creates or updates the complete object objref
func (*OverrideObjref) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*OverrideObjref) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s/usedby", ref)
}
