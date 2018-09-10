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

var _ sophos.Endpoint = &Hotspot{}

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

// HotspotGroups is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotGroups []HotspotGroup

// HotspotGroup represents a UTM group
type HotspotGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &HotspotGroup{}

// GetPath implements sophos.RestObject and returns the HotspotGroups GET path
// Returns all available hotspot/group objects
func (*HotspotGroups) GetPath() string { return "/api/objects/hotspot/group/" }

// RefRequired implements sophos.RestObject
func (*HotspotGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HotspotGroups GET path
// Returns all available group types
func (h *HotspotGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HotspotGroup) RefRequired() (string, bool) { return h.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s/usedby", ref)
}

// HotspotPortals is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotPortals []HotspotPortal

// HotspotPortal represents a UTM hotspot
type HotspotPortal struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// HostnameType can be one of: []string{"none", "custom"}
	// HostnameType default value is "none"
	HostnameType string `json:"hostname_type"`
	Name         string `json:"name"`
	SmsText      string `json:"sms_text"`
	Terms        string `json:"terms"`
	// VoucherQrcode default value is false
	VoucherQrcode bool   `json:"voucher_qrcode"`
	Description   string `json:"description"`
	// FiasCodeset can be one of: []string{"cp850", "cp1252"}
	// FiasCodeset default value is "cp850"
	FiasCodeset string `json:"fias_codeset"`
	// FiasServer description: REF(network/host), REF(network/dns_host)
	// FiasServer default value is ""
	FiasServer string        `json:"fias_server"`
	Vouchers   []interface{} `json:"vouchers"`
	Maclimit   int           `json:"maclimit"`
	AdminUsers []interface{} `json:"admin_users"`
	// Logo default value is ""
	Logo string `json:"logo"`
	// LogoResize default value is false
	LogoResize bool `json:"logo_resize"`
	// CustomAssets description: (HASH)
	CustomAssets interface{} `json:"custom_assets"`
	// Pagesize default value is "a4"
	Pagesize     string        `json:"pagesize"`
	Title        string        `json:"title"`
	HotspotUsers []interface{} `json:"hotspot_users"`
	// SslRedirect default value is false
	SslRedirect bool `json:"ssl_redirect"`
	// VoucherTemplate description: (HASH)
	VoucherTemplate interface{}   `json:"voucher_template"`
	VouchersPerPage int           `json:"vouchers_per_page"`
	Comment         string        `json:"comment"`
	Interfaces      []interface{} `json:"interfaces"`
	Mail            []interface{} `json:"mail"`
	// PwTime description: (TIME)
	PwTime      string `json:"pw_time"`
	RedirectUrl string `json:"redirect_url"`
	// SyncPsk default value is false
	SyncPsk bool `json:"sync_psk"`
	// Type can be one of: []string{"terms", "password", "voucher", "backend_auth", "sms", "fias"}
	Type string `json:"type"`
	// CustomizationType can be one of: []string{"basic", "full"}
	// CustomizationType default value is "basic"
	CustomizationType string `json:"customization_type"`
	// FiasPort description: REF(service/tcp)
	// FiasPort default value is ""
	FiasPort string `json:"fias_port"`
	// LogoFilename default value is "default_logo.png"
	LogoFilename string `json:"logo_filename"`
	Expiry       int    `json:"expiry"`
	// Hostname description: REF(network/dns_host)
	// Hostname default value is ""
	Hostname string `json:"hostname"`
	// Template description: (HASH)
	Template interface{} `json:"template"`
}

var _ sophos.RestGetter = &HotspotPortal{}

// GetPath implements sophos.RestObject and returns the HotspotPortals GET path
// Returns all available hotspot/portal objects
func (*HotspotPortals) GetPath() string { return "/api/objects/hotspot/portal/" }

// RefRequired implements sophos.RestObject
func (*HotspotPortals) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HotspotPortals GET path
// Returns all available portal types
func (h *HotspotPortal) GetPath() string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HotspotPortal) RefRequired() (string, bool) { return h.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotPortal) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s/usedby", ref)
}

// HotspotVouchers is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotVouchers []HotspotVoucher

// HotspotVoucher represents a UTM voucher definiton
type HotspotVoucher struct {
	Locked       string `json:"_locked"`
	ObjectType   string `json:"_type"`
	Reference    string `json:"_ref"`
	Name         string `json:"name"`
	Timequota    int    `json:"timequota"`
	Trafficlimit int    `json:"trafficlimit"`
	Comment      string `json:"comment"`
	Expiry       int    `json:"expiry"`
}

var _ sophos.RestGetter = &HotspotVoucher{}

// GetPath implements sophos.RestObject and returns the HotspotVouchers GET path
// Returns all available hotspot/voucher objects
func (*HotspotVouchers) GetPath() string { return "/api/objects/hotspot/voucher/" }

// RefRequired implements sophos.RestObject
func (*HotspotVouchers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HotspotVouchers GET path
// Returns all available voucher types
func (h *HotspotVoucher) GetPath() string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HotspotVoucher) RefRequired() (string, bool) { return h.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*HotspotVoucher) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s/usedby", ref)
}
