package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// IpfixConnection is a generated struct representing the Sophos IpfixConnection Endpoint
// GET /api/nodes/ipfix_connection
type IpfixConnection struct {
	IpfixConnectionGroup           IpfixConnectionGroup           `json:"ipfix_connection_group"`
	IpfixConnectionIpfixConnection IpfixConnectionIpfixConnection `json:"ipfix_connection_ipfix_connection"`
}

var defsIpfixConnection = map[string]sophos.RestObject{
	"IpfixConnectionGroup":           &IpfixConnectionGroup{},
	"IpfixConnectionIpfixConnection": &IpfixConnectionIpfixConnection{},
}

// RestObjects implements the sophos.Node interface and returns a map of IpfixConnection's Objects
func (IpfixConnection) RestObjects() map[string]sophos.RestObject { return defsIpfixConnection }

// GetPath implements sophos.RestGetter
func (*IpfixConnection) GetPath() string { return "/api/nodes/ipfix_connection" }

// RefRequired implements sophos.RestGetter
func (*IpfixConnection) RefRequired() (string, bool) { return "", false }

var defIpfixConnection = &sophos.Definition{Description: "ipfix_connection", Name: "ipfix_connection", Link: "/api/definitions/ipfix_connection"}

// Definition returns the /api/definitions struct of IpfixConnection
func (IpfixConnection) Definition() sophos.Definition { return *defIpfixConnection }

// ApiRoutes returns all known IpfixConnection Paths
func (IpfixConnection) ApiRoutes() []string {
	return []string{
		"/api/objects/ipfix_connection/group/",
		"/api/objects/ipfix_connection/group/{ref}",
		"/api/objects/ipfix_connection/group/{ref}/usedby",
		"/api/objects/ipfix_connection/ipfix_connection/",
		"/api/objects/ipfix_connection/ipfix_connection/{ref}",
		"/api/objects/ipfix_connection/ipfix_connection/{ref}/usedby",
	}
}

// References returns the IpfixConnection's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (IpfixConnection) References() []string {
	return []string{
		"REF_IpfixConnectionGroup",
		"REF_IpfixConnectionIpfixConnection",
	}
}

// IpfixConnectionGroup is an Sophos Endpoint subType and implements sophos.RestObject
type IpfixConnectionGroup []interface{}

// GetPath implements sophos.RestObject and returns the IpfixConnectionGroup GET path
// Returns all available ipfix_connection/group objects
func (*IpfixConnectionGroup) GetPath() string { return "/api/objects/ipfix_connection/group/" }

// RefRequired implements sophos.RestObject
func (*IpfixConnectionGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpfixConnectionGroup DELETE path
// Creates or updates the complete object group
func (*IpfixConnectionGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpfixConnectionGroup PATCH path
// Changes to parts of the object group types
func (*IpfixConnectionGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpfixConnectionGroup POST path
// Create a new ipfix_connection/group object
func (*IpfixConnectionGroup) PostPath() string {
	return "/api/objects/ipfix_connection/group/"
}

// PutPath implements sophos.RestObject and returns the IpfixConnectionGroup PUT path
// Creates or updates the complete object group
func (*IpfixConnectionGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpfixConnectionGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/group/%s/usedby", ref)
}

// IpfixConnectionIpfixConnection is an Sophos Endpoint subType and implements sophos.RestObject
type IpfixConnectionIpfixConnection []interface{}

// GetPath implements sophos.RestObject and returns the IpfixConnectionIpfixConnection GET path
// Returns all available ipfix_connection/ipfix_connection objects
func (*IpfixConnectionIpfixConnection) GetPath() string {
	return "/api/objects/ipfix_connection/ipfix_connection/"
}

// RefRequired implements sophos.RestObject
func (*IpfixConnectionIpfixConnection) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpfixConnectionIpfixConnection DELETE path
// Creates or updates the complete object ipfix_connection
func (*IpfixConnectionIpfixConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/ipfix_connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpfixConnectionIpfixConnection PATCH path
// Changes to parts of the object ipfix_connection types
func (*IpfixConnectionIpfixConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/ipfix_connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpfixConnectionIpfixConnection POST path
// Create a new ipfix_connection/ipfix_connection object
func (*IpfixConnectionIpfixConnection) PostPath() string {
	return "/api/objects/ipfix_connection/ipfix_connection/"
}

// PutPath implements sophos.RestObject and returns the IpfixConnectionIpfixConnection PUT path
// Creates or updates the complete object ipfix_connection
func (*IpfixConnectionIpfixConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/ipfix_connection/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpfixConnectionIpfixConnection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipfix_connection/ipfix_connection/%s/usedby", ref)
}
