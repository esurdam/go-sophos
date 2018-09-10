package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Snmp is a generated struct representing the Sophos Snmp Endpoint
// GET /api/nodes/snmp
type Snmp struct {
	AllowedNetworks []interface{} `json:"allowed_networks"`
	AuthPassword    string        `json:"auth_password"`
	AuthType        string        `json:"auth_type"`
	Community       string        `json:"community"`
	DeviceAdmin     string        `json:"device_admin"`
	DeviceLocation  string        `json:"device_location"`
	DeviceName      string        `json:"device_name"`
	EncryptPassword string        `json:"encrypt_password"`
	EncryptType     string        `json:"encrypt_type"`
	Status          int64         `json:"status"`
	Traps           []interface{} `json:"traps"`
	TrapsUseOldOids int64         `json:"traps_use_old_oids"`
	Username        string        `json:"username"`
	Version         string        `json:"version"`
}

var _ sophos.Endpoint = &Snmp{}

var defsSnmp = map[string]sophos.RestObject{
	"SnmpGroup": &SnmpGroup{},
	"SnmpTrap":  &SnmpTrap{},
}

// RestObjects implements the sophos.Node interface and returns a map of Snmp's Objects
func (Snmp) RestObjects() map[string]sophos.RestObject { return defsSnmp }

// GetPath implements sophos.RestGetter
func (*Snmp) GetPath() string { return "/api/nodes/snmp" }

// RefRequired implements sophos.RestGetter
func (*Snmp) RefRequired() (string, bool) { return "", false }

var defSnmp = &sophos.Definition{Description: "snmp", Name: "snmp", Link: "/api/definitions/snmp"}

// Definition returns the /api/definitions struct of Snmp
func (Snmp) Definition() sophos.Definition { return *defSnmp }

// ApiRoutes returns all known Snmp Paths
func (Snmp) ApiRoutes() []string {
	return []string{
		"/api/objects/snmp/group/",
		"/api/objects/snmp/group/{ref}",
		"/api/objects/snmp/group/{ref}/usedby",
		"/api/objects/snmp/trap/",
		"/api/objects/snmp/trap/{ref}",
		"/api/objects/snmp/trap/{ref}/usedby",
	}
}

// References returns the Snmp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Snmp) References() []string {
	return []string{
		"REF_SnmpGroup",
		"REF_SnmpTrap",
	}
}

// SnmpGroups is an Sophos Endpoint subType and implements sophos.RestObject
type SnmpGroups []SnmpGroup

// SnmpGroup represents a UTM group
type SnmpGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &SnmpGroup{}

// GetPath implements sophos.RestObject and returns the SnmpGroups GET path
// Returns all available snmp/group objects
func (*SnmpGroups) GetPath() string { return "/api/objects/snmp/group/" }

// RefRequired implements sophos.RestObject
func (*SnmpGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SnmpGroups GET path
// Returns all available group types
func (s *SnmpGroup) GetPath() string { return fmt.Sprintf("/api/objects/snmp/group/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *SnmpGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SnmpGroup DELETE path
// Creates or updates the complete object group
func (*SnmpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SnmpGroup PATCH path
// Changes to parts of the object group types
func (*SnmpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SnmpGroup POST path
// Create a new snmp/group object
func (*SnmpGroup) PostPath() string {
	return "/api/objects/snmp/group/"
}

// PutPath implements sophos.RestObject and returns the SnmpGroup PUT path
// Creates or updates the complete object group
func (*SnmpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SnmpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/group/%s/usedby", ref)
}

// SnmpTraps is an Sophos Endpoint subType and implements sophos.RestObject
type SnmpTraps []SnmpTrap

// SnmpTrap represents a UTM SNMP trap
type SnmpTrap struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	// EncryptType can be one of: []string{"None", "DES", "AES"}
	// EncryptType default value is "None"
	EncryptType string `json:"encrypt_type"`
	// Host description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Host string `json:"host"`
	// Username description: (SNMPSTRING)
	// Username default value is ""
	Username string `json:"username"`
	// Version can be one of: []string{"v2c", "v3"}
	// Version default value is "v2c"
	Version string `json:"version"`
	// AuthPassword description: (SNMPSTRING)
	// AuthPassword default value is ""
	AuthPassword string `json:"auth_password"`
	// AuthType can be one of: []string{"MD5", "SHA"}
	// AuthType default value is "MD5"
	AuthType string `json:"auth_type"`
	// Community description: (SNMPSTRING)
	// Community default value is "public"
	Community string `json:"community"`
	// EncryptPassword description: (SNMPSTRING)
	// EncryptPassword default value is ""
	EncryptPassword string `json:"encrypt_password"`
	// Engineid default value is ""
	Engineid string `json:"engineid"`
	Name     string `json:"name"`
	// Status default value is true
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &SnmpTrap{}

// GetPath implements sophos.RestObject and returns the SnmpTraps GET path
// Returns all available snmp/trap objects
func (*SnmpTraps) GetPath() string { return "/api/objects/snmp/trap/" }

// RefRequired implements sophos.RestObject
func (*SnmpTraps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SnmpTraps GET path
// Returns all available trap types
func (s *SnmpTrap) GetPath() string { return fmt.Sprintf("/api/objects/snmp/trap/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *SnmpTrap) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SnmpTrap DELETE path
// Creates or updates the complete object trap
func (*SnmpTrap) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/trap/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SnmpTrap PATCH path
// Changes to parts of the object trap types
func (*SnmpTrap) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/trap/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SnmpTrap POST path
// Create a new snmp/trap object
func (*SnmpTrap) PostPath() string {
	return "/api/objects/snmp/trap/"
}

// PutPath implements sophos.RestObject and returns the SnmpTrap PUT path
// Creates or updates the complete object trap
func (*SnmpTrap) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/trap/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SnmpTrap) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/snmp/trap/%s/usedby", ref)
}
