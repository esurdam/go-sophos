package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Dhcp is a generated struct representing the Sophos Dhcp Endpoint
// GET /api/nodes/dhcp
type Dhcp struct {
	Relay struct {
		DhcpServer string        `json:"dhcp_server"`
		Interfaces []interface{} `json:"interfaces"`
		Status     int64         `json:"status"`
	} `json:"relay"`
	Relay6 struct {
		ItfsFacingClients []interface{} `json:"itfs_facing_clients"`
		ItfsFacingServer6 []interface{} `json:"itfs_facing_server6"`
		Status            int64         `json:"status"`
	} `json:"relay6"`
	Server struct {
		Custom4 string   `json:"custom4"`
		Custom6 string   `json:"custom6"`
		Servers []string `json:"servers"`
	} `json:"server"`
}

var _ sophos.Endpoint = &Dhcp{}

var defsDhcp = map[string]sophos.RestObject{
	"DhcpGroup":     &DhcpGroup{},
	"DhcpOption":    &DhcpOption{},
	"DhcpOption6":   &DhcpOption6{},
	"DhcpServer":    &DhcpServer{},
	"DhcpServer6":   &DhcpServer6{},
	"DhcpStateless": &DhcpStateless{},
}

// RestObjects implements the sophos.Node interface and returns a map of Dhcp's Objects
func (Dhcp) RestObjects() map[string]sophos.RestObject { return defsDhcp }

// GetPath implements sophos.RestGetter
func (*Dhcp) GetPath() string { return "/api/nodes/dhcp" }

// RefRequired implements sophos.RestGetter
func (*Dhcp) RefRequired() (string, bool) { return "", false }

var defDhcp = &sophos.Definition{Description: "dhcp", Name: "dhcp", Link: "/api/definitions/dhcp"}

// Definition returns the /api/definitions struct of Dhcp
func (Dhcp) Definition() sophos.Definition { return *defDhcp }

// ApiRoutes returns all known Dhcp Paths
func (Dhcp) ApiRoutes() []string {
	return []string{
		"/api/objects/dhcp/group/",
		"/api/objects/dhcp/group/{ref}",
		"/api/objects/dhcp/group/{ref}/usedby",
		"/api/objects/dhcp/option/",
		"/api/objects/dhcp/option/{ref}",
		"/api/objects/dhcp/option/{ref}/usedby",
		"/api/objects/dhcp/option6/",
		"/api/objects/dhcp/option6/{ref}",
		"/api/objects/dhcp/option6/{ref}/usedby",
		"/api/objects/dhcp/server/",
		"/api/objects/dhcp/server/{ref}",
		"/api/objects/dhcp/server/{ref}/usedby",
		"/api/objects/dhcp/server6/",
		"/api/objects/dhcp/server6/{ref}",
		"/api/objects/dhcp/server6/{ref}/usedby",
		"/api/objects/dhcp/stateless/",
		"/api/objects/dhcp/stateless/{ref}",
		"/api/objects/dhcp/stateless/{ref}/usedby",
	}
}

// References returns the Dhcp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Dhcp) References() []string {
	return []string{
		"REF_DhcpGroup",
		"REF_DhcpOption",
		"REF_DhcpOption6",
		"REF_DhcpServer",
		"REF_DhcpServer6",
		"REF_DhcpStateless",
	}
}

// DhcpGroups is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpGroups []DhcpGroup

// DhcpGroup represents a UTM group
type DhcpGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &DhcpGroup{}

// GetPath implements sophos.RestObject and returns the DhcpGroups GET path
// Returns all available dhcp/group objects
func (*DhcpGroups) GetPath() string { return "/api/objects/dhcp/group/" }

