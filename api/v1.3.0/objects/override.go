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

var _ sophos.Endpoint = &Override{}

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

// OverrideGroups is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideGroups []OverrideGroup

// OverrideGroup represents a UTM group
type OverrideGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &OverrideGroup{}

// GetPath implements sophos.RestObject and returns the OverrideGroups GET path
// Returns all available override/group objects
func (*OverrideGroups) GetPath() string { return "/api/objects/override/group/" }

// RefRequired implements sophos.RestObject
func (*OverrideGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OverrideGroups GET path
// Returns all available group types
func (o *OverrideGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/override/group/%s", o.Reference)
}

// RefRequired implements sophos.RestObject
func (o *OverrideGroup) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OverrideGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s/usedby", ref)
}

// OverrideObjrefs is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideObjrefs []OverrideObjref

// OverrideObjref represents a UTM monitoring action
type OverrideObjref struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Attr       string `json:"attr"`
	Comment    string `json:"comment"`
	// Condition description: REF(condition/*)
	Condition string `json:"condition"`
	Name      string `json:"name"`
	// Ref description: REF(/*)
	Ref   string `json:"ref"`
	Value string `json:"value"`
}

var _ sophos.RestGetter = &OverrideObjref{}

// GetPath implements sophos.RestObject and returns the OverrideObjrefs GET path
// Returns all available override/objref objects
func (*OverrideObjrefs) GetPath() string { return "/api/objects/override/objref/" }

// RefRequired implements sophos.RestObject
func (*OverrideObjrefs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OverrideObjrefs GET path
// Returns all available objref types
func (o *OverrideObjref) GetPath() string {
	return fmt.Sprintf("/api/objects/override/objref/%s", o.Reference)
}

// RefRequired implements sophos.RestObject
func (o *OverrideObjref) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OverrideObjref) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s/usedby", ref)
}
