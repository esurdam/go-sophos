package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Network is a generated struct representing the Sophos Network Endpoint
// GET /api/nodes/network
type Network struct {
	NetworkAaa                NetworkAaa                `json:"network_aaa"`
	NetworkAny                NetworkAny                `json:"network_any"`
	NetworkAvailabilityGroup  NetworkAvailabilityGroup  `json:"network_availability_group"`
	NetworkDnsGroup           NetworkDnsGroup           `json:"network_dns_group"`
	NetworkDnsHost            NetworkDnsHost            `json:"network_dns_host"`
	NetworkGroup              NetworkGroup              `json:"network_group"`
	NetworkHost               NetworkHost               `json:"network_host"`
	NetworkInterfaceAddress   NetworkInterfaceAddress   `json:"network_interface_address"`
	NetworkInterfaceBroadcast NetworkInterfaceBroadcast `json:"network_interface_broadcast"`
	NetworkInterfaceNetwork   NetworkInterfaceNetwork   `json:"network_interface_network"`
	NetworkMulticast          NetworkMulticast          `json:"network_multicast"`
	NetworkNetwork            NetworkNetwork            `json:"network_network"`
	NetworkRange              NetworkRange              `json:"network_range"`
}

var _ sophos.Endpoint = &Network{}

var defsNetwork = map[string]sophos.RestObject{
	"NetworkAaa":                &NetworkAaa{},
	"NetworkAny":                &NetworkAny{},
	"NetworkAvailabilityGroup":  &NetworkAvailabilityGroup{},
	"NetworkDnsGroup":           &NetworkDnsGroup{},
	"NetworkDnsHost":            &NetworkDnsHost{},
	"NetworkGroup":              &NetworkGroup{},
	"NetworkHost":               &NetworkHost{},
	"NetworkInterfaceAddress":   &NetworkInterfaceAddress{},
	"NetworkInterfaceBroadcast": &NetworkInterfaceBroadcast{},
	"NetworkInterfaceNetwork":   &NetworkInterfaceNetwork{},
	"NetworkMulticast":          &NetworkMulticast{},
	"NetworkNetwork":            &NetworkNetwork{},
	"NetworkRange":              &NetworkRange{},
}

// RestObjects implements the sophos.Node interface and returns a map of Network's Objects
func (Network) RestObjects() map[string]sophos.RestObject { return defsNetwork }

// GetPath implements sophos.RestGetter
func (*Network) GetPath() string { return "/api/nodes/network" }

// RefRequired implements sophos.RestGetter
func (*Network) RefRequired() (string, bool) { return "", false }

var defNetwork = &sophos.Definition{Description: "network", Name: "network", Link: "/api/definitions/network"}

// Definition returns the /api/definitions struct of Network
func (Network) Definition() sophos.Definition { return *defNetwork }

// ApiRoutes returns all known Network Paths
func (Network) ApiRoutes() []string {
	return []string{
		"/api/objects/network/aaa/",
		"/api/objects/network/aaa/{ref}",
		"/api/objects/network/aaa/{ref}/usedby",
		"/api/objects/network/any/",
		"/api/objects/network/any/{ref}",
		"/api/objects/network/any/{ref}/usedby",
		"/api/objects/network/availability_group/",
		"/api/objects/network/availability_group/{ref}",
		"/api/objects/network/availability_group/{ref}/usedby",
		"/api/objects/network/dns_group/",
		"/api/objects/network/dns_group/{ref}",
		"/api/objects/network/dns_group/{ref}/usedby",
		"/api/objects/network/dns_host/",
		"/api/objects/network/dns_host/{ref}",
		"/api/objects/network/dns_host/{ref}/usedby",
		"/api/objects/network/group/",
		"/api/objects/network/group/{ref}",
		"/api/objects/network/group/{ref}/usedby",
		"/api/objects/network/host/",
		"/api/objects/network/host/{ref}",
		"/api/objects/network/host/{ref}/usedby",
		"/api/objects/network/interface_address/",
		"/api/objects/network/interface_address/{ref}",
		"/api/objects/network/interface_address/{ref}/usedby",
		"/api/objects/network/interface_broadcast/",
		"/api/objects/network/interface_broadcast/{ref}",
		"/api/objects/network/interface_broadcast/{ref}/usedby",
		"/api/objects/network/interface_network/",
		"/api/objects/network/interface_network/{ref}",
		"/api/objects/network/interface_network/{ref}/usedby",
		"/api/objects/network/multicast/",
		"/api/objects/network/multicast/{ref}",
		"/api/objects/network/multicast/{ref}/usedby",
		"/api/objects/network/network/",
		"/api/objects/network/network/{ref}",
		"/api/objects/network/network/{ref}/usedby",
		"/api/objects/network/range/",
		"/api/objects/network/range/{ref}",
		"/api/objects/network/range/{ref}/usedby",
	}
}