// RefRequired implements sophos.RestObject
func (*DhcpGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpGroups GET path
// Returns all available group types
func (d *DhcpGroup) GetPath() string { return fmt.Sprintf("/api/objects/dhcp/group/%s", d.Reference) }

// RefRequired implements sophos.RestObject
func (d *DhcpGroup) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpGroup DELETE path
// Creates or updates the complete object group
func (*DhcpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpGroup PATCH path
// Changes to parts of the object group types
func (*DhcpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpGroup POST path
// Create a new dhcp/group object
func (*DhcpGroup) PostPath() string {
	return "/api/objects/dhcp/group/"
}

// PutPath implements sophos.RestObject and returns the DhcpGroup PUT path
// Creates or updates the complete object group
func (*DhcpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/group/%s/usedby", ref)
}

// DhcpOptions is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpOptions []DhcpOption

// DhcpOption is a generated Sophos object
type DhcpOption struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	ObjectType string        `json:"_type"`
	Address    string        `json:"address"`
	Code       int64         `json:"code"`
	Comment    string        `json:"comment"`
	DhcpName   string        `json:"dhcp_name"`
	Host       []interface{} `json:"host"`
	Integer    int64         `json:"integer"`
	Mac        string        `json:"mac"`
	Name       string        `json:"name"`
	Scope      string        `json:"scope"`
	Server     []interface{} `json:"server"`
	Status     bool          `json:"status"`
	String     string        `json:"string"`
	Text       string        `json:"text"`
	Type       string        `json:"type"`
	Vendor     string        `json:"vendor"`
}

var _ sophos.RestGetter = &DhcpOption{}

// GetPath implements sophos.RestObject and returns the DhcpOptions GET path
// Returns all available dhcp/option objects
func (*DhcpOptions) GetPath() string { return "/api/objects/dhcp/option/" }

// RefRequired implements sophos.RestObject
func (*DhcpOptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpOptions GET path
// Returns all available option types
func (d *DhcpOption) GetPath() string { return fmt.Sprintf("/api/objects/dhcp/option/%s", d.Reference) }

// RefRequired implements sophos.RestObject
func (d *DhcpOption) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpOption DELETE path
// Creates or updates the complete object option
func (*DhcpOption) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpOption PATCH path
// Changes to parts of the object option types
func (*DhcpOption) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpOption POST path
// Create a new dhcp/option object
func (*DhcpOption) PostPath() string {
	return "/api/objects/dhcp/option/"
}

// PutPath implements sophos.RestObject and returns the DhcpOption PUT path
// Creates or updates the complete object option
func (*DhcpOption) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpOption) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option/%s/usedby", ref)
}

// GetType implements sophos.Object
func (d *DhcpOption) GetType() string { return d.ObjectType }

// DhcpOption6s is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpOption6s []DhcpOption6

// DhcpOption6 represents a UTM DHCPv6 option
type DhcpOption6 struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// String description: (HEXSTRING)
	String string `json:"string"`
	// Address description: REF(network/interface_address), REF(network/host), REF(network/dns_host), REF(network/dns_group), REF(network/availability_group), REF(network/group)
	Address string        `json:"address"`
	Host    []interface{} `json:"host"`
	// Mac description: (HEXSTRING)
	Mac    string        `json:"mac"`
	Server []interface{} `json:"server"`
	// Status default value is false
	Status bool   `json:"status"`
	Text   string `json:"text"`
	// Code description: Constraints: 7, 10-12, 15-18, 21-255
	Code     int    `json:"code"`
	Comment  string `json:"comment"`
	DhcpName string `json:"dhcp_name"`
	Integer  int    `json:"integer"`
	Name     string `json:"name"`
	// Scope can be one of: []string{"global", "server", "host", "mac", "vendor"}
	Scope string `json:"scope"`
	// Type can be one of: []string{"ip-address", "text", "string", "integer"}
	Type   string `json:"type"`
	Vendor string `json:"vendor"`
}

var _ sophos.RestGetter = &DhcpOption6{}

// GetPath implements sophos.RestObject and returns the DhcpOption6s GET path
// Returns all available dhcp/option6 objects
func (*DhcpOption6s) GetPath() string { return "/api/objects/dhcp/option6/" }

