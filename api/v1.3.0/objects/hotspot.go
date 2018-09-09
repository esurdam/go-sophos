package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Hotspot is a generated struct representing the Sophos Hotspot Endpoint
// GET /api/nodes/hotspot
type Hotspot struct {
	Cert            string        `json:"cert"`
	DeleteDays      int64         `json:"delete_days"`
	SslPortal       int64         `json:"ssl_portal"`
	Status          int64         `json:"status"`
	TransparentSkip []interface{} `json:"transparent_skip"`
}

var defsHotspot = map[string]sophos.RestObject{
	"HotspotGroup":   &HotspotGroup{},
	"HotspotPortal":  &HotspotPortal{},
	"HotspotVoucher": &HotspotVoucher{},
}

// RestObjects implements the sophos.Node interface and returns a map of Hotspot's Objects
func (Hotspot) RestObjects() map[string]sophos.RestObject { return defsHotspot }

// GetPath implements sophos.RestGetter
func (*Hotspot) GetPath() string { return "/api/nodes/hotspot" }

// RefRequired implements sophos.RestGetter
func (*Hotspot) RefRequired() (string, bool) { return "", false }

var defHotspot = &sophos.Definition{Description: "hotspot", Name: "hotspot", Link: "/api/definitions/hotspot"}

// Definition returns the /api/definitions struct of Hotspot
func (Hotspot) Definition() sophos.Definition { return *defHotspot }

// ApiRoutes returns all known Hotspot Paths
func (Hotspot) ApiRoutes() []string {
	return []string{
		"/api/objects/hotspot/group/",
		"/api/objects/hotspot/group/{ref}",
		"/api/objects/hotspot/group/{ref}/usedby",
		"/api/objects/hotspot/portal/",
		"/api/objects/hotspot/portal/{ref}",
		"/api/objects/hotspot/portal/{ref}/usedby",
		"/api/objects/hotspot/voucher/",
		"/api/objects/hotspot/voucher/{ref}",
		"/api/objects/hotspot/voucher/{ref}/usedby",
	}
}

// References returns the Hotspot's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Hotspot) References() []string {
	return []string{
		"REF_HotspotGroup",
		"REF_HotspotPortal",
		"REF_HotspotVoucher",
	}
}

// HotspotGroup is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotGroup []interface{}

// GetPath implements sophos.RestObject and returns the HotspotGroup GET path
// Returns all available hotspot/group objects
func (*HotspotGroup) GetPath() string { return "/api/objects/hotspot/group/" }

// RefRequired implements sophos.RestObject
func (*HotspotGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotGroup DELETE path
// Creates or updates the complete object group
func (*HotspotGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotGroup PATCH path
// Changes to parts of the object group types
func (*HotspotGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotGroup POST path
// Create a new hotspot/group object
func (*HotspotGroup) PostPath() string {
	return "/api/objects/hotspot/group/"
}

// PutPath implements sophos.RestObject and returns the HotspotGroup PUT path
// Creates or updates the complete object group
func (*HotspotGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s/usedby", ref)
}

// HotspotPortal is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotPortal []interface{}

// GetPath implements sophos.RestObject and returns the HotspotPortal GET path
// Returns all available hotspot/portal objects
func (*HotspotPortal) GetPath() string { return "/api/objects/hotspot/portal/" }

// RefRequired implements sophos.RestObject
func (*HotspotPortal) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotPortal DELETE path
// Creates or updates the complete object portal
func (*HotspotPortal) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotPortal PATCH path
// Changes to parts of the object portal types
func (*HotspotPortal) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotPortal POST path
// Create a new hotspot/portal object
func (*HotspotPortal) PostPath() string {
	return "/api/objects/hotspot/portal/"
}

// PutPath implements sophos.RestObject and returns the HotspotPortal PUT path
// Creates or updates the complete object portal
func (*HotspotPortal) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotPortal) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s/usedby", ref)
}

// HotspotVoucher is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotVoucher []interface{}

// GetPath implements sophos.RestObject and returns the HotspotVoucher GET path
// Returns all available hotspot/voucher objects
func (*HotspotVoucher) GetPath() string { return "/api/objects/hotspot/voucher/" }

// RefRequired implements sophos.RestObject
func (*HotspotVoucher) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotVoucher DELETE path
// Creates or updates the complete object voucher
func (*HotspotVoucher) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotVoucher PATCH path
// Changes to parts of the object voucher types
func (*HotspotVoucher) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotVoucher POST path
// Create a new hotspot/voucher object
func (*HotspotVoucher) PostPath() string {
	return "/api/objects/hotspot/voucher/"
}

// PutPath implements sophos.RestObject and returns the HotspotVoucher PUT path
// Creates or updates the complete object voucher
func (*HotspotVoucher) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotVoucher) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s/usedby", ref)
}