// References returns the Network's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Network) References() []string {
	return []string{
		"REF_NetworkAaa",
		"REF_NetworkAny",
		"REF_NetworkAvailabilityGroup",
		"REF_NetworkDnsGroup",
		"REF_NetworkDnsHost",
		"REF_NetworkGroup",
		"REF_NetworkHost",
		"REF_NetworkInterfaceAddress",
		"REF_NetworkInterfaceBroadcast",
		"REF_NetworkInterfaceNetwork",
		"REF_NetworkMulticast",
		"REF_NetworkNetwork",
		"REF_NetworkRange",
	}
}

// NetworkAaas is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkAaas []NetworkAaa

// NetworkAaa is a generated Sophos object
type NetworkAaa struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	_type      string        `json:"_type"`
	Addresses  []interface{} `json:"addresses"`
	Addresses6 []interface{} `json:"addresses6"`
	Comment    string        `json:"comment"`
	Name       string        `json:"name"`
	Resolved   bool          `json:"resolved"`
	Resolved6  bool          `json:"resolved6"`
}

var _ sophos.RestGetter = &NetworkAaa{}

// GetPath implements sophos.RestObject and returns the NetworkAaas GET path
// Returns all available network/aaa objects
func (*NetworkAaas) GetPath() string { return "/api/objects/network/aaa/" }