// RefRequired implements sophos.RestObject
func (*DhcpOption6s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpOption6s GET path
// Returns all available option6 types
func (d *DhcpOption6) GetPath() string {
	return fmt.Sprintf("/api/objects/dhcp/option6/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *DhcpOption6) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpOption6 DELETE path
// Creates or updates the complete object option6
func (*DhcpOption6) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option6/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpOption6 PATCH path
// Changes to parts of the object option6 types
func (*DhcpOption6) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option6/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpOption6 POST path
// Create a new dhcp/option6 object
func (*DhcpOption6) PostPath() string {
	return "/api/objects/dhcp/option6/"
}

// PutPath implements sophos.RestObject and returns the DhcpOption6 PUT path
// Creates or updates the complete object option6
func (*DhcpOption6) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option6/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpOption6) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/option6/%s/usedby", ref)
}

// DhcpServers is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpServers []DhcpServer

// DhcpServer is a generated Sophos object
type DhcpServer struct {
	Locked          string   `json:"_locked"`
	Reference       string   `json:"_ref"`
	ObjectType      string   `json:"_type"`
	Address         string   `json:"address"`
	Comment         string   `json:"comment"`
	Custom          string   `json:"custom"`
	DefaultGateway  string   `json:"default_gateway"`
	DenyUnknown     bool     `json:"deny_unknown"`
	DNS1            string   `json:"dns1"`
	DNS2            string   `json:"dns2"`
	Domain          string   `json:"domain"`
	Interface       string   `json:"interface"`
	LeaseTime       int64    `json:"lease_time"`
	Mappings        []string `json:"mappings"`
	Name            string   `json:"name"`
	Netmask         int64    `json:"netmask"`
	ProxyAutoconfig bool     `json:"proxy_autoconfig"`
	RangeEnd        string   `json:"range_end"`
	RangeStart      string   `json:"range_start"`
	RelayMode       bool     `json:"relay_mode"`
	Status          bool     `json:"status"`
	Wins            string   `json:"wins"`
	WinsNodeType    string   `json:"wins_node_type"`
}

var _ sophos.RestGetter = &DhcpServer{}

// GetPath implements sophos.RestObject and returns the DhcpServers GET path
// Returns all available dhcp/server objects
func (*DhcpServers) GetPath() string { return "/api/objects/dhcp/server/" }

// RefRequired implements sophos.RestObject
func (*DhcpServers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpServers GET path
// Returns all available server types
func (d *DhcpServer) GetPath() string { return fmt.Sprintf("/api/objects/dhcp/server/%s", d.Reference) }

// RefRequired implements sophos.RestObject
func (d *DhcpServer) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpServer DELETE path
// Creates or updates the complete object server
func (*DhcpServer) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpServer PATCH path
// Changes to parts of the object server types
func (*DhcpServer) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpServer POST path
// Create a new dhcp/server object
func (*DhcpServer) PostPath() string {
	return "/api/objects/dhcp/server/"
}

// PutPath implements sophos.RestObject and returns the DhcpServer PUT path
// Creates or updates the complete object server
func (*DhcpServer) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpServer) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server/%s/usedby", ref)
}

// GetType implements sophos.Object
func (d *DhcpServer) GetType() string { return d.ObjectType }

// DhcpServer6s is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpServer6s []DhcpServer6

// DhcpServer6 represents a UTM DHCPv6 server
type DhcpServer6 struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	DefaultLft int    `json:"default_lft"`
	// RangeEnd description: (IP6ADDR)
	RangeEnd string `json:"range_end"`
	// Address description: REF(itfparams/*)
	Address string `json:"address"`
	// Domain default value is ""
	Domain string `json:"domain"`
	// Interface description: REF(interface/ethernet), REF(interface/vlan), REF(interface/bridge)
	Interface string `json:"interface"`
	// OnLink default value is false
	OnLink   bool `json:"on_link"`
	PrefdLft int  `json:"prefd_lft"`
	// RelayMode default value is false
	RelayMode bool `json:"relay_mode"`
	// Custom default value is ""
	Custom string `json:"custom"`
	// Dns2 description: (IP6ADDR)
	// Dns2 default value is "::"
	Dns2     string        `json:"dns2"`
	Mappings []interface{} `json:"mappings"`
	// Mtu description: Constraints: 0, 1280-9000
	Mtu int `json:"mtu"`
	// RangeStart description: (IP6ADDR)
	RangeStart string `json:"range_start"`
	// Status default value is false
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
	// Dns1 description: (IP6ADDR)
	// Dns1 default value is "::"
	Dns1     string `json:"dns1"`
	Name     string `json:"name"`
	Netmask6 int    `json:"netmask6"`
	// ProxyAutoconfig default value is false
	ProxyAutoconfig bool `json:"proxy_autoconfig"`
	ValidLft        int  `json:"valid_lft"`
	// DenyUnknown default value is false
	DenyUnknown bool `json:"deny_unknown"`
}

