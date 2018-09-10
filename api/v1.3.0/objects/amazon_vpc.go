package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// AmazonVpc is a generated struct representing the Sophos AmazonVpc Endpoint
// GET /api/nodes/amazon_vpc
type AmazonVpc struct {
	AutoPfrule  int64    `json:"auto_pfrule"`
	Connections []string `json:"connections"`
	Networks    []string `json:"networks"`
	Status      int64    `json:"status"`
}

var _ sophos.Endpoint = &AmazonVpc{}

var defsAmazonVpc = map[string]sophos.RestObject{
	"AmazonVpcConnection": &AmazonVpcConnection{},
	"AmazonVpcGroup":      &AmazonVpcGroup{},
	"AmazonVpcTunnel":     &AmazonVpcTunnel{},
}

// RestObjects implements the sophos.Node interface and returns a map of AmazonVpc's Objects
func (AmazonVpc) RestObjects() map[string]sophos.RestObject { return defsAmazonVpc }

// GetPath implements sophos.RestGetter
func (*AmazonVpc) GetPath() string { return "/api/nodes/amazon_vpc" }

// RefRequired implements sophos.RestGetter
func (*AmazonVpc) RefRequired() (string, bool) { return "", false }

var defAmazonVpc = &sophos.Definition{Description: "amazon_vpc", Name: "amazon_vpc", Link: "/api/definitions/amazon_vpc"}

// Definition returns the /api/definitions struct of AmazonVpc
func (AmazonVpc) Definition() sophos.Definition { return *defAmazonVpc }

// ApiRoutes returns all known AmazonVpc Paths
func (AmazonVpc) ApiRoutes() []string {
	return []string{
		"/api/objects/amazon_vpc/connection/",
		"/api/objects/amazon_vpc/connection/{ref}",
		"/api/objects/amazon_vpc/connection/{ref}/usedby",
		"/api/objects/amazon_vpc/group/",
		"/api/objects/amazon_vpc/group/{ref}",
		"/api/objects/amazon_vpc/group/{ref}/usedby",
		"/api/objects/amazon_vpc/tunnel/",
		"/api/objects/amazon_vpc/tunnel/{ref}",
		"/api/objects/amazon_vpc/tunnel/{ref}/usedby",
	}
}

// References returns the AmazonVpc's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (AmazonVpc) References() []string {
	return []string{
		"REF_AmazonVpcConnection",
		"REF_AmazonVpcGroup",
		"REF_AmazonVpcTunnel",
	}
}

// AmazonVpcConnections is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcConnections []AmazonVpcConnection

// AmazonVpcConnection is a generated Sophos object
type AmazonVpcConnection struct {
	Locked     string   `json:"_locked"`
	Reference  string   `json:"_ref"`
	ObjectType string   `json:"_type"`
	Comment    string   `json:"comment"`
	Dev        string   `json:"dev"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Region     string   `json:"region"`
	Status     bool     `json:"status"`
	Tunnel     []string `json:"tunnel"`
	VpcGateway string   `json:"vpc_gateway"`
	VpcID      string   `json:"vpc_id"`
	VpcNetmask int64    `json:"vpc_netmask"`
	VpcNetwork string   `json:"vpc_network"`
}

var _ sophos.RestGetter = &AmazonVpcConnection{}

// GetPath implements sophos.RestObject and returns the AmazonVpcConnections GET path
// Returns all available amazon_vpc/connection objects
func (*AmazonVpcConnections) GetPath() string { return "/api/objects/amazon_vpc/connection/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AmazonVpcConnections GET path
// Returns all available connection types
func (a *AmazonVpcConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AmazonVpcConnection) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AmazonVpcConnection DELETE path
// Creates or updates the complete object connection
func (*AmazonVpcConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcConnection PATCH path
// Changes to parts of the object connection types
func (*AmazonVpcConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcConnection POST path
// Create a new amazon_vpc/connection object
func (*AmazonVpcConnection) PostPath() string {
	return "/api/objects/amazon_vpc/connection/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcConnection PUT path
// Creates or updates the complete object connection
func (*AmazonVpcConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AmazonVpcConnection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/connection/%s/usedby", ref)
}

// GetType implements sophos.Object
func (a *AmazonVpcConnection) GetType() string { return a.ObjectType }

// AmazonVpcGroups is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcGroups []AmazonVpcGroup

// AmazonVpcGroup represents a UTM group
type AmazonVpcGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &AmazonVpcGroup{}

// GetPath implements sophos.RestObject and returns the AmazonVpcGroups GET path
// Returns all available amazon_vpc/group objects
func (*AmazonVpcGroups) GetPath() string { return "/api/objects/amazon_vpc/group/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AmazonVpcGroups GET path
// Returns all available group types
func (a *AmazonVpcGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AmazonVpcGroup) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AmazonVpcGroup DELETE path
// Creates or updates the complete object group
func (*AmazonVpcGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcGroup PATCH path
// Changes to parts of the object group types
func (*AmazonVpcGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcGroup POST path
// Create a new amazon_vpc/group object
func (*AmazonVpcGroup) PostPath() string {
	return "/api/objects/amazon_vpc/group/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcGroup PUT path
// Creates or updates the complete object group
func (*AmazonVpcGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AmazonVpcGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/group/%s/usedby", ref)
}

// AmazonVpcTunnels is an Sophos Endpoint subType and implements sophos.RestObject
type AmazonVpcTunnels []AmazonVpcTunnel

// AmazonVpcTunnel is a generated Sophos object
type AmazonVpcTunnel struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Address    string `json:"address"`
	Bgp        string `json:"bgp"`
	Comment    string `json:"comment"`
	Ipsec      string `json:"ipsec"`
	Name       string `json:"name"`
	Netmask    int64  `json:"netmask"`
}

var _ sophos.RestGetter = &AmazonVpcTunnel{}

// GetPath implements sophos.RestObject and returns the AmazonVpcTunnels GET path
// Returns all available amazon_vpc/tunnel objects
func (*AmazonVpcTunnels) GetPath() string { return "/api/objects/amazon_vpc/tunnel/" }

// RefRequired implements sophos.RestObject
func (*AmazonVpcTunnels) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AmazonVpcTunnels GET path
// Returns all available tunnel types
func (a *AmazonVpcTunnel) GetPath() string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AmazonVpcTunnel) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AmazonVpcTunnel DELETE path
// Creates or updates the complete object tunnel
func (*AmazonVpcTunnel) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AmazonVpcTunnel PATCH path
// Changes to parts of the object tunnel types
func (*AmazonVpcTunnel) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AmazonVpcTunnel POST path
// Create a new amazon_vpc/tunnel object
func (*AmazonVpcTunnel) PostPath() string {
	return "/api/objects/amazon_vpc/tunnel/"
}

// PutPath implements sophos.RestObject and returns the AmazonVpcTunnel PUT path
// Creates or updates the complete object tunnel
func (*AmazonVpcTunnel) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AmazonVpcTunnel) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/amazon_vpc/tunnel/%s/usedby", ref)
}

// GetType implements sophos.Object
func (a *AmazonVpcTunnel) GetType() string { return a.ObjectType }
