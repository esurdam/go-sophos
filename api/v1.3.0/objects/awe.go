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

var _ sophos.Endpoint = &Awe{}

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

// AweClients is an Sophos Endpoint subType and implements sophos.RestObject
type AweClients []AweClient

// AweClient represents a UTM wireless client
type AweClient struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Lastseen   int    `json:"lastseen"`
	// Mac description: (MACADDR)
	// Mac default value is "00:00:00:00:00:00"
	Mac  string `json:"mac"`
	Name string `json:"name"`
	// Vendor default value is "unknown"
	Vendor string `json:"vendor"`
}

var _ sophos.RestGetter = &AweClient{}

// GetPath implements sophos.RestObject and returns the AweClients GET path
// Returns all available awe/client objects
func (*AweClients) GetPath() string { return "/api/objects/awe/client/" }

// RefRequired implements sophos.RestObject
func (*AweClients) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AweClients GET path
// Returns all available client types
func (a *AweClient) GetPath() string { return fmt.Sprintf("/api/objects/awe/client/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AweClient) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AweClient) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/client/%s/usedby", ref)
}

// AweDevices is an Sophos Endpoint subType and implements sophos.RestObject
type AweDevices []AweDevice

// AweDevice represents a UTM wireless access point
type AweDevice struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	ObjectType  string `json:"_type"`
	AutoChannel int    `json:"auto_channel"`
	// Band can be one of: []string{"g", "a"}
	// Band default value is ""
	Band string `json:"band"`
	// Id default value is "Remote Wifi Device"
	Id string `json:"id"`
	// LanMac description: (MACADDR)
	// LanMac default value is "00:00:00:00:00:00"
	LanMac            string        `json:"lan_mac"`
	AllowedChannels   []interface{} `json:"allowed_channels"`
	AllowedCountries  []interface{} `json:"allowed_countries"`
	ApLocaldebuglevel int           `json:"ap_localdebuglevel"`
	ApVlantag         int           `json:"ap_vlantag"`
	// ChannelWidth can be one of: []string{"HT20", "HT40"}
	// ChannelWidth default value is "HT20"
	ChannelWidth string `json:"channel_width"`
	// LastIp description: (IPADDR)
	// LastIp default value is ""
	LastIp string `json:"last_ip"`
	// MeshAbility11G default value is false
	MeshAbility11G bool `json:"mesh_ability11g"`
	// R0KhSecret default value is ""
	R0KhSecret string `json:"r0kh_secret"`
	// TimeScheduling11A default value is false
	TimeScheduling11A bool          `json:"time_scheduling11a"`
	TimeSelect11A     []interface{} `json:"time_select11a"`
	// AcAbility default value is false
	AcAbility      bool          `json:"ac_ability"`
	ActiveChannels []interface{} `json:"active_channels"`
	AutoChannel11A int           `json:"auto_channel11a"`
	Channel        int           `json:"channel"`
	// DfsAbility default value is false
	DfsAbility           bool `json:"dfs_ability"`
	ScanInterval11A      int  `json:"scan_interval11a"`
	SchedScanInterval11A int  `json:"sched_scan_interval11a"`
	// TimeScheduling default value is false
	TimeScheduling bool `json:"time_scheduling"`
	// Vlantagging default value is false
	Vlantagging bool          `json:"vlantagging"`
	BridgeModes []interface{} `json:"bridge_modes"`
	Channel11A  int           `json:"channel11a"`
	// MeshAbility default value is false
	MeshAbility  bool `json:"mesh_ability"`
	ScanInterval int  `json:"scan_interval"`
	// ChannelWidth11A can be one of: []string{"HT20", "HT40", "VHT20", "VHT40", "VHT80"}
	// ChannelWidth11A default value is "HT20"
	ChannelWidth11A string `json:"channel_width11a"`
	// Country description: (REGEX)
	// Country default value is ""
	Country string `json:"country"`
	// Key default value is ""
	Key  string `json:"key"`
	Name string `json:"name"`
	// Stp default value is false
	Stp bool `json:"stp"`
	// TxPowerControl default value is false
	TxPowerControl bool          `json:"tx_power_control"`
	MaxSsids       int           `json:"max_ssids"`
	TimeSelect     []interface{} `json:"time_select"`
	// TunnelId default value is ""
	TunnelId   string `json:"tunnel_id"`
	Txpower    int    `json:"txpower"`
	Txpower11A int    `json:"txpower11a"`
	Comment    string `json:"comment"`
	// Enabled default value is false
	Enabled  bool          `json:"enabled"`
	Networks []interface{} `json:"networks"`
	// Status default value is false
	Status bool `json:"status"`
	// Interface description: REF(interface/*)
	// Interface default value is ""
	Interface string `json:"interface"`
	// Location default value is ""
	Location string `json:"location"`
	// MeshAbility11A default value is false
	MeshAbility11A    bool `json:"mesh_ability11a"`
	SchedScanInterval int  `json:"sched_scan_interval"`
	// Type default value is ""
	Type string `json:"type"`
	// WifiMac description: (MACADDR)
	// WifiMac default value is "00:00:00:00:00:00"
	WifiMac string `json:"wifi_mac"`
}