var _ sophos.RestGetter = &DhcpServer6{}

// GetPath implements sophos.RestObject and returns the DhcpServer6s GET path
// Returns all available dhcp/server6 objects
func (*DhcpServer6s) GetPath() string { return "/api/objects/dhcp/server6/" }

// RefRequired implements sophos.RestObject
func (*DhcpServer6s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpServer6s GET path
// Returns all available server6 types
func (d *DhcpServer6) GetPath() string {
	return fmt.Sprintf("/api/objects/dhcp/server6/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *DhcpServer6) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpServer6 DELETE path
// Creates or updates the complete object server6
func (*DhcpServer6) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server6/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpServer6 PATCH path
// Changes to parts of the object server6 types
func (*DhcpServer6) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server6/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpServer6 POST path
// Create a new dhcp/server6 object
func (*DhcpServer6) PostPath() string {
	return "/api/objects/dhcp/server6/"
}

// PutPath implements sophos.RestObject and returns the DhcpServer6 PUT path
// Creates or updates the complete object server6
func (*DhcpServer6) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server6/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpServer6) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/server6/%s/usedby", ref)
}

// DhcpStatelesss is an Sophos Endpoint subType and implements sophos.RestObject
type DhcpStatelesss []DhcpStateless

// DhcpStateless represents a UTM IPv6 prefix advertisement
type DhcpStateless struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// StatelessServerStatus default value is true
	StatelessServerStatus bool `json:"stateless_server_status"`
	// Dns2 description: (IP6ADDR)
	// Dns2 default value is "::"
	Dns2 string `json:"dns2"`
	// Domain default value is ""
	Domain string `json:"domain"`
	// ManagedFlag default value is false
	ManagedFlag bool   `json:"managed_flag"`
	Name        string `json:"name"`
	Comment     string `json:"comment"`
	// Interface description: REF(interface/ethernet), REF(interface/vlan), REF(interface/bridge)
	Interface string `json:"interface"`
	PrefdLft  int    `json:"prefd_lft"`
	// ProxyAutoconfig default value is false
	ProxyAutoconfig bool `json:"proxy_autoconfig"`
	DefaultLft      int  `json:"default_lft"`
	// Dns1 description: (IP6ADDR)
	// Dns1 default value is "::"
	Dns1 string `json:"dns1"`
	// OtherConfig default value is false
	OtherConfig bool `json:"other_config"`
	// Status default value is false
	Status   bool `json:"status"`
	ValidLft int  `json:"valid_lft"`
	// Address description: REF(itfparams/*)
	Address string `json:"address"`
	// Custom default value is ""
	Custom string `json:"custom"`
	// Mtu description: Constraints: 0, 1280-9000
	Mtu int `json:"mtu"`
	// OnLink default value is false
	OnLink bool `json:"on_link"`
}

var _ sophos.RestGetter = &DhcpStateless{}

// GetPath implements sophos.RestObject and returns the DhcpStatelesss GET path
// Returns all available dhcp/stateless objects
func (*DhcpStatelesss) GetPath() string { return "/api/objects/dhcp/stateless/" }

// RefRequired implements sophos.RestObject
func (*DhcpStatelesss) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DhcpStatelesss GET path
// Returns all available stateless types
func (d *DhcpStateless) GetPath() string {
	return fmt.Sprintf("/api/objects/dhcp/stateless/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *DhcpStateless) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DhcpStateless DELETE path
// Creates or updates the complete object stateless
func (*DhcpStateless) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/stateless/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DhcpStateless PATCH path
// Changes to parts of the object stateless types
func (*DhcpStateless) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/stateless/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DhcpStateless POST path
// Create a new dhcp/stateless object
func (*DhcpStateless) PostPath() string {
	return "/api/objects/dhcp/stateless/"
}

// PutPath implements sophos.RestObject and returns the DhcpStateless PUT path
// Creates or updates the complete object stateless
func (*DhcpStateless) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/stateless/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DhcpStateless) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dhcp/stateless/%s/usedby", ref)
}
