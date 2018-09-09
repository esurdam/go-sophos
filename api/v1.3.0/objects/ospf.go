package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ospf is a generated struct representing the Sophos Ospf Endpoint
// GET /api/nodes/ospf
type Ospf struct {
	OspfArea             OspfArea             `json:"ospf_area"`
	OspfGroup            OspfGroup            `json:"ospf_group"`
	OspfInterface        OspfInterface        `json:"ospf_interface"`
	OspfMessageDigestKey OspfMessageDigestKey `json:"ospf_message_digest_key"`
}

var defsOspf = map[string]sophos.RestObject{
	"OspfArea":             &OspfArea{},
	"OspfGroup":            &OspfGroup{},
	"OspfInterface":        &OspfInterface{},
	"OspfMessageDigestKey": &OspfMessageDigestKey{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ospf's Objects
func (Ospf) RestObjects() map[string]sophos.RestObject { return defsOspf }

// GetPath implements sophos.RestGetter
func (*Ospf) GetPath() string { return "/api/nodes/ospf" }

// RefRequired implements sophos.RestGetter
func (*Ospf) RefRequired() (string, bool) { return "", false }

var defOspf = &sophos.Definition{Description: "ospf", Name: "ospf", Link: "/api/definitions/ospf"}

// Definition returns the /api/definitions struct of Ospf
func (Ospf) Definition() sophos.Definition { return *defOspf }

// ApiRoutes returns all known Ospf Paths
func (Ospf) ApiRoutes() []string {
	return []string{
		"/api/objects/ospf/area/",
		"/api/objects/ospf/area/{ref}",
		"/api/objects/ospf/area/{ref}/usedby",
		"/api/objects/ospf/group/",
		"/api/objects/ospf/group/{ref}",
		"/api/objects/ospf/group/{ref}/usedby",
		"/api/objects/ospf/interface/",
		"/api/objects/ospf/interface/{ref}",
		"/api/objects/ospf/interface/{ref}/usedby",
		"/api/objects/ospf/message_digest_key/",
		"/api/objects/ospf/message_digest_key/{ref}",
		"/api/objects/ospf/message_digest_key/{ref}/usedby",
	}
}

// References returns the Ospf's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ospf) References() []string {
	return []string{
		"REF_OspfArea",
		"REF_OspfGroup",
		"REF_OspfInterface",
		"REF_OspfMessageDigestKey",
	}
}

// OspfArea is an Sophos Endpoint subType and implements sophos.RestObject
type OspfArea []interface{}

// GetPath implements sophos.RestObject and returns the OspfArea GET path
// Returns all available ospf/area objects
func (*OspfArea) GetPath() string { return "/api/objects/ospf/area/" }

// RefRequired implements sophos.RestObject
func (*OspfArea) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfArea DELETE path
// Creates or updates the complete object area
func (*OspfArea) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfArea PATCH path
// Changes to parts of the object area types
func (*OspfArea) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfArea POST path
// Create a new ospf/area object
func (*OspfArea) PostPath() string {
	return "/api/objects/ospf/area/"
}

// PutPath implements sophos.RestObject and returns the OspfArea PUT path
// Creates or updates the complete object area
func (*OspfArea) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*OspfArea) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s/usedby", ref)
}

// OspfGroup is an Sophos Endpoint subType and implements sophos.RestObject
type OspfGroup []interface{}

// GetPath implements sophos.RestObject and returns the OspfGroup GET path
// Returns all available ospf/group objects
func (*OspfGroup) GetPath() string { return "/api/objects/ospf/group/" }

// RefRequired implements sophos.RestObject
func (*OspfGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfGroup DELETE path
// Creates or updates the complete object group
func (*OspfGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfGroup PATCH path
// Changes to parts of the object group types
func (*OspfGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfGroup POST path
// Create a new ospf/group object
func (*OspfGroup) PostPath() string {
	return "/api/objects/ospf/group/"
}

// PutPath implements sophos.RestObject and returns the OspfGroup PUT path
// Creates or updates the complete object group
func (*OspfGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*OspfGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s/usedby", ref)
}

// OspfInterface is an Sophos Endpoint subType and implements sophos.RestObject
type OspfInterface []interface{}

// GetPath implements sophos.RestObject and returns the OspfInterface GET path
// Returns all available ospf/interface objects
func (*OspfInterface) GetPath() string { return "/api/objects/ospf/interface/" }

// RefRequired implements sophos.RestObject
func (*OspfInterface) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfInterface DELETE path
// Creates or updates the complete object interface
func (*OspfInterface) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfInterface PATCH path
// Changes to parts of the object interface types
func (*OspfInterface) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfInterface POST path
// Create a new ospf/interface object
func (*OspfInterface) PostPath() string {
	return "/api/objects/ospf/interface/"
}

// PutPath implements sophos.RestObject and returns the OspfInterface PUT path
// Creates or updates the complete object interface
func (*OspfInterface) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*OspfInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s/usedby", ref)
}

// OspfMessageDigestKey is an Sophos Endpoint subType and implements sophos.RestObject
type OspfMessageDigestKey []interface{}

// GetPath implements sophos.RestObject and returns the OspfMessageDigestKey GET path
// Returns all available ospf/message_digest_key objects
func (*OspfMessageDigestKey) GetPath() string { return "/api/objects/ospf/message_digest_key/" }

// RefRequired implements sophos.RestObject
func (*OspfMessageDigestKey) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfMessageDigestKey DELETE path
// Creates or updates the complete object message_digest_key
func (*OspfMessageDigestKey) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfMessageDigestKey PATCH path
// Changes to parts of the object message_digest_key types
func (*OspfMessageDigestKey) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfMessageDigestKey POST path
// Create a new ospf/message_digest_key object
func (*OspfMessageDigestKey) PostPath() string {
	return "/api/objects/ospf/message_digest_key/"
}

// PutPath implements sophos.RestObject and returns the OspfMessageDigestKey PUT path
// Creates or updates the complete object message_digest_key
func (*OspfMessageDigestKey) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*OspfMessageDigestKey) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s/usedby", ref)
}
