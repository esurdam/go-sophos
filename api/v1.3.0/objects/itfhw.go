package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Itfhw is a generated struct representing the Sophos Itfhw Endpoint
// GET /api/nodes/itfhw
type Itfhw struct {
	ItfhwAweNetwork      ItfhwAweNetwork      `json:"itfhw_awe_network"`
	ItfhwAweNetworkGroup ItfhwAweNetworkGroup `json:"itfhw_awe_network_group"`
	ItfhwBridge          ItfhwBridge          `json:"itfhw_bridge"`
	ItfhwEthernet        ItfhwEthernet        `json:"itfhw_ethernet"`
	ItfhwGroup           ItfhwGroup           `json:"itfhw_group"`
	ItfhwLag             ItfhwLag             `json:"itfhw_lag"`
	ItfhwRedClient       ItfhwRedClient       `json:"itfhw_red_client"`
	ItfhwRedServer       ItfhwRedServer       `json:"itfhw_red_server"`
	ItfhwSerial          ItfhwSerial          `json:"itfhw_serial"`
	ItfhwUsbserial       ItfhwUsbserial       `json:"itfhw_usbserial"`
	ItfhwVirtual         ItfhwVirtual         `json:"itfhw_virtual"`
}

var _ sophos.Endpoint = &Itfhw{}

var defsItfhw = map[string]sophos.RestObject{
	"ItfhwAweNetwork":      &ItfhwAweNetwork{},
	"ItfhwAweNetworkGroup": &ItfhwAweNetworkGroup{},
	"ItfhwBridge":          &ItfhwBridge{},
	"ItfhwEthernet":        &ItfhwEthernet{},
	"ItfhwGroup":           &ItfhwGroup{},
	"ItfhwLag":             &ItfhwLag{},
	"ItfhwRedClient":       &ItfhwRedClient{},
	"ItfhwRedServer":       &ItfhwRedServer{},
	"ItfhwSerial":          &ItfhwSerial{},
	"ItfhwUsbserial":       &ItfhwUsbserial{},
	"ItfhwVirtual":         &ItfhwVirtual{},
}

// RestObjects implements the sophos.Node interface and returns a map of Itfhw's Objects
func (Itfhw) RestObjects() map[string]sophos.RestObject { return defsItfhw }

// GetPath implements sophos.RestGetter
func (*Itfhw) GetPath() string { return "/api/nodes/itfhw" }

// RefRequired implements sophos.RestGetter
func (*Itfhw) RefRequired() (string, bool) { return "", false }

var defItfhw = &sophos.Definition{Description: "itfhw", Name: "itfhw", Link: "/api/definitions/itfhw"}

// Definition returns the /api/definitions struct of Itfhw
func (Itfhw) Definition() sophos.Definition { return *defItfhw }

// ApiRoutes returns all known Itfhw Paths
func (Itfhw) ApiRoutes() []string {
	return []string{
		"/api/objects/itfhw/awe_network/",
		"/api/objects/itfhw/awe_network/{ref}",
		"/api/objects/itfhw/awe_network/{ref}/usedby",
		"/api/objects/itfhw/awe_network_group/",
		"/api/objects/itfhw/awe_network_group/{ref}",
		"/api/objects/itfhw/awe_network_group/{ref}/usedby",
		"/api/objects/itfhw/bridge/",
		"/api/objects/itfhw/bridge/{ref}",
		"/api/objects/itfhw/bridge/{ref}/usedby",
		"/api/objects/itfhw/ethernet/",
		"/api/objects/itfhw/ethernet/{ref}",
		"/api/objects/itfhw/ethernet/{ref}/usedby",
		"/api/objects/itfhw/group/",
		"/api/objects/itfhw/group/{ref}",
		"/api/objects/itfhw/group/{ref}/usedby",
		"/api/objects/itfhw/lag/",
		"/api/objects/itfhw/lag/{ref}",
		"/api/objects/itfhw/lag/{ref}/usedby",
		"/api/objects/itfhw/red_client/",
		"/api/objects/itfhw/red_client/{ref}",
		"/api/objects/itfhw/red_client/{ref}/usedby",
		"/api/objects/itfhw/red_server/",
		"/api/objects/itfhw/red_server/{ref}",
		"/api/objects/itfhw/red_server/{ref}/usedby",
		"/api/objects/itfhw/serial/",
		"/api/objects/itfhw/serial/{ref}",
		"/api/objects/itfhw/serial/{ref}/usedby",
		"/api/objects/itfhw/usbserial/",
		"/api/objects/itfhw/usbserial/{ref}",
		"/api/objects/itfhw/usbserial/{ref}/usedby",
		"/api/objects/itfhw/virtual/",
		"/api/objects/itfhw/virtual/{ref}",
		"/api/objects/itfhw/virtual/{ref}/usedby",
	}
}

