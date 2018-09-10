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
	ObjectType        string        `json:"_type"`
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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwAweNetwork) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwAweNetwork) GetType() string { return i.ObjectType }

// ItfhwAweNetworkGroups is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwAweNetworkGroups []ItfhwAweNetworkGroup

// ItfhwAweNetworkGroup represents a UTM wireless access point group
type ItfhwAweNetworkGroup struct {
	Locked     string        `json:"_locked"`
	ObjectType string        `json:"_type"`
	Reference  string        `json:"_ref"`
	ApVlantag  int           `json:"ap_vlantag"`
	Comment    string        `json:"comment"`
	Members    []interface{} `json:"members"`
	Name       string        `json:"name"`
	// Status default value is false
	Status bool `json:"status"`
	// Vlantagging default value is false
	Vlantagging bool `json:"vlantagging"`
}

var _ sophos.RestGetter = &ItfhwAweNetworkGroup{}

// GetPath implements sophos.RestObject and returns the ItfhwAweNetworkGroups GET path
// Returns all available itfhw/awe_network_group objects
func (*ItfhwAweNetworkGroups) GetPath() string { return "/api/objects/itfhw/awe_network_group/" }

// RefRequired implements sophos.RestObject
func (*ItfhwAweNetworkGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwAweNetworkGroups GET path
// Returns all available awe_network_group types
func (i *ItfhwAweNetworkGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwAweNetworkGroup) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwAweNetworkGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/awe_network_group/%s/usedby", ref)
}

// ItfhwBridges is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwBridges []ItfhwBridge

