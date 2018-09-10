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

var _ sophos.Endpoint = &Ospf{}

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

// OspfAreas is an Sophos Endpoint subType and implements sophos.RestObject
type OspfAreas []OspfArea

// OspfArea represents a UTM OSPF area
type OspfArea struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Id description: (IPADDR)
	Id         string        `json:"id"`
	Interfaces []interface{} `json:"interfaces"`
	Name       string        `json:"name"`
	// Type can be one of: []string{"normal", "stub", "nssa", "stub no-summary", "nssa no-summary"}
	// Type default value is "normal"
	Type         string        `json:"type"`
	VirtualLinks []interface{} `json:"virtual_links"`
	// Authentication can be one of: []string{"message-digest", "plain-text", "null"}
	Authentication string `json:"authentication"`
	Comment        string `json:"comment"`
	DefaultCost    int    `json:"default_cost"`
}

var _ sophos.RestGetter = &OspfArea{}

// GetPath implements sophos.RestObject and returns the OspfAreas GET path
// Returns all available ospf/area objects
func (*OspfAreas) GetPath() string { return "/api/objects/ospf/area/" }

// RefRequired implements sophos.RestObject
func (*OspfAreas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OspfAreas GET path
// Returns all available area types
func (o *OspfArea) GetPath() string { return fmt.Sprintf("/api/objects/ospf/area/%s", o.Reference) }

// RefRequired implements sophos.RestObject
func (o *OspfArea) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OspfArea) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s/usedby", ref)
}

// OspfGroups is an Sophos Endpoint subType and implements sophos.RestObject
type OspfGroups []OspfGroup

// OspfGroup represents a UTM group
type OspfGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &OspfGroup{}

// GetPath implements sophos.RestObject and returns the OspfGroups GET path
// Returns all available ospf/group objects
func (*OspfGroups) GetPath() string { return "/api/objects/ospf/group/" }

// RefRequired implements sophos.RestObject
func (*OspfGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OspfGroups GET path
// Returns all available group types
func (o *OspfGroup) GetPath() string { return fmt.Sprintf("/api/objects/ospf/group/%s", o.Reference) }

// RefRequired implements sophos.RestObject
func (o *OspfGroup) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OspfGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s/usedby", ref)
}

// OspfInterfaces is an Sophos Endpoint subType and implements sophos.RestObject
type OspfInterfaces []OspfInterface

// OspfInterface represents a UTM OSPF interface
type OspfInterface struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// RetransmitInterval description: Constraints: 0, 3-65535
	RetransmitInterval int `json:"retransmit_interval"`
	// Authentication can be one of: []string{"message-digest", "plain-text", "null"}
	Authentication string `json:"authentication"`
	// AuthenticationKey description: (REGEX)
	AuthenticationKey string `json:"authentication_key"`
	Comment           string `json:"comment"`
	// HelloInterval description: Constraints: 0, 1-65535
	HelloInterval int `json:"hello_interval"`
	// Interface description: REF(interface/*)
	Interface         string        `json:"interface"`
	MessageDigestKeys []interface{} `json:"message_digest_keys"`
	Priority          int           `json:"priority"`
	// TransmitDelay description: Constraints: 0, 1-65535
	TransmitDelay int `json:"transmit_delay"`
	Cost          int `json:"cost"`
	// DeadInterval description: Constraints: 0, 1-65535
	DeadInterval int    `json:"dead_interval"`
	Name         string `json:"name"`
}

var _ sophos.RestGetter = &OspfInterface{}

// GetPath implements sophos.RestObject and returns the OspfInterfaces GET path
// Returns all available ospf/interface objects
func (*OspfInterfaces) GetPath() string { return "/api/objects/ospf/interface/" }

// RefRequired implements sophos.RestObject
func (*OspfInterfaces) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OspfInterfaces GET path
// Returns all available interface types
func (o *OspfInterface) GetPath() string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", o.Reference)
}

// RefRequired implements sophos.RestObject
func (o *OspfInterface) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OspfInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s/usedby", ref)
}

// OspfMessageDigestKeys is an Sophos Endpoint subType and implements sophos.RestObject
type OspfMessageDigestKeys []OspfMessageDigestKey

// OspfMessageDigestKey represents a UTM OSPF message digest key
type OspfMessageDigestKey struct {
	Locked             string `json:"_locked"`
	ObjectType         string `json:"_type"`
	Reference          string `json:"_ref"`
	MessageDigestKeyId int    `json:"message_digest_key_id"`
	Name               string `json:"name"`
	Comment            string `json:"comment"`
	// MessageDigestKey description: (REGEX)
	MessageDigestKey string `json:"message_digest_key"`
}

var _ sophos.RestGetter = &OspfMessageDigestKey{}

// GetPath implements sophos.RestObject and returns the OspfMessageDigestKeys GET path
// Returns all available ospf/message_digest_key objects
func (*OspfMessageDigestKeys) GetPath() string { return "/api/objects/ospf/message_digest_key/" }

// RefRequired implements sophos.RestObject
func (*OspfMessageDigestKeys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the OspfMessageDigestKeys GET path
// Returns all available message_digest_key types
func (o *OspfMessageDigestKey) GetPath() string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", o.Reference)
}

// RefRequired implements sophos.RestObject
func (o *OspfMessageDigestKey) RefRequired() (string, bool) { return o.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*OspfMessageDigestKey) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s/usedby", ref)
}