// References returns the Itfhw's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Itfhw) References() []string {
	return []string{
		"REF_ItfhwAweNetwork",
		"REF_ItfhwAweNetworkGroup",
		"REF_ItfhwBridge",
		"REF_ItfhwEthernet",
		"REF_ItfhwGroup",
		"REF_ItfhwLag",
		"REF_ItfhwRedClient",
		"REF_ItfhwRedServer",
		"REF_ItfhwSerial",
		"REF_ItfhwUsbserial",
		"REF_ItfhwVirtual",
	}
}

// ItfhwAweNetworks is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwAweNetworks []ItfhwAweNetwork

// ItfhwAweNetwork is a generated Sophos object
type ItfhwAweNetwork struct {
	Locked            string        `json:"_locked"`
	Reference         string        `json:"_ref"`
	_type             string        `json:"_type"`
	ApBridgemode      string        `json:"ap_bridgemode"`
	Bridge            string        `json:"bridge"`
	ClientIsolation   bool          `json:"client_isolation"`
	Comment           string        `json:"comment"`
	CryptoAlg         string        `json:"crypto_alg"`
	Description       string        `json:"description"`
	Dot11r            bool          `json:"dot11r"`
	DynamicVlan       int64         `json:"dynamic_vlan"`
	EncryptionMode    string        `json:"encryption_mode"`
	FreqBands         string        `json:"freq_bands"`
	Hardware          string        `json:"hardware"`
	HideSsid          bool          `json:"hide_ssid"`
	InterfaceName     string        `json:"interface_name"`
	Mac               string        `json:"mac"`
	MacFilter         string        `json:"mac_filter"`
	MacList           string        `json:"mac_list"`
	MeshID            string        `json:"mesh_id"`
	MeshMode          string        `json:"mesh_mode"`
	MeshSubtag        string        `json:"mesh_subtag"`
	Name              string        `json:"name"`
	NetworkMode       string        `json:"network_mode"`
	NetworkName       string        `json:"network_name"`
	Psk               string        `json:"psk"`
	R0khSecret        string        `json:"r0kh_secret"`
	Ssid              string        `json:"ssid"`
	SsidVlantag       interface{}   `json:"ssid_vlantag"`
	Status            bool          `json:"status"`
	TimeScheduling    bool          `json:"time_scheduling"`
	TimeSelect        []interface{} `json:"time_select"`
	Uapsd             bool          `json:"uapsd"`
	Utf8Ssid          bool          `json:"utf8_ssid"`
	Vlantag           string        `json:"vlantag"`
	Wep128            string        `json:"wep128"`
	WepAuthentication string        `json:"wep_authentication"`
}

var _ sophos.RestGetter = &ItfhwAweNetwork{}

// GetPath implements sophos.RestObject and returns the ItfhwAweNetworks GET path
// Returns all available itfhw/awe_network objects
func (*ItfhwAweNetworks) GetPath() string { return "/api/objects/itfhw/awe_network/" }

// RefRequired implements sophos.RestObject
func (*ItfhwAweNetworks) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwAweNetworks GET path
// Returns all available awe_network types
func (i *ItfhwAweNetwork) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwAweNetwork) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the ItfhwAweNetwork DELETE path
// Creates or updates the complete object awe_network
func (*ItfhwAweNetwork) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwAweNetwork PATCH path
// Changes to parts of the object awe_network types
func (*ItfhwAweNetwork) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwAweNetwork POST path
// Create a new itfhw/awe_network object
func (*ItfhwAweNetwork) PostPath() string {
	return "/api/objects/itfhw/awe_network/"
}