// ItfhwBridge represents a UTM bridge interface
type ItfhwBridge struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Mac description: (MACADDR)
	// Mac default value is "00:00:00:00:00:00"
	Mac     string `json:"mac"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	// Description default value is "Bridge"
	Description string `json:"description"`
	// Hardware description: (REGEX)
	Hardware string `json:"hardware"`
}

var _ sophos.RestGetter = &ItfhwBridge{}

// GetPath implements sophos.RestObject and returns the ItfhwBridges GET path
// Returns all available itfhw/bridge objects
func (*ItfhwBridges) GetPath() string { return "/api/objects/itfhw/bridge/" }

// RefRequired implements sophos.RestObject
func (*ItfhwBridges) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwBridges GET path
// Returns all available bridge types
func (i *ItfhwBridge) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/bridge/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwBridge) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
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
	ObjectType            string `json:"_type"`
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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwEthernet) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/ethernet/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwEthernet) GetType() string { return i.ObjectType }

// ItfhwGroups is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwGroups []ItfhwGroup

// ItfhwGroup represents a UTM group
type ItfhwGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &ItfhwGroup{}

// GetPath implements sophos.RestObject and returns the ItfhwGroups GET path
// Returns all available itfhw/group objects
func (*ItfhwGroups) GetPath() string { return "/api/objects/itfhw/group/" }

// RefRequired implements sophos.RestObject
func (*ItfhwGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwGroups GET path
// Returns all available group types
func (i *ItfhwGroup) GetPath() string { return fmt.Sprintf("/api/objects/itfhw/group/%s", i.Reference) }

// RefRequired implements sophos.RestObject
func (i *ItfhwGroup) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
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
	ObjectType     string `json:"_type"`
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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwLag) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/lag/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfhwLag) GetType() string { return i.ObjectType }

// ItfhwRedClients is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwRedClients []ItfhwRedClient

// ItfhwRedClient represents a UTM RED client interface
type ItfhwRedClient struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	LocalKey   string `json:"local_key"`
	// Mac description: (MACADDR)
	// Mac default value is "00:00:00:00:00:00"
	Mac string `json:"mac"`
	// Status default value is false
	Status   bool   `json:"status"`
	TunnelId int    `json:"tunnel_id"`
	Comment  string `json:"comment"`
	// HubHost description: REF(network/host), REF(network/dns_host)
	HubHost   string `json:"hub_host"`
	LocalCert string `json:"local_cert"`
	Name      string `json:"name"`
	// TunnelCompression default value is false
	TunnelCompression bool `json:"tunnel_compression"`
	// TunnelCompressionAlgorithm can be one of: []string{"deflate", "lzo", "gzip"}
	// TunnelCompressionAlgorithm default value is "lzo"
	TunnelCompressionAlgorithm string `json:"tunnel_compression_algorithm"`
	// Description default value is "Remote Ethernet Client Device"
	Description string `json:"description"`
	// Hardware description: (REGEX)
	Hardware string `json:"hardware"`
	HubCa    string `json:"hub_ca"`
}

var _ sophos.RestGetter = &ItfhwRedClient{}

// GetPath implements sophos.RestObject and returns the ItfhwRedClients GET path
// Returns all available itfhw/red_client objects
func (*ItfhwRedClients) GetPath() string { return "/api/objects/itfhw/red_client/" }

// RefRequired implements sophos.RestObject
func (*ItfhwRedClients) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwRedClients GET path
// Returns all available red_client types
func (i *ItfhwRedClient) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwRedClient) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwRedClient) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_client/%s/usedby", ref)
}

// ItfhwRedServers is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwRedServers []ItfhwRedServer

// ItfhwRedServer represents a UTM RED server interface
type ItfhwRedServer struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// PrevUnlockCode default value is ""
	PrevUnlockCode string `json:"prev_unlock_code"`
	// RouteMode can be one of: []string{"default", "split", "fullbr"}
	// RouteMode default value is "default"
	RouteMode string `json:"route_mode"`
	// Password default value is ""
	Password string `json:"password"`
	// ActivateModem default value is false
	ActivateModem bool `json:"activate_modem"`
	// BridgeAddress description: (IPADDR)
	// BridgeAddress default value is "0.0.0.0"
	BridgeAddress string `json:"bridge_address"`
	// FailoverDirect default value is true
	FailoverDirect bool `json:"failover_direct"`
	// Lan1Vids default value is ""
	Lan1Vids string `json:"lan1_vids"`
	// Lan3Vids default value is ""
	Lan3Vids       string `json:"lan3_vids"`
	Manual2Netmask int    `json:"manual2_netmask"`
	// ManualDns description: (IPADDR)
	// ManualDns default value is "0.0.0.0"
	ManualDns string `json:"manual_dns"`
	// RemoteCert description: REF(ca/host_key_cert)
	// RemoteCert default value is ""
	RemoteCert            string        `json:"remote_cert"`
	FullbrDomains         []interface{} `json:"fullbr_domains"`
	LocalNetworks         []interface{} `json:"local_networks"`
	MacFilterEntriesRed15 int           `json:"mac_filter_entries_red15"`
	MacFilterEntriesRed50 int           `json:"mac_filter_entries_red50"`
	ManualNetmask         int           `json:"manual_netmask"`
	// Status default value is false
	Status bool `json:"status"`
	// TunnelState default value is false
	TunnelState bool `json:"tunnel_state"`
	// BridgeProto can be one of: []string{"dhcp", "static", "none"}
	// BridgeProto default value is "none"
	BridgeProto string `json:"bridge_proto"`
	// Manual2Defgw description: (IPADDR)
	// Manual2Defgw default value is "0.0.0.0"
	Manual2Defgw string `json:"manual2_defgw"`
	// ManualDefgw description: (IPADDR)
	// ManualDefgw default value is "0.0.0.0"
	ManualDefgw string `json:"manual_defgw"`
	Pin         int    `json:"pin"`
	// TunnelCompressionAlgorithm can be one of: []string{"deflate", "lzo", "gzip"}
	// TunnelCompressionAlgorithm default value is "lzo"
	TunnelCompressionAlgorithm string `json:"tunnel_compression_algorithm"`
	// Type can be one of: []string{"red", "red15", "red15w", "red50", "asg", "software"}
	// Type default value is "asg"
	Type    string `json:"type"`
	Comment string `json:"comment"`
	// FastFailover default value is false
	FastFailover bool `json:"fast_failover"`
	// Lan1Mode can be one of: []string{"tagged", "untagged", "untagged_drop_tagged", "unused"}
	// Lan1Mode default value is "unused"
	Lan1Mode               string `json:"lan1_mode"`
	MacFilterEntriesRed15W int    `json:"mac_filter_entries_red15w"`
	// Manual2Dns description: (IPADDR)
	// Manual2Dns default value is "0.0.0.0"
	Manual2Dns string `json:"manual2_dns"`
	Name       string `json:"name"`
	// PinAsString description: (REGEX)
	// PinAsString default value is ""
	PinAsString string `json:"pin_as_string"`
	// Username default value is ""
	Username string `json:"username"`
	// Authorized default value is false
	Authorized bool `json:"authorized"`
	// Description default value is "Remote Ethernet Server Device"
	Description string `json:"description"`
	// Hub2Hostname default value is ""
	Hub2Hostname string `json:"hub2_hostname"`
	// MobileNetwork can be one of: []string{"gsm", "cdma"}
	// MobileNetwork default value is "gsm"
	MobileNetwork string `json:"mobile_network"`
	// State can be one of: []string{"initializing", "runnable", "notbound"}
	// State default value is "initializing"
	State string `json:"state"`
	// DeploymentMode can be one of: []string{"online", "offline"}
	// DeploymentMode default value is "online"
	DeploymentMode string `json:"deployment_mode"`
	// Lan3Mode can be one of: []string{"tagged", "untagged", "untagged_drop_tagged", "unused"}
	// Lan3Mode default value is "unused"
	Lan3Mode string `json:"lan3_mode"`
	// LocalNetworksTarget description: REF(network/host), REF(network/dns_host), REF(network/interface_address)
	// LocalNetworksTarget default value is ""
	LocalNetworksTarget string `json:"local_networks_target"`
	// Manual2Address description: (IPADDR)
	// Manual2Address default value is "0.0.0.0"
	Manual2Address string `json:"manual2_address"`
	RedId          string `json:"red_id"`
	// Uplink2Mode can be one of: []string{"dhcp", "manual"}
	// Uplink2Mode default value is "dhcp"
	Uplink2Mode string `json:"uplink2_mode"`
	// Hardware description: (REGEX)
	Hardware string `json:"hardware"`
	// Lan4Mode can be one of: []string{"tagged", "untagged", "untagged_drop_tagged", "unused"}
	// Lan4Mode default value is "unused"
	Lan4Mode string `json:"lan4_mode"`
	// UplinkMode can be one of: []string{"dhcp", "manual"}
	// UplinkMode default value is "dhcp"
	UplinkMode string `json:"uplink_mode"`
	// Apn default value is ""
	Apn           string `json:"apn"`
	BridgeNetmask int    `json:"bridge_netmask"`
	// DialString default value is "*99#"
	DialString string `json:"dial_string"`
	// Lan2Vids default value is ""
	Lan2Vids string `json:"lan2_vids"`
	// LanportMode can be one of: []string{"switch", "vlan"}
	// LanportMode default value is "switch"
	LanportMode string `json:"lanport_mode"`
	// MacFilterType can be one of: []string{"none", "whitelist", "blacklist"}
	// MacFilterType default value is "none"
	MacFilterType string `json:"mac_filter_type"`
	// ManualAddress description: (IPADDR)
	// ManualAddress default value is "0.0.0.0"
	ManualAddress string `json:"manual_address"`
	DebugLevel    int    `json:"debug_level"`
	// HubHostname default value is ""
	HubHostname string `json:"hub_hostname"`
	// Lan4Vids default value is ""
	Lan4Vids string `json:"lan4_vids"`
	// Mac description: (MACADDR)
	// Mac default value is "00:00:00:00:00:00"
	Mac           string        `json:"mac"`
	SplitNetworks []interface{} `json:"split_networks"`
	TunnelId      int           `json:"tunnel_id"`
	// FullbrDns description: REF(network/host), REF(network/dns_host), REF(network/interface_address)
	// FullbrDns default value is ""
	FullbrDns string `json:"fullbr_dns"`
	// HostnameBalancing can be one of: []string{"balance", "failover"}
	// HostnameBalancing default value is "failover"
	HostnameBalancing string `json:"hostname_balancing"`
	// Lan2Mode can be one of: []string{"tagged", "untagged", "untagged_drop_tagged", "unused"}
	// Lan2Mode default value is "unused"
	Lan2Mode string `json:"lan2_mode"`
	// MacFilterList description: REF(mac_list/*)
	// MacFilterList default value is ""
	MacFilterList string `json:"mac_filter_list"`
	// UnlockCode default value is ""
	UnlockCode string `json:"unlock_code"`
	// UplinkBalancing can be one of: []string{"balance", "failover"}
	// UplinkBalancing default value is "failover"
	UplinkBalancing       string `json:"uplink_balancing"`
	MacFilterEntriesRed10 int    `json:"mac_filter_entries_red10"`
	// TunnelCompression default value is false
	TunnelCompression bool `json:"tunnel_compression"`
	// UmtsState can be one of: []string{"READY", "PIN", "PUK"}
	// UmtsState default value is "READY"
	UmtsState string `json:"umts_state"`
}

var _ sophos.RestGetter = &ItfhwRedServer{}

// GetPath implements sophos.RestObject and returns the ItfhwRedServers GET path
// Returns all available itfhw/red_server objects
func (*ItfhwRedServers) GetPath() string { return "/api/objects/itfhw/red_server/" }

// RefRequired implements sophos.RestObject
func (*ItfhwRedServers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwRedServers GET path
// Returns all available red_server types
func (i *ItfhwRedServer) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwRedServer) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwRedServer) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/red_server/%s/usedby", ref)
}

// ItfhwSerials is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwSerials []ItfhwSerial

// ItfhwSerial represents a UTM serial interface
type ItfhwSerial struct {
	Locked      string `json:"_locked"`
	ObjectType  string `json:"_type"`
	Reference   string `json:"_ref"`
	Description string `json:"description"`
	// Hardware description: (REGEX)
	Hardware string `json:"hardware"`
	Irq      int    `json:"irq"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	// Baud default value is ""
	Baud    string `json:"baud"`
	Comment string `json:"comment"`
}