// RefRequired implements sophos.RestObject
func (*NetworkAaas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkAaas GET path
// Returns all available aaa types
func (n *NetworkAaa) GetPath() string { return fmt.Sprintf("/api/objects/network/aaa/%s", n.Reference) }

// RefRequired implements sophos.RestObject
func (n *NetworkAaa) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkAaa DELETE path
// Creates or updates the complete object aaa
func (*NetworkAaa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/aaa/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkAaa PATCH path
// Changes to parts of the object aaa types
func (*NetworkAaa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/aaa/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkAaa POST path
// Create a new network/aaa object
func (*NetworkAaa) PostPath() string {
	return "/api/objects/network/aaa/"
}

// PutPath implements sophos.RestObject and returns the NetworkAaa PUT path
// Creates or updates the complete object aaa
func (*NetworkAaa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/aaa/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkAaa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/aaa/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkAaa) GetType() string { return n._type }

// NetworkAnys is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkAnys []NetworkAny

// NetworkAny is a generated Sophos object
type NetworkAny struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Address6  string `json:"address6"`
	Comment   string `json:"comment"`
	Interface string `json:"interface"`
	Name      string `json:"name"`
	Netmask   int64  `json:"netmask"`
	Netmask6  int64  `json:"netmask6"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
}

var _ sophos.RestGetter = &NetworkAny{}

// GetPath implements sophos.RestObject and returns the NetworkAnys GET path
// Returns all available network/any objects
func (*NetworkAnys) GetPath() string { return "/api/objects/network/any/" }

// RefRequired implements sophos.RestObject
func (*NetworkAnys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkAnys GET path
// Returns all available any types
func (n *NetworkAny) GetPath() string { return fmt.Sprintf("/api/objects/network/any/%s", n.Reference) }

// RefRequired implements sophos.RestObject
func (n *NetworkAny) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkAny DELETE path
// Creates or updates the complete object any
func (*NetworkAny) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/any/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkAny PATCH path
// Changes to parts of the object any types
func (*NetworkAny) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/any/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkAny POST path
// Create a new network/any object
func (*NetworkAny) PostPath() string {
	return "/api/objects/network/any/"
}

// PutPath implements sophos.RestObject and returns the NetworkAny PUT path
// Creates or updates the complete object any
func (*NetworkAny) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/any/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkAny) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/any/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkAny) GetType() string { return n._type }

// NetworkAvailabilityGroup is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkAvailabilityGroup []interface{}

var _ sophos.RestObject = &NetworkAvailabilityGroup{}

// GetPath implements sophos.RestObject and returns the NetworkAvailabilityGroup GET path
// Returns all available network/availability_group objects
func (*NetworkAvailabilityGroup) GetPath() string { return "/api/objects/network/availability_group/" }

// RefRequired implements sophos.RestObject
func (*NetworkAvailabilityGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the NetworkAvailabilityGroup DELETE path
// Creates or updates the complete object availability_group
func (*NetworkAvailabilityGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/availability_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkAvailabilityGroup PATCH path
// Changes to parts of the object availability_group types
func (*NetworkAvailabilityGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/availability_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkAvailabilityGroup POST path
// Create a new network/availability_group object
func (*NetworkAvailabilityGroup) PostPath() string {
	return "/api/objects/network/availability_group/"
}

// PutPath implements sophos.RestObject and returns the NetworkAvailabilityGroup PUT path
// Creates or updates the complete object availability_group
func (*NetworkAvailabilityGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/availability_group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkAvailabilityGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/availability_group/%s/usedby", ref)
}

// NetworkDnsGroups is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkDnsGroups []NetworkDnsGroup

// NetworkDnsGroup is a generated Sophos object
type NetworkDnsGroup struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	_type      string        `json:"_type"`
	Addresses  []string      `json:"addresses"`
	Addresses6 []interface{} `json:"addresses6"`
	Comment    string        `json:"comment"`
	Hostname   string        `json:"hostname"`
	Interface  string        `json:"interface"`
	Name       string        `json:"name"`
	Resolved   bool          `json:"resolved"`
	Resolved6  bool          `json:"resolved6"`
	Timeout    int64         `json:"timeout"`
}

var _ sophos.RestGetter = &NetworkDnsGroup{}

// GetPath implements sophos.RestObject and returns the NetworkDnsGroups GET path
// Returns all available network/dns_group objects
func (*NetworkDnsGroups) GetPath() string { return "/api/objects/network/dns_group/" }

// RefRequired implements sophos.RestObject
func (*NetworkDnsGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkDnsGroups GET path
// Returns all available dns_group types
func (n *NetworkDnsGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/network/dns_group/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkDnsGroup) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkDnsGroup DELETE path
// Creates or updates the complete object dns_group
func (*NetworkDnsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkDnsGroup PATCH path
// Changes to parts of the object dns_group types
func (*NetworkDnsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkDnsGroup POST path
// Create a new network/dns_group object
func (*NetworkDnsGroup) PostPath() string {
	return "/api/objects/network/dns_group/"
}

// PutPath implements sophos.RestObject and returns the NetworkDnsGroup PUT path
// Creates or updates the complete object dns_group
func (*NetworkDnsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkDnsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkDnsGroup) GetType() string { return n._type }

// NetworkDnsHosts is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkDnsHosts []NetworkDnsHost

// NetworkDnsHost is a generated Sophos object
type NetworkDnsHost struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Address6  string `json:"address6"`
	Comment   string `json:"comment"`
	Hostname  string `json:"hostname"`
	Interface string `json:"interface"`
	Name      string `json:"name"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
	Timeout   int64  `json:"timeout"`
}

var _ sophos.RestGetter = &NetworkDnsHost{}

// GetPath implements sophos.RestObject and returns the NetworkDnsHosts GET path
// Returns all available network/dns_host objects
func (*NetworkDnsHosts) GetPath() string { return "/api/objects/network/dns_host/" }

// RefRequired implements sophos.RestObject
func (*NetworkDnsHosts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkDnsHosts GET path
// Returns all available dns_host types
func (n *NetworkDnsHost) GetPath() string {
	return fmt.Sprintf("/api/objects/network/dns_host/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkDnsHost) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkDnsHost DELETE path
// Creates or updates the complete object dns_host
func (*NetworkDnsHost) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_host/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkDnsHost PATCH path
// Changes to parts of the object dns_host types
func (*NetworkDnsHost) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_host/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkDnsHost POST path
// Create a new network/dns_host object
func (*NetworkDnsHost) PostPath() string {
	return "/api/objects/network/dns_host/"
}

// PutPath implements sophos.RestObject and returns the NetworkDnsHost PUT path
// Creates or updates the complete object dns_host
func (*NetworkDnsHost) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_host/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkDnsHost) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/dns_host/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkDnsHost) GetType() string { return n._type }

// NetworkGroups is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkGroups []NetworkGroup

// NetworkGroup is a generated Sophos object
type NetworkGroup struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	Members   []string `json:"members"`
	Name      string   `json:"name"`
	Types     []string `json:"types"`
}

var _ sophos.RestGetter = &NetworkGroup{}

// GetPath implements sophos.RestObject and returns the NetworkGroups GET path
// Returns all available network/group objects
func (*NetworkGroups) GetPath() string { return "/api/objects/network/group/" }

// RefRequired implements sophos.RestObject
func (*NetworkGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkGroups GET path
// Returns all available group types
func (n *NetworkGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/network/group/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkGroup) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkGroup DELETE path
// Creates or updates the complete object group
func (*NetworkGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkGroup PATCH path
// Changes to parts of the object group types
func (*NetworkGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkGroup POST path
// Create a new network/group object
func (*NetworkGroup) PostPath() string {
	return "/api/objects/network/group/"
}

// PutPath implements sophos.RestObject and returns the NetworkGroup PUT path
// Creates or updates the complete object group
func (*NetworkGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkGroup) GetType() string { return n._type }

// NetworkHosts is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkHosts []NetworkHost

// NetworkHost is a generated Sophos object
type NetworkHost struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	_type      string        `json:"_type"`
	Address    string        `json:"address"`
	Address6   string        `json:"address6"`
	Comment    string        `json:"comment"`
	Duids      []interface{} `json:"duids"`
	Hostnames  []string      `json:"hostnames"`
	Interface  string        `json:"interface"`
	Macs       []interface{} `json:"macs"`
	Name       string        `json:"name"`
	Resolved   bool          `json:"resolved"`
	Resolved6  bool          `json:"resolved6"`
	ReverseDNS bool          `json:"reverse_dns"`
}

var _ sophos.RestGetter = &NetworkHost{}

// GetPath implements sophos.RestObject and returns the NetworkHosts GET path
// Returns all available network/host objects
func (*NetworkHosts) GetPath() string { return "/api/objects/network/host/" }

// RefRequired implements sophos.RestObject
func (*NetworkHosts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkHosts GET path
// Returns all available host types
func (n *NetworkHost) GetPath() string {
	return fmt.Sprintf("/api/objects/network/host/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkHost) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkHost DELETE path
// Creates or updates the complete object host
func (*NetworkHost) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/host/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkHost PATCH path
// Changes to parts of the object host types
func (*NetworkHost) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/host/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkHost POST path
// Create a new network/host object
func (*NetworkHost) PostPath() string {
	return "/api/objects/network/host/"
}

// PutPath implements sophos.RestObject and returns the NetworkHost PUT path
// Creates or updates the complete object host
func (*NetworkHost) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/host/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkHost) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/host/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkHost) GetType() string { return n._type }

// NetworkInterfaceAddresss is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkInterfaceAddresss []NetworkInterfaceAddress

// NetworkInterfaceAddress is a generated Sophos object
type NetworkInterfaceAddress struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Address6  string `json:"address6"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
}

var _ sophos.RestGetter = &NetworkInterfaceAddress{}

// GetPath implements sophos.RestObject and returns the NetworkInterfaceAddresss GET path
// Returns all available network/interface_address objects
func (*NetworkInterfaceAddresss) GetPath() string { return "/api/objects/network/interface_address/" }

// RefRequired implements sophos.RestObject
func (*NetworkInterfaceAddresss) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkInterfaceAddresss GET path
// Returns all available interface_address types
func (n *NetworkInterfaceAddress) GetPath() string {
	return fmt.Sprintf("/api/objects/network/interface_address/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkInterfaceAddress) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkInterfaceAddress DELETE path
// Creates or updates the complete object interface_address
func (*NetworkInterfaceAddress) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_address/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkInterfaceAddress PATCH path
// Changes to parts of the object interface_address types
func (*NetworkInterfaceAddress) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_address/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkInterfaceAddress POST path
// Create a new network/interface_address object
func (*NetworkInterfaceAddress) PostPath() string {
	return "/api/objects/network/interface_address/"
}

// PutPath implements sophos.RestObject and returns the NetworkInterfaceAddress PUT path
// Creates or updates the complete object interface_address
func (*NetworkInterfaceAddress) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_address/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkInterfaceAddress) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_address/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkInterfaceAddress) GetType() string { return n._type }

// NetworkInterfaceBroadcasts is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkInterfaceBroadcasts []NetworkInterfaceBroadcast

// NetworkInterfaceBroadcast is a generated Sophos object
type NetworkInterfaceBroadcast struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Resolved  bool   `json:"resolved"`
}

var _ sophos.RestGetter = &NetworkInterfaceBroadcast{}

// GetPath implements sophos.RestObject and returns the NetworkInterfaceBroadcasts GET path
// Returns all available network/interface_broadcast objects
func (*NetworkInterfaceBroadcasts) GetPath() string {
	return "/api/objects/network/interface_broadcast/"
}

// RefRequired implements sophos.RestObject
func (*NetworkInterfaceBroadcasts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkInterfaceBroadcasts GET path
// Returns all available interface_broadcast types
func (n *NetworkInterfaceBroadcast) GetPath() string {
	return fmt.Sprintf("/api/objects/network/interface_broadcast/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkInterfaceBroadcast) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkInterfaceBroadcast DELETE path
// Creates or updates the complete object interface_broadcast
func (*NetworkInterfaceBroadcast) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_broadcast/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkInterfaceBroadcast PATCH path
// Changes to parts of the object interface_broadcast types
func (*NetworkInterfaceBroadcast) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_broadcast/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkInterfaceBroadcast POST path
// Create a new network/interface_broadcast object
func (*NetworkInterfaceBroadcast) PostPath() string {
	return "/api/objects/network/interface_broadcast/"
}

// PutPath implements sophos.RestObject and returns the NetworkInterfaceBroadcast PUT path
// Creates or updates the complete object interface_broadcast
func (*NetworkInterfaceBroadcast) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_broadcast/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkInterfaceBroadcast) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_broadcast/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkInterfaceBroadcast) GetType() string { return n._type }

// NetworkInterfaceNetworks is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkInterfaceNetworks []NetworkInterfaceNetwork

// NetworkInterfaceNetwork is a generated Sophos object
type NetworkInterfaceNetwork struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Address6  string `json:"address6"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Netmask   int64  `json:"netmask"`
	Netmask6  int64  `json:"netmask6"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
}

var _ sophos.RestGetter = &NetworkInterfaceNetwork{}

// GetPath implements sophos.RestObject and returns the NetworkInterfaceNetworks GET path
// Returns all available network/interface_network objects
func (*NetworkInterfaceNetworks) GetPath() string { return "/api/objects/network/interface_network/" }

// RefRequired implements sophos.RestObject
func (*NetworkInterfaceNetworks) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkInterfaceNetworks GET path
// Returns all available interface_network types
func (n *NetworkInterfaceNetwork) GetPath() string {
	return fmt.Sprintf("/api/objects/network/interface_network/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkInterfaceNetwork) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkInterfaceNetwork DELETE path
// Creates or updates the complete object interface_network
func (*NetworkInterfaceNetwork) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_network/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkInterfaceNetwork PATCH path
// Changes to parts of the object interface_network types
func (*NetworkInterfaceNetwork) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_network/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkInterfaceNetwork POST path
// Create a new network/interface_network object
func (*NetworkInterfaceNetwork) PostPath() string {
	return "/api/objects/network/interface_network/"
}

// PutPath implements sophos.RestObject and returns the NetworkInterfaceNetwork PUT path
// Creates or updates the complete object interface_network
func (*NetworkInterfaceNetwork) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_network/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkInterfaceNetwork) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/interface_network/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkInterfaceNetwork) GetType() string { return n._type }

// NetworkMulticast is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkMulticast []interface{}

var _ sophos.RestObject = &NetworkMulticast{}

// GetPath implements sophos.RestObject and returns the NetworkMulticast GET path
// Returns all available network/multicast objects
func (*NetworkMulticast) GetPath() string { return "/api/objects/network/multicast/" }

// RefRequired implements sophos.RestObject
func (*NetworkMulticast) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the NetworkMulticast DELETE path
// Creates or updates the complete object multicast
func (*NetworkMulticast) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/multicast/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkMulticast PATCH path
// Changes to parts of the object multicast types
func (*NetworkMulticast) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/multicast/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkMulticast POST path
// Create a new network/multicast object
func (*NetworkMulticast) PostPath() string {
	return "/api/objects/network/multicast/"
}

// PutPath implements sophos.RestObject and returns the NetworkMulticast PUT path
// Creates or updates the complete object multicast
func (*NetworkMulticast) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/multicast/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkMulticast) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/multicast/%s/usedby", ref)
}

// NetworkNetworks is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkNetworks []NetworkNetwork

// NetworkNetwork is a generated Sophos object
type NetworkNetwork struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Address   string `json:"address"`
	Address6  string `json:"address6"`
	Comment   string `json:"comment"`
	Interface string `json:"interface"`
	Name      string `json:"name"`
	Netmask   int64  `json:"netmask"`
	Netmask6  int64  `json:"netmask6"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
}

var _ sophos.RestGetter = &NetworkNetwork{}

// GetPath implements sophos.RestObject and returns the NetworkNetworks GET path
// Returns all available network/network objects
func (*NetworkNetworks) GetPath() string { return "/api/objects/network/network/" }

// RefRequired implements sophos.RestObject
func (*NetworkNetworks) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkNetworks GET path
// Returns all available network types
func (n *NetworkNetwork) GetPath() string {
	return fmt.Sprintf("/api/objects/network/network/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkNetwork) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkNetwork DELETE path
// Creates or updates the complete object network
func (*NetworkNetwork) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/network/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkNetwork PATCH path
// Changes to parts of the object network types
func (*NetworkNetwork) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/network/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkNetwork POST path
// Create a new network/network object
func (*NetworkNetwork) PostPath() string {
	return "/api/objects/network/network/"
}

// PutPath implements sophos.RestObject and returns the NetworkNetwork PUT path
// Creates or updates the complete object network
func (*NetworkNetwork) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/network/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkNetwork) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/network/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkNetwork) GetType() string { return n._type }

// NetworkRanges is an Sophos Endpoint subType and implements sophos.RestObject
type NetworkRanges []NetworkRange

// NetworkRange is a generated Sophos object
type NetworkRange struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	From      string `json:"from"`
	From6     string `json:"from6"`
	Interface string `json:"interface"`
	Name      string `json:"name"`
	Resolved  bool   `json:"resolved"`
	Resolved6 bool   `json:"resolved6"`
	To        string `json:"to"`
	To6       string `json:"to6"`
}

var _ sophos.RestGetter = &NetworkRange{}

// GetPath implements sophos.RestObject and returns the NetworkRanges GET path
// Returns all available network/range objects
func (*NetworkRanges) GetPath() string { return "/api/objects/network/range/" }

// RefRequired implements sophos.RestObject
func (*NetworkRanges) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NetworkRanges GET path
// Returns all available range types
func (n *NetworkRange) GetPath() string {
	return fmt.Sprintf("/api/objects/network/range/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NetworkRange) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NetworkRange DELETE path
// Creates or updates the complete object range
func (*NetworkRange) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/network/range/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NetworkRange PATCH path
// Changes to parts of the object range types
func (*NetworkRange) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/range/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NetworkRange POST path
// Create a new network/range object
func (*NetworkRange) PostPath() string {
	return "/api/objects/network/range/"
}

// PutPath implements sophos.RestObject and returns the NetworkRange PUT path
// Creates or updates the complete object range
func (*NetworkRange) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/range/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NetworkRange) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/network/range/%s/usedby", ref)
}

// GetType implements sophos.Object
func (n *NetworkRange) GetType() string { return n._type }
