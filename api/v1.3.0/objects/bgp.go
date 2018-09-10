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
	ObjectType   string   `json:"_type"`
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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpAmazonVpc) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/amazon_vpc/%s/usedby", ref)
}

// GetType implements sophos.Object
func (b *BgpAmazonVpc) GetType() string { return b.ObjectType }

// BgpFilters is an Sophos Endpoint subType and implements sophos.RestObject
type BgpFilters []BgpFilter

// BgpFilter represents a UTM BGP filter list
type BgpFilter struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// Action can be one of: []string{"permit", "deny"}
	Action  string        `json:"action"`
	Address []interface{} `json:"address"`
	AsRegex []interface{} `json:"as_regex"`
	Comment string        `json:"comment"`
	Name    string        `json:"name"`
	// Type can be one of: []string{"as_number", "ip_address"}
	Type string `json:"type"`
}

var _ sophos.RestGetter = &BgpFilter{}

// GetPath implements sophos.RestObject and returns the BgpFilters GET path
// Returns all available bgp/filter objects
func (*BgpFilters) GetPath() string { return "/api/objects/bgp/filter/" }

// RefRequired implements sophos.RestObject
func (*BgpFilters) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpFilters GET path
// Returns all available filter types
func (b *BgpFilter) GetPath() string { return fmt.Sprintf("/api/objects/bgp/filter/%s", b.Reference) }

// RefRequired implements sophos.RestObject
func (b *BgpFilter) RefRequired() (string, bool) { return b.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpFilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/filter/%s/usedby", ref)
}

// BgpGroups is an Sophos Endpoint subType and implements sophos.RestObject
type BgpGroups []BgpGroup

// BgpGroup represents a UTM group
type BgpGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &BgpGroup{}

// GetPath implements sophos.RestObject and returns the BgpGroups GET path
// Returns all available bgp/group objects
func (*BgpGroups) GetPath() string { return "/api/objects/bgp/group/" }

// RefRequired implements sophos.RestObject
func (*BgpGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpGroups GET path
// Returns all available group types
func (b *BgpGroup) GetPath() string { return fmt.Sprintf("/api/objects/bgp/group/%s", b.Reference) }

// RefRequired implements sophos.RestObject
func (b *BgpGroup) RefRequired() (string, bool) { return b.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/group/%s/usedby", ref)
}

// BgpNeighbors is an Sophos Endpoint subType and implements sophos.RestObject
type BgpNeighbors []BgpNeighbor

// BgpNeighbor represents a UTM BGP neighbor
type BgpNeighbor struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// FilterIn description: REF(bgp/filter)
	FilterIn string `json:"filter_in"`
	Name     string `json:"name"`
	Asn      int    `json:"asn"`
	// DefaultOriginate default value is false
	DefaultOriginate bool `json:"default_originate"`
	// Host description: REF(network/host)
	Host string `json:"host"`
	// NextHopSelf default value is false
	NextHopSelf bool `json:"next_hop_self"`
	// Password description: (REGEX)
	Password string `json:"password"`
	// Status default value is true
	Status bool `json:"status"`
	Weight int  `json:"weight"`
	// Authentication can be one of: []string{"null", "password"}
	// Authentication default value is "null"
	Authentication string `json:"authentication"`
	// RouteIn description: REF(bgp/route_map)
	RouteIn string `json:"route_in"`
	// SoftReconfiguration default value is true
	SoftReconfiguration bool `json:"soft_reconfiguration"`
	// FilterOut description: REF(bgp/filter)
	FilterOut string `json:"filter_out"`
	// Multihop default value is false
	Multihop bool `json:"multihop"`
	// RouteOut description: REF(bgp/route_map)
	RouteOut string `json:"route_out"`
	Comment  string `json:"comment"`
}

var _ sophos.RestGetter = &BgpNeighbor{}

// GetPath implements sophos.RestObject and returns the BgpNeighbors GET path
// Returns all available bgp/neighbor objects
func (*BgpNeighbors) GetPath() string { return "/api/objects/bgp/neighbor/" }

// RefRequired implements sophos.RestObject
func (*BgpNeighbors) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpNeighbors GET path
// Returns all available neighbor types
func (b *BgpNeighbor) GetPath() string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s", b.Reference)
}

// RefRequired implements sophos.RestObject
func (b *BgpNeighbor) RefRequired() (string, bool) { return b.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpNeighbor) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/neighbor/%s/usedby", ref)
}

// BgpRouteMaps is an Sophos Endpoint subType and implements sophos.RestObject
type BgpRouteMaps []BgpRouteMap

// BgpRouteMap represents a UTM route_map
type BgpRouteMap struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// Prepend description: (REGEX)
	// Prepend default value is ""
	Prepend    string        `json:"prepend"`
	AsRegex    []interface{} `json:"as_regex"`
	Name       string        `json:"name"`
	Metric     int           `json:"metric"`
	Preference int           `json:"preference"`
	// Type can be one of: []string{"as_number", "ip_address"}
	Type    string        `json:"type"`
	Weight  int           `json:"weight"`
	Address []interface{} `json:"address"`
	Comment string        `json:"comment"`
}

var _ sophos.RestGetter = &BgpRouteMap{}

// GetPath implements sophos.RestObject and returns the BgpRouteMaps GET path
// Returns all available bgp/route_map objects
func (*BgpRouteMaps) GetPath() string { return "/api/objects/bgp/route_map/" }

// RefRequired implements sophos.RestObject
func (*BgpRouteMaps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpRouteMaps GET path
// Returns all available route_map types
func (b *BgpRouteMap) GetPath() string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s", b.Reference)
}

// RefRequired implements sophos.RestObject
func (b *BgpRouteMap) RefRequired() (string, bool) { return b.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpRouteMap) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/route_map/%s/usedby", ref)
}

// BgpSystems is an Sophos Endpoint subType and implements sophos.RestObject
type BgpSystems []BgpSystem

// BgpSystem represents a UTM BGP system
type BgpSystem struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Asn        int    `json:"asn"`
	// Id description: (IPADDR)
	Id string `json:"id"`
	// InstallRoutes default value is true
	InstallRoutes bool   `json:"install_routes"`
	MaximumPaths  int    `json:"maximum_paths"`
	Name          string `json:"name"`
	Comment       string `json:"comment"`
	// Custom default value is ""
	Custom   string        `json:"custom"`
	Neighbor []interface{} `json:"neighbor"`
	Network  []interface{} `json:"network"`
	// Status default value is true
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &BgpSystem{}

// GetPath implements sophos.RestObject and returns the BgpSystems GET path
// Returns all available bgp/system objects
func (*BgpSystems) GetPath() string { return "/api/objects/bgp/system/" }

// RefRequired implements sophos.RestObject
func (*BgpSystems) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the BgpSystems GET path
// Returns all available system types
func (b *BgpSystem) GetPath() string { return fmt.Sprintf("/api/objects/bgp/system/%s", b.Reference) }

// RefRequired implements sophos.RestObject
func (b *BgpSystem) RefRequired() (string, bool) { return b.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*BgpSystem) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/bgp/system/%s/usedby", ref)
}