var _ sophos.RestGetter = &ItfhwSerial{}

// GetPath implements sophos.RestObject and returns the ItfhwSerials GET path
// Returns all available itfhw/serial objects
func (*ItfhwSerials) GetPath() string { return "/api/objects/itfhw/serial/" }

// RefRequired implements sophos.RestObject
func (*ItfhwSerials) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwSerials GET path
// Returns all available serial types
func (i *ItfhwSerial) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwSerial) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwSerial) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/serial/%s/usedby", ref)
}

// ItfhwUsbserials is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwUsbserials []ItfhwUsbserial

// ItfhwUsbserial represents a UTM USB serial interface
type ItfhwUsbserial struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Name       string `json:"name"`
	// Product default value is ""
	Product string `json:"product"`
	// Vendor default value is ""
	Vendor  string `json:"vendor"`
	Comment string `json:"comment"`
	// Control description: (REGEX)
	// Control default value is ""
	Control     string `json:"control"`
	Description string `json:"description"`
	// Hardware description: (REGEX)
	Hardware string `json:"hardware"`
}

var _ sophos.RestGetter = &ItfhwUsbserial{}

// GetPath implements sophos.RestObject and returns the ItfhwUsbserials GET path
// Returns all available itfhw/usbserial objects
func (*ItfhwUsbserials) GetPath() string { return "/api/objects/itfhw/usbserial/" }