var _ sophos.RestGetter = &AweDevice{}

// GetPath implements sophos.RestObject and returns the AweDevices GET path
// Returns all available awe/device objects
func (*AweDevices) GetPath() string { return "/api/objects/awe/device/" }

// RefRequired implements sophos.RestObject
func (*AweDevices) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AweDevices GET path
// Returns all available device types
func (a *AweDevice) GetPath() string { return fmt.Sprintf("/api/objects/awe/device/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AweDevice) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AweDevice) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/device/%s/usedby", ref)
}

// AweGroups is an Sophos Endpoint subType and implements sophos.RestObject
type AweGroups []AweGroup

// AweGroup represents a UTM group
type AweGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &AweGroup{}

// GetPath implements sophos.RestObject and returns the AweGroups GET path
// Returns all available awe/group objects
func (*AweGroups) GetPath() string { return "/api/objects/awe/group/" }

// RefRequired implements sophos.RestObject
func (*AweGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AweGroups GET path
// Returns all available group types
func (a *AweGroup) GetPath() string { return fmt.Sprintf("/api/objects/awe/group/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AweGroup) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AweGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/group/%s/usedby", ref)
}

// AweLocals is an Sophos Endpoint subType and implements sophos.RestObject
type AweLocals []AweLocal

// AweLocal represents a UTM SG wifi
type AweLocal struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// Id default value is "Remote Wifi Device"
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Networks []interface{} `json:"networks"`
	// Status default value is false
	Status         bool          `json:"status"`
	ActiveChannels []interface{} `json:"active_channels"`
	// Band can be one of: []string{"g", "a"}
	// Band default value is "g"
	Band            string        `json:"band"`
	Channel         int           `json:"channel"`
	Comment         string        `json:"comment"`
	TimeSelect      []interface{} `json:"time_select"`
	AllowedChannels []interface{} `json:"allowed_channels"`
	ScanInterval    int           `json:"scan_interval"`
	Txpower         int           `json:"txpower"`
	// MeshAbility11G default value is false
	MeshAbility11G bool `json:"mesh_ability11g"`
	// TimeScheduling default value is false
	TimeScheduling bool `json:"time_scheduling"`
	// WifiMac description: (MACADDR)
	// WifiMac default value is "00:00:00:00:00:00"
	WifiMac           string `json:"wifi_mac"`
	ApLocaldebuglevel int    `json:"ap_localdebuglevel"`
	AutoChannel       int    `json:"auto_channel"`
	MaxSsids          int    `json:"max_ssids"`
	// MeshAbility default value is false
	MeshAbility bool `json:"mesh_ability"`
	// TxPowerControl default value is false
	TxPowerControl bool `json:"tx_power_control"`
	// Type default value is ""
	Type        string        `json:"type"`
	BridgeModes []interface{} `json:"bridge_modes"`
	// DfsAbility default value is false
	DfsAbility bool `json:"dfs_ability"`
	// MeshAbility11A default value is false
	MeshAbility11A    bool `json:"mesh_ability11a"`
	SchedScanInterval int  `json:"sched_scan_interval"`
}

