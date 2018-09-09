package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// MacList is a generated struct representing the Sophos MacList Endpoint
// GET /api/nodes/mac_list
type MacList struct {
	MacListGroup   MacListGroup   `json:"mac_list_group"`
	MacListMacList MacListMacList `json:"mac_list_mac_list"`
}

var defsMacList = map[string]sophos.RestObject{
	"MacListGroup":   &MacListGroup{},
	"MacListMacList": &MacListMacList{},
}

// RestObjects implements the sophos.Node interface and returns a map of MacList's Objects
func (MacList) RestObjects() map[string]sophos.RestObject { return defsMacList }

// GetPath implements sophos.RestGetter
func (*MacList) GetPath() string { return "/api/nodes/mac_list" }

// RefRequired implements sophos.RestGetter
func (*MacList) RefRequired() (string, bool) { return "", false }

var defMacList = &sophos.Definition{Description: "mac_list", Name: "mac_list", Link: "/api/definitions/mac_list"}

// Definition returns the /api/definitions struct of MacList
func (MacList) Definition() sophos.Definition { return *defMacList }

// ApiRoutes returns all known MacList Paths
func (MacList) ApiRoutes() []string {
	return []string{
		"/api/objects/mac_list/group/",
		"/api/objects/mac_list/group/{ref}",
		"/api/objects/mac_list/group/{ref}/usedby",
		"/api/objects/mac_list/mac_list/",
		"/api/objects/mac_list/mac_list/{ref}",
		"/api/objects/mac_list/mac_list/{ref}/usedby",
	}
}

// References returns the MacList's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (MacList) References() []string {
	return []string{
		"REF_MacListGroup",
		"REF_MacListMacList",
	}
}

// MacListGroup is an Sophos Endpoint subType and implements sophos.RestObject
type MacListGroup []interface{}

// GetPath implements sophos.RestObject and returns the MacListGroup GET path
// Returns all available mac_list/group objects
func (*MacListGroup) GetPath() string { return "/api/objects/mac_list/group/" }

// RefRequired implements sophos.RestObject
func (*MacListGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the MacListGroup DELETE path
// Creates or updates the complete object group
func (*MacListGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the MacListGroup PATCH path
// Changes to parts of the object group types
func (*MacListGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the MacListGroup POST path
// Create a new mac_list/group object
func (*MacListGroup) PostPath() string {
	return "/api/objects/mac_list/group/"
}

// PutPath implements sophos.RestObject and returns the MacListGroup PUT path
// Creates or updates the complete object group
func (*MacListGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*MacListGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/group/%s/usedby", ref)
}

// MacListMacList is an Sophos Endpoint subType and implements sophos.RestObject
type MacListMacList []interface{}

// GetPath implements sophos.RestObject and returns the MacListMacList GET path
// Returns all available mac_list/mac_list objects
func (*MacListMacList) GetPath() string { return "/api/objects/mac_list/mac_list/" }

// RefRequired implements sophos.RestObject
func (*MacListMacList) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the MacListMacList DELETE path
// Creates or updates the complete object mac_list
func (*MacListMacList) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/mac_list/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the MacListMacList PATCH path
// Changes to parts of the object mac_list types
func (*MacListMacList) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/mac_list/%s", ref)
}

// PostPath implements sophos.RestObject and returns the MacListMacList POST path
// Create a new mac_list/mac_list object
func (*MacListMacList) PostPath() string {
	return "/api/objects/mac_list/mac_list/"
}

// PutPath implements sophos.RestObject and returns the MacListMacList PUT path
// Creates or updates the complete object mac_list
func (*MacListMacList) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/mac_list/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*MacListMacList) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/mac_list/mac_list/%s/usedby", ref)
}