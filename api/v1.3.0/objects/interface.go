package objects

// THIS FILE IS AUTO GENERATED BY bin/gen.go! DO NOT EDIT!

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Interface is a generated struct representing the Sophos Interface Endpoint
// GET /api/nodes/interface
type Interface struct {
	InterfaceBridge   InterfaceBridge   `json:"interface_bridge"`
	InterfaceEthernet InterfaceEthernet `json:"interface_ethernet"`
	InterfaceGroup    InterfaceGroup    `json:"interface_group"`
	InterfacePpp3G    InterfacePpp3G    `json:"interface_ppp3g"`
	InterfacePppmodem InterfacePppmodem `json:"interface_pppmodem"`
	InterfacePppoa    InterfacePppoa    `json:"interface_pppoa"`
	InterfacePppoe    InterfacePppoe    `json:"interface_pppoe"`
	InterfaceTunnel   InterfaceTunnel   `json:"interface_tunnel"`
	InterfaceVlan     InterfaceVlan     `json:"interface_vlan"`
}

var _ sophos.Endpoint = &Interface{}

var defsInterface = map[string]sophos.RestObject{
	"InterfaceBridge":   &InterfaceBridge{},
	"InterfaceEthernet": &InterfaceEthernet{},
	"InterfaceGroup":    &InterfaceGroup{},
	"InterfacePpp3G":    &InterfacePpp3G{},
	"InterfacePppmodem": &InterfacePppmodem{},
	"InterfacePppoa":    &InterfacePppoa{},
	"InterfacePppoe":    &InterfacePppoe{},
	"InterfaceTunnel":   &InterfaceTunnel{},
	"InterfaceVlan":     &InterfaceVlan{},
}

// RestObjects implements the sophos.Node interface and returns a map of Interface's Objects
func (Interface) RestObjects() map[string]sophos.RestObject { return defsInterface }

// GetPath implements sophos.RestGetter
func (*Interface) GetPath() string { return "/api/nodes/interface" }

// RefRequired implements sophos.RestGetter
func (*Interface) RefRequired() (string, bool) { return "", false }

var defInterface = &sophos.Definition{Description: "interface", Name: "interface", Link: "/api/definitions/interface"}

// Definition returns the /api/definitions struct of Interface
func (Interface) Definition() sophos.Definition { return *defInterface }

// ApiRoutes returns all known Interface Paths
func (Interface) ApiRoutes() []string {
	return []string{
		"/api/objects/interface/bridge/",
		"/api/objects/interface/bridge/{ref}",
		"/api/objects/interface/bridge/{ref}/usedby",
		"/api/objects/interface/ethernet/",
		"/api/objects/interface/ethernet/{ref}",
		"/api/objects/interface/ethernet/{ref}/usedby",
		"/api/objects/interface/group/",
		"/api/objects/interface/group/{ref}",
		"/api/objects/interface/group/{ref}/usedby",
		"/api/objects/interface/ppp3g/",
		"/api/objects/interface/ppp3g/{ref}",
		"/api/objects/interface/ppp3g/{ref}/usedby",
		"/api/objects/interface/pppmodem/",
		"/api/objects/interface/pppmodem/{ref}",
		"/api/objects/interface/pppmodem/{ref}/usedby",
		"/api/objects/interface/pppoa/",
		"/api/objects/interface/pppoa/{ref}",
		"/api/objects/interface/pppoa/{ref}/usedby",
		"/api/objects/interface/pppoe/",
		"/api/objects/interface/pppoe/{ref}",
		"/api/objects/interface/pppoe/{ref}/usedby",
		"/api/objects/interface/tunnel/",
		"/api/objects/interface/tunnel/{ref}",
		"/api/objects/interface/tunnel/{ref}/usedby",
		"/api/objects/interface/vlan/",
		"/api/objects/interface/vlan/{ref}",
		"/api/objects/interface/vlan/{ref}/usedby",
	}
}

// References returns the Interface's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Interface) References() []string {
	return []string{
		"REF_InterfaceBridge",
		"REF_InterfaceEthernet",
		"REF_InterfaceGroup",
		"REF_InterfacePpp3G",
		"REF_InterfacePppmodem",
		"REF_InterfacePppoa",
		"REF_InterfacePppoe",
		"REF_InterfaceTunnel",
		"REF_InterfaceVlan",
	}
}

