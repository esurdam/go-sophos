package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Bgp is a generated struct representing the Sophos Bgp Endpoint
// GET /api/nodes/bgp
type Bgp struct {
	BgpAmazonVpc BgpAmazonVpc `json:"bgp_amazon_vpc"`
	BgpFilter    BgpFilter    `json:"bgp_filter"`
	BgpGroup     BgpGroup     `json:"bgp_group"`
	BgpNeighbor  BgpNeighbor  `json:"bgp_neighbor"`
	BgpRouteMap  BgpRouteMap  `json:"bgp_route_map"`
	BgpSystem    BgpSystem    `json:"bgp_system"`
}

var _ sophos.Endpoint = &Bgp{}

var defsBgp = map[string]sophos.RestObject{
	"BgpAmazonVpc": &BgpAmazonVpc{},
	"BgpFilter":    &BgpFilter{},
	"BgpGroup":     &BgpGroup{},
	"BgpNeighbor":  &BgpNeighbor{},
	"BgpRouteMap":  &BgpRouteMap{},
	"BgpSystem":    &BgpSystem{},
}

// RestObjects implements the sophos.Node interface and returns a map of Bgp's Objects
func (Bgp) RestObjects() map[string]sophos.RestObject { return defsBgp }

// GetPath implements sophos.RestGetter
func (*Bgp) GetPath() string { return "/api/nodes/bgp" }

// RefRequired implements sophos.RestGetter
func (*Bgp) RefRequired() (string, bool) { return "", false }

var defBgp = &sophos.Definition{Description: "bgp", Name: "bgp", Link: "/api/definitions/bgp"}

// Definition returns the /api/definitions struct of Bgp
func (Bgp) Definition() sophos.Definition { return *defBgp }

// ApiRoutes returns all known Bgp Paths
func (Bgp) ApiRoutes() []string {
	return []string{
		"/api/objects/bgp/amazon_vpc/",
		"/api/objects/bgp/amazon_vpc/{ref}",
		"/api/objects/bgp/amazon_vpc/{ref}/usedby",
		"/api/objects/bgp/filter/",
		"/api/objects/bgp/filter/{ref}",
		"/api/objects/bgp/filter/{ref}/usedby",
		"/api/objects/bgp/group/",
		"/api/objects/bgp/group/{ref}",
		"/api/objects/bgp/group/{ref}/usedby",
		"/api/objects/bgp/neighbor/",
		"/api/objects/bgp/neighbor/{ref}",
		"/api/objects/bgp/neighbor/{ref}/usedby",
		"/api/objects/bgp/route_map/",
		"/api/objects/bgp/route_map/{ref}",
		"/api/objects/bgp/route_map/{ref}/usedby",
		"/api/objects/bgp/system/",
		"/api/objects/bgp/system/{ref}",
		"/api/objects/bgp/system/{ref}/usedby",
	}
}

// References returns the Bgp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Bgp) References() []string {
	return []string{
		"REF_BgpAmazonVpc",
		"REF_BgpFilter",
		"REF_BgpGroup",
		"REF_BgpNeighbor",
		"REF_BgpRouteMap",
		"REF_BgpSystem",
	}
}

// BgpAmazonVpcs is an Sophos Endpoint subType and implements sophos.RestObject
type BgpAmazonVpcs []BgpAmazonVpc

// BgpAmazonVpc is a generated Sophos object
type BgpAmazonVpc struct {
	Locked       string   `json:"_locked"`
	Reference    string   `json:"_ref"`
	_type        string   `json:"_type"`
	Comment      string   `json:"comment"`
	Custom       string   `json:"custom"`
	Host         string   `json:"host"`
	ID           string   `json:"id"`
	LocalAsn     int64    `json:"local_asn"`
	MaximumPaths int64    `json:"maximum_paths"`
	Name         string   `json:"name"`
	Network      []string `json:"network"`
	RemoteAsn    int64    `json:"remote_asn"`
}

var _ sophos.RestGetter = &BgpAmazonVpc{}

// GetPath implements sophos.RestObject and returns the BgpAmazonVpcs GET path
// Returns all available bgp/amazon_vpc objects
func (*BgpAmazonVpcs) GetPath() string { return "/api/objects/bgp/amazon_vpc/" }

// RefRequired implements sophos.RestObject
func (*BgpAmazonVpcs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpAmazonVpcs GET path
// Returns all available amazon_vpc types
func (b *BgpAmazonVpc) GetPath() string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s", b.Reference)
}

// RefRequired implements sophos.RestObject
func (b *BgpAmazonVpc) RefRequired() (string, bool) { return b.Reference, true }

// DeletePath implements sophos.RestObject and returns the BgpAmazonVpc DELETE path
// Creates or updates the complete object amazon_vpc
func (*BgpAmazonVpc) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpAmazonVpc PATCH path
// Changes to parts of the object amazon_vpc types
func (*BgpAmazonVpc) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpAmazonVpc POST path
// Create a new bgp/amazon_vpc object
func (*BgpAmazonVpc) PostPath() string {
	return "/api/objects/bgp/amazon_vpc/"
}

// PutPath implements sophos.RestObject and returns the BgpAmazonVpc PUT path
// Creates or updates the complete object amazon_vpc
func (*BgpAmazonVpc) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpAmazonVpc) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s/usedby", ref)
}

