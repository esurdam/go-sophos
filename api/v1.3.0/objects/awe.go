package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Awe is a generated struct representing the Sophos Awe Endpoint
// GET /api/nodes/awe
type Awe struct {
	AllowedInterfaces []interface{} `json:"allowed_interfaces"`
	Clients           []interface{} `json:"clients"`
	Devices           []interface{} `json:"devices"`
	Global            struct {
		ApAutoaccept        int64  `json:"ap_autoaccept"`
		ApDebuglevel        int64  `json:"ap_debuglevel"`
		ApSoftlimit         int64  `json:"ap_softlimit"`
		ApVlantag           int64  `json:"ap_vlantag"`
		AweStatus           int64  `json:"awe_status"`
		BridgeUpdateKickout int64  `json:"bridge_update_kickout"`
		InitialSetup        int64  `json:"initial_setup"`
		LogLevel            int64  `json:"log_level"`
		MagicIP             string `json:"magic_ip"`
		NotificationTimeout int64  `json:"notification_timeout"`
		RadiusConf          string `json:"radius_conf"`
		Rootpw              string `json:"rootpw"`
		StayOnline          int64  `json:"stay_online"`
		StoreBssStats       int64  `json:"store_bss_stats"`
		TunnelIDOffset      int64  `json:"tunnel_id_offset"`
		Vlantagging         int64  `json:"vlantagging"`
	} `json:"global"`
	Networks []string `json:"networks"`
}

var defsAwe = map[string]sophos.RestObject{
	"AweClient": &AweClient{},
	"AweDevice": &AweDevice{},
	"AweGroup":  &AweGroup{},
	"AweLocal":  &AweLocal{},
	"AweRed":    &AweRed{},
}

// RestObjects implements the sophos.Node interface and returns a map of Awe's Objects
func (Awe) RestObjects() map[string]sophos.RestObject { return defsAwe }

// GetPath implements sophos.RestGetter
func (*Awe) GetPath() string { return "/api/nodes/awe" }

// RefRequired implements sophos.RestGetter
func (*Awe) RefRequired() (string, bool) { return "", false }

var defAwe = &sophos.Definition{Description: "awe", Name: "awe", Link: "/api/definitions/awe"}

// Definition returns the /api/definitions struct of Awe
func (Awe) Definition() sophos.Definition { return *defAwe }

// ApiRoutes returns all known Awe Paths
func (Awe) ApiRoutes() []string {
	return []string{
		"/api/objects/awe/client/",
		"/api/objects/awe/client/{ref}",
		"/api/objects/awe/client/{ref}/usedby",
		"/api/objects/awe/device/",
		"/api/objects/awe/device/{ref}",
		"/api/objects/awe/device/{ref}/usedby",
		"/api/objects/awe/group/",
		"/api/objects/awe/group/{ref}",
		"/api/objects/awe/group/{ref}/usedby",
		"/api/objects/awe/local/",
		"/api/objects/awe/local/{ref}",
		"/api/objects/awe/local/{ref}/usedby",
		"/api/objects/awe/red/",
		"/api/objects/awe/red/{ref}",
		"/api/objects/awe/red/{ref}/usedby",
	}
}

// References returns the Awe's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Awe) References() []string {
	return []string{
		"REF_AweClient",
		"REF_AweDevice",
		"REF_AweGroup",
		"REF_AweLocal",
		"REF_AweRed",
	}
}

// AweClient is an Sophos Endpoint subType and implements sophos.RestObject
type AweClient []interface{}

// GetPath implements sophos.RestObject and returns the AweClient GET path
// Returns all available awe/client objects
func (*AweClient) GetPath() string { return "/api/objects/awe/client/" }

// RefRequired implements sophos.RestObject
func (*AweClient) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweClient DELETE path
// Creates or updates the complete object client
func (*AweClient) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/client/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweClient PATCH path
// Changes to parts of the object client types
func (*AweClient) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/client/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweClient POST path
// Create a new awe/client object
func (*AweClient) PostPath() string {
	return "/api/objects/awe/client/"
}

// PutPath implements sophos.RestObject and returns the AweClient PUT path
// Creates or updates the complete object client
func (*AweClient) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/client/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweClient) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/client/%s/usedby", ref)
}

// AweDevice is an Sophos Endpoint subType and implements sophos.RestObject
type AweDevice []interface{}

// GetPath implements sophos.RestObject and returns the AweDevice GET path
// Returns all available awe/device objects
func (*AweDevice) GetPath() string { return "/api/objects/awe/device/" }

// RefRequired implements sophos.RestObject
func (*AweDevice) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweDevice DELETE path
// Creates or updates the complete object device
func (*AweDevice) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/device/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweDevice PATCH path
// Changes to parts of the object device types
func (*AweDevice) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/device/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweDevice POST path
// Create a new awe/device object
func (*AweDevice) PostPath() string {
	return "/api/objects/awe/device/"
}

// PutPath implements sophos.RestObject and returns the AweDevice PUT path
// Creates or updates the complete object device
func (*AweDevice) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/device/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweDevice) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/device/%s/usedby", ref)
}

// AweGroup is an Sophos Endpoint subType and implements sophos.RestObject
type AweGroup []interface{}

// GetPath implements sophos.RestObject and returns the AweGroup GET path
// Returns all available awe/group objects
func (*AweGroup) GetPath() string { return "/api/objects/awe/group/" }

// RefRequired implements sophos.RestObject
func (*AweGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweGroup DELETE path
// Creates or updates the complete object group
func (*AweGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweGroup PATCH path
// Changes to parts of the object group types
func (*AweGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweGroup POST path
// Create a new awe/group object
func (*AweGroup) PostPath() string {
	return "/api/objects/awe/group/"
}

// PutPath implements sophos.RestObject and returns the AweGroup PUT path
// Creates or updates the complete object group
func (*AweGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/group/%s/usedby", ref)
}

// AweLocal is an Sophos Endpoint subType and implements sophos.RestObject
type AweLocal []interface{}

// GetPath implements sophos.RestObject and returns the AweLocal GET path
// Returns all available awe/local objects
func (*AweLocal) GetPath() string { return "/api/objects/awe/local/" }

// RefRequired implements sophos.RestObject
func (*AweLocal) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweLocal DELETE path
// Creates or updates the complete object local
func (*AweLocal) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/local/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweLocal PATCH path
// Changes to parts of the object local types
func (*AweLocal) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/local/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweLocal POST path
// Create a new awe/local object
func (*AweLocal) PostPath() string {
	return "/api/objects/awe/local/"
}

// PutPath implements sophos.RestObject and returns the AweLocal PUT path
// Creates or updates the complete object local
func (*AweLocal) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/local/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweLocal) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/local/%s/usedby", ref)
}

// AweRed is an Sophos Endpoint subType and implements sophos.RestObject
type AweRed []interface{}

// GetPath implements sophos.RestObject and returns the AweRed GET path
// Returns all available awe/red objects
func (*AweRed) GetPath() string { return "/api/objects/awe/red/" }

// RefRequired implements sophos.RestObject
func (*AweRed) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AweRed DELETE path
// Creates or updates the complete object red
func (*AweRed) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/red/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AweRed PATCH path
// Changes to parts of the object red types
func (*AweRed) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/red/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AweRed POST path
// Create a new awe/red object
func (*AweRed) PostPath() string {
	return "/api/objects/awe/red/"
}

// PutPath implements sophos.RestObject and returns the AweRed PUT path
// Creates or updates the complete object red
func (*AweRed) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/red/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*AweRed) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/red/%s/usedby", ref)
}