// PutPath implements sophos.RestObject and returns the ItfhwAweNetwork PUT path
// Creates or updates the complete object awe_network
func (*ItfhwAweNetwork) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwAweNetwork) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwAweNetwork) GetType() string { return i._type }

// ItfhwAweNetworkGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwAweNetworkGroup []interface{}

var _ sophos.RestObject = &ItfhwAweNetworkGroup{}

// GetPath implements sophos.RestObject and returns the ItfhwAweNetworkGroup GET path
// Returns all available itfhw/awe_network_group objects
func (*ItfhwAweNetworkGroup) GetPath() string { return "/api/objects/itfhw/awe_network_group/" }

// RefRequired implements sophos.RestObject
func (*ItfhwAweNetworkGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwAweNetworkGroup DELETE path
// Creates or updates the complete object awe_network_group
func (*ItfhwAweNetworkGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwAweNetworkGroup PATCH path
// Changes to parts of the object awe_network_group types
func (*ItfhwAweNetworkGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwAweNetworkGroup POST path
// Create a new itfhw/awe_network_group object
func (*ItfhwAweNetworkGroup) PostPath() string {
	return "/api/objects/itfhw/awe_network_group/"
}

// PutPath implements sophos.RestObject and returns the ItfhwAweNetworkGroup PUT path
// Creates or updates the complete object awe_network_group
func (*ItfhwAweNetworkGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwAweNetworkGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s/usedby", ref)
}

// ItfhwBridge is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwBridge []interface{}

var _ sophos.RestObject = &ItfhwBridge{}

// GetPath implements sophos.RestObject and returns the ItfhwBridge GET path
// Returns all available itfhw/bridge objects
func (*ItfhwBridge) GetPath() string { return "/api/objects/itfhw/bridge/" }

// RefRequired implements sophos.RestObject
func (*ItfhwBridge) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwBridge DELETE path
// Creates or updates the complete object bridge
func (*ItfhwBridge) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/bridge/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwBridge PATCH path
// Changes to parts of the object bridge types
func (*ItfhwBridge) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/bridge/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwBridge POST path
// Create a new itfhw/bridge object
func (*ItfhwBridge) PostPath() string {
	return "/api/objects/itfhw/bridge/"
}

// PutPath implements sophos.RestObject and returns the ItfhwBridge PUT path
// Creates or updates the complete object bridge
func (*ItfhwBridge) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/bridge/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwBridge) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/bridge/%s/usedby", ref)
}

// ItfhwEthernets is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwEthernets []ItfhwEthernet

// ItfhwEthernet is a generated Sophos object
type ItfhwEthernet struct {
	Locked                string `json:"_locked"`
	Reference             string `json:"_ref"`
	_type                 string `json:"_type"`
	AutoNegotiation       bool   `json:"auto_negotiation"`
	AutoNegotiationStatus bool   `json:"auto_negotiation_status"`
	Comment               string `json:"comment"`
	Description           string `json:"description"`
	Duplex                string `json:"duplex"`
	Hardware              string `json:"hardware"`
	Irq                   int64  `json:"irq"`
	LinkMonitoring        bool   `json:"link_monitoring"`
	Mac                   string `json:"mac"`
	Mii                   bool   `json:"mii"`
	Name                  string `json:"name"`
	Pcidev                string `json:"pcidev"`
	PoeEnabled            bool   `json:"poe_enabled"`
	PoeStatus             string `json:"poe_status"`
	Slot                  string `json:"slot"`
	Speed                 string `json:"speed"`
	SupportedLinkModes    string `json:"supported_link_modes"`
	VirtualMac            string `json:"virtual_mac"`
}

var _ sophos.RestGetter = &ItfhwEthernet{}

// GetPath implements sophos.RestObject and returns the ItfhwEthernets GET path
// Returns all available itfhw/ethernet objects
func (*ItfhwEthernets) GetPath() string { return "/api/objects/itfhw/ethernet/" }

// RefRequired implements sophos.RestObject
func (*ItfhwEthernets) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwEthernets GET path
// Returns all available ethernet types
func (i *ItfhwEthernet) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwEthernet) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the ItfhwEthernet DELETE path
// Creates or updates the complete object ethernet
func (*ItfhwEthernet) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwEthernet PATCH path
// Changes to parts of the object ethernet types
func (*ItfhwEthernet) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwEthernet POST path
// Create a new itfhw/ethernet object
func (*ItfhwEthernet) PostPath() string {
	return "/api/objects/itfhw/ethernet/"
}