// InterfaceBridges is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceBridges []InterfaceBridge

// InterfaceBridge represents a UTM bridge
type InterfaceBridge struct {
	Locked              string   `json:"_locked"`
	ObjectType          string   `json:"_type"`
	Reference           string   `json:"_ref"`
	AdditionalAddresses []string `json:"additional_addresses"`
	Ageing              int      `json:"ageing"`
	// ArpBcast default value is false
	ArpBcast bool   `json:"arp_bcast"`
	Comment  string `json:"comment"`
	// ConvertedFromHw description: REF(itfhw/*)
	// ConvertedFromHw default value is ""
	ConvertedFromHw     string   `json:"converted_from_hw"`
	ForwardedEthertypes []string `json:"forwarded_ethertypes"`
	// Itfhw description: REF(itfhw/bridge)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	Mtu  int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool     `json:"mtu_auto_discovery"`
	Name             string   `json:"name"`
	Ports            []string `json:"ports"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// Proxyarp default value is false
	Proxyarp bool `json:"proxyarp"`
	// Proxyndp default value is false
	Proxyndp bool `json:"proxyndp"`
	// Status default value is false
	Status    bool `json:"status"`
	StpFd     int  `json:"stp_fd"`
	StpHello  int  `json:"stp_hello"`
	StpMaxage int  `json:"stp_maxage"`
	StpPrio   int  `json:"stp_prio"`
	// StpStatus default value is false
	StpStatus bool `json:"stp_status"`
	// UseDhcp default value is false
	UseDhcp bool `json:"use_dhcp"`
	// UseDhcpv6 default value is false
	UseDhcpv6 bool `json:"use_dhcpv6"`
	// VirtualMac description: (MACADDR)
	// VirtualMac default value is "00:00:00:00:00:00"
	VirtualMac string `json:"virtual_mac"`
}

var _ sophos.RestGetter = &InterfaceBridge{}

// GetPath implements sophos.RestObject and returns the InterfaceBridges GET path
// Returns all available interface/bridge objects
func (*InterfaceBridges) GetPath() string { return "/api/objects/interface/bridge/" }

// RefRequired implements sophos.RestObject
func (*InterfaceBridges) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceBridges GET path
// Returns all available bridge types
func (i *InterfaceBridge) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceBridge) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceBridge DELETE path
// Creates or updates the complete object bridge
func (*InterfaceBridge) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceBridge PATCH path
// Changes to parts of the object bridge types
func (*InterfaceBridge) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceBridge POST path
// Create a new interface/bridge object
func (*InterfaceBridge) PostPath() string {
	return "/api/objects/interface/bridge/"
}

// PutPath implements sophos.RestObject and returns the InterfaceBridge PUT path
// Creates or updates the complete object bridge
func (*InterfaceBridge) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceBridge) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s/usedby", ref)
}

// InterfaceEthernets is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceEthernets []InterfaceEthernet

// InterfaceEthernet represents a UTM ethernet standard interface
type InterfaceEthernet struct {
	Locked              string   `json:"_locked"`
	ObjectType          string   `json:"_type"`
	Reference           string   `json:"_ref"`
	AdditionalAddresses []string `json:"additional_addresses"`
	Bandwidth           int      `json:"bandwidth"`
	Comment             string   `json:"comment"`
	Inbandwidth         int      `json:"inbandwidth"`
	// Itfhw description: REF(itfhw/ethernet), REF(itfhw/bridge), REF(itfhw/lag), REF(itfhw/red_server), REF(itfhw/red_client), REF(itfhw/awe_network)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	Mtu  int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool   `json:"mtu_auto_discovery"`
	Name             string `json:"name"`
	Outbandwidth     int    `json:"outbandwidth"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// Proxyarp default value is false
	Proxyarp bool `json:"proxyarp"`
	// Proxyndp default value is false
	Proxyndp bool `json:"proxyndp"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &InterfaceEthernet{}

// GetPath implements sophos.RestObject and returns the InterfaceEthernets GET path
// Returns all available interface/ethernet objects
func (*InterfaceEthernets) GetPath() string { return "/api/objects/interface/ethernet/" }

// RefRequired implements sophos.RestObject
func (*InterfaceEthernets) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceEthernets GET path
// Returns all available ethernet types
func (i *InterfaceEthernet) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceEthernet) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceEthernet DELETE path
// Creates or updates the complete object ethernet
func (*InterfaceEthernet) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceEthernet PATCH path
// Changes to parts of the object ethernet types
func (*InterfaceEthernet) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceEthernet POST path
// Create a new interface/ethernet object
func (*InterfaceEthernet) PostPath() string {
	return "/api/objects/interface/ethernet/"
}

// PutPath implements sophos.RestObject and returns the InterfaceEthernet PUT path
// Creates or updates the complete object ethernet
func (*InterfaceEthernet) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceEthernet) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s/usedby", ref)
}

// InterfaceGroups is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceGroups []InterfaceGroup

// InterfaceGroup represents a UTM interface group
type InterfaceGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	// Link default value is true
	Link    bool     `json:"link"`
	Members []string `json:"members"`
	Name    string   `json:"name"`
	// PrimaryAddresses description: REF(network/interface_address), REF(network/group)
	// PrimaryAddresses default value is ""
	PrimaryAddresses string `json:"primary_addresses"`
}

var _ sophos.RestGetter = &InterfaceGroup{}

// GetPath implements sophos.RestObject and returns the InterfaceGroups GET path
// Returns all available interface/group objects
func (*InterfaceGroups) GetPath() string { return "/api/objects/interface/group/" }

// RefRequired implements sophos.RestObject
func (*InterfaceGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceGroups GET path
// Returns all available group types
func (i *InterfaceGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/group/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceGroup) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceGroup DELETE path
// Creates or updates the complete object group
func (*InterfaceGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceGroup PATCH path
// Changes to parts of the object group types
func (*InterfaceGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceGroup POST path
// Create a new interface/group object
func (*InterfaceGroup) PostPath() string {
	return "/api/objects/interface/group/"
}

// PutPath implements sophos.RestObject and returns the InterfaceGroup PUT path
// Creates or updates the complete object group
func (*InterfaceGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s/usedby", ref)
}

// InterfacePpp3Gs is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePpp3Gs []InterfacePpp3G

// InterfacePpp3G represents a UTM ppp3g
type InterfacePpp3G struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Apn default value is "unknown"
	Apn string `json:"apn"`
	// ApnAuto default value is true
	ApnAuto   bool   `json:"apn_auto"`
	Bandwidth int    `json:"bandwidth"`
	Comment   string `json:"comment"`
	// Custom default value is ""
	Custom string `json:"custom"`
	// DialString default value is "*99#"
	DialString string `json:"dial_string"`
	// IdleTime default value is ""
	IdleTime    string `json:"idle_time"`
	Inbandwidth int    `json:"inbandwidth"`
	// InitString default value is "ATZ"
	InitString string `json:"init_string"`
	// Itfhw description: REF(itfhw/usbserial)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	// MobileNetwork can be one of: []string{"gsm", "cdma", "lte"}
	// MobileNetwork default value is "gsm"
	MobileNetwork string `json:"mobile_network"`
	Mtu           int    `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool `json:"mtu_auto_discovery"`
	// Multilink default value is false
	Multilink    bool   `json:"multilink"`
	Name         string `json:"name"`
	Outbandwidth int    `json:"outbandwidth"`
	// Password default value is ""
	Password string `json:"password"`
	// Pin default value is ""
	Pin string `json:"pin"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// ResetString default value is "ATZ"
	ResetString string `json:"reset_string"`
	Signal      int    `json:"signal"`
	// Status default value is false
	Status   bool   `json:"status"`
	Username string `json:"username"`
	// VirtualDevice description: (REGEX)
	// VirtualDevice default value is ""
	VirtualDevice string `json:"virtual_device"`
}

var _ sophos.RestGetter = &InterfacePpp3G{}

// GetPath implements sophos.RestObject and returns the InterfacePpp3Gs GET path
// Returns all available interface/ppp3g objects
func (*InterfacePpp3Gs) GetPath() string { return "/api/objects/interface/ppp3g/" }

// RefRequired implements sophos.RestObject
func (*InterfacePpp3Gs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfacePpp3Gs GET path
// Returns all available ppp3g types
func (i *InterfacePpp3G) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfacePpp3G) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfacePpp3G DELETE path
// Creates or updates the complete object ppp3g
func (*InterfacePpp3G) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePpp3G PATCH path
// Changes to parts of the object ppp3g types
func (*InterfacePpp3G) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePpp3G POST path
// Create a new interface/ppp3g object
func (*InterfacePpp3G) PostPath() string {
	return "/api/objects/interface/ppp3g/"
}

// PutPath implements sophos.RestObject and returns the InterfacePpp3G PUT path
// Creates or updates the complete object ppp3g
func (*InterfacePpp3G) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePpp3G) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s/usedby", ref)
}

// InterfacePppmodems is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppmodems []InterfacePppmodem

// InterfacePppmodem represents a UTM PPP modem interface
type InterfacePppmodem struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Bandwidth  int    `json:"bandwidth"`
	Comment    string `json:"comment"`
	// Custom default value is ""
	Custom string `json:"custom"`
	// DialString default value is ""
	DialString string `json:"dial_string"`
	// FlowControl can be one of: []string{"hardware", "software"}
	// FlowControl default value is "hardware"
	FlowControl string `json:"flow_control"`
	// IdleTime default value is ""
	IdleTime    string `json:"idle_time"`
	Inbandwidth int    `json:"inbandwidth"`
	// InitString default value is "ATZ"
	InitString string `json:"init_string"`
	// Itfhw description: REF(itfhw/serial)
	Itfhw string `json:"itfhw"`
	// LineSpeed can be one of: []string{"9600", "14400", "19200", "26400", "31200", "38400", "57600", "115200", "230400"}
	// LineSpeed default value is "115200"
	LineSpeed string `json:"line_speed"`
	// Link default value is true
	Link bool `json:"link"`
	Mtu  int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool `json:"mtu_auto_discovery"`
	// Multilink default value is false
	Multilink    bool   `json:"multilink"`
	Name         string `json:"name"`
	Outbandwidth int    `json:"outbandwidth"`
	// Password default value is ""
	Password string `json:"password"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// ResetString default value is "ATZ"
	ResetString string `json:"reset_string"`
	// Status default value is false
	Status   bool   `json:"status"`
	Username string `json:"username"`
	// VirtualDevice description: (REGEX)
	// VirtualDevice default value is ""
	VirtualDevice string `json:"virtual_device"`
}