// RefRequired implements sophos.RestObject
func (*ItfhwUsbserials) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwUsbserials GET path
// Returns all available usbserial types
func (i *ItfhwUsbserial) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwUsbserial) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwUsbserial) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/usbserial/%s/usedby", ref)
}

// ItfhwVirtuals is an Sophos Endpoint subType and implements sophos.RestObject
type ItfhwVirtuals []ItfhwVirtual

// ItfhwVirtual represents a UTM virtual interface
type ItfhwVirtual struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	// Description default value is "IPv6 Tunnel"
	Description string `json:"description"`
	// Hardware can be one of: []string{"6to4", "aiccu", "tspc", "teredo", "he.net"}
	// Hardware default value is "teredo"
	Hardware string `json:"hardware"`
}

var _ sophos.RestGetter = &ItfhwVirtual{}

// GetPath implements sophos.RestObject and returns the ItfhwVirtuals GET path
// Returns all available itfhw/virtual objects
func (*ItfhwVirtuals) GetPath() string { return "/api/objects/itfhw/virtual/" }

// RefRequired implements sophos.RestObject
func (*ItfhwVirtuals) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfhwVirtuals GET path
// Returns all available virtual types
func (i *ItfhwVirtual) GetPath() string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfhwVirtual) RefRequired() (string, bool) { return i.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfhwVirtual) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfhw/virtual/%s/usedby", ref)
}