// GetType implements sophos.Object
func (b *BgpAmazonVpc) GetType() string { return b._type }

// BgpFilter is an Sophos Endpoint subType and implements sophos.RestObject
type BgpFilter []interface{}

var _ sophos.RestObject = &BgpFilter{}

// GetPath implements sophos.RestObject and returns the BgpFilter GET path
// Returns all available bgp/filter objects
func (*BgpFilter) GetPath() string { return "/api/objects/bgp/filter/" }

// RefRequired implements sophos.RestObject
func (*BgpFilter) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the BgpFilter DELETE path
// Creates or updates the complete object filter
func (*BgpFilter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/filter/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpFilter PATCH path
// Changes to parts of the object filter types
func (*BgpFilter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/filter/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpFilter POST path
// Create a new bgp/filter object
func (*BgpFilter) PostPath() string {
	return "/api/objects/bgp/filter/"
}

// PutPath implements sophos.RestObject and returns the BgpFilter PUT path
// Creates or updates the complete object filter
func (*BgpFilter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/filter/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpFilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/filter/%s/usedby", ref)
}

// BgpGroup is an Sophos Endpoint subType and implements sophos.RestObject
type BgpGroup []interface{}

var _ sophos.RestObject = &BgpGroup{}

// GetPath implements sophos.RestObject and returns the BgpGroup GET path
// Returns all available bgp/group objects
func (*BgpGroup) GetPath() string { return "/api/objects/bgp/group/" }

// RefRequired implements sophos.RestObject
func (*BgpGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the BgpGroup DELETE path
// Creates or updates the complete object group
func (*BgpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpGroup PATCH path
// Changes to parts of the object group types
func (*BgpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpGroup POST path
// Create a new bgp/group object
func (*BgpGroup) PostPath() string {
	return "/api/objects/bgp/group/"
}

// PutPath implements sophos.RestObject and returns the BgpGroup PUT path
// Creates or updates the complete object group
func (*BgpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/group/%s/usedby", ref)
}

// BgpNeighbor is an Sophos Endpoint subType and implements sophos.RestObject
type BgpNeighbor []interface{}

var _ sophos.RestObject = &BgpNeighbor{}

// GetPath implements sophos.RestObject and returns the BgpNeighbor GET path
// Returns all available bgp/neighbor objects
func (*BgpNeighbor) GetPath() string { return "/api/objects/bgp/neighbor/" }

// RefRequired implements sophos.RestObject
func (*BgpNeighbor) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the BgpNeighbor DELETE path
// Creates or updates the complete object neighbor
func (*BgpNeighbor) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpNeighbor PATCH path
// Changes to parts of the object neighbor types
func (*BgpNeighbor) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpNeighbor POST path
// Create a new bgp/neighbor object
func (*BgpNeighbor) PostPath() string {
	return "/api/objects/bgp/neighbor/"
}

// PutPath implements sophos.RestObject and returns the BgpNeighbor PUT path
// Creates or updates the complete object neighbor
func (*BgpNeighbor) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpNeighbor) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s/usedby", ref)
}

// BgpRouteMap is an Sophos Endpoint subType and implements sophos.RestObject
type BgpRouteMap []interface{}

var _ sophos.RestObject = &BgpRouteMap{}

// GetPath implements sophos.RestObject and returns the BgpRouteMap GET path
// Returns all available bgp/route_map objects
func (*BgpRouteMap) GetPath() string { return "/api/objects/bgp/route_map/" }

// RefRequired implements sophos.RestObject
func (*BgpRouteMap) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the BgpRouteMap DELETE path
// Creates or updates the complete object route_map
func (*BgpRouteMap) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpRouteMap PATCH path
// Changes to parts of the object route_map types
func (*BgpRouteMap) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpRouteMap POST path
// Create a new bgp/route_map object
func (*BgpRouteMap) PostPath() string {
	return "/api/objects/bgp/route_map/"
}

// PutPath implements sophos.RestObject and returns the BgpRouteMap PUT path
// Creates or updates the complete object route_map
func (*BgpRouteMap) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpRouteMap) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s/usedby", ref)
}

// BgpSystem is an Sophos Endpoint subType and implements sophos.RestObject
type BgpSystem []interface{}

var _ sophos.RestObject = &BgpSystem{}

// GetPath implements sophos.RestObject and returns the BgpSystem GET path
// Returns all available bgp/system objects
func (*BgpSystem) GetPath() string { return "/api/objects/bgp/system/" }

// RefRequired implements sophos.RestObject
func (*BgpSystem) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the BgpSystem DELETE path
// Creates or updates the complete object system
func (*BgpSystem) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/system/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the BgpSystem PATCH path
// Changes to parts of the object system types
func (*BgpSystem) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/system/%s", ref)
}

// PostPath implements sophos.RestObject and returns the BgpSystem POST path
// Create a new bgp/system object
func (*BgpSystem) PostPath() string {
	return "/api/objects/bgp/system/"
}

// PutPath implements sophos.RestObject and returns the BgpSystem PUT path
// Creates or updates the complete object system
func (*BgpSystem) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/system/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpSystem) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/system/%s/usedby", ref)
}
