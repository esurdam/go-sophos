package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// PimSm is a generated struct representing the Sophos PimSm Endpoint
// GET /api/nodes/pim_sm
type PimSm struct {
	AutoPfOut       string        `json:"auto_pf_out"`
	AutoPfrule      int64         `json:"auto_pfrule"`
	Debug           int64         `json:"debug"`
	Interfaces      []interface{} `json:"interfaces"`
	RpRouters       []interface{} `json:"rp_routers"`
	SptSwitchBytes  int64         `json:"spt_switch_bytes"`
	SptSwitchStatus int64         `json:"spt_switch_status"`
	Status          int64         `json:"status"`
}

var _ sophos.Endpoint = &PimSm{}

var defsPimSm = map[string]sophos.RestObject{
	"PimSmGroup":     &PimSmGroup{},
	"PimSmInterface": &PimSmInterface{},
	"PimSmRoute":     &PimSmRoute{},
	"PimSmRpRouter":  &PimSmRpRouter{},
}

// RestObjects implements the sophos.Node interface and returns a map of PimSm's Objects
func (PimSm) RestObjects() map[string]sophos.RestObject { return defsPimSm }

// GetPath implements sophos.RestGetter
func (*PimSm) GetPath() string { return "/api/nodes/pim_sm" }

// RefRequired implements sophos.RestGetter
func (*PimSm) RefRequired() (string, bool) { return "", false }

var defPimSm = &sophos.Definition{Description: "pim_sm", Name: "pim_sm", Link: "/api/definitions/pim_sm"}

// Definition returns the /api/definitions struct of PimSm
func (PimSm) Definition() sophos.Definition { return *defPimSm }

// ApiRoutes returns all known PimSm Paths
func (PimSm) ApiRoutes() []string {
	return []string{
		"/api/objects/pim_sm/group/",
		"/api/objects/pim_sm/group/{ref}",
		"/api/objects/pim_sm/group/{ref}/usedby",
		"/api/objects/pim_sm/interface/",
		"/api/objects/pim_sm/interface/{ref}",
		"/api/objects/pim_sm/interface/{ref}/usedby",
		"/api/objects/pim_sm/route/",
		"/api/objects/pim_sm/route/{ref}",
		"/api/objects/pim_sm/route/{ref}/usedby",
		"/api/objects/pim_sm/rp_router/",
		"/api/objects/pim_sm/rp_router/{ref}",
		"/api/objects/pim_sm/rp_router/{ref}/usedby",
	}
}

// References returns the PimSm's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (PimSm) References() []string {
	return []string{
		"REF_PimSmGroup",
		"REF_PimSmInterface",
		"REF_PimSmRoute",
		"REF_PimSmRpRouter",
	}
}

// PimSmGroups is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmGroups []PimSmGroup

// PimSmGroup represents a UTM group
type PimSmGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &PimSmGroup{}

// GetPath implements sophos.RestObject and returns the PimSmGroups GET path
// Returns all available pim_sm/group objects
func (*PimSmGroups) GetPath() string { return "/api/objects/pim_sm/group/" }