// PutPath implements sophos.RestObject and returns the ItfhwEthernet PUT path
// Creates or updates the complete object ethernet
func (*ItfhwEthernet) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwEthernet) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwEthernet) GetType() string { return i._type }

// ItfhwGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwGroup []interface{}

var _ sophos.RestObject = &ItfhwGroup{}

// GetPath implements sophos.RestObject and returns the ItfhwGroup GET path
// Returns all available itfhw/group objects
func (*ItfhwGroup) GetPath() string { return "/api/objects/itfhw/group/" }

// RefRequired implements sophos.RestObject
func (*ItfhwGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwGroup DELETE path
// Creates or updates the complete object group
func (*ItfhwGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwGroup PATCH path
// Changes to parts of the object group types
func (*ItfhwGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwGroup POST path
// Create a new itfhw/group object
func (*ItfhwGroup) PostPath() string {
	return "/api/objects/itfhw/group/"
}

// PutPath implements sophos.RestObject and returns the ItfhwGroup PUT path
// Creates or updates the complete object group
func (*ItfhwGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/group/%s/usedby", ref)
}

// ItfhwLags is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwLags []ItfhwLag

// ItfhwLag is a generated Sophos object
type ItfhwLag struct {
	Locked         string `json:"_locked"`
	Reference      string `json:"_ref"`
	_type          string `json:"_type"`
	Comment        string `json:"comment"`
	Description    string `json:"description"`
	Hardware       string `json:"hardware"`
	LinkMonitoring bool   `json:"link_monitoring"`
	Mac            string `json:"mac"`
	Name           string `json:"name"`
}

var _ sophos.RestGetter = &ItfhwLag{}

// GetPath implements sophos.RestObject and returns the ItfhwLags GET path
// Returns all available itfhw/lag objects
func (*ItfhwLags) GetPath() string { return "/api/objects/itfhw/lag/" }

// RefRequired implements sophos.RestObject
func (*ItfhwLags) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwLags GET path
// Returns all available lag types
func (i *ItfhwLag) GetPath() string { return fmt.Sprintf("/api/objects/itfhw/lag/%s", i.Reference) }

// RefRequired implements sophos.RestObject
func (i *ItfhwLag) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the ItfhwLag DELETE path
// Creates or updates the complete object lag
func (*ItfhwLag) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/lag/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwLag PATCH path
// Changes to parts of the object lag types
func (*ItfhwLag) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/lag/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwLag POST path
// Create a new itfhw/lag object
func (*ItfhwLag) PostPath() string {
	return "/api/objects/itfhw/lag/"
}

// PutPath implements sophos.RestObject and returns the ItfhwLag PUT path
// Creates or updates the complete object lag
func (*ItfhwLag) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/lag/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwLag) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/lag/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwLag) GetType() string { return i._type }

// ItfhwRedClient is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwRedClient []interface{}

var _ sophos.RestObject = &ItfhwRedClient{}

// GetPath implements sophos.RestObject and returns the ItfhwRedClient GET path
// Returns all available itfhw/red_client objects
func (*ItfhwRedClient) GetPath() string { return "/api/objects/itfhw/red_client/" }

// RefRequired implements sophos.RestObject
func (*ItfhwRedClient) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwRedClient DELETE path
// Creates or updates the complete object red_client
func (*ItfhwRedClient) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwRedClient PATCH path
// Changes to parts of the object red_client types
func (*ItfhwRedClient) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwRedClient POST path
// Create a new itfhw/red_client object
func (*ItfhwRedClient) PostPath() string {
	return "/api/objects/itfhw/red_client/"
}

// PutPath implements sophos.RestObject and returns the ItfhwRedClient PUT path
// Creates or updates the complete object red_client
func (*ItfhwRedClient) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwRedClient) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s/usedby", ref)
}