var _ sophos.RestGetter = &AweLocal{}

// GetPath implements sophos.RestObject and returns the AweLocals GET path
// Returns all available awe/local objects
func (*AweLocals) GetPath() string { return "/api/objects/awe/local/" }

// RefRequired implements sophos.RestObject
func (*AweLocals) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AweLocals GET path
// Returns all available local types
func (a *AweLocal) GetPath() string { return fmt.Sprintf("/api/objects/awe/local/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AweLocal) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AweLocal) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/local/%s/usedby", ref)
}

// AweReds is an Sophos Endpoint subType and implements sophos.RestObject
type AweReds []AweRed

// AweRed represents a UTM RED wifi
type AweRed struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// ChannelWidth can be one of: []string{"HT20", "HT40"}
	// ChannelWidth default value is "HT20"
	ChannelWidth string `json:"channel_width"`
	// Country description: (REGEX)
	// Country default value is ""
	Country string `json:"country"`
	// Interface description: REF(interface/*)
	// Interface default value is ""
	Interface string `json:"interface"`
	// LanMac description: (MACADDR)
	// LanMac default value is "00:00:00:00:00:00"
	LanMac string `json:"lan_mac"`
	// LastIp description: (IPADDR)
	// LastIp default value is ""
	LastIp string `json:"last_ip"`
	// Status default value is false
	Status     bool          `json:"status"`
	TimeSelect []interface{} `json:"time_select"`
	// WifiMac description: (MACADDR)
	// WifiMac default value is "00:00:00:00:00:00"
	WifiMac          string        `json:"wifi_mac"`
	AllowedCountries []interface{} `json:"allowed_countries"`
	Channel          int           `json:"channel"`
	// ForcedCountry default value is ""
	ForcedCountry string `json:"forced_country"`
	MaxSsids      int    `json:"max_ssids"`
	// R0KhSecret default value is ""
	R0KhSecret string `json:"r0kh_secret"`
	// TunnelId default value is ""
	TunnelId string `json:"tunnel_id"`
	// Vlantagging default value is false
	Vlantagging bool `json:"vlantagging"`
	// AcAbility default value is false
	AcAbility         bool          `json:"ac_ability"`
	AllowedChannels   []interface{} `json:"allowed_channels"`
	ApLocaldebuglevel int           `json:"ap_localdebuglevel"`
	// Band can be one of: []string{"g", "a"}
	// Band default value is "g"
	Band        string        `json:"band"`
	BridgeModes []interface{} `json:"bridge_modes"`
	// Enabled default value is false
	Enabled bool `json:"enabled"`
	// TimeScheduling default value is false
	TimeScheduling bool `json:"time_scheduling"`
	// TxPowerControl default value is false
	TxPowerControl bool `json:"tx_power_control"`
	// DfsAbility default value is false
	DfsAbility  bool `json:"dfs_ability"`
	ApVlantag   int  `json:"ap_vlantag"`
	AutoChannel int  `json:"auto_channel"`
	// Location default value is ""
	Location       string        `json:"location"`
	ScanInterval   int           `json:"scan_interval"`
	Txpower        int           `json:"txpower"`
	ActiveChannels []interface{} `json:"active_channels"`
	// MeshAbility default value is false
	MeshAbility bool          `json:"mesh_ability"`
	Networks    []interface{} `json:"networks"`
	// Type default value is ""
	Type    string `json:"type"`
	Comment string `json:"comment"`
	// Id default value is "Remote Wifi Device"
	Id string `json:"id"`
	// Key default value is ""
	Key               string `json:"key"`
	Name              string `json:"name"`
	SchedScanInterval int    `json:"sched_scan_interval"`
}

var _ sophos.RestGetter = &AweRed{}

// GetPath implements sophos.RestObject and returns the AweReds GET path
// Returns all available awe/red objects
func (*AweReds) GetPath() string { return "/api/objects/awe/red/" }

// RefRequired implements sophos.RestObject
func (*AweReds) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AweReds GET path
// Returns all available red types
func (a *AweRed) GetPath() string { return fmt.Sprintf("/api/objects/awe/red/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AweRed) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AweRed) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awe/red/%s/usedby", ref)
}