// RefRequired implements sophos.RestObject
func (*PimSmGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PimSmGroups GET path
// Returns all available group types
func (p *PimSmGroup) GetPath() string { return fmt.Sprintf("/api/objects/pim_sm/group/%s", p.Reference) }

// RefRequired implements sophos.RestObject
func (p *PimSmGroup) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PimSmGroup DELETE path
// Creates or updates the complete object group
func (*PimSmGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PimSmGroup PATCH path
// Changes to parts of the object group types
func (*PimSmGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PimSmGroup POST path
// Create a new pim_sm/group object
func (*PimSmGroup) PostPath() string {
	return "/api/objects/pim_sm/group/"
}

// PutPath implements sophos.RestObject and returns the PimSmGroup PUT path
// Creates or updates the complete object group
func (*PimSmGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/group/%s/usedby", ref)
}

// PimSmInterfaces is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmInterfaces []PimSmInterface

// PimSmInterface represents a UTM multicast routing interface
type PimSmInterface struct {
	Locked       string        `json:"_locked"`
	ObjectType   string        `json:"_type"`
	Reference    string        `json:"_ref"`
	DrPriority   int           `json:"dr_priority"`
	IgmpVersions []interface{} `json:"igmp_versions"`
	// Interface description: REF(interface/*)
	Interface string `json:"interface"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
}

var _ sophos.RestGetter = &PimSmInterface{}

// GetPath implements sophos.RestObject and returns the PimSmInterfaces GET path
// Returns all available pim_sm/interface objects
func (*PimSmInterfaces) GetPath() string { return "/api/objects/pim_sm/interface/" }

// RefRequired implements sophos.RestObject
func (*PimSmInterfaces) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PimSmInterfaces GET path
// Returns all available interface types
func (p *PimSmInterface) GetPath() string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PimSmInterface) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PimSmInterface DELETE path
// Creates or updates the complete object interface
func (*PimSmInterface) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PimSmInterface PATCH path
// Changes to parts of the object interface types
func (*PimSmInterface) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PimSmInterface POST path
// Create a new pim_sm/interface object
func (*PimSmInterface) PostPath() string {
	return "/api/objects/pim_sm/interface/"
}

// PutPath implements sophos.RestObject and returns the PimSmInterface PUT path
// Creates or updates the complete object interface
func (*PimSmInterface) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s/usedby", ref)
}

// PimSmRoutes is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmRoutes []PimSmRoute

// PimSmRoute represents a UTM multicast route
type PimSmRoute struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	// Gateway description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	// Gateway default value is ""
	Gateway string `json:"gateway"`
	// Interface description: REF(interface/*)
	// Interface default value is ""
	Interface string `json:"interface"`
	Name      string `json:"name"`
	// Network description: REF(network/*)
	Network string `json:"network"`
	// Status default value is false
	Status bool `json:"status"`
	// Type can be one of: []string{"gateway", "interface"}
	Type string `json:"type"`
}

var _ sophos.RestGetter = &PimSmRoute{}

// GetPath implements sophos.RestObject and returns the PimSmRoutes GET path
// Returns all available pim_sm/route objects
func (*PimSmRoutes) GetPath() string { return "/api/objects/pim_sm/route/" }

// RefRequired implements sophos.RestObject
func (*PimSmRoutes) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PimSmRoutes GET path
// Returns all available route types
func (p *PimSmRoute) GetPath() string { return fmt.Sprintf("/api/objects/pim_sm/route/%s", p.Reference) }

// RefRequired implements sophos.RestObject
func (p *PimSmRoute) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PimSmRoute DELETE path
// Creates or updates the complete object route
func (*PimSmRoute) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/route/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PimSmRoute PATCH path
// Changes to parts of the object route types
func (*PimSmRoute) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/route/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PimSmRoute POST path
// Create a new pim_sm/route object
func (*PimSmRoute) PostPath() string {
	return "/api/objects/pim_sm/route/"
}

// PutPath implements sophos.RestObject and returns the PimSmRoute PUT path
// Creates or updates the complete object route
func (*PimSmRoute) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/route/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmRoute) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/route/%s/usedby", ref)
}

// PimSmRpRouters is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmRpRouters []PimSmRpRouter

// PimSmRpRouter represents a UTM rendezvous point router
type PimSmRpRouter struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	// Host description: REF(network/host), REF(network/dns_host), REF(network/interface_address)
	Host            string        `json:"host"`
	MulticastGroups []interface{} `json:"multicast_groups"`
	Name            string        `json:"name"`
	RpPriority      int           `json:"rp_priority"`
}

var _ sophos.RestGetter = &PimSmRpRouter{}

// GetPath implements sophos.RestObject and returns the PimSmRpRouters GET path
// Returns all available pim_sm/rp_router objects
func (*PimSmRpRouters) GetPath() string { return "/api/objects/pim_sm/rp_router/" }

// RefRequired implements sophos.RestObject
func (*PimSmRpRouters) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PimSmRpRouters GET path
// Returns all available rp_router types
func (p *PimSmRpRouter) GetPath() string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PimSmRpRouter) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PimSmRpRouter DELETE path
// Creates or updates the complete object rp_router
func (*PimSmRpRouter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PimSmRpRouter PATCH path
// Changes to parts of the object rp_router types
func (*PimSmRpRouter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PimSmRpRouter POST path
// Create a new pim_sm/rp_router object
func (*PimSmRpRouter) PostPath() string {
	return "/api/objects/pim_sm/rp_router/"
}

// PutPath implements sophos.RestObject and returns the PimSmRpRouter PUT path
// Creates or updates the complete object rp_router
func (*PimSmRpRouter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmRpRouter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s/usedby", ref)
}