var _ sophos.RestGetter = &InterfacePppmodem{}

// GetPath implements sophos.RestObject and returns the InterfacePppmodems GET path
// Returns all available interface/pppmodem objects
func (*InterfacePppmodems) GetPath() string { return "/api/objects/interface/pppmodem/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppmodems) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfacePppmodems GET path
// Returns all available pppmodem types
func (i *InterfacePppmodem) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfacePppmodem) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfacePppmodem DELETE path
// Creates or updates the complete object pppmodem
func (*InterfacePppmodem) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppmodem PATCH path
// Changes to parts of the object pppmodem types
func (*InterfacePppmodem) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppmodem POST path
// Create a new interface/pppmodem object
func (*InterfacePppmodem) PostPath() string {
	return "/api/objects/interface/pppmodem/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppmodem PUT path
// Creates or updates the complete object pppmodem
func (*InterfacePppmodem) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppmodem) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s/usedby", ref)
}

// InterfacePppoas is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppoas []InterfacePppoa

// InterfacePppoa represents a UTM PPPoA/PPTP DSL interface
type InterfacePppoa struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Bandwidth  int    `json:"bandwidth"`
	Comment    string `json:"comment"`
	// Custom default value is ""
	Custom      string `json:"custom"`
	Inbandwidth int    `json:"inbandwidth"`
	// Itfhw description: REF(itfhw/ethernet)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	// ModemAddress description: (IPADDR)
	ModemAddress string `json:"modem_address"`
	Mtu          int    `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool `json:"mtu_auto_discovery"`
	// Multilink default value is false
	Multilink bool   `json:"multilink"`
	Name      string `json:"name"`
	// NicAddress description: (IPADDR)
	NicAddress   string `json:"nic_address"`
	NicNetmask   int    `json:"nic_netmask"`
	Outbandwidth int    `json:"outbandwidth"`
	// Password default value is ""
	Password string `json:"password"`
	// PingAddress description: (IPADDR)
	// PingAddress default value is ""
	PingAddress string `json:"ping_address"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// ReconnectDaily description: (TIME)
	// ReconnectDaily default value is ""
	ReconnectDaily   string `json:"reconnect_daily"`
	ReconnectTimeout int    `json:"reconnect_timeout"`
	// Status default value is false
	Status   bool   `json:"status"`
	Username string `json:"username"`
	// VirtualDevice description: (REGEX)
	// VirtualDevice default value is ""
	VirtualDevice string `json:"virtual_device"`
}