// ItfhwRedServer is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwRedServer []interface{}

var _ sophos.RestObject = &ItfhwRedServer{}

// GetPath implements sophos.RestObject and returns the ItfhwRedServer GET path
// Returns all available itfhw/red_server objects
func (*ItfhwRedServer) GetPath() string { return "/api/objects/itfhw/red_server/" }

// RefRequired implements sophos.RestObject
func (*ItfhwRedServer) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwRedServer DELETE path
// Creates or updates the complete object red_server
func (*ItfhwRedServer) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwRedServer PATCH path
// Changes to parts of the object red_server types
func (*ItfhwRedServer) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwRedServer POST path
// Create a new itfhw/red_server object
func (*ItfhwRedServer) PostPath() string {
	return "/api/objects/itfhw/red_server/"
}

// PutPath implements sophos.RestObject and returns the ItfhwRedServer PUT path
// Creates or updates the complete object red_server
func (*ItfhwRedServer) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwRedServer) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s/usedby", ref)
}

// ItfhwSerial is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwSerial []interface{}

var _ sophos.RestObject = &ItfhwSerial{}

// GetPath implements sophos.RestObject and returns the ItfhwSerial GET path
// Returns all available itfhw/serial objects
func (*ItfhwSerial) GetPath() string { return "/api/objects/itfhw/serial/" }

// RefRequired implements sophos.RestObject
func (*ItfhwSerial) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwSerial DELETE path
// Creates or updates the complete object serial
func (*ItfhwSerial) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwSerial PATCH path
// Changes to parts of the object serial types
func (*ItfhwSerial) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwSerial POST path
// Create a new itfhw/serial object
func (*ItfhwSerial) PostPath() string {
	return "/api/objects/itfhw/serial/"
}

// PutPath implements sophos.RestObject and returns the ItfhwSerial PUT path
// Creates or updates the complete object serial
func (*ItfhwSerial) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwSerial) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s/usedby", ref)
}

// ItfhwUsbserial is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwUsbserial []interface{}

var _ sophos.RestObject = &ItfhwUsbserial{}

// GetPath implements sophos.RestObject and returns the ItfhwUsbserial GET path
// Returns all available itfhw/usbserial objects
func (*ItfhwUsbserial) GetPath() string { return "/api/objects/itfhw/usbserial/" }

// RefRequired implements sophos.RestObject
func (*ItfhwUsbserial) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwUsbserial DELETE path
// Creates or updates the complete object usbserial
func (*ItfhwUsbserial) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwUsbserial PATCH path
// Changes to parts of the object usbserial types
func (*ItfhwUsbserial) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwUsbserial POST path
// Create a new itfhw/usbserial object
func (*ItfhwUsbserial) PostPath() string {
	return "/api/objects/itfhw/usbserial/"
}

// PutPath implements sophos.RestObject and returns the ItfhwUsbserial PUT path
// Creates or updates the complete object usbserial
func (*ItfhwUsbserial) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwUsbserial) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s/usedby", ref)
}

// ItfhwVirtual is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwVirtual []interface{}

var _ sophos.RestObject = &ItfhwVirtual{}

// GetPath implements sophos.RestObject and returns the ItfhwVirtual GET path
// Returns all available itfhw/virtual objects
func (*ItfhwVirtual) GetPath() string { return "/api/objects/itfhw/virtual/" }

// RefRequired implements sophos.RestObject
func (*ItfhwVirtual) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfhwVirtual DELETE path
// Creates or updates the complete object virtual
func (*ItfhwVirtual) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfhwVirtual PATCH path
// Changes to parts of the object virtual types
func (*ItfhwVirtual) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfhwVirtual POST path
// Create a new itfhw/virtual object
func (*ItfhwVirtual) PostPath() string {
	return "/api/objects/itfhw/virtual/"
}

// PutPath implements sophos.RestObject and returns the ItfhwVirtual PUT path
// Creates or updates the complete object virtual
func (*ItfhwVirtual) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwVirtual) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s/usedby", ref)
}
