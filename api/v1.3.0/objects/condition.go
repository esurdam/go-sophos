package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Condition is a generated struct representing the Sophos Condition Endpoint
// GET /api/nodes/condition
type Condition struct {
	ConditionGroup  ConditionGroup  `json:"condition_group"`
	ConditionObjref ConditionObjref `json:"condition_objref"`
}

var _ sophos.Endpoint = &Condition{}

var defsCondition = map[string]sophos.RestObject{
	"ConditionGroup":  &ConditionGroup{},
	"ConditionObjref": &ConditionObjref{},
}

// RestObjects implements the sophos.Node interface and returns a map of Condition's Objects
func (Condition) RestObjects() map[string]sophos.RestObject { return defsCondition }

// GetPath implements sophos.RestGetter
func (*Condition) GetPath() string { return "/api/nodes/condition" }

// RefRequired implements sophos.RestGetter
func (*Condition) RefRequired() (string, bool) { return "", false }

var defCondition = &sophos.Definition{Description: "condition", Name: "condition", Link: "/api/definitions/condition"}

// Definition returns the /api/definitions struct of Condition
func (Condition) Definition() sophos.Definition { return *defCondition }

// ApiRoutes returns all known Condition Paths
func (Condition) ApiRoutes() []string {
	return []string{
		"/api/objects/condition/group/",
		"/api/objects/condition/group/{ref}",
		"/api/objects/condition/group/{ref}/usedby",
		"/api/objects/condition/objref/",
		"/api/objects/condition/objref/{ref}",
		"/api/objects/condition/objref/{ref}/usedby",
	}
}

// References returns the Condition's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Condition) References() []string {
	return []string{
		"REF_ConditionGroup",
		"REF_ConditionObjref",
	}
}

// ConditionGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ConditionGroup []interface{}

var _ sophos.RestObject = &ConditionGroup{}

// GetPath implements sophos.RestObject and returns the ConditionGroup GET path
// Returns all available condition/group objects
func (*ConditionGroup) GetPath() string { return "/api/objects/condition/group/" }

// RefRequired implements sophos.RestObject
func (*ConditionGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ConditionGroup DELETE path
// Creates or updates the complete object group
func (*ConditionGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ConditionGroup PATCH path
// Changes to parts of the object group types
func (*ConditionGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ConditionGroup POST path
// Create a new condition/group object
func (*ConditionGroup) PostPath() string {
	return "/api/objects/condition/group/"
}

// PutPath implements sophos.RestObject and returns the ConditionGroup PUT path
// Creates or updates the complete object group
func (*ConditionGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ConditionGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/group/%s/usedby", ref)
}

// ConditionObjrefs is an Sophos Endpoint subType and implements sophos.RestObject
type ConditionObjrefs []ConditionObjref

// ConditionObjref is a generated Sophos object
type ConditionObjref struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Attr      string `json:"attr"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Operator  string `json:"operator"`
	Ref       string `json:"ref"`
	Value     string `json:"value"`
}

var _ sophos.RestGetter = &ConditionObjref{}

// GetPath implements sophos.RestObject and returns the ConditionObjrefs GET path
// Returns all available condition/objref objects
func (*ConditionObjrefs) GetPath() string { return "/api/objects/condition/objref/" }

// RefRequired implements sophos.RestObject
func (*ConditionObjrefs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ConditionObjrefs GET path
// Returns all available objref types
func (c *ConditionObjref) GetPath() string {
	return fmt.Sprintf("/api/objects/condition/objref/%s", c.Reference)
}

// RefRequired implements sophos.RestObject
func (c *ConditionObjref) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the ConditionObjref DELETE path
// Creates or updates the complete object objref
func (*ConditionObjref) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/objref/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ConditionObjref PATCH path
// Changes to parts of the object objref types
func (*ConditionObjref) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/objref/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ConditionObjref POST path
// Create a new condition/objref object
func (*ConditionObjref) PostPath() string {
	return "/api/objects/condition/objref/"
}

// PutPath implements sophos.RestObject and returns the ConditionObjref PUT path
// Creates or updates the complete object objref
func (*ConditionObjref) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/objref/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ConditionObjref) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/condition/objref/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *ConditionObjref) GetType() string { return c._type }
