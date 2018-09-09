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

// PimSmGroup is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmGroup []interface{}

// GetPath implements sophos.RestObject and returns the PimSmGroup GET path
// Returns all available pim_sm/group objects
func (*PimSmGroup) GetPath() string { return "/api/objects/pim_sm/group/" }

// RefRequired implements sophos.RestObject
func (*PimSmGroup) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/group/%s/usedby", ref)
}

// PimSmInterface is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmInterface []interface{}

// GetPath implements sophos.RestObject and returns the PimSmInterface GET path
// Returns all available pim_sm/interface objects
func (*PimSmInterface) GetPath() string { return "/api/objects/pim_sm/interface/" }

// RefRequired implements sophos.RestObject
func (*PimSmInterface) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/interface/%s/usedby", ref)
}

// PimSmRoute is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmRoute []interface{}

// GetPath implements sophos.RestObject and returns the PimSmRoute GET path
// Returns all available pim_sm/route objects
func (*PimSmRoute) GetPath() string { return "/api/objects/pim_sm/route/" }

// RefRequired implements sophos.RestObject
func (*PimSmRoute) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmRoute) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/route/%s/usedby", ref)
}

// PimSmRpRouter is an Sophos Endpoint subType and implements sophos.RestObject
type PimSmRpRouter []interface{}

// GetPath implements sophos.RestObject and returns the PimSmRpRouter GET path
// Returns all available pim_sm/rp_router objects
func (*PimSmRpRouter) GetPath() string { return "/api/objects/pim_sm/rp_router/" }

// RefRequired implements sophos.RestObject
func (*PimSmRpRouter) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*PimSmRpRouter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pim_sm/rp_router/%s/usedby", ref)
}