var _ sophos.RestGetter = &InterfacePppoa{}

// GetPath implements sophos.RestObject and returns the InterfacePppoas GET path
// Returns all available interface/pppoa objects
func (*InterfacePppoas) GetPath() string { return "/api/objects/interface/pppoa/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppoas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfacePppoas GET path
// Returns all available pppoa types
func (i *InterfacePppoa) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfacePppoa) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfacePppoa DELETE path
// Creates or updates the complete object pppoa
func (*InterfacePppoa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppoa PATCH path
// Changes to parts of the object pppoa types
func (*InterfacePppoa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppoa POST path
// Create a new interface/pppoa object
func (*InterfacePppoa) PostPath() string {
	return "/api/objects/interface/pppoa/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppoa PUT path
// Creates or updates the complete object pppoa
func (*InterfacePppoa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppoa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s/usedby", ref)
}

// InterfacePppoes is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppoes []InterfacePppoe

// InterfacePppoe represents a UTM PPPoE DSL interface
type InterfacePppoe struct {
	Locked              string   `json:"_locked"`
	ObjectType          string   `json:"_type"`
	Reference           string   `json:"_ref"`
	AdditionalAddresses []string `json:"additional_addresses"`
	Bandwidth           int      `json:"bandwidth"`
	Comment             string   `json:"comment"`
	// Custom default value is ""
	Custom      string `json:"custom"`
	Inbandwidth int    `json:"inbandwidth"`
	// Itfhw description: REF(itfhw/ethernet)
	Itfhw       string   `json:"itfhw"`
	ItfhwSlaves []string `json:"itfhw_slaves"`
	// Link default value is true
	Link bool `json:"link"`
	// Macvlan default value is false
	Macvlan bool `json:"macvlan"`
	Mtu     int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool `json:"mtu_auto_discovery"`
	// Multilink default value is false
	Multilink bool `json:"multilink"`
	// MultilinkStatus description: (HASH)
	MultilinkStatus interface{} `json:"multilink_status"`
	Name            string      `json:"name"`
	Outbandwidth    int         `json:"outbandwidth"`
	// Password default value is ""
	Password string `json:"password"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// ReconnectDaily description: (TIME)
	// ReconnectDaily default value is ""
	ReconnectDaily   string `json:"reconnect_daily"`
	ReconnectTimeout int    `json:"reconnect_timeout"`
	// Status default value is false
	Status   bool   `json:"status"`
	Username string `json:"username"`
	// VirtualDevice description: (REGEX)
	// VirtualDevice default value is ""
	VirtualDevice string `json:"virtual_device"`
	Vlantag       int    `json:"vlantag"`
}

var _ sophos.RestGetter = &InterfacePppoe{}

// GetPath implements sophos.RestObject and returns the InterfacePppoes GET path
// Returns all available interface/pppoe objects
func (*InterfacePppoes) GetPath() string { return "/api/objects/interface/pppoe/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppoes) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfacePppoes GET path
// Returns all available pppoe types
func (i *InterfacePppoe) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfacePppoe) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfacePppoe DELETE path
// Creates or updates the complete object pppoe
func (*InterfacePppoe) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppoe PATCH path
// Changes to parts of the object pppoe types
func (*InterfacePppoe) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppoe POST path
// Create a new interface/pppoe object
func (*InterfacePppoe) PostPath() string {
	return "/api/objects/interface/pppoe/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppoe PUT path
// Creates or updates the complete object pppoe
func (*InterfacePppoe) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppoe) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s/usedby", ref)
}

// InterfaceTunnels is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceTunnels []InterfaceTunnel

// InterfaceTunnel represents a UTM tunnel interface
type InterfaceTunnel struct {
	Locked              string   `json:"_locked"`
	ObjectType          string   `json:"_type"`
	Reference           string   `json:"_ref"`
	AdditionalAddresses []string `json:"additional_addresses"`
	Bandwidth           int      `json:"bandwidth"`
	Comment             string   `json:"comment"`
	Inbandwidth         int      `json:"inbandwidth"`
	// Itfhw description: REF(itfhw/virtual)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	Mtu  int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool   `json:"mtu_auto_discovery"`
	Name             string `json:"name"`
	Outbandwidth     int    `json:"outbandwidth"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &InterfaceTunnel{}

// GetPath implements sophos.RestObject and returns the InterfaceTunnels GET path
// Returns all available interface/tunnel objects
func (*InterfaceTunnels) GetPath() string { return "/api/objects/interface/tunnel/" }

// RefRequired implements sophos.RestObject
func (*InterfaceTunnels) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceTunnels GET path
// Returns all available tunnel types
func (i *InterfaceTunnel) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceTunnel) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceTunnel DELETE path
// Creates or updates the complete object tunnel
func (*InterfaceTunnel) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceTunnel PATCH path
// Changes to parts of the object tunnel types
func (*InterfaceTunnel) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceTunnel POST path
// Create a new interface/tunnel object
func (*InterfaceTunnel) PostPath() string {
	return "/api/objects/interface/tunnel/"
}

// PutPath implements sophos.RestObject and returns the InterfaceTunnel PUT path
// Creates or updates the complete object tunnel
func (*InterfaceTunnel) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceTunnel) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s/usedby", ref)
}

// InterfaceVlans is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceVlans []InterfaceVlan

// InterfaceVlan represents a UTM ethernet VLAN interface
type InterfaceVlan struct {
	Locked              string   `json:"_locked"`
	ObjectType          string   `json:"_type"`
	Reference           string   `json:"_ref"`
	AdditionalAddresses []string `json:"additional_addresses"`
	Bandwidth           int      `json:"bandwidth"`
	Comment             string   `json:"comment"`
	Inbandwidth         int      `json:"inbandwidth"`
	// Itfhw description: REF(itfhw/ethernet), REF(itfhw/bridge), REF(itfhw/lag), REF(itfhw/red_server), REF(itfhw/red_client), REF(itfhw/awe_network)
	Itfhw string `json:"itfhw"`
	// Link default value is true
	Link bool `json:"link"`
	// Macvlan default value is false
	Macvlan bool `json:"macvlan"`
	Mtu     int  `json:"mtu"`
	// MtuAutoDiscovery default value is false
	MtuAutoDiscovery bool   `json:"mtu_auto_discovery"`
	Name             string `json:"name"`
	Outbandwidth     int    `json:"outbandwidth"`
	// PrimaryAddress description: REF(itfparams/primary)
	// PrimaryAddress default value is ""
	PrimaryAddress string `json:"primary_address"`
	// Proxyarp default value is false
	Proxyarp bool `json:"proxyarp"`
	// Proxyndp default value is false
	Proxyndp bool `json:"proxyndp"`
	// Status default value is false
	Status  bool `json:"status"`
	Vlantag int  `json:"vlantag"`
}

var _ sophos.RestGetter = &InterfaceVlan{}

// GetPath implements sophos.RestObject and returns the InterfaceVlans GET path
// Returns all available interface/vlan objects
func (*InterfaceVlans) GetPath() string { return "/api/objects/interface/vlan/" }

// RefRequired implements sophos.RestObject
func (*InterfaceVlans) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceVlans GET path
// Returns all available vlan types
func (i *InterfaceVlan) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceVlan) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceVlan DELETE path
// Creates or updates the complete object vlan
func (*InterfaceVlan) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceVlan PATCH path
// Changes to parts of the object vlan types
func (*InterfaceVlan) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceVlan POST path
// Create a new interface/vlan object
func (*InterfaceVlan) PostPath() string {
	return "/api/objects/interface/vlan/"
}

// PutPath implements sophos.RestObject and returns the InterfaceVlan PUT path
// Creates or updates the complete object vlan
func (*InterfaceVlan) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceVlan) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s/usedby", ref)
}
